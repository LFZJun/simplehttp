# HTTP Client

eg. search github
```go
import "github.com/ljun20160606/simplehttp"

func main() {
     simplehttp.
            Get("https://github.com/search").
            Query("utf8", "✓").
            Query("q", "httpclient").
            Send().
            String()
}
```

see [example/github.go](./example/github.go)

