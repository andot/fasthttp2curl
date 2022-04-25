# fasthttp2curl
Convert fasthttp.Request to CURL command line


## Example

```go
import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/andot/fasthttp2curl"
)

req := fasthttp.AcquireRequest()
req.Header.SetMethod(fasthttp.MethodPut)
req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
req.SetBodyString(`{"hello":"world","answer":42}`)
req.Header.SetContentType("application/json")
command := GetCurlCommand(req)
fasthttp.ReleaseRequest(req)
fmt.Println(command)

// Output:
// curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
```

## Install

```bash
go get github.com/andot/fasthttp2curl
```

## License

© 2022 [Ma Bingyao](https://coolcode.org)

Licensed under the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE)).
