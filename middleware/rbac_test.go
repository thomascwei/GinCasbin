package middleware

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRBAC(t *testing.T) {
	result := RBAC("admin", "data", "GET")
	require.True(t, result)
}
