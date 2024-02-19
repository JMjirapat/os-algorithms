package PFR

import "fmt"

type LRU struct {
	queue    []int
	capacity int
	refStr   []int
}

func NewLRU(cap int, refStr []int) PFR {
	return &LRU{
		queue:    make([]int, 0),
		capacity: cap,
		refStr:   refStr,
	}
}

func (l *LRU) Append(el int) {
	if len(l.queue) >= l.capacity {
		l.queue = l.queue[1:]
	}
	l.queue = append(l.queue, el)
}

func (l *LRU) Remove() {
	l.queue = l.queue[1:]
}

func (l LRU) IsExisted(el int) bool {
	for _, e := range l.queue {
		if el == e {
			return true
		}
	}
	return false
}

func (l *LRU) Print() {
	fmt.Print("[ ")
	for idx, el := range l.queue {
		fmt.Printf("%v", el)
		if idx != len(l.queue)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}

func (l *LRU) Run() (page_fault int) {
	page_fault = 0
	for _, i := range l.refStr {
		if !l.IsExisted(i) {
			l.Append(i)
			page_fault = page_fault + 1
		} else {
			l.MoveToBack(i)
		}
	}
	return page_fault
}

func (l *LRU) MoveToBack(el int) {
	prev := l.queue
	for idx, e := range prev {
		if e == el {
			if idx != len(l.queue)-1 {
				copy(prev[idx:], prev[idx+1:])
			}
			l.queue[len(l.queue)-1] = e
			break
		}
	}
}
