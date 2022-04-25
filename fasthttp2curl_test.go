package fasthttp2curl

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/valyala/fasthttp"
)

func ExampleGetCurlCommand() {
	form := url.Values{}
	form.Add("age", "10")
	form.Add("name", "Hudson")
	body := form.Encode()
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI("http://foo.com/cats")
	req.SetBodyString(body)
	req.Header.Set("API_KEY", "123")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)
	// Output:
	// curl -X 'POST' -d 'age=10&name=Hudson' -H 'Api_key: 123' 'http://foo.com/cats'
}

func ExampleGetCurlCommand_json() {
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
}

func ExampleGetCurlCommand_noBody() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPut)
	req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output:
	// curl -X 'PUT' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_emptyStringBody() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPut)
	req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	req.SetBodyString("")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output:
	// curl -X 'PUT' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_newlineInBody() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	req.SetBodyString("hello\nworld")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output:
	// curl -X 'POST' -d 'hello
	// world' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_specialCharsInBody() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	req.SetBodyString(`Hello $123 o'neill -"-`)
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output:
	// curl -X 'POST' -d 'Hello $123 o'\''neill -"-' -H 'Content-Type: application/json' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_other() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPut)
	req.SetRequestURI("http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	req.SetBodyString(`{"hello":"world","answer":42}`)
	req.Header.Set("X-Auth-Token", "private-token")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output: curl -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' -H 'X-Auth-Token: private-token' 'http://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

func ExampleGetCurlCommand_https() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPut)
	req.SetRequestURI("https://www.example.com/abc/def.ghi?jlk=mno&pqr=stu")
	req.Header.SetContentType("application/json")
	req.SetBodyString(`{"hello":"world","answer":42}`)
	req.Header.Set("X-Auth-Token", "private-token")
	command := GetCurlCommand(req)
	fasthttp.ReleaseRequest(req)
	fmt.Println(command)

	// Output: curl -k -X 'PUT' -d '{"hello":"world","answer":42}' -H 'Content-Type: application/json' -H 'X-Auth-Token: private-token' 'https://www.example.com/abc/def.ghi?jlk=mno&pqr=stu'
}

// Benchmark test for GetCurlCommand
func BenchmarkGetCurlCommand(b *testing.B) {
	form := url.Values{}
	form.Add("number", "1")
	body := form.Encode()
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI("http://foo.com")
	req.SetBodyString(body)
	for i := 0; i <= b.N; i++ {
		_ = GetCurlCommand(req)

	}
}
