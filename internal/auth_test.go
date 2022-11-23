package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_auth(t *testing.T) {
	got := auth("thomas", "123456")
	require.True(t, got)
	got = auth("abc", "123456")
	require.False(t, got)
}
