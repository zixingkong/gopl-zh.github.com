//练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
//并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：
//func LimitReader(r io.Reader, n int64) io.Reader

// ex7.5 provides a LimitReader that reports EOF at a given offset.
package reader

import (
	"io"
)

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
