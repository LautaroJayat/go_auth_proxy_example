# Strategy package

Here you can put any kind of strategy you want. The only constraints is that you must return an `httprouter.Hanlder` which is just a function with the following signature `func (w http.ResponseWriter, r *http.Request, _ httprouter.Params)`.

Inside, that function it can do any kind of checks:
* Check cookies
* Use some service to check in a persistence layer
* Do some signature checkings
* What you wish

If you want to give an early response, just use the `http.ResponseWriter`. If you want to fordward the response to the `httputil.ReverseProxy` just pass it as an argument to the high order function and then usit in the returned function.

The most basic implementation is the following one:

```go
// We pass a reference to our ReverseProxy
func BasicOne(rp *httputil.ReverseProxy) httprouter.Handle{
	// We return a new function that contains a copy of that pointer
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		// We serve the request usign the Reverse Proxy
		rp.ServeHTTP(w, r)
	}
}
```

