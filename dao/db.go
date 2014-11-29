package dao

import (
	"database/sql"
	"fmt"
	"log"

	"backend/models"
	"backend/utils"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

// var pqWrapper *PostGresWrapper = newPostGresDB()

// //PostGres Wrapper
// type PostGresWrapper struct {
// 	user     string
// 	password string
// 	dbName   string
// 	sslMode  string
// 	host     string
// }

// func newPostGresDB() *PostGresWrapper {
// 	wrapper := &PostGresWrapper{
// 		user:     utils.Conf.GetString("postgres.user"),
// 		password: utils.Conf.GetString("postgres.password"),
// 		dbName:   utils.Conf.GetString("postgres.database"),
// 		sslMode:  utils.Conf.GetString("postgres.sslmode"),
// 		host:     utils.Conf.GetString("postgres.host"),
// 	}
// 	return wrapper
// }

func getDbMap() (*gorp.DbMap, error) {
	postgresUser := utils.Conf.GetString("postgres.user")
	postgresPassword := utils.Conf.GetString("postgres.password")
	postgresDBName := utils.Conf.GetString("postgres.dbname")
	postgresSSLMode := utils.Conf.GetString("postgres.sslmode")

	postgresConnection := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		postgresUser,
		postgresPassword,
		postgresDBName,
		postgresSSLMode)

	db, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		return nil, err
	}

	// test the connection before using it
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// construct a gorp DbMap
	dbMap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
		// TypeConverter: models.TypeConverter{},
	}

	dbMap.AddTableWithName(models.Player{}, kPLAYER_TABLE).SetKeys(true, "Id")
	dbMap.AddTableWithName(models.Match{}, kMATCH_TABLE).SetKeys(true, "Id")
	dbMap.AddTableWithName(models.Turn{}, kTURN_TABLE).SetKeys(true, "Id")

	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalf("Failed to create tables with error: %v\n", err)
	}

	return dbMap, nil
}
