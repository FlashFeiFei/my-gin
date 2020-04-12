package config

import (
	"errors"
	"github.com/spf13/viper"
)

//直接封装了viper，好像这个东西不是线程安全，所以set操作可能会有问题，所以不要set操作
//在我封装的gin框架中，Config只用来做读取操作
type Config struct {
	*Options
	*viper.Viper //匿名字段，采用类型名作为字段
}

func NewConfig(options ...Option) (*Config, error) {
	if len(options) != 3 {
		return nil, errors.New("文件名、路径、文件类型设置不能为空")
	}
	//设置选项
	Options := new(Options)
	for _, op := range options {
		op(Options)
	}
	//创建一个配置
	Viper := viper.New()
	Viper.SetConfigName(Options.FileName)
	Viper.AddConfigPath(Options.FilePath)
	Viper.SetConfigType(Options.FileType)
	err := Viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Options: Options,
		Viper:   Viper,
	}, nil
}
