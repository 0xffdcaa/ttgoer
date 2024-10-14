package main

import (
	"os"
	"os/signal"
	"syscall"

	"ttgoer/bot"
	"ttgoer/log"
	"ttgoer/tiktok"
)

func main() {
	log.S().Info("starting")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	tiktok.Init()
	bot.Start()

	_ = tiktok.SetCookiesFromJson(`[
    {
        "domain": ".tiktok.com",
        "expirationDate": 1729719015.420247,
        "hostOnly": false,
        "httpOnly": false,
        "name": "msToken",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "stJlK_LAWaDPKlYS3rw2ph2qnDSGiPJce2SmtMsnUwiuozjvZmFPUi3kAYk1oD83JrG_NpsSWKJRCstk-r-ljucvzQky2gV7GcoLgTHqRLOqI0OdFE8UlwFCzP2IgX95n8EGf8A4z9S0HoYW5k3tCTyX"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1759957092.774149,
        "hostOnly": false,
        "httpOnly": true,
        "name": "sid_guard",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1da5a202f24859ed0f0978a2450c95ab%7C1728853092%7C15552000%7CFri%2C+11-Apr-2025+20%3A58%3A12+GMT"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1760389092.774054,
        "hostOnly": false,
        "httpOnly": true,
        "name": "ttwid",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1%7C32ng0IZH5jzVcittNnp-HkFTCq4vVJ4vqs50MXE-epQ%7C1728853081%7Cd1bad3ee360e92aaa50df4f8821d48bfa453b155bde7d43ccddf5e062ba4641e"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774173,
        "hostOnly": false,
        "httpOnly": true,
        "name": "uid_tt",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "ead586be520f591ad05f996e4613c371b1c67379f4ec844399984e6716a6e350"
    },
    {
        "domain": ".tiktok.com",
        "hostOnly": false,
        "httpOnly": false,
        "name": "s_v_web_id",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": true,
        "storeId": null,
        "value": "verify_m282hpst_zNsHLy2g_KU5z_4uxh_Bppw_0bmFGSHSDl49"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774323,
        "hostOnly": false,
        "httpOnly": true,
        "name": "ssid_ucp_v1",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1.0.0-KDk4Y2I0NGY2MzE2ZjRmMjE1ZDk2ZWY4ZmUxOWMwZjZkYzg2ODdiYzcKIQiGiLyE-_mUhWcQ5OiwuAYYswsgDDD7p6m4BjgIQBJIBBADGgZtYWxpdmEiIDFkYTVhMjAyZjI0ODU5ZWQwZjA5NzhhMjQ1MGM5NWFi"
    },
    {
        "domain": ".www.tiktok.com",
        "expirationDate": 1754774352,
        "hostOnly": false,
        "httpOnly": false,
        "name": "tiktok_webapp_theme",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "dark"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1728860282.446082,
        "hostOnly": false,
        "httpOnly": false,
        "name": "bm_sv",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "7E5C5F43AB1EA0EF0F21D0B32B75EF82~YAAQFMUTAiY5zIGSAQAAxRvIhxn4ynJsTNQ8Qwjt+4v4I2i82Hd8HGKKarI2if+9Ue2YoqcPYXIyoH849y4yT7pOJp+qyMjFj7Bt4qMKMiijFrIbSbwAFh0IKm9nMseecY+ce8cirwSQjRANZDc5afg9ZOXz5LNg5eQSTlqttXdyJAZtFPzJ3QsbaQ1wGYzm8ycnEGcsS2sKYW1Tl9M4/08pve7a6kEu07ZEb1pz5Ia6ghi1MXeFQo7yfNjtj88+~1"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1734037092.774029,
        "hostOnly": false,
        "httpOnly": true,
        "name": "cmpl_token",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "AgQQAPP4F-RO0rc0vWg4vd0__W2teOYav4dzYNQzfg"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1734037092.773881,
        "hostOnly": false,
        "httpOnly": true,
        "name": "multi_sids",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "7424839087159182342%3A1da5a202f24859ed0f0978a2450c95ab"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1731445092.774119,
        "hostOnly": false,
        "httpOnly": true,
        "name": "passport_auth_status_ss",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "0679a253bd9a023b9a9e40ba69330cee%2C"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1734037084.949498,
        "hostOnly": false,
        "httpOnly": false,
        "name": "passport_csrf_token",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "e690de81eaeb61e1fe8d22e3225194ef"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774251,
        "hostOnly": false,
        "httpOnly": true,
        "name": "sessionid",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1da5a202f24859ed0f0978a2450c95ab"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774273,
        "hostOnly": false,
        "httpOnly": true,
        "name": "sessionid_ss",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1da5a202f24859ed0f0978a2450c95ab"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774225,
        "hostOnly": false,
        "httpOnly": true,
        "name": "sid_tt",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1da5a202f24859ed0f0978a2450c95ab"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.774298,
        "hostOnly": false,
        "httpOnly": true,
        "name": "sid_ucp_v1",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "1.0.0-KDk4Y2I0NGY2MzE2ZjRmMjE1ZDk2ZWY4ZmUxOWMwZjZkYzg2ODdiYzcKIQiGiLyE-_mUhWcQ5OiwuAYYswsgDDD7p6m4BjgIQBJIBBADGgZtYWxpdmEiIDFkYTVhMjAyZjI0ODU5ZWQwZjA5NzhhMjQ1MGM5NWFi"
    },
    {
        "domain": ".www.tiktok.com",
        "expirationDate": 1754774352,
        "hostOnly": false,
        "httpOnly": false,
        "name": "tiktok_webapp_theme_source",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "auto"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744406891.445999,
        "hostOnly": false,
        "httpOnly": true,
        "name": "tt_chain_token",
        "path": "/",
        "sameSite": null,
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "xEPIgFPx5OaYX8Fe11x12w=="
    },
    {
        "domain": ".tiktok.com",
        "hostOnly": false,
        "httpOnly": true,
        "name": "tt_csrf_token",
        "path": "/",
        "sameSite": "lax",
        "secure": true,
        "session": true,
        "storeId": null,
        "value": "X16Q69vD-0jSyoU_iIiURn0TTFdurPQKU7Uk"
    },
    {
        "domain": ".tiktok.com",
        "expirationDate": 1744405092.7742,
        "hostOnly": false,
        "httpOnly": true,
        "name": "uid_tt_ss",
        "path": "/",
        "sameSite": "no_restriction",
        "secure": true,
        "session": false,
        "storeId": null,
        "value": "ead586be520f591ad05f996e4613c371b1c67379f4ec844399984e6716a6e350"
    }
]`)

	<-quit
	tiktok.GracefulShutdown()
	bot.Stop()
	log.S().Info("stopping")
}
