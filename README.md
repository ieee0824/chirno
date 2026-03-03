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

A wrapper tool that adds a health check API (`/health_check`) to any application.
It runs a specified command as a child process while providing an HTTP endpoint for health checking.
Optionally acts as a reverse proxy when `backend_url` is set.

[Japanese (日本語)](README.ja.md)

## Installation

```
go install github.com/ieee0824/chirno@latest
```

## Usage

```
chirno -f conf.json
```

| Flag | Default | Description |
|---|---|---|
| `-f` | `./conf.json` | Path to config file |

## Configuration

Uses a JSON config file.

| Field | Type | Required | Description |
|---|---|---|---|
| `command` | `string[]` | yes | Command and arguments to execute |
| `port` | `string` | yes | Port for chirno to listen on |
| `backend_url` | `string` | no | Enables reverse proxy when set |

### Example: Add health check to an application

```json
{
    "command": ["go", "run", "app.go"],
    "port": "9090"
}
```

`GET /health_check` returns `It works!`.

### Example: Reverse proxy + health check

```json
{
    "backend_url": "http://localhost:8080",
    "command": ["go", "run", "server.go"],
    "port": "9090"
}
```

Requests other than `/health_check` are forwarded to `backend_url`.
