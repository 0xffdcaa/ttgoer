package tiktok

import (
	"encoding/json"

	"ttgoer/log"

	"github.com/playwright-community/playwright-go"
)

func parseCookies(input string) ([]playwright.OptionalCookie, error) {
	var cookies []playwright.OptionalCookie

	if err := json.Unmarshal([]byte(input), &cookies); err != nil {
		log.S().Errorf("failed to parse cookies: %v", err)
		return nil, err
	}

	for i := range cookies {
		if cookies[i].SameSite == nil {
			cookies[i].SameSite = playwright.SameSiteAttributeNone
		}

		switch *cookies[i].SameSite {
		case "none", "no_restriction":
			cookies[i].SameSite = playwright.SameSiteAttributeNone
		case "lax":
			cookies[i].SameSite = playwright.SameSiteAttributeLax
		case "strict":
			cookies[i].SameSite = playwright.SameSiteAttributeStrict
		}
	}

	return cookies, nil
}
