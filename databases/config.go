// This file stores all configurations of our app `myapp`
//  usees Viper to search for and read a configuration file ( in our case, app.env).

package databases

import (
	"log"

	"github.com/spf13/viper"
)

// The Config struct will hold config variables of our app that can be read from file or environment variables.
// We have two environment variables for now: DB_USERNAME and DB_POSTGRES_URL
type ConfigStruct struct {
	DB_USERNAME     string `mapstructure:"DB_DRIVER"`
	DB_POSTGRES_URL string `mapstructure:"DB_SOURCE"`
	// Viper uses the mapstructure package to unmarshall values, and get those values of the variables
}

//The LoadConfig function will read configurations from a config file `app.env` or environment variables inside the path,
// and then LoadConfig loads the config into the `ConfigStruct`` Struct defined earlier
func LoadConfig() (config ConfigStruct, err error) {

	viper.SetConfigName("app") // name of config file (without extension).The config file is `app.env`, so `app`.
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name

	/* Add config path - there are a few methods:
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	*/
	viper.AddConfigPath(".") // The dot in "." means the current working directory

	// AutomaticEnv method aids working with ENV. Viper treats ENV variables as case sensitive.
	// Read values from environment variables.
	viper.AutomaticEnv()

	//  Find and read the config file
	err = viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatal("Fatal error config file:", err)
	}

	// If error != nil, then return else unmarshal the values into the target `config` object
	// and return the config object and any error if it occurs.
	err = viper.Unmarshal(&config)

	return
}
