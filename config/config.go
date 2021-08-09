package config

import (
	"fmt"

	_ "github.com/ghodss/yaml"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func init() {
	setDefaults()
	//PrintConfig()
}

type ConfigInfo struct {
	base string
}

func setDefaults() {
	//system
	viper.SetDefault("system.address", "0.0.0.0:18000")
	viper.SetDefault("system.passwd", "")
	viper.SetDefault("system.model", 0)

	//result
	viper.SetDefault("result.tofs", 1)
	viper.SetDefault("result.fshost", "10.130.29.88")
	viper.SetDefault("result.fsport", "8021")
	viper.SetDefault("result.fspasswd", "ClueCon")

	//log
	viper.SetDefault("log.path", "./log/")
	viper.SetDefault("log.maxsize", 500)
	viper.SetDefault("log.maxrotate", 100)
	viper.SetDefault("log.maxdays", 30)
	viper.SetDefault("log.compress", false)
	viper.SetDefault("log.level", 0)

	//bak file
	//workmodel=1
	viper.SetDefault("filebak.switch", 0)
	viper.SetDefault("filebak.bakRoot", "./bak")
	viper.SetDefault("filebak.bakExt", "")
	viper.SetDefault("filebak.rmMix", false)     //remove filename suffix .mix.*"
	viper.SetDefault("filebak.addWavHead", true) //"add HEAD for *.wav file."
	//viper.SetDefault(key string, value interface{})
	//viper.SetDefault("filebak.addWavHead", false) // "add HEAD for *.wav file."

	//tohttp
	viper.SetDefault("tohttp.switch", 1)
	viper.SetDefault("tohttp.address", "https://0.0.0.0:19000/v1/addaudio")
	viper.SetDefault("tohttp.srvinfo", "{\"scene_id\":\"100011xxxx\"}")
	viper.SetDefault("tohttp.contenttype", "multipart/form-data;boundary=------WebKitFormBoundarypQQgHEMvpxJSMTei")
	viper.SetDefault("tohttp.boundary", "----WebKitFormBoundarypQQgHEMvpxJSMTei")
	viper.SetDefault("tohttp.idleconn", 300)
	viper.SetDefault("tohttp.maxconn", 1000)
	viper.SetDefault("tohttp.sendmss", 8000)
	viper.SetDefault("tohttp.totaltimeout", 5)
	viper.SetDefault("tohttp.resptimeout", 2)
	viper.SetDefault("tohttp.enabletls", 1)
	viper.SetDefault("tohttp.clientcrt", "")
	viper.SetDefault("tohttp.clientkey", "")
	viper.SetDefault("tohttp.cacrt", "./crt/server.crt")

	//websocket
	viper.SetDefault("towebsocket.switch", 0)
	viper.SetDefault("towebsocket.address", "0.0.0.0:19010")
	viper.SetDefault("towebsocket.path", "ws")
	viper.SetDefault("towebsocket.sendresult", 1)
	viper.SetDefault("towebsocket.meta", 0)

}

func GetConfigString() string {
	c := viper.AllSettings()
	gConfig, _ := yaml.Marshal(c)
	return string(gConfig)
}

func PrintConfig() {
	fmt.Print(GetConfigString())
}

func DebugConfigPath() bool {
	return true
}

func GetDebugSwitch() bool {
	return false
}
