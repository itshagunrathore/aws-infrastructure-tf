package models

type Configurations struct {
	DbDriver string   `mapstructure:"dbDriver"`
	DbConfig DbConfig `mapstructure:"dbConfig"`
}
type DbConfig struct {
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Port       int
	Host       string
	DbName     string
	SchemaName string
	SSLEnabled bool
}

// func (d DbConfig) Host() string {
// 	return d.host
// }

// func (d DbConfig) Port() int {
// 	return d.port
// }

// func (d DbConfig) Username() string {
// 	return d.username
// }
// func (d DbConfig) Password() string {
// 	return d.password
// }
// func (d DbConfig) SchemaName() string {
// 	return d.schemaName
// }
// func (d DbConfig) DatabaseName() string {
// 	return d.dbName
// }

// func (d DbConfig) SSLEnabled() bool {
// 	return d.sslEnalbed
// }
