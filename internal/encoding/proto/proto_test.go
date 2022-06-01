package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	c := new(codec)
	assert.Equal(t, c.Name(), "proto")
}

func TestCodec(t *testing.T) {
	_ = new(codec)
}
