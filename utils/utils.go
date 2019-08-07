package utils

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnvVariables <struct>
// is used to descrive environment variables structure of a project
type EnvVariables struct {
	DebugMode bool

	MongoDBName string
	MongoDBURL  string
}

var variables *EnvVariables

// InitEnvVars <function>
// is used to initialize environment variables of a project
func InitEnvVars() {
	variables = new(EnvVariables)
	err := envconfig.Process("api", variables)
	if err != nil {
		log.Fatal(err)
	}
}

// GetEnvVars <function>
// is used to return current environment variables
func GetEnvVars() *EnvVariables {
	return variables
}
