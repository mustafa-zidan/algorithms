package util

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
    assert.Equal(t, 1, Min(1, 2))
    assert.Equal(t, 1, Min(2, 1))
    assert.Equal(t, 3, Min(3, 3))
}

func TestMax(t *testing.T) {
    assert.Equal(t, 2, Max(1, 2))
    assert.Equal(t, 2, Max(2, 1))
    assert.Equal(t, 3, Max(3, 3))
}
