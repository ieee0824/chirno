# chirno

`chirno` provides health_check api for any application.

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

