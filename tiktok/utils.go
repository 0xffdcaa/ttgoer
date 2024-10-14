package tiktok

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/playwright-community/playwright-go"
)

const ExampleTikTokURL = "tiktok.com/@example/video/74510392"

var tikTokURLRegex = regexp.MustCompile(`^(https:\/\/www\.)?tiktok\.com\/@[^\/]+\/video\/\d+.*$`)

func IsValidTikTokURL(url string) bool {
	return strings.Contains(url, "vm.tiktok.com") || tikTokURLRegex.MatchString(url)
}

func extractUsernameFromUrl(url string) (string, error) {
	parts := strings.Split(url, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, "@") {
			return part, nil
		}
	}
	return "", fmt.Errorf("no username found in url: %v", url)
}

func notifyOnRelevantMediaFound(page playwright.Page) <-chan playwright.Response {
	found := false
	urlFoundCh := make(chan playwright.Response, 1)

	page.On("response", func(response playwright.Response) {
		if !found && response.Request().ResourceType() == "media" {
			if !isRelevantVideoURL(response.URL()) {
				return
			}
			found = true
			urlFoundCh <- response
			close(urlFoundCh)
		}
	})

	return urlFoundCh
}

func isRelevantVideoURL(url string) bool {
	return strings.Contains(url, "mime_type=video_mp4")
}
