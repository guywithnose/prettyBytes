package prettyBytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	assert.Equal(t, "70B", Format(70))
}

func TestKilobytes(t *testing.T) {
	assert.Equal(t, "68.36KB", Format(70000))
}

func TestMegabytes(t *testing.T) {
	assert.Equal(t, "66.76MB", Format(70000000))
}

func TestGigabytes(t *testing.T) {
	assert.Equal(t, "65.19GB", Format(70000000000))
}

func TestTerrabytes(t *testing.T) {
	assert.Equal(t, "63.66TB", Format(70000000000000))
}

func TestPetabytes(t *testing.T) {
	assert.Equal(t, "62.17PB", Format(70000000000000000))
}

func TestYottabytes(t *testing.T) {
	assert.Equal(t, "60.72YB", Format(70000000000000000000))
}
