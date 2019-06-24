package lazy

type singleton struct {
}

var ins *singleton = &singleton{}

func GetInstance() *singleton {
	return ins
}
