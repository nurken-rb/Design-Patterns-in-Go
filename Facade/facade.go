package facade

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height,
		make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// a facade over buffers and viewports
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)

	fmt.Printf("u: %v\n", u)
}

//{"threads":[{"position":452,"start":0,"end":451,"connection":"closed"},{"position":452,"start":452,"end":901,"connection":"open"}],"url":"https://a.udemycdn.com/2019-12-18_10-05-13-2f2ce3a4f2a94b665c3129275ae1cb99/original.txt?nva=20200311123216&token=0b8434fac63885b929429","method":"GET","port":443,"downloadSize":901,"headers":{"date":"Wed, 11 Mar 2020 08:08:37 GMT","content-type":"text/plain","content-length":"901","connection":"close","etag":"\"6d0f91e510adf6746a70ee24d245622e\"","last-modified":"Wed, 18 Dec 2019 10:05:15 GMT","server":"AmazonS3","x-amz-id-2":"b816TaZj4hIt3ZxzbWbUHG6XzEdh1l7yiBIxn4cLFW+kF4Q1Tv+6bdhBnaAUgx+Jm797UiQx6sU=","x-amz-meta-qqfilename":"structural.facade.facade.go.txt","x-amz-replication-status":"COMPLETED","x-amz-request-id":"796EFC20C3E2C972","x-amz-version-id":"sxrsDEv0ierJSk.VjrD3RoKI9nf0r93T","access-control-allow-origin":"*","age":"0","accept-ranges":"bytes"}}
