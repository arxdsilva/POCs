package funcaspar

func RandomFunc(s string, f func(string) string) string {
	return f(s)
}

func RandomFuncFF(s string, f ff) string {
	return f(s)
}

type ff func(string) string
