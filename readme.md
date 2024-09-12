# Golang Vagas Bot para BlueSky

Este é um bot desenvolvido em Go para repostar automaticamente vagas de emprego relacionadas a Golang no BlueSky. Ele busca por postagens que contenham as seguintes hashtags:

- #govagas
- #golangvagas
- #vagasgolang
- #vagasgo
- #gojobs

## Funcionalidades

1. **Repostagem Automática**: Sempre que o bot encontrar uma postagem com uma das hashtags acima, ele fará um repost automaticamente.
2. **Curtir Postagem**: Além de repostar, o bot também dará um "like" na postagem original.

## Como Rodar o Bot

Para rodar o bot localmente, siga os seguintes passos:

1. Certifique-se de ter o Go instalado em sua máquina. Você pode verificar isso rodando o comando:

    ```bash
      go version
    ```

2. Clone o repositório e entre no diretório:

    ```bash
      git clone https://github.com/seu-usuario/golang-vagas-bot.git
      cd golang-vagas-bot
    ```

3. Execute o bot:

    ```bash
      go run cmd/main.go
    ```

4. Crie um arquivo `.env` com as seguintes credenciais:

    ```bash
      BLUESKY_IDENTIFIER=<seu_identificador>
      BLUESKY_PASSWORD=<seu_app_password>
    ```

> **Atenção**: **Nunca use a senha da sua conta no BlueSky, utilize um App Password**.

O bot utiliza Websocket para monitorar as postagens no [BlueSky](https://bsky.app).

Para alterar as hashtags monitoradas altere em `internal/utils/filter-term.go`

<img src="https://cdn.bsky.app/img/avatar_thumbnail/plain/did:plc:hf37h3zvhdcw7jjik6rd43ws/bafkreifr7wqzf5fagkpbcwcc27cpiploqqslkc3dut255ja46hoxiudnse@jpeg" alt="Golamg Jobs" width="50"/> [Perfil](https://bsky.app/profile/govagas.bsky.social) do Bot no BlueSky
