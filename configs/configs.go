package configs

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

const DefaultConfigPath = "./config.yaml"

var Conf = new(Config)

type Config struct {
	MySqlConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
}

type MySqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type KafkaConfig struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

func Init() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "config", DefaultConfigPath, "配置文件绝对路径或相对路径")
	flag.Parse()

	logrus.Printf("===> config path is: %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("config file has changed")
		if err := viper.Unmarshal(&Conf); err != nil {
			logrus.Errorf("failed at unmarshal config file after change, err: %v", err)
			panic(fmt.Sprintf("failed at init config: %v", err))
		}
		// 如果有环境变量就覆盖，适用于本地开发使用文件，实际运行使用环境变量的场景
		DBHost := os.Getenv("DB_HOST")
		if DBHost != "" {
			Conf.MySqlConfig.Host = DBHost
		}
		DBPassword := os.Getenv("DB_PASSWORD")
		if DBPassword != "" {
			Conf.MySqlConfig.Password = DBPassword
		}
		RedisHost := os.Getenv("REDIS_HOST")
		if RedisHost != "" {
			Conf.RedisConfig.Host = RedisHost
		}
		RedisPassword := os.Getenv("REDIS_PASSWORD")
		if RedisPassword != "" {
			Conf.RedisConfig.Password = RedisPassword
		}
		logrus.Infof("%+v", Conf)
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}
	// 解析到变量中
	if err := viper.Unmarshal(&Conf); err != nil {
		logrus.Errorf("failed at Unmarshal config file, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}

	// 如果有环境变量就覆盖，适用于本地开发使用文件，实际运行使用环境变量的场景
	DBHost := os.Getenv("DB_HOST")
	if DBHost != "" {
		Conf.MySqlConfig.Host = DBHost
	}
	DBPassword := os.Getenv("DB_PASSWORD")
	if DBPassword != "" {
		Conf.MySqlConfig.Password = DBPassword
	}
	RedisHost := os.Getenv("REDIS_HOST")
	if RedisHost != "" {
		Conf.RedisConfig.Host = RedisHost
	}
	RedisPassword := os.Getenv("REDIS_PASSWORD")
	if RedisPassword != "" {
		Conf.RedisConfig.Password = RedisPassword
	}
	logrus.Infof("%+v", Conf)
}
