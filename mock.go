package prohermes

import "github.com/stretchr/testify/mock"

type PromelvenMock struct {
	mock.Mock
}

func (m *PromelvenMock) NewPrometheusHook() (*PrometheusHook, error) {
	args := m.Called()
	return args.Get(0).(*PrometheusHook), args.Error(1)
}
