package config

import (
	"log"
	"github.com/go-yaml/yaml"
)

type Config struct {
	URL      string     `yaml:"url"`       // 需要监控的网站URL
	Timeout  int64      `yaml:"timeout"`   // 连接URL超时的时间(s)
	CronTime string     `yaml:"cron_time"` // 循环运行的时间(分)
	Do       [][]string `yaml:"do"`        // 需要执行的操作
	Email    Email      `yaml:"email"`     // 邮件
}

type Email struct {
	Subject  string `yaml:"subject"`
	SMTP     string `yaml:"smtp"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	To       string `yaml:"to"`
}

func NewConfig() {
	config := Config{}
	// 解析配置文件
	if config, err := ioutil.ReadFile(./config.yaml); err != nil {
		log.Fatal("配置文件读取失败！", err)
	}
	return config
}
