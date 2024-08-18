//练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
//实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）

// ex7.4 provides a simple string reader.
package reader

import (
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
