# Teleman
Telegram cli tool for bot notifications
I used telegram botfather for this cli app


# export env variables


export TELEGRAM_API_TOKEN="apiToken"

export TELEGRAM_CHAT_ID="chat_id"


## How to Use ?

```
echo "hello" | teleman
```
```
cat subdomains.txt | nuclei -t ./nuclei-templates -silent | teleman
```

You will get a notification in your telegram bot if you setup the token and chat it correctly

# Install

go get github.com/balook/teleman

## I have used bot father of Telegram

[Generate Access Token and Chat_id](https://blog.r0b.re/automation/bash/2020/06/30/setup-telegram-notifications-for-your-shell.html)
