package conf

// 配置
// type Listen struct {
// 	Port string `yaml:"port"`
// 	Host string `yaml:"host"`
// }

// type Log struct {
// 	File  string `yaml:"file"`
// 	Level uint32 `yaml:"level"` // 0-panic, 1-fatal, 2-error, 3-warn, 4-info, 5-debug, 6-trace
// }

type MysqlConfig struct {
	Name         string `yaml:"name"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxOpenConn  int    `yaml:"max_open_conn"`
	MaxIdleConn  int    `yaml:"max_idle_conn"`
	ConnLifetime int    `yaml:"conn_lifetime"`
	LogMode      bool   `yaml:"log_mode"`
	Charset      string `yaml:"charset"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Db       int    `yaml:"db"`
	PassWord string `yaml:"password"`
}

type NamiConfig struct {
	Domain           string `yaml:"domain"`
	WsDomain         string `yaml:"ws_domain"`
	VideoDomain      string `yaml:"video_domain"`
	OldAccountUser   string `yaml:"account_user"`
	OldAccountSecret string `yaml:"account_secret"`
	Localhost        string `yaml:"localhost"` // 仅仅api服务需要
}

//type NamiConfig struct {
//	Domain             string `yaml:"domain"`
//	WsDomain           string `yaml:"ws_domain"`
//	VideoDomain        string `yaml:"video_domain"`
//	AccountUser        string `yaml:"account_user"`
//	AccountSecret      string `yaml:"account_secret"`
//	VideoAccountUser   string `yaml:"video_account_user"`
//	VideoAccountSecret string `yaml:"video_account_secret"`
//	Localhost          string `yaml:"localhost"` // 仅仅api服务需要
//}

// type TencentCloudConfig struct {
// 	Cos       TencentCloudCosConfig `yaml:"cos"`
// 	Css       TencentCloudCssConfig `yaml:"css"`
// 	AppID     string                `yaml:"app_id"`
// 	SecretID  string                `yaml:"secret_id"`
// 	SecretKey string                `yaml:"secret_key"`
// }

// type TencentCloudCosConfig struct {
// 	BucketName string `yaml:"bucket_name"`
// 	Region     string `yaml:"region"`
// 	UrlFormat  string `yaml:"url_format"`
// }

// type TencentCloudCssConfig struct {
// 	PushUrl         string `yaml:"push_url"`
// 	PushSecret      string `yaml:"push_secret"`
// 	PullUrl         string `yaml:"pull_url"`
// 	PullSecret      string `yaml:"pull_secret"`
// 	AppName         string `yaml:"app_name"`
// 	StreamKeyPrefix string `yaml:"stream_key_prefix"`
// 	CallbackKey     string `yaml:"callback_key"`
// }

type AliyunCloudConfig struct {
	Live                   AliyunCloudLiveConfig `yaml:"live"`
	Oss                    AliyunCloudOssConfig  `yaml:"oss"`
	AccessKeyID            string                `yaml:"access_key_id"`
	AccessKeySecret        string                `yaml:"access_key_secret"`
	PicAddress             string                `yaml:"pic_address"`
	Snapshot               string                `yaml:"snapshot"`
	SenderId               string                `yaml:"sender_id"`
	ChannelName            string                `yaml:"channel_name"`
	EgyBallAccessKeyID     string                `yaml:"egyball_access_key_id"`
	EgyBallAccessKeySecret string                `yaml:"egyball_access_key_secret"`
}

type AliyunCloudLiveConfig struct {
	DomainName          string `yaml:"domain_name"`
	AppName             string `yaml:"app_name"`
	AnchorName          string `yaml:"anchor_name"`
	VodTranscodeGroupID string `yaml:"vod_transcode_group_id"`
	StorageLocation     string `yaml:"storage_location"`
	CycleDuration       int    `yaml:"cycle_duration"`
	PushUrl             string `yaml:"push_url"`
	PlaySecret          string `yaml:"play_secret"`
	PushSecret          string `yaml:"push_secret"`
	PullUrl             string `yaml:"pull_url"`
	StreamKeyPrefix     string `yaml:"stream_key_prefix"`
	CallbackKey         string `yaml:"callback_key"`
}

type AliyunCloudOssConfig struct {
	Endpoint   string `yaml:"endpoint"`
	BucketName string `yaml:"bucketName"`
}
type SmsConfig struct {
	Buka Buka `yaml:"buka"`
	CL   CL   `yaml:"cl"`
}
type CL struct {
	SmsAccount  string `yaml:"sms_account"`
	SmsPassword string `yaml:"sms_password"`

	WaAccount  string `yaml:"wa_account"`
	WaPassword string `yaml:"wa_password"`
}
type Buka struct {
	ApiKey string `yaml:"api_key"`
	ApiPwd string `yaml:"api_pwd"`
	AppId  string `yaml:"app_id"`
}
