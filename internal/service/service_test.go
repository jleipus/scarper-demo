package service_test

import (
	"context"
	"scaper-demo/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	counter := 0
	mockRun := func(ctx context.Context) error {
		counter++
		return nil
	}

	service.Run(mockRun)

	require.Equal(t, 1, counter)
}
