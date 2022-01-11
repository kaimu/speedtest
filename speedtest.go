package speedtest

import (
	"github.com/kaimu/speedtest/providers/netflix"
	"github.com/kaimu/speedtest/providers/ookla"
)

// Netflix returns Fast.com download result in Mb/s
func Netflix() (down float64, err error) {
	return netflix.Fetch()
}

// Ookla returns speedtest.net download and upload results in Mb/s
func Ookla() (down float64, up float64, err error) {
	return ookla.Fetch()
}
