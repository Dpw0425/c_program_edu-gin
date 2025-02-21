package config

import (
	"c_program_edu-gin/utils/encrypt"
	"c_program_edu-gin/utils/generator"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	sid        string      // 服务运行 ID
	App        *App        `json_utils:"app" yaml:"app" mapstructure:"app"`
	Server     *Server     `json_utils:"server" yaml:"server" mapstructure:"server"`
	Log        *Log        `json_utils:"log" yaml:"log" mapstructure:"log"`
	Mysql      *Mysql      `json_utils:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Jwt        *Jwt        `json_utils:"jwt" yaml:"jwt" mapstructure:"jwt"`
	Redis      *Redis      `json_utils:"redis" yaml:"redis" mapstructure:"redis"`
	Email      *Email      `json_utils:"email" yaml:"email" mapstructure:"email"`
	FileSystem *FileSystem `json:"file_system" yaml:"file_system" mapstructure:"file_system"`
}

func Load(filename string) *Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		// logger.Panicf("Read Config Error: %v!", err)
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		// logger.Panicf("Parsing Config Failed: %v!", err)
		panic(fmt.Sprintf("解析 config.yaml 失败: %v!", err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), generator.Random(6)))

	return &conf
}
