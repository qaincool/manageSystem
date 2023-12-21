package model

// 配置文件
type Config struct {
	Port     string   `yaml:"port"`
	FilePath string   `yaml:"file_path"`
	Database Database `yaml:"database"`
}

type Database struct {
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
