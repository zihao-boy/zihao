package config

import (
	"fmt"
	"sync"

	io "io/ioutil"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
)

var mutex sync.Mutex

// global var
var (
	G_AppConfig AppConfig
	G_DBConfig  DBConfig
)

const (
	//WorkSpace string = "/zihao/master/"
	WorkSpace string = "/Users/wuxuewen"
	Slave  int = 7001
	Remote_Images_Url string = "http://bbs.homecommunity.cn/app/zihaoApp.listZihaoApp"
	Remote_My_Images_Url string = "http://bbs.homecommunity.cn/app/zihaoApp.myListZihaoAppCmd"
	Remote_Apply_Publish_App_Url string = "http://bbs.homecommunity.cn/app/zihaoApp.applyPublishZihaoAppCmd"
	Remote_Apply_Publish_App_Version_Url string = "http://bbs.homecommunity.cn/app/zihaoApp.applyPublishZihaoAppVersion"

	Remote_get_images_version_Url string = "http://bbs.homecommunity.cn/app/zihaoAppImagesVersion.listZihaoAppImagesVersion"


	Remote_Offline_App_Url string = "http://bbs.homecommunity.cn/app/zihaoApp.offlineZihaoApp"

	Remote_Save_Publisher string = "http://bbs.homecommunity.cn/app/zihaoAppPublisher.saveZihaoAppPublisher"
	Remote_Update_Publisher string = "http://bbs.homecommunity.cn/app/zihaoAppPublisher.updateZihaoAppPublisher"
	Hc_cloud_app_id string = "102021120963240004"
	Kafka_switch_on string ="ON"
)

// 全局配置文件对应的结构体
type (
	// app
	AppConfig struct {
		iris.Configuration `yaml:"Configuration"`
		Own                `yaml: "own"`
	}
	Own struct {
		Separate      bool     `yaml:"separate"` // 是否前后端分离
		Port          int      `yaml:"port"`
		IgnoreURLs    []string `yaml:"ignore_urls,flow"`
		InterceptURLs []string `yaml:"intercept_urls,flow"`
		JWTTimeout    int      `yaml:"jwt_timeout"`
		LogLevel      string   `yaml:"log_level"`
		Secret        string   `yaml:"secret"`
		WebsocketPool int      `yaml:"websocket_pool"`
		Domain        string   `yaml:"domain"`
		Db            string   `yaml:"db"`
		Cache         string   `yaml:"cache"`
		DataPath      string   `yaml:"data_path"`
		SqlitePath      string   `yaml:"sqlite_path"`
		ContainerScheduling      string   `yaml:"container_scheduling"`
		KafkaIp      string   `yaml:"kafka_ip"`
		KafkaPort      string   `yaml:"kafka_port"`
		KafkaGroup      string   `yaml:"kafka_group"`
		KafkaTopic      string   `yaml:"kafka_topic"`
		KafkaSwitch      string   `yaml:"kafka_switch"`
		IpData      string   `yaml:"ip_data"`
		QqwryUrl      string   `yaml:"qqwry_url"`
		ServerIpUrl   string `yaml:"server_ip_url"`
	}

	// db
	DBConfig struct {
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
			PoolSize int    `yaml:"poolSize"`
		}
		Mysql struct {
			Dialect      string `yaml:"dialect"`
			User         string `yaml:"user"`
			Password     string `yaml:"password"`
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			Database     string `yaml:"database"`
			Charset      string `yaml:"charset"`
			ShowSql      bool   `yaml:"showSql"`
			LogLevel     string `yaml:"logLevel"`
			MaxOpenConns int    `yaml:"maxOpenConns"`
			MaxIdleConns int    `yaml:"maxIdleConns"`
			//ParseTime       bool   `yaml:"parseTime"`
			//MaxIdleConns    int    `yaml:"maxIdleConns"`
			//MaxOpenConns    int    `yaml:"maxOpenConns"`
			//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
			//Sslmode         string `yaml:"sslmode"`
		}
	}
)

func (conf DBConfig) DBConnUrl() string {
	var info = conf.Mysql
	//"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
	//return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
}

func loadConfig(filename string) ([]byte, error) {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
	}()
	data, err := io.ReadFile(filename)
	return data, err
}

func InitConfig() {
	var (
		app  = AppConfig{}
		db   = DBConfig{}
		data []byte
		err  error
	)
	// app
	data, err = loadConfig("conf/app.yaml")
	if err != nil {
		goto ERR
	}
	if err = yaml.Unmarshal(data, &app); err != nil {
		goto ERR
	}
	G_AppConfig = app
	golog.Infof("[app config]=> %v", G_AppConfig)

	// db
	data, err = loadConfig("conf/db.yaml")
	if err != nil {
		goto ERR
	}
	if err = yaml.Unmarshal(data, &db); err != nil {
		goto ERR
	}
	G_DBConfig = db
	golog.Infof("[db  config]=> %v", G_DBConfig)
	return
ERR:
	golog.Fatalf("~~> 解析配置文件错误,原因:%s", err.Error())
}
