package config

//Option设计模式
type Options struct {
	FileName string //文件名
	FilePath     string //文件路径
	FileType     string //文件类型,json|xml|yaml
}

//选项
type Option func(*Options)

//设置配置的文件名
func ConfigFileNameOption(fileName string) Option {
	return func(options *Options) {
		options.FileName = fileName
	}
}

//只是配置文件的路径
func ConfigFilePathOption(filePath string) Option {
	return func(options *Options) {
		options.FilePath = filePath
	}
}

func ConfigFileType(fileType string) Option {
	return func(options *Options) {
		options.FileType = fileType
	}
}
