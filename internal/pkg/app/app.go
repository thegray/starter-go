package app

import (
	"sync"
	"time"
)

type App interface {
	Start()
	Stop()
}

func AppController(apps ...App) (stopFn func(duration time.Duration)) {
	var wg sync.WaitGroup

	for _, t := range apps {
		wg.Add(1)
		go t.Start()
	}

	return func(duration time.Duration) {
		for _, t := range apps {
			go func(t App) {
				t.Stop()
				wg.Done()
			}(t)
		}

		select {
		case <-time.After(duration):
			return
		case <-func() chan struct{} {
			c := make(chan struct{})
			go func() {
				wg.Wait()
				close(c)
			}()
			return c
		}():
			return
		}
	}
}
