package netflix

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_netflix "github.com/kaimu/speedtest/providers/netflix/mock"
	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	var (
		req            = require.New(t)
		mockController = gomock.NewController(t)
		testErr        = errors.New("test error")
	)

	tests := []struct {
		name    string
		setup   func() Netflix
		want    float64
		wantErr bool
	}{
		{
			name: "Init error",
			setup: func() Netflix {
				mock := mock_netflix.NewMockNetflix(mockController)
				mock.EXPECT().Init().Return(testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "GetUrls error",
			setup: func() Netflix {
				mock := mock_netflix.NewMockNetflix(mockController)
				mock.EXPECT().Init()
				mock.EXPECT().GetUrls().Return(nil, testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "Measure error",
			setup: func() Netflix {
				mock := mock_netflix.NewMockNetflix(mockController)
				mock.EXPECT().Init()
				mock.EXPECT().GetUrls()
				mock.EXPECT().Measure(gomock.Any(), gomock.Any()).Return(testErr)
				return mock
			},
			wantErr: true,
		},
		{
			name: "OK",
			setup: func() Netflix {
				mock := mock_netflix.NewMockNetflix(mockController)
				mock.EXPECT().Init()
				mock.EXPECT().GetUrls()
				mock.EXPECT().Measure(gomock.Any(), gomock.Any()).DoAndReturn(func(urls []string, passes chan<- float64) error {
					passes <- 10 * 1024
					passes <- 15 * 1024 // expected result
					passes <- 5 * 1024
					close(passes)
					return nil
				})
				return mock
			},
			want:    15,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.setup())

			if !tt.wantErr {
				req.NoError(err)
				req.Equal(tt.want, got)
			} else {
				req.Error(err)
				req.EqualError(err, testErr.Error())
			}
		})
	}
}
