# chirno

`chirno` provides health_check api for any application.

```
　　　　／^? 　 　 　,.へ___
　　　/　　　＞''´￣￣｀'''ヽ7
　　　|　 ／´ _　　　　　　　_'ヽ､
　　　〉 /　/´　 / 　,　 ､ 　､　ヽ〉
　　/　 i ｲ 　レレ ハノ! Ji i　　i
　 └ｒｲ　レｲ 　 ┃　　┃　|　　 ｎ⌒i
　　く_ﾉ　 〉　i　　 　　 　 　 iﾊ _,,.!　ノ
　　　ﾊ.　i　ハ､ 　 　o　 　 人|´　／
　 i?ﾚﾍハﾚへ〉'=r--r='i´Vヽy'´
　 ヽ､,_｀ヽ,r'´　｀ﾄ､ ∞」 i　　ノ
　 ＜　 ￣〉､___ノ　 ￣　 Y／
　　　＞_/　／〉　,　 , 　 ､!_ゝ
　 　 ｀(⊆ノ／　/　　! 　 ハ
　　　 　　くﾍ,.へ_,.へ__,.ﾍ,.ﾍ
　　　 　　　｀'r､__ﾊ___ﾊ__!ン
　　　 　　　　ﾄ_ﾝ　　　ﾄ_ﾉ
```

# installation
```
$ go get github.com/ieee0824/chirno
```

# configuration
using json.

## options

### backend_url
Specifying this parameter operates as a reverse proxy.

### command
Execution command of the application to add health_api.

### port
api's listen port.

## example

```json
{
    "backend_url": "http://localhost:8081",
    "command": [
        "writing",
        "exec",
        "command",
        "options"
    ],
    "port": "8080"
}
```

