package adt

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	println(Timestamp())
}

func TestUnixTime(t *testing.T) {
	time.AfterFunc(time.Second*1, func() {
		println(UnixTime())
	})
}
