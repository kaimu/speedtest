package ookla

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kaimu/speedtest/providers/ookla/adapter"
	mock_adapter "github.com/kaimu/speedtest/providers/ookla/adapter/mock"
	mock_ookla "github.com/kaimu/speedtest/providers/ookla/mock"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	var (
		req            = require.New(t)
		mockController = gomock.NewController(t)
		testErr        = errors.New("test error")
	)

	tests := []struct {
		name     string
		setup    func() Ookla
		wantDown float64
		wantUp   float64
		wantErr  bool
	}{
		{
			name: "FetchUserInfo error",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				mock.EXPECT().FetchUserInfo().Return(nil, testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "FetchServerList error",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				mock.EXPECT().FetchUserInfo()
				mock.EXPECT().FetchServerList(gomock.Any()).Return(speedtest.ServerList{}, testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "FindServer error",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				mock.EXPECT().FetchUserInfo()
				mock.EXPECT().FetchServerList(gomock.Any())
				mock.EXPECT().FindServer(gomock.Any(), gomock.Any()).Return(nil, testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "DownloadTest error",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				mock.EXPECT().FetchUserInfo()
				mock.EXPECT().FetchServerList(gomock.Any())
				mockServer := mock_adapter.NewMockOoklaServer(mockController)
				mock.EXPECT().FindServer(gomock.Any(), gomock.Any()).Return([]adapter.OoklaServer{mockServer}, nil)
				mockServer.EXPECT().DownloadTest(false).Return(testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "UploadTest error",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				mock.EXPECT().FetchUserInfo()
				mock.EXPECT().FetchServerList(gomock.Any())
				mockServer := mock_adapter.NewMockOoklaServer(mockController)
				mock.EXPECT().FindServer(gomock.Any(), gomock.Any()).Return([]adapter.OoklaServer{mockServer}, nil)
				mockServer.EXPECT().DownloadTest(false)
				mockServer.EXPECT().UploadTest(false).Return(testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "OK",
			setup: func() Ookla {
				mock := mock_ookla.NewMockOokla(mockController)
				user := speedtest.User{}
				mock.EXPECT().FetchUserInfo().Return(&user, nil)
				servers := speedtest.ServerList{}
				mock.EXPECT().FetchServerList(&user).Return(servers, nil)
				mockServer := mock_adapter.NewMockOoklaServer(mockController)
				mock.EXPECT().FindServer(servers, []int{0}).Return([]adapter.OoklaServer{mockServer}, nil)
				mockServer.EXPECT().DownloadTest(false)
				mockServer.EXPECT().UploadTest(false)
				mockServer.EXPECT().Results().Return(10.0, 9.0)
				return mock
			},
			wantDown: 10,
			wantUp:   9,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDown, gotUp, err := run(tt.setup())

			if !tt.wantErr {
				req.NoError(err)
				req.Equal(tt.wantUp, gotUp)
				req.Equal(tt.wantDown, gotDown)
			} else {
				req.Error(err)
				req.EqualError(err, testErr.Error())
			}
		})
	}
}
