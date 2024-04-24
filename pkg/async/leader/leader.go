package leader

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
)

type Leader[J any] struct {
	ctx          context.Context
	deliveryChan api.ChanAPI[J]
}

func (b *Leader[J]) Do(job J) {
	b.deliveryChan.Send(job)
}

func (b *Leader[J]) AddWorker(worker *worker.Worker[J]) {
	worker.WithContext(b.ctx)
	worker.WithGraceful(true)
	worker.WithTrigger(b.deliveryChan)
	worker.Run()
}

func New[J any](ctx context.Context, sendTimeout, recvTimeout time.Duration) *Leader[J] {
	var bucket Leader[J]
	bucket.ctx = ctx
	bucket.deliveryChan = gchan.Make[J](0)
	bucket.deliveryChan.Timeout(sendTimeout, recvTimeout)
	return &bucket
}
