package config

import (
	"errors"
	"fmt"
)

//单例模式
var configManage *ConfigManage

func init() {
	if configManage == nil {
		configManage = &ConfigManage{
			configMap: make(map[string]*Config),
		}
	}
}

//configManage管理
type ConfigManage struct {
	configMap map[string]*Config
}

//提过一个对外的函数加载配置文件进入
func LoadConfig(fileName, filePath, fileType string) error {
	config, err := NewConfig(
		ConfigFileNameOption(fileName),
		ConfigFilePathOption(filePath),
		ConfigFileType(fileType),
	)
	if err != nil {
		return err
	}
	configManage.configMap[fileName] = config
	return nil
}

//获取配置文件
func GetConfig(fileName string) (*Config, error) {
	if config, ok := configManage.configMap[fileName]; ok {
		return config, nil
	}
	return nil, errors.New(fmt.Sprintf("没有找到配置文件: %s", fileName))
}
