package config

type server struct {
	Host         string
	Port         int64
	AllowOrigins []string
	AllowMethods []string
}

type database struct {
	Host     string
	Port     int64
	Username string
	Password string
}

// Config representes de aplication configuration
type Config struct {
	Server   server
	Database database
}
