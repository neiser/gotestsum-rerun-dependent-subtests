package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sayHello(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		require.NoError(t, sayHello(""))
	})
	t.Run("some name", func(t *testing.T) {
		require.NoError(t, sayHello("Buddy"))
	})
	t.Run("Voldemort makes it fail every second time", func(t *testing.T) {
		require.NoError(t, sayHello("Voldemort"))
	})
}
