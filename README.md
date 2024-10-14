# ttgoer

This is personal project, a bot capable of downloading videos from popular platforms such as TT.
Currently, it only supports TT, but it is planned to support YT as well.

## Usage
In order to use it, you will need to deploy it somewhere, provide a Telegram bot token and install Playwright.

The executable expects **config.yaml** to be put in the same directory.

### config.yaml
```yaml
bot:
  token: "<TOKEN-FROM-BOT-FATHER>"
  poller_timeout: 1s
  admin_id: 1234564
  # Optional. The user with admin_id is always allowed
  allowed_user_ids:
    - 1233556
tik_tok:
  download_timeout: 3s
  shutdown_timeout: 20s
  download_max_retries: 5
```

## How It Works (internally)
Currently, the flow is simple and unreliable:
1. Start Chrome WebDriver (via Playwright)
2. Go to TT page
3. Try to extract the media (video) URL
4. Download the video and send it to the requester

It's the simplest flow I could come up with, yet it works in
most cases. **ttgoer** will automatically retry video download
up to a configured amount of times. Usually, 5 retries is enough.

