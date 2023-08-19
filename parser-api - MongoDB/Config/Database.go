package Config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

// DBConfig represents db configuration
type DBConfig struct {
	ConnectionString string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		ConnectionString: "mongodb+srv://anjul2002:UqB36Vw08pxQjH9J@experimental.nlebqja.mongodb.net/",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return dbConfig.ConnectionString
}
