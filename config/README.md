# microservice-config

Load config from a json file. Those configs are uses frequently to run a micro service.

# How to use it 

```go
package main

import (
    microconfig "github.com/alvarios/microconfig"
    "github.com/kataras/iris/v12"
)

func main(){
    config := microconfig.NewConfigFromEnv()
    app := iris.New()
    app.Handle("GET", "/health", func(ctx iris.Context){
       ctx.Json([]byte("I'm ok")) 
    })
    app.Listen(config.Server.Port)
}
```
NewConfigFromEnv  create a Configuration instance from a json file. The path to the json file is load from
env an variable CONFIG_PATH


# Description

```go 
/*Database struct build to load all the database config variables
type  struct {
	Dialect   string `default:"MangoDB"` // Dialect the name of the SGBD MangoDB by default
	Debug     bool   `default:"false"`   // Debug Mode can be debug or release
	Username  string `required:"true" default:"root"` // Username is the database owner id
	Password  string `required:"true"`	// Password is the database (identified by Dbname) password
	Host      string `required:"true"` // Host localhost or custom value
	DebugHost string `required:"true"` //DebugHost is similar to Host, it's an optionnal value use by the server when it's in DEBUG Mode
	Port      string `required:"true"` // Port is the db service connection port
	Dbname    string `required:"true"` // Dbname is the database name
	SSLMode   bool   `default:"true"`  // SSLMode is a boolean. It can be set to true if you want to activate the SSL encryption mode
}
*/
type Database struct {
	Dialect   string `default:"MangoDB"`
	Debug     bool   `default:"false"`
	Username  string `required:"true" default:"root"`
	Password  string `required:"true"`
	Host      string `required:"true"`
	DebugHost string `required:"true"`
	Port      string `required:"true"`
	Dbname    string `required:"true"`
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

```
