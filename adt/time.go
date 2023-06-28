package adt

import (
	"fmt"
	"time"
)

// 计时器
type timer struct {
	start time.Time
}

func TimerStart() *timer {
	t := timer{}
	t.start = time.Now()
	return &t
}

/*
获取花费时间
*/
func (t *timer) Cost() int64 {
	cost := time.Now().Sub(t.start)
	t.start = time.Now()
	return cost.Milliseconds()
}

func (t *timer) End() {
	end := time.Now()
	fmt.Println("花费时间为：", end.Sub(t.start))
}
