package config

import (
	"os"

	"github.com/w1png/htmx/errors"
)

type Config struct {
	StorageType    string
	SqliteFileName string

	ServerPort string
}

var ConfigInstance Config

func InitConfig() error {
	var ok bool
	ConfigInstance = Config{SqliteFileName: "htmx.db"}

	ConfigInstance.StorageType, ok = os.LookupEnv("STORAGE_TYPE")
	if !ok {
		return errors.NewEnvironmentVariableNotFoundError("STORAGE_TYPE")
	}

	ConfigInstance.SqliteFileName, ok = os.LookupEnv("SQLITE_FILENAME")
	if !ok {
		return errors.NewEnvironmentVariableNotFoundError("SQLITE_FILENAME")
	}

	ConfigInstance.ServerPort, ok = os.LookupEnv("SERVER_PORT")
	if !ok {
		return errors.NewEnvironmentVariableNotFoundError("SERVER_PORT")
	}

	return nil
}
