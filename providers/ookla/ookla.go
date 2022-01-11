package ookla

import (
	"github.com/kaimu/speedtest/providers/ookla/adapter"
	"github.com/showwin/speedtest-go/speedtest"
)

//go:generate mockgen -destination=mock/mock_ookla.go . Ookla
type Ookla interface {
	FetchUserInfo() (*speedtest.User, error)
	FetchServerList(*speedtest.User) (speedtest.ServerList, error)
	FindServer(speedtest.ServerList, []int) ([]adapter.OoklaServer, error)
}

func Fetch() (down float64, up float64, err error) {
	return run(adapter.Ookla{})
}

func run(o Ookla) (down float64, up float64, err error) {
	user, err := o.FetchUserInfo()
	if err != nil {
		return
	}

	serverList, err := o.FetchServerList(user)
	if err != nil {
		return
	}

	// using only the first server in the list
	targets, err := o.FindServer(serverList, []int{0})
	if err != nil {
		return
	}

	for _, s := range targets {
		err = s.DownloadTest(false)
		if err != nil {
			return
		}

		err = s.UploadTest(false)
		if err != nil {
			return
		}

		down, up = s.Results()
	}

	return
}
