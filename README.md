# discord-notion-expander

Discord内に投稿した非公開なNotionページのEmbedを表示するためのBotです。

## Dockerおよびdocker-composeを使ったセットアップ

```bash
docker-compose build
docker-compose up -d
```

で一発で上がるようになっています。

## 自分でビルドしてセットアップ

go 1.16をインストールした上でビルドする。

```bash
go build
./discord-notion-expander
```

## .envや環境変数周り

```dotenv
BASE_NOTION_URL=https://www.notion.so/hogehoge/
NOTION_TOKEN=<NOTION_TOKEN_V2>
DISCORD_TOKEN=<DISCORD_BOT_TOKEN>
```

を設定してください。

`NOTION_TOKEN` にはNotionのAPIを使用するため、Notionを開いているブラウザから取れるCookieの `token_v2` を設定してください。

`DISCORD_TOKEN` にはDiscordのBotの管理ページから閲覧できる、Build-A-Botの中にある `TOKEN` を設定してください。

## DiscordのBOTのサーバへの追加

```
https://discordapp.com/oauth2/authorize?client_id=<CLIENT_ID>&scope=bot&permissions=121856
```

上記URLをブラウザから開き、サーバへ追加してください。 

`CLIENT_ID` にはDiscordのBotの管理ページから確認できる、OAuth2の `CLIENT ID`を指定してください。

**また、Botの管理画面で忘れずに PUBLIC BOT の設定をオフにするようにしてください。** 

意図せぬサーバに追加されると、プライベートなページの中身の一部が見られてしまう可能性があります。必ず自分だけがサーバにBotを追加できる状態にしてください。

## 仕様
- Discordの仕様によりURLが貼り付けられた投稿そのもののEmbedを書き換えることは出来ません。

これは他人が投稿したメッセージを編集することが出来ない仕様によるものです。そのため、Embedは別のメッセージとして投稿されます。
