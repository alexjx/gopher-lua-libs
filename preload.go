package libs

import (
	cloudwatch "github.com/alexjx/gopher-lua-libs/aws/cloudwatch"
	"github.com/alexjx/gopher-lua-libs/base64"
	cert_util "github.com/alexjx/gopher-lua-libs/cert_util"
	chef "github.com/alexjx/gopher-lua-libs/chef"
	cmd "github.com/alexjx/gopher-lua-libs/cmd"
	crypto "github.com/alexjx/gopher-lua-libs/crypto"
	db "github.com/alexjx/gopher-lua-libs/db"
	filepath "github.com/alexjx/gopher-lua-libs/filepath"
	goos "github.com/alexjx/gopher-lua-libs/goos"
	http "github.com/alexjx/gopher-lua-libs/http"
	humanize "github.com/alexjx/gopher-lua-libs/humanize"
	inspect "github.com/alexjx/gopher-lua-libs/inspect"
	ioutil "github.com/alexjx/gopher-lua-libs/ioutil"
	json "github.com/alexjx/gopher-lua-libs/json"
	log "github.com/alexjx/gopher-lua-libs/log"
	pb "github.com/alexjx/gopher-lua-libs/pb"
	plugin "github.com/alexjx/gopher-lua-libs/plugin"
	pprof "github.com/alexjx/gopher-lua-libs/pprof"
	prometheus "github.com/alexjx/gopher-lua-libs/prometheus/client"
	regexp "github.com/alexjx/gopher-lua-libs/regexp"
	runtime "github.com/alexjx/gopher-lua-libs/runtime"
	"github.com/alexjx/gopher-lua-libs/shellescape"
	"github.com/alexjx/gopher-lua-libs/stats"
	storage "github.com/alexjx/gopher-lua-libs/storage"
	strings "github.com/alexjx/gopher-lua-libs/strings"
	tac "github.com/alexjx/gopher-lua-libs/tac"
	tcp "github.com/alexjx/gopher-lua-libs/tcp"
	telegram "github.com/alexjx/gopher-lua-libs/telegram"
	template "github.com/alexjx/gopher-lua-libs/template"
	time "github.com/alexjx/gopher-lua-libs/time"
	xmlpath "github.com/alexjx/gopher-lua-libs/xmlpath"
	yaml "github.com/alexjx/gopher-lua-libs/yaml"
	zabbix "github.com/alexjx/gopher-lua-libs/zabbix"

	lua "github.com/yuin/gopher-lua"
)

// Preload preload all gopher lua packages
func Preload(L *lua.LState) {
	base64.Preload(L)
	time.Preload(L)
	strings.Preload(L)
	filepath.Preload(L)
	ioutil.Preload(L)
	http.Preload(L)
	regexp.Preload(L)
	tac.Preload(L)
	inspect.Preload(L)
	yaml.Preload(L)
	plugin.Preload(L)
	cmd.Preload(L)
	json.Preload(L)
	tcp.Preload(L)
	xmlpath.Preload(L)
	db.Preload(L)
	cert_util.Preload(L)
	runtime.Preload(L)
	shellescape.Preload(L)
	telegram.Preload(L)
	zabbix.Preload(L)
	pprof.Preload(L)
	prometheus.Preload(L)
	pb.Preload(L)
	crypto.Preload(L)
	goos.Preload(L)
	storage.Preload(L)
	humanize.Preload(L)
	chef.Preload(L)
	template.Preload(L)
	cloudwatch.Preload(L)
	log.Preload(L)
	stats.Preload(L)
}
