package client

import (
	"fmt"
	"time"

	"github.com/leliuga/cdk/types"
)

func NewProgress(name string, total uint64) *Progress {
	return &Progress{
		name:      name,
		total:     total,
		current:   0,
		startTime: time.Now(),
		stopCh:    nil,
	}
}

func (p *Progress) Start(frequency time.Duration) {
	p.stopCh = make(chan struct{}, 1)
	p.wg.Add(1)
	go func() {
		stopCh := p.stopCh

		for {
			select {
			case <-stopCh:
				p.report()
				close(stopCh)
				p.wg.Done()
				return
			case <-time.After(frequency):
				p.report()
			}
		}
	}()
}

func (p *Progress) Stop() {
	p.stopCh <- struct{}{}
	p.wg.Wait()
	p.stopCh = nil
}

func (p *Progress) Write(b []byte) (int, error) {
	n := len(b)
	p.current += uint64(n)

	return n, nil
}

func (p *Progress) report() {
	elapsed := time.Since(p.startTime).Seconds()
	speed := float64(p.current) / 1024 / 1024 / elapsed
	percent := float64(p.current*100) / float64(p.total)

	switch {
	case p.current == p.total:
		fmt.Printf("\rDownloaded %s successfully in %.0f seconds.", p.name, elapsed)
	default:
		fmt.Printf("\rDownloading %s %.2f MiB/sec, %s of %s (%.2f%%)", p.name, speed, types.BytesSize(p.current), types.BytesSize(p.total), percent)
	}

}
