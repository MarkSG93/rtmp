package rtmp_test

import (
	"testing"

	"github.com/MarkSG93/rtmp"
	"github.com/MarkSG93/rtmp/server"
	"github.com/stretchr/testify/assert"
)

func TestNewServerMakesNewServers(t *testing.T) {
	s, err := rtmp.NewServer(":1935")

	assert.IsType(t, new(server.Server), s)
	assert.Nil(t, err)
}
