package fasthttp2curl

import (
	"bytes"

	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

var (
	https             = []byte(`https`)
	singleQuote       = []byte(`'`)
	escapeSingleQuote = []byte(`'\''`)
	pool              = bytebufferpool.Pool{}
)

func GetCurlCommand(req *fasthttp.Request) (string, error) {
	buf := pool.Get()
	defer pool.Put(buf)

	buf.WriteString("curl")

	uri := req.RequestURI()
	if bytes.HasPrefix(uri, https) {
		buf.WriteString(" -k")
	}

	buf.WriteString(" -X '")
	buf.Write(req.Header.Method())
	buf.WriteByte('\'')

	if body := req.Body(); len(body) > 0 {
		buf.WriteString(" -d '")
		buf.Write(bytes.Replace(body, singleQuote, escapeSingleQuote, -1))
		buf.WriteByte('\'')
	}

	req.Header.VisitAll(func(key, value []byte) {
		buf.WriteString(" -H '")
		buf.Write(bytes.Replace(key, singleQuote, escapeSingleQuote, -1))
		buf.WriteString(": ")
		buf.Write(bytes.Replace(value, singleQuote, escapeSingleQuote, -1))
		buf.WriteByte('\'')
	})

	buf.WriteString(" '")
	buf.Write(bytes.Replace(uri, singleQuote, escapeSingleQuote, -1))
	buf.WriteByte('\'')

	return buf.String(), nil
}
