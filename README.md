# go-secure-cookies
A secure cookie module for golang

```go

// securekey must be at least 16 char long
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
cookies, err := NewCookies(key)

var ck http.Cookie
ck.Name = "test1"
ck.Value = "this-is-a-test-and-more"

//write----
w  http.ResponseWriter
err := cookies.write(w, ck)


var name = "test1"

//read----
r    *http.Request
cookieValue, err := cookies.Read(r, name)


```