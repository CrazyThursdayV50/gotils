package cron

func (c *Cron) Run() {
	c.worker.Run()
	c.runOnStart()
	if c.waitAfterRun {
		c.timerRun()
	} else {
		c.tickRun()
	}
}

func (c *Cron) Stop()                 { c.worker.Stop() }
func (c *Cron) OnStart(f func())      { c.worker.OnStart(f) }
func (c *Cron) OnExit(f func())       { c.worker.OnExit(f) }
func (c *Cron) Done() <-chan struct{} { return c.worker.Done() }
