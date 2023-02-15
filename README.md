# go-secure-cookies
A secure cookie module for golang

```go
//write----
var key = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"
cookies, err := NewCookies(key)

var ck http.Cookie
ck.Name = "test1"
ck.Value = "this-is-a-test-and-more"

   w  http.ResponseWriter
   err := cookies.write(w, ck)


    var name = "test1"

    r    *http.Request
    cookieValue, err := cookies.Read(r, name)


```