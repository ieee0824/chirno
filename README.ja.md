# chirno

```
　　　　／^? 　 　 　,.へ___
　　　/　　　＞''´￣￣｀'''ヽ7
　　　|　 ／´ _　　　　　　　_'ヽ､
　　　〉 /　/´　 / 　,　 ､ 　､　ヽ〉
　　/　 i ｲ 　レレ ハノ! Ji i　　i
　 └ｒｲ　レｲ 　 ┃　　┃　|　　 ｎ⌒i
　　く_ﾉ　 〉　i　　 　　 　 　 iﾊ _,,.!　ノ
　　　ﾊ.　i　ハ､ 　 　o　 　 人|´　／
　 i?ﾚﾘハﾚへ〉'=r--r='i´Vヽy'´
　 ヽ､,_｀ヽ,r'´　｀ﾄ､ ∞」 i　　ノ
　＜　 ￣〉､___ノ　 ￣　 Y／
　　　＞_/　／〉　,　 , 　 ､!_ゝ
　 　 ｀(⊆ノ／　/　　! 　 ハ
　　　 　　くﾍ,.へ_,.へ__,.ﾍ,.ﾍ
　　　 　　　｀'r､__ﾊ___ﾊ__!ン
　　　 　　　　ﾄ_ﾝ　　　ﾄ_ﾉ
```

任意のアプリケーションにヘルスチェック API (`/health_check`) を追加するラッパーツールです。
指定したコマンドを子プロセスとして実行しつつ、ヘルスチェック用の HTTP エンドポイントを提供します。
`backend_url` を指定すればリバースプロキシとしても動作します。

[English](README.md)

## Installation

```
go install github.com/ieee0824/chirno@latest
```

## Usage

```
chirno -f conf.json
```

| フラグ | デフォルト | 説明 |
|---|---|---|
| `-f` | `./conf.json` | 設定ファイルのパス |

## Configuration

JSON 形式の設定ファイルを使用します。

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `command` | `string[]` | yes | 実行するコマンドと引数 |
| `port` | `string` | yes | chirno がリッスンするポート |
| `backend_url` | `string` | no | 指定するとリバースプロキシとして動作 |

### Example: アプリケーションにヘルスチェックを追加

```json
{
    "command": ["go", "run", "app.go"],
    "port": "9090"
}
```

`GET /health_check` で `It works!` が返ります。

### Example: リバースプロキシ + ヘルスチェック

```json
{
    "backend_url": "http://localhost:8080",
    "command": ["go", "run", "server.go"],
    "port": "9090"
}
```

`/health_check` 以外のリクエストは `backend_url` に転送されます。
