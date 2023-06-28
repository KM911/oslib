package adt

// method ?  for []string
type StringSlice []string

func (s StringSlice) Pop() {
	lens := len(s)
	if lens == 0 {
		return
	} else {
		s = s[:lens-1]
	}
}

func (s StringSlice) Out() {
	lens := len(s)
	if lens == 0 {
		return
	} else {
		s = s[1:]
	}
}

func (s StringSlice) Remove(str string) []string {
	lens := len(s)
	if lens == 0 {
		return s
	} else {
		index := -1
		for i, v := range s {
			if v == str {
				index = i
				break
			}
		}
		println("index is ", index)
		switch index {
		case -1:
			return s
		case 0:
			return s[1:]
		case lens - 1:
			return s[:lens-1]
		default:
			return append(s[:index], s[index+1:]...)
		}
	}
}
