# Golang Auth proxy example

This is just an example on how to use the `httputil.ReverseProxy` type provided in the golang standard library.

1. Open a terminal in the root folder and run `go run target/main.go`. This will start the target server. It will only respond to request from `localhost:8081`

2. Open another terminal in the root folder and run the following commands:

```sh
# first export some env variables
export SCHEMA="http"
export HOST="localhost:8080"
export Auth="123456"

# then run proxy server
run go proxy/main.go
```

This will start our auth proxy in `localhost:8081` and will redirect traffic to `http://localhost:8080` only if there is an header like `"Auth: 123456"`

3. Open another terminal and start making some requests

4. Have fun

