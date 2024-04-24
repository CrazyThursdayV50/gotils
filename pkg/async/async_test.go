package async

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/consumer"
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/logger"
	defaultLogger "github.com/CrazyThursdayV50/gotils/pkg/logger/default"
)

func TestAsync(t *testing.T) {
	l := defaultLogger.DefaultLogger()
	l.Debug("test start")
	type testData struct {
		name  string
		value int
	}

	ctx := context.TODO()
	ch := gchan.Make[*testData](0)
	conf := consumer.NewConfig[*testData]("")
	conf.SetChannel(ch)
	conf.SetContext(ctx)
	conf.SetLogger(l)
	conf.SetHandler(func(e *testData, logger logger.Logger) {
		logger.Info("receive data: %v", e)
		time.Sleep(time.Second * 1)
	})

	consumer := conf.Build()
	consumer.Run()
	goo.Goo(func() {
		for i := 0; i < 1000; i++ {
			go func(i int) {
				ch.Send(&testData{name: fmt.Sprintf("%X", i), value: i})
			}(i)
		}
	}, func(err error) {
		if err == nil {
			return
		}
		l.Error("%v", err)
	})

	time.Sleep(time.Microsecond * 100)
	l.Info("exit")
	ch.Close()
	consumer.Stop()
	time.Sleep(time.Second * 10)
}
