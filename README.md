# livecam-map-backend

## 概要

[ライブカメラ地図](https://livecam-map.vercel.app/)のバックエンド(Go+Gin+Postgres)

[フロントエンド](https://github.com/otsuneko/livecam-map)

- ~~日本中~~世界中のライブカメラ([日本](https://livecam.asia/)、[海外](https://www.cametan.com/world/))を地図上にプロットする
- 雨雲レーダーを重畳する
- 雨が降ってる地点のライブカメラ映像を眺めると落ち着くので良い

## ToDo

- [ ] gin+postgresでライブカメラ情報をJSONで返すRestAPIサーバ実装
- [ ] 雨雲レーダー画像生成処理作成
- [ ] 雨雲に覆われたエリアに含まれるライブカメラ一覧の抽出機能作成
- [ ] YouTube Liveの配信切れURL検知+配信切れ動画の新URLを自動で再登録する機能作成  
⇒YouTube Data APIを使ってチャンネルに含まれる動画一覧を取得し、類似タイトルの動画があればそのURLに書き換え、といった処理で実現可能と想像

## 参考
- [Gin Web Framework](https://gin-gonic.com/ja/docs/)
- [YouTube Data APIを使ってチャンネルに含まれる動画を取得する流れ](https://zenn.dev/yorifuji/articles/youtube-data-api)
