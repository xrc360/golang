package gjson

import (
	"testing"

	"github.com/xrcn/cg/test/gtest"
)

func Test_checkDataType(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
bb           = """
                   dig := dig;                         END;"""
`)
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`)
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
"cg.gvalid.rule.required"             = "The :attribute field is required"
"cg.gvalid.rule.required-if"          = "The :attribute field is required"
"cg.gvalid.rule.required-unless"      = "The :attribute field is required"
"cg.gvalid.rule.required-with"        = "The :attribute field is required"
"cg.gvalid.rule.required-with-all"    = "The :attribute field is required"
"cg.gvalid.rule.required-without"     = "The :attribute field is required"
"cg.gvalid.rule.required-without-all" = "The :attribute field is required"
"cg.gvalid.rule.date"                 = "The :attribute value is not a valid date"
"cg.gvalid.rule.date-format"          = "The :attribute value does not match the format :format"
"cg.gvalid.rule.email"                = "The :attribute value must be a valid email address"
"cg.gvalid.rule.phone"                = "The :attribute value must be a valid phone number"
"cg.gvalid.rule.telephone"            = "The :attribute value must be a valid telephone number"
"cg.gvalid.rule.passport"             = "The :attribute value is not a valid passport format"
"cg.gvalid.rule.password"             = "The :attribute value is not a valid passport format"
"cg.gvalid.rule.password2"            = "The :attribute value is not a valid passport format"
"cg.gvalid.rule.password3"            = "The :attribute value is not a valid passport format"
"cg.gvalid.rule.postcode"             = "The :attribute value is not a valid passport format"
"cg.gvalid.rule.resident-id"          = "The :attribute value is not a valid resident id number"
"cg.gvalid.rule.bank-card"            = "The :attribute value must be a valid bank card number"
"cg.gvalid.rule.qq"                   = "The :attribute value must be a valid QQ number"
"cg.gvalid.rule.ip"                   = "The :attribute value must be a valid IP address"
"cg.gvalid.rule.ipv4"                 = "The :attribute value must be a valid IPv4 address"
"cg.gvalid.rule.ipv6"                 = "The :attribute value must be a valid IPv6 address"
"cg.gvalid.rule.mac"                  = "The :attribute value must be a valid MAC address"
"cg.gvalid.rule.url"                  = "The :attribute value must be a valid URL address"
"cg.gvalid.rule.domain"               = "The :attribute value must be a valid domain format"
"cg.gvalid.rule.length"               = "The :attribute value length must be between :min and :max"
"cg.gvalid.rule.min-length"           = "The :attribute value length must be equal or greater than :min"
"cg.gvalid.rule.max-length"           = "The :attribute value length must be equal or lesser than :max"
"cg.gvalid.rule.between"              = "The :attribute value must be between :min and :max"
"cg.gvalid.rule.min"                  = "The :attribute value must be equal or greater than :min"
"cg.gvalid.rule.max"                  = "The :attribute value must be equal or lesser than :max"
"cg.gvalid.rule.json"                 = "The :attribute value must be a valid JSON string"
"cg.gvalid.rule.xml"                  = "The :attribute value must be a valid XML string"
"cg.gvalid.rule.array"                = "The :attribute value must be an array"
"cg.gvalid.rule.integer"              = "The :attribute value must be an integer"
"cg.gvalid.rule.float"                = "The :attribute value must be a float"
"cg.gvalid.rule.boolean"              = "The :attribute value field must be true or false"
"cg.gvalid.rule.same"                 = "The :attribute value must be the same as field :field"
"cg.gvalid.rule.different"            = "The :attribute value must be different from field :field"
"cg.gvalid.rule.in"                   = "The :attribute value is not in acceptable range"
"cg.gvalid.rule.not-in"               = "The :attribute value is not in acceptable range"
"cg.gvalid.rule.regex"                = "The :attribute value is invalid"
`)
		// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*".+"`, data))
		// fmt.Println(gregex.IsMatch(`^[\s\t\n\r]*[\w\-]+\s*:\s*\w+`, data))
		// fmt.Println(gregex.IsMatch(`[\s\t\n\r]+[\w\-]+\s*:\s*".+"`, data))
		// fmt.Println(gregex.IsMatch(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, data))
		// fmt.Println(gregex.MatchString(`[\n\r]+[\w\-\s\t]+\s*:\s*\w+`, string(data)))
		t.Assert(checkDataType(data), "toml")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
[default]
db.engine         = mysql
db.max.idle.conns = 5
db.max.open.conns = 100
allow_ips         =
api.key           =
api.secret        =
enable_tls        = false
concurrency.queue = 500
auth_secret       = 63358e6f3daf0e5775ec3fb4d2516b01d41530bf30960aa76972f6ce7e08552f
ca_file           =
cert_file         =
key_file          =
host_port         = 8088
log_path          = /Users/zhaosuji/go/src/git.medlinker.com/foundations/gocron/log
#k8s-api地址(只提供内网访问)
k8s-inner-api = http://127.0.0.1:8081/kube/add
conf_dir = ./config
app_conf = ./config/app.ini
`)
		t.Assert(checkDataType(data), "ini")
	})

	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
# API Server
[server]
    address = ":8199"

# Jenkins
[jenkins]
    url          = "https://jenkins-swimlane.com"
    nodeJsStaticBuildCmdTpl = """
npm i --registry=https://registry.npm.taobao.org
wget http://consul.infra:8500/v1/kv/app_{{.SwimlaneName}}/{{.RepoName}}/.env.qa?raw=true -O ./env.qa
npm run build:qa
"""
`)
		t.Assert(checkDataType(data), "toml")
	})
}
