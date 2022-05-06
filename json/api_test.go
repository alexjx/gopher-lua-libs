package json

import (
	"testing"

	"github.com/alexjx/gopher-lua-libs/strings"
	"github.com/alexjx/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"

	inspect "github.com/alexjx/gopher-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		Preload,
		inspect.Preload,
		strings.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
