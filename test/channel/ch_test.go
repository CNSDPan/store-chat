package channel

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type S struct {
	bs []*Bucket
}

type Bucket struct {
	l  []*List
	ch chan string
}
type List struct {
	name string
	ch   chan string
}

var Serve *S

func TestCh(t *testing.T) {
	var s = new(S)
	var bs = make([]*Bucket, 0)
	for i := 0; i < 4; i++ {
		bs = append(bs, NewB(i))
	}
	s.bs = bs
	Serve = s
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				for _, tB := range Serve.bs {
					tB.ch <- "hello"
				}
				log.Printf("结束了\n\n")
			}
		}
	}()

	select {}
}
func NewB(i int) *Bucket {
	b := &Bucket{
		l:  make([]*List, 0),
		ch: make(chan string, 100),
	}
	for ii := 0; ii < 3; ii++ {
		b.l = append(b.l, NewL(i, ii))
	}
	b.BCh()
	return b
}
func (b *Bucket) BCh() {
	go func() {
		for {
			select {
			case s := <-b.ch:
				go func() {
					for _, l := range b.l {
						l.ch <- s
					}
				}()

			}
		}
	}()
}
func NewL(i int, ii int) *List {
	l := &List{
		name: fmt.Sprintf("%d-%d", i, ii),
		ch:   make(chan string, 100),
	}
	l.LCh()
	return l
}
func (l *List) LCh() {
	go func() {
		for {
			select {
			case s := <-l.ch:
				log.Printf("%s:%s", l.name, s)
			}
		}
	}()
}
