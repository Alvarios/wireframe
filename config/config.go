package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

const configPathKey = "CONFIG_PATH"

//ServerModeRelease - value for release mode
const ServerModeRelease = "release"

//ServerModeDebug - value for debug mode
const ServerModeDebug = "debug"

/*Database struct build to load all the database config variables
type  struct {
	Dialect   string `default:"MangoDB"` // Dialect the name of the SGBD MangoDB by default
	Debug     bool   `default:"false"`   // Debug Mode can be debug or release
	Username  string `required:"true" default:"root"` // Username is the database owner id
	Password  string `required:"true"`	// Password is the database (identified by Dbname) password
	Host      string `required:"true"` // Host localhost or custom value
	DebugHost string `required:"true"` //DebugHost is similar to Host, it's an optional value use by the server when it's in DEBUG Mode
	Port      string `required:"true"` // Port is the db service connection port
	Dbname    string `required:"true"` // Dbname is the database name
	SSLMode   bool   `default:"true"`  // SSLMode is a boolean. It can be set to true if you want to activate the SSL encryption mode
}
*/
type Database struct {
	Dialect   string `default:"MangoDB"`
	Debug     bool   `default:"false"`
	Username  string `required:"false" default:"root"`
	Password  string `required:"false"`
	Host      string `required:"required"`
	DebugHost string `required:"false"`
	Port      string `required:"false"`
	Dbname    string `required:"false"`
	SSLMode   bool   `default:"true"`
}

// JWT struct contains the JWT token secret
type JWT struct {
	Secret string `required:"true"`
}

/*Server - struct use to load server config
type  struct {
	Port    string `required:"true" default:":8080"` // Port server port
	DebugPort string `json:"debug_port"`
	Domain  string `required:"false"`	// Domain
	Host    string `required:"false"`   // Host can be localhost or custom value
	LogFile string `required:"false"`	// LogFile file where the log will be stored
	Mode string "required:true" 		// server mode can be release or dev
	Name    string `json:"name"` 		// server name
}
*/
type Server struct {
	Port      string `required:"true" default:":8080"`
	DebugPort string `json:"debug_port"`
	Domain    string `required:"false"`
	Host      string `required:"false"`
	LogFile   string `required:"false"`
	Mode      string `json:"mode"`
	Name      string `json:"name"`
}

/*Configuration is used to read all the config from the config.Env.json file. This config variable are used to
set up the server.
type  struct {
	Database Database `required:"true"` // Database
	JWT      JWT      `required:"true"` // JWT
	Server   Server   `required:"true"`	// Server
}
*/
type Configuration struct {
	Database Database `required:"false"`
	JWT      JWT      `required:"false"`
	Server   Server   `required:"true"`
}

/*NewConfigFromEnv  create a Configuration instance from a json file. The path to the json file is load from
env an variable CONFIG_PATH
*/
func NewConfigFromEnv() *Configuration {
	envPathVar := os.Getenv(configPathKey)
	if envPathVar == "" {
		panic(configPathKey + " is empty. this env variable must be init with the path to the json config file")
	}
	config := &Configuration{}
	if err := configor.Load(config, envPathVar); err != nil {
		fmt.Printf("Config error : %v || path find in env was : %s \n", err.Error(), envPathVar)
		panic(err.Error())
	}
	return config
}

/*NewConfig  create a Configuration instance from a json file. configPath is the path to the json config file
 */
func NewConfig(configPath string) *Configuration {
	config := &Configuration{}
	if err := configor.Load(config, configPath); err != nil {
		panic(err.Error())
	}
	return config
}
