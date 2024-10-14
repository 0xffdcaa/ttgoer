package tiktok

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"ttgoer/cfg"
	"ttgoer/log"
	"ttgoer/utils"

	"github.com/playwright-community/playwright-go"
	"go.uber.org/zap"
)

var pw *playwright.Playwright
var browser playwright.Browser
var httpClient = &http.Client{}
var downloadTimeout = cfg.Get().TikTok.DownloadTimeout
var shutdownTimeout = cfg.Get().TikTok.ShutdownTimeout
var fetchMaxTries = cfg.Get().TikTok.DownloadMaxRetries
var tracker = newTracker()
var shutdownInitiated = false
var cookies []playwright.OptionalCookie

func Init() {
	var err error

	if err := playwright.Install(&playwright.RunOptions{
		DriverDirectory: "driver/",
		Browsers:        []string{"chromium"},
	}); err != nil {
		log.S().Panicf("could not install Chromium: %v", err)
	}

	pw, err = playwright.Run()
	if err != nil {
		log.S().Panicf("could not start Playwright: %v", err)
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args: []string{
			"--disable-notifications",
			"--disable-geolocation",
		},
	})
	if err != nil {
		log.S().Panicf("could not launch browser: %v", err)
	}
}

func SetCookiesFromJson(json string) error {
	newCookies, err := parseCookies(json)

	if err != nil {
		return err
	}

	cookies = newCookies

	log.S().Info("Updating global cookies")

	return nil
}

func GracefulShutdown() {
	log.S().Infow("attempting graceful shutdown", zap.Int("inProgress", tracker.inProgressCount()))
	logInProgressRequests()

	shutdownInitiated = true

	waitForRequestsToFinish()

	if err := browser.Close(); err != nil {
		log.S().Warnf("could not close the browser: %v", err)
	}

	if err := pw.Stop(); err != nil {
		log.S().Warnf("could not stop Playwright: %v", err)
	}
}

func logInProgressRequests() {
	inProgress := tracker.getInProgress()
	if len(inProgress) > 0 {
		log.S().Infof("waiting for %d requests to finish:", len(inProgress))
		for i, req := range inProgress {
			log.S().Infof("[%d] - %s (%d) is in-progress for %v",
				i,
				req.tikTokURL,
				req.username,
				time.Since(req.addedAt),
			)
		}
		log.S().Infoln()
	}
}

func waitForRequestsToFinish() {
	timeoutCh := time.After(shutdownTimeout)
	allDoneCh := make(chan bool)
	waiterCancelled := false
	go func() {
		for tracker.inProgressCount() > 0 {
			time.Sleep(2 * time.Second)

			if waiterCancelled {
				return
			}

			logInProgressRequests()
		}

		log.S().Info("all requests have been finished")
		allDoneCh <- true
	}()

	select {
	case <-allDoneCh:
		// Nothing
	case <-timeoutCh:
		waiterCancelled = true
		log.S().Warnf("shutdown timeout of %v reached, %d still in-progress, forcing shutdown...",
			shutdownTimeout, tracker.inProgressCount())
	}
}

func Fetch(tikTokURL string, username string) (*TikTok, error) {
	if shutdownInitiated {
		return nil, fmt.Errorf("shutdown initiated")
	}

	log.S().
		With(zap.String("username", username)).
		With(zap.String("tikTokURL", tikTokURL)).
		With(zap.String("goroutineName", utils.GetGoroutineName())).
		Infof("start fetching TikTok")

	tracker.track(tikTokURL, username)
	defer tracker.untrack(tikTokURL)

	var tryCount uint = 1
	startTime := time.Now()
	tt, err := innerFetch(tikTokURL)

	for err != nil && tryCount < fetchMaxTries {
		log.S().
			With(zap.String("username", username)).
			With(zap.String("tikTokURL", tikTokURL)).
			With(zap.Uint("triesLeft", fetchMaxTries-tryCount)).
			Errorf("retrying fetch due to: %v", err)
		tt, err = innerFetch(tikTokURL)
		tryCount++
	}

	if err == nil {
		log.S().
			With(zap.String("took", time.Since(startTime).String())).
			With(zap.Float32("sizeMiB", tt.FileSizeMiB)).
			With(zap.String("tikTokURL", tikTokURL)).
			Infof("fetched TikTok after %d tries", tryCount)
	} else {
		log.S().
			With(zap.String("took", time.Since(startTime).String())).
			With(zap.String("tikTokURL", tikTokURL)).
			Errorf("failed to fetched TikTok after %d tries", tryCount)
	}

	return tt, err
}

func innerFetch(tiktokURL string) (*TikTok, error) {
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: randomUserAgent(),
	})
	defer context.Close()

	if err != nil {
		log.S().Fatalf("could not create new context: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.S().Fatalf("could not create new page: %v", err)
	}

	if cookies != nil {
		// TODO: Cookies
		// if err = page.Context().AddCookies(cookies); err != nil {
		// 	log.S().Warnf("failed to add cookies: %v", err)
		// }
	}

	if _, err := page.Goto(tiktokURL); err != nil {
		return nil, fmt.Errorf("could not navigate to page: %v", err)
	}

	var videoDownloadURL string
	var videoRequest playwright.Request
	mediaURLCh := notifyOnRelevantMediaFound(page)
	captchaCh := notifyOnCaptcha(page)
	timeoutCh := time.After(downloadTimeout)

	select {
	case <-captchaCh:
		page.Pause()
		return nil, fmt.Errorf("captcha")
	case response := <-mediaURLCh:
		videoDownloadURL = response.URL()
		videoRequest = response.Request()
	case <-timeoutCh:
		page.Pause()
		return nil, fmt.Errorf("timeout")
	}

	if videoDownloadURL == "" {
		return nil, fmt.Errorf("video URL not found: %s", tiktokURL)
	}

	headers, err := videoRequest.AllHeaders()
	if err != nil {
		return nil, err
	}

	buf, err := downloadFile(videoDownloadURL, headers)
	if err != nil {
		return nil, fmt.Errorf("could not request video: %v", err)
	}

	// TODO?
	// user, _ := extractUsernameFromUrl(tiktokURL)

	tikTok := TikTok{
		// TODO: Extract description and username
		Description: "",
		Url:         tiktokURL,
		User:        "",
		FileBytes:   buf,
		FileSizeMiB: float32(len(buf)) / 1024 / 1024,
	}

	return &tikTok, nil
}

func notifyOnCaptcha(page playwright.Page) <-chan bool {
	var ch = make(chan bool, 1)
	go func() {
		defer close(ch)

		err := page.Locator("div.captcha-verify-container").WaitFor(playwright.LocatorWaitForOptions{
			State:   playwright.WaitForSelectorStateVisible,
			Timeout: playwright.Float(float64(downloadTimeout.Milliseconds())),
		})
		if err != nil {
			return
		}

		ch <- true
	}()
	return ch
}

func downloadFile(url string, headers map[string]string) ([]byte, error) {
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		if !strings.HasPrefix(key, ":") {
			httpReq.Header.Add(key, value)
		}
	}

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return nil, fmt.Errorf("unexpected status '%s' for url: %s", resp.Status, url)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
