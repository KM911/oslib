package pprof

import "github.com/pkg/profile"

// 就是这样存在一个问题 不方便我们进行测试

func CpuProfile(mode uint) interface{ Stop() } {
	switch mode {
	case 1:
		return profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	case 2:
		return profile.Start(profile.MemProfile, profile.ProfilePath("."))
	case 3:
		return profile.Start(profile.BlockProfile, profile.ProfilePath("."))
	case 4:
		return profile.Start(profile.MutexProfile, profile.ProfilePath("."))
	default:
		return profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	}
}
