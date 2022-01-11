package adapter

import (
	"github.com/showwin/speedtest-go/speedtest"
)

//go:generate mockgen -destination=mock/mock_adapter.go . OoklaServer
type OoklaServer interface {
	DownloadTest(bool) error
	UploadTest(bool) error
	Results() (float64, float64)
}

// Ookla encapsulates speedtest-go package top-level functions to implement parent `Ookla` interface (acts as an adapter to allow mock-testing)
type Ookla struct{}

func (Ookla) FetchUserInfo() (*speedtest.User, error) {
	return speedtest.FetchUserInfo()
}

func (Ookla) FetchServerList(user *speedtest.User) (speedtest.ServerList, error) {
	return speedtest.FetchServerList(user)
}

func (Ookla) FindServer(sl speedtest.ServerList, serverID []int) (result []OoklaServer, err error) {
	servers, err := sl.FindServer(serverID)
	if err != nil {
		return nil, err
	}
	for _, s := range servers {
		result = append(result, ooklaServer{s})
	}
	return
}

// ooklaServer wraps original speedtest-go struct to implement `OoklaServer` interface
type ooklaServer struct {
	*speedtest.Server
}

func (o ooklaServer) Results() (dl float64, ul float64) {
	return o.DLSpeed, o.ULSpeed
}
