package hungry

import "sync"

type singleton struct {
}

var ins *singleton
var once sync.Once

func GetInstance() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

func GetInstance2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
