package oslib

import (
	"fmt"
	"time"
)

type TimeCounter struct {
	start time.Time
}

func (t *TimeCounter) Start() *TimeCounter {
	t.start = time.Now()
	return t
}

func (t *TimeCounter) End() {
	end := time.Now()
	fmt.Println("花费时间为：", end.Sub(t.start))
}
