package zabbix_test

import (
	"io/ioutil"
	"testing"

	http "github.com/alexjx/gopher-lua-libs/http"
	inspect "github.com/alexjx/gopher-lua-libs/inspect"
	zabbix "github.com/alexjx/gopher-lua-libs/zabbix"
	lua "github.com/yuin/gopher-lua"
)

func TestApi(t *testing.T) {
	data, err := ioutil.ReadFile("./test/test_api.lua")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
	state := lua.NewState()
	zabbix.Preload(state)
	http.Preload(state)
	inspect.Preload(state)
	if err := state.DoString(string(data)); err != nil {
		t.Fatalf("execute test: %s\n", err.Error())
	}
}
