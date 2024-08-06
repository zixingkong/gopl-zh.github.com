//练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
//返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针

// ex7.2 wraps a writer to count written words.
package counter

import (
	"io"
)

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.written += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.written
}
