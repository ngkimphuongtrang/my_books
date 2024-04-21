package db

const (
	mysqlOption                           = "charset=utf8&parseTime=True&loc=Local&multiStatements=True&maxAllowedPacket=0"
	defaultMySQLConnectionLifetimeSeconds = 300
)

type MySQLConfig struct {
	DSN                       string `json:"dsn" yaml:"dsn"`
	Server                    string `json:"server" yaml:"server"`
	Schema                    string `json:"schema" yaml:"schema"`
	User                      string `json:"user" yaml:"user"`
	Password                  string `json:"password" yaml:"password"`
	Option                    string `json:"option" yaml:"option"`
	ConnectionLifetimeSeconds int    `json:"connection_lifetime_seconds" yaml:"connection_lifetime_seconds"`
	MaxIdleConnections        int    `json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConnections        int    `json:"max_open_connections" yaml:"max_open_connections"`
}
