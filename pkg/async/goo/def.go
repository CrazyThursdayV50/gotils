package goo

import (
	"fmt"
	"runtime"
	"strings"
)

func try(f func()) (err error) {
	defer func() {
		er := recover()
		if er == nil {
			return
		}

		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		location := locatePanic(buf[:n])
		err = fmt.Errorf("%v, %s", er, strings.Join(location, " <- "))
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
