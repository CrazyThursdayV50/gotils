package goo

import "fmt"

func goError(f func()) (err error) {
	defer func() {
		er := recover()
		if er == nil {
			return
		}

		var ok bool
		err, ok = er.(error)
		if ok {
			return
		}

		err = fmt.Errorf("%v", er)
	}()

	go f()
	return
}

func Go(f func()) {
	_ = goError(f)
}

func Goo(f func(), h func(error)) {
	h(goError(f))
}
