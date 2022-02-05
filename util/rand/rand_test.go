package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	ss := RandStringCharset(10, "A")
	if ss != "AAAAAAAAAA" {
		t.Errorf("Expected 10 As, but got %q", ss)
	}
	ss = RandStringCharset(5, "ABC123")
	if len(ss) != 5 {
		t.Errorf("Expected random string of length 10, but got %q", ss)
	}
}

func TestRandInt63(t *testing.T) {
	i1 := RandInt63()
	i2 := RandInt63()
	assert.NotEqual(t, i1, i2)
}
