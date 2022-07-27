package config

var Cfg = New()

type Config interface {
	GetString(name string) string
}

func New() Config {
	config := &Instance{
		Package: NewViper(),
	}

	return config
}

func GetProductionStatus() bool {
	return Cfg.GetString("ENV") == "production"
}
