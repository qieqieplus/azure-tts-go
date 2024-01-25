package azure_tts

import "context"

type SynthesizeTask struct {
	ctx    context.Context
	cancel context.CancelFunc
	Audio  chan []byte
	Event  chan any
	Error  chan error
}

func newSynthesizeTask(parent context.Context) SynthesizeTask {
	ctx, cancel := context.WithCancel(parent)
	return SynthesizeTask{
		ctx:    ctx,
		cancel: cancel,
		Audio:  make(chan []byte),
		Event:  make(chan any),
		Error:  make(chan error),
	}
}

func (t SynthesizeTask) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t SynthesizeTask) Close() {
	t.cancel()
}
