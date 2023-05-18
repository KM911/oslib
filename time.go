package oslib

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

func (t *timer) End() {
	end := time.Now()
	fmt.Println("花费时间为：", end.Sub(t.start))

}
