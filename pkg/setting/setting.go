package setting

//读取conf 配置文件 ini文件
import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	JwtSecret    string
)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini") //加载ini 配置文件
	if err != nil {
		//panic(err)
		fmt.Printf("读取配置文件错误 %v\n", err.Error())

		return
	}
	LoadBase()
	LoadServer()
}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	var err error
	sec, err := Cfg.GetSection("server")
	if err != nil {
		//panic(err)
		fmt.Printf("%v\n", err.Error())
		return
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
