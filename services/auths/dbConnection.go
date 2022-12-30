package auths

import (
	"database/sql"

	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/go-sql-driver/mysql"
)

type DbConnectable interface {
	GetConnection(smng secretsmanageriface.SecretsManagerAPI) (*sql.DB, error)
}

type dbConnect struct {
	dbName     string
	endpoint   string
	secretName string
}

func NewDBConnector(dbName, endpoint, secretName string) DbConnectable {
	return &dbConnect{
		dbName:     dbName,
		endpoint:   endpoint,
		secretName: secretName,
	}
}

func (r *dbConnect) GetConnection(smng secretsmanageriface.SecretsManagerAPI) (*sql.DB, error) {
	config := mysql.Config{
		User:                 "userName",
		Passwd:               "pword",
		DBName:               r.dbName,
		Net:                  "tcp",
		Addr:                 "school-db-instance-1.cqoecukejjum.eu-west-2.rds.amazonaws.com:3306",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", config.FormatDSN()+`&time_zone="UTC"`)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
