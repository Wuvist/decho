package conf

// App stores application configuration
type App struct {
	_       string `prefix:"app"`
	Address string
	Mysql   string
}
