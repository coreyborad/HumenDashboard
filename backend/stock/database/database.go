package database

import (
	"fmt"
	"stock/config"

	"github.com/jinzhu/gorm"

	// mssql
	_ "github.com/jinzhu/gorm/dialects/mssql"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// Get Data Source Name
func getDSN(driver string) (dsn string, err error) {
	switch driver {
	case "mysql":
		// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
		dsn = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
			config.Database.Mysql.Username,
			config.Database.Mysql.Password,
			config.Database.Mysql.Host,
			config.Database.Mysql.Port,
			config.Database.Mysql.Dbname,
			config.Database.Mysql.Charset,
			"True",
			"Local",
		)
	case "pgsql":
		// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			config.Database.Pgsql.Host,
			config.Database.Pgsql.Port,
			config.Database.Pgsql.Username,
			config.Database.Pgsql.Dbname,
			config.Database.Pgsql.Password,
		)
	case "sqlite":
		// "/tmp/gorm.db"
		dsn = config.Database.Sqlite.Dbname
	case "sqlsrv":
		// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			config.Database.Sqlsrv.Host,
			config.Database.Sqlsrv.Port,
			config.Database.Sqlsrv.Username,
			config.Database.Sqlsrv.Dbname,
			config.Database.Sqlsrv.Password,
		)
	default:
		return "", fmt.Errorf("Unknown database type: %s", driver)
	}

	return
}

// ConnectDB Connect Database
func ConnectDB(drivers ...string) (connect *gorm.DB, err error) {
	var driver string

	if len(drivers) > 0 {
		driver = drivers[0]
	} else {
		driver = config.Database.Default
	}

	dsn, err := getDSN(driver)
	if err != nil {
		return
	}

	connect, err = gorm.Open(driver, dsn)

	return
}

// GetDB GetDB
func GetDB() *gorm.DB {
	return db
}

// Init Init
func Init() (err error) {
	if err := config.Check(); err != nil {
		return err
	}

	db, err = ConnectDB()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	db.LogMode(config.App.Debug)

	return nil
}
