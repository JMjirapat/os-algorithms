package PFR

import "fmt"

type FIFO struct {
	queue    []int
	capacity int
	refStr   []int
}

func NewFIFO(cap int, refStr []int) PFR {
	return &FIFO{
		queue:    make([]int, 0),
		capacity: cap,
		refStr:   refStr,
	}
}

func (f *FIFO) Append(el int) {
	if len(f.queue) >= f.capacity {
		f.queue = f.queue[1:]
	}
	f.queue = append(f.queue, el)
}

func (f *FIFO) Remove() {
	f.queue = f.queue[1:]
}

func (f *FIFO) Run() (page_fault int) {
	page_fault = 0
	for _, i := range f.refStr {
		if !f.IsExisted(i) {
			f.Append(i)
			page_fault = page_fault + 1
		}
		f.Print()
	}
	return page_fault
}

func (f FIFO) IsExisted(el int) bool {
	for _, e := range f.queue {
		if el == e {
			return true
		}
	}
	return false
}

func (f *FIFO) Print() {
	fmt.Print("[ ")
	for idx, el := range f.queue {
		fmt.Printf("%v", el)
		if idx != len(f.queue)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}
