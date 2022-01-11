package netflix

import (
	"sync"

	"github.com/ddo/go-fast"
)

//go:generate mockgen -destination=mock/mock_netflix.go . Netflix
type Netflix interface {
	Init() error
	GetUrls() ([]string, error)
	Measure([]string, chan<- float64) error
}

func Fetch() (float64, error) {
	nf := fast.New()
	return run(nf)
}

func run(nf Netflix) (float64, error) {
	err := nf.Init()
	if err != nil {
		return 0, err
	}

	urls, err := nf.GetUrls()
	if err != nil {
		return 0, err
	}

	passes := make(chan float64, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	var resultKbps float64
	go func() {
		for p := range passes {
			// selecting the highest speed of all passes
			if resultKbps < p {
				resultKbps = p
			}
		}

		wg.Done()
	}()

	err = nf.Measure(urls, passes)
	if err != nil {
		return 0, err
	}

	wg.Wait()

	return resultKbps / 1024, nil
}
