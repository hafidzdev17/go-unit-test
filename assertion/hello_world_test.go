package assertion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestSubTest(t *testing.T) {
	t.Run("hafid", func(t *testing.T) {
		result := HelloWorld("hafid")
		assert.Equal(t, "hello hafid", result, "result must be 'hello hafid'")
	})
	t.Run("deddy", func(t *testing.T) {
		result := HelloWorld("deddy")
		assert.Equal(t, "hello deddy", result, "result must be 'hello deddy'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run on Mac Os")
	}
	result := HelloWorld("hafid")
	assert.Equal(t, "hello hafid", result, "result must be 'hello hafid' ")
}

func TestAssert(t *testing.T) {
	result := HelloWorld("hafid")
	assert.Equal(t, "hello hafid", result, "result must be 'hello hafid' ")
	fmt.Println("test hello world done")
}

func TestRequire(t *testing.T) {
	result := HelloWorld("hafid")
	require.Equal(t, "hello hafid", result, "result must be 'hello hafid' ")
	fmt.Println("test hello can't be printed")
}
