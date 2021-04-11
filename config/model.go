package config

type Config struct {
	Debug                bool
	DatabaseURI          string
	DatabaseReadTimeout  uint
	DatabaseWriteTimeout uint
	DatabaseUserName     string
	DatabasePassword     string
}
