package conf

import (
	"common/xstr"
	"fmt"

	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	. "common"
)

var (
	//  config
	Conf = &BaseConfig{}
)

// Config
type BaseConfig struct {
	//Listen       Listen
	//Log          Log

	Mysql MysqlConfig
	Redis RedisConfig
	Nami  NamiConfig
	//TencentCloud TencentCloudConfig `yaml:"tencent_cloud"`
	AliyunCloud AliyunCloudConfig `yaml:"aliyun_cloud"`
	SmsConfig   SmsConfig         `yaml:"sms"`
	SetupMode   string            //运行模式 prod | dev | local
	RootPath    string            //根目录
	EtcPath     string            //配置文件路径
	//LogPath      string
}

func LoadYamlCfg(cfgFilePath string, cfg interface{}) {
	var data []byte
	data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		panic(fmt.Sprint("Fail to read file: %s", err.Error()))
	}
	//读取配置
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(fmt.Sprint("Fail to read yaml: %s", err.Error()))
	}
}

func LoadBaseConfig() *BaseConfig {
	// 读取程序目录
	dir, _ := os.Executable()
	binPath := filepath.Dir(dir)
	dirSep := string(os.PathSeparator)
	Conf.RootPath = xstr.Substr(binPath, 0, strings.LastIndex(binPath, dirSep))

	// 初始化log
	//_, _ = time.LoadLocation("Asia/Shanghai")
	_, _ = time.LoadLocation("Europe/Istanbul")
	//nowStr := time.Now().Format("20060102")0
	LogDir := Conf.RootPath + "/logs/"
	Logger.Start(LogDir, CurrentServerName)
	//Conf.LogPath = LogDir + CurrentServerName + "_" + nowStr + ".log"
	//os.MkdirAll(LogDir, 0666)
	//f, _ := os.OpenFile(Conf.LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//log.SetOutput(io.MultiWriter(os.Stdout, f))

	// 加载配置文件
	Conf.SetupMode = os.Getenv("liveRunMode")
	if Conf.SetupMode == "" {
		//Conf.SetupMode = "local" // 测试环境
		Conf.SetupMode = "dev" // 测试环境
	}
	Conf.EtcPath = Conf.RootPath + "/etc/"
	//Conf.EtcPath = "D:\\trball\\etc\\"
	cfgFilePath := fmt.Sprintf("%scommon_%s.yaml", Conf.EtcPath, Conf.SetupMode)

	Info("now=%v", time.Now())
	Info("Conf filename=%s", cfgFilePath)
	Info("Conf.SetupMode=%s", Conf.SetupMode)
	Info("Conf.RootPath=%s", Conf.RootPath)
	Info("Conf.EtcPath=%s", Conf.EtcPath)
	LoadYamlCfg(cfgFilePath, &Conf)
	return Conf
}
