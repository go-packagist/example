package config

type App struct {
	Name    string
	Key     string
	EnvPath string
}

var app = &App{
	Name:    "example",
	Key:     "EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM",
	EnvPath: "./.env",
}
