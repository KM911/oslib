package test

import (
	"github.com/KM911/oslib/pprof"
	"testing"
	"time"
)

func Test_cpu(t *testing.T) {
	t.Log("cpu")
	defer pprof.CpuProfile(1).Stop()
	time.Sleep(1 * time.Second)

}
