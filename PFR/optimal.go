package PFR

import "fmt"

type Optimal struct {
	queue    []int
	capacity int
	refStr   []int
}

func NewOptimal(cap int, refStr []int) PFR {
	return &Optimal{
		queue:    make([]int, 0),
		capacity: cap,
		refStr:   refStr,
	}
}

func (o *Optimal) Append(el int) {
	o.queue = append(o.queue, el)
}

func (o *Optimal) Predice(curr_idx int) (pred_idx int) {
	pred_idx = 0
	farthest := curr_idx
	for i, fv := range o.queue {
		for j := curr_idx; j < len(o.refStr); j++ {
			if o.refStr[j] == fv {
				if j > farthest {
					farthest = j
					pred_idx = i
				}
				break
			}
			if j == len(o.refStr)-1 {
				return i
			}
		}
	}
	return pred_idx
}

func (o *Optimal) Remove() {
	o.queue = o.queue[1:]
}

func (o *Optimal) Run() (page_fault int) {
	page_fault = 0
	for curr_idx, i := range o.refStr {
		if !o.IsExisted(i) {
			if len(o.queue) < o.capacity {
				o.Append(i)
			} else {
				pred_idx := o.Predice(curr_idx)
				o.queue[pred_idx] = i
			}
			page_fault = page_fault + 1
		}
		o.Print()
	}
	return page_fault
}

func (o Optimal) IsExisted(el int) bool {
	for _, e := range o.queue {
		if el == e {
			return true
		}
	}
	return false
}

func (o *Optimal) Print() {
	fmt.Print("[ ")
	for idx, el := range o.queue {
		fmt.Printf("%v", el)
		if idx != len(o.queue)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}
