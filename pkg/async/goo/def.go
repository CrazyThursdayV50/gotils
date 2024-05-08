package goo

import "fmt"

func try(f func()) (err error) {
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

	f()
	return
}

func Go(f func()) {
	go func() {
		try(f)
	}()
}

func Goo(f func(), h func(error)) {
	go func() {
		h(try(f))
	}()
}

func Try(f func()) {
	_ = try(f)
}

func TryE(f func()) error {
	return try(f)
}
