package conf

type app struct {
	Address string
	Mysql   string
}

// Config struct stores configuration data
type Config struct {
	App app
}
