package conf

import (
	"gopkg.in/ini.v1"
	//"gopkg.in/macaron.v1"
	"os"
)

// Cfg
var Cfg *ini.File

// find configuration file
func init() {

	os.Setenv("oauth_key", "")
	os.Setenv("max_conn", "1")
	os.Setenv("idle_conn", "1")
	os.Setenv("cache_adapter", "memory")
	os.Setenv("cache_adapter_config", "")
	os.Setenv("http_port", "8181")
	os.Setenv("force_local_http_port", "false")

	//os.Setenv("mongo_uri", "mongodb://localhost:27017")
	//os.Setenv("mongo_uri", "mongodb://192.168.0.13:27017")
	os.Setenv("mongo_db", "curriculum")

	// var err error
	// Cfg, err = macaron.SetConfig("./conf/app.ini")
	// if err != nil {
	// 	panic(err)
	// }
}
