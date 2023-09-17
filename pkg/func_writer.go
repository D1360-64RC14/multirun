package pkg

type FuncWriter func(p []byte)

func (f FuncWriter) Write(p []byte) (int, error) {
	f(p)
	return len(p), nil
}
