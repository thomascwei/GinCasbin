package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenToken(t *testing.T) {
	got, err := GenToken("thomas")
	require.NoError(t, err)
	Logger.Info(got)
}
