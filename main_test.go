package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sayHello(t *testing.T) {
	someFlag := false
	t.Run("empty name", func(t *testing.T) {
		require.NoError(t, sayHello(""))
	})
	t.Run("some name", func(t *testing.T) {
		require.NoError(t, sayHello("Buddy"))
	})
	t.Run("Voldemort makes it fail every second time", func(t *testing.T) {
		require.NoError(t, sayHello("Voldemort"))
		someFlag = true
	})
	t.Run("check flag", func(t *testing.T) {
		require.True(t, someFlag)
	})
}
