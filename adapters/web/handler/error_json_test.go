package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)
	require.Equal(t, `{"message":"Hello json"}`, string(result))
}
