package db

import (
	"fmt"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type DatabaseConfigurator interface {
// 	Host() string
// 	Port() int
// 	Username() string
// 	Password() string
// 	DatabaseName() string
// 	SchemaName() string
// 	SSLEnabled() bool
// }

type DbConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int
	Host     string
	DbName   string
	//SchemaName string
	SSLEnabled bool
}

type PostgresDB interface {
	DB() *gorm.DB
}

type postgresDB struct {
	db *gorm.DB
}

func (p postgresDB) DB() *gorm.DB {
	return p.db
}

func NewDBConnection(dbCfg DbConfig) PostgresDB {
	dsn := prepareDsn(dbCfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Error While Connecting to database")
	} else {
		fmt.Println("Database Connected")
	}

	return &postgresDB{db: db}
}

func prepareDsn(dbCfg DbConfig) string {
	mode := ""
	if !dbCfg.SSLEnabled {
		mode = "sslmode=disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s TimeZone=Asia/Shanghai",
		dbCfg.Host, dbCfg.Username, dbCfg.Password, dbCfg.DbName, dbCfg.Port, mode)

	return dsn
}
