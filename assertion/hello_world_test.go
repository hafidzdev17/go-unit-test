package assertion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

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
