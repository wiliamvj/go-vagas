# Golang Jobs Bot for BlueSky

This is a Go bot designed to automatically repost job openings related to Golang on BlueSky.

It searches for posts containing the following hashtags:

- #govagas
- #golangvagas
- #vagasgolang
- #vagasgo
- #gojobs

## Features

1. **Automatic Reposting**: Whenever the bot finds a post with one of the hashtags above, it will automatically repost it.
2. **Like Post**: In addition to reposting, the bot will also like the original post.

## How to Run the Bot

To run the bot locally, follow these steps:

1. Make sure you have Go installed on your machine (if you don't use Docker). You can check this by running the command:

```bash
    go version
```

1. Clone the repository and navigate to the directory:

    ```bash
        git clone https://github.com/your-username/golang-vagas-bot.git
        cd golang-vagas-bot
    ```

2. Run the bot:

    ```bash
      go run cmd/main.go
    ```

3. Create a `.env` file with the following credentials:

    ```bash
      BLUESKY_IDENTIFIER=<your_identifier>
      BLUESKY_PASSWORD=<your_app_password>
    ```

> **Note**: **Never use your BlueSky account password, use an App Password instead**.

The bot uses Websocket to monitor posts on [BlueSky](https://bsky.app).

To change the monitored hashtags, modify the `internal/utils/filter-term.go` file.

<img src="https://cdn.bsky.app/img/avatar_thumbnail/plain/did:plc:hf37h3zvhdcw7jjik6rd43ws/bafkreifr7wqzf5fagkpbcwcc27cpiploqqslkc3dut255ja46hoxiudnse@jpeg" alt="Golamg Jobs" width="50"/> [Bot Profile](https://bsky.app/profile/govagas.bsky.social) on BlueSky
