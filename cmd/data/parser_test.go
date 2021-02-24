package data_test

import (
	"github.com/MarkSG93/rtmp/chunk"
	"github.com/MarkSG93/rtmp/cmd/data"
	"github.com/stretchr/testify/mock"
)

type MockParser struct {
	mock.Mock
}

var _ data.Parser = new(MockParser)

func (m *MockParser) Parse(c *chunk.Chunk) (data.Data, error) {
	args := m.Called(c)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(data.Data), args.Error(1)
}
