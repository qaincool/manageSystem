package model

type Config struct {
	Port     string   `yaml:"port"`
	Database Database `yaml:"database"`
}

type Database struct {
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
