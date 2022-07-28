package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

type config struct {
	Verbose bool
	Logging struct {
		Level string
	}
	Name string
}

var C config

func main() {
	fmt.Println("** Viper-Cobra-Playground started **")

	// The viper.AddConfigPath() line sets the config file path. You can specify multiple paths
	// by adding multiple AddConfigPath methods, and Viper will read them in the order you provide.
	viper.AddConfigPath("./configs")

	viper.SetConfigFile("tada.env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println("tada.env - Port is " + strconv.Itoa(viper.GetInt("PORT")))

	// now its json files
	viper.SetConfigName("cfg1") // Register config file name (no extension)
	viper.SetConfigType("json") // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	port := viper.GetInt("prod.port")
	fmt.Println("cfg1.json - Prod-Port is " + strconv.Itoa(port))

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
	fmt.Println("cfg1.json - File is from " + C.Name)
	fmt.Println("cfg1.json - Logging.level is " + C.Logging.Level)
	fmt.Println("cfg1.json - Verbose-Mode is " + strconv.FormatBool(C.Verbose))

	// now its yaml files
	viper.SetConfigName("application-config") // Register config file name (no extension)
	viper.SetConfigType("yml")                // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	port = viper.GetInt("prod.port")
	fmt.Println("yml - Prod-Port is " + strconv.Itoa(port))

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
	fmt.Println("yml - File is from " + C.Name)
	fmt.Println("yml - Logging.level is " + C.Logging.Level)
	fmt.Println("yml - Verbose-Mode is " + strconv.FormatBool(C.Verbose))

	viper.Reset()
	// The viper.AddConfigPath() line sets the config file path. You can specify multiple paths
	// by adding multiple AddConfigPath methods, and Viper will read them in the order you provide.
	viper.AddConfigPath("./config1")
	viper.AddConfigPath("./configs")

	viper.SetConfigName("application-config") // Register config file name (no extension)
	viper.SetConfigType("yml")                // Look for specific type
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	port = viper.GetInt("prod.port")
	fmt.Println("yml - Prod-Port is " + strconv.Itoa(port))

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
	fmt.Println("yml - File is from " + C.Name)
	fmt.Println("yml - Logging.level is " + C.Logging.Level)
	fmt.Println("yml - Verbose-Mode is " + strconv.FormatBool(C.Verbose))
}
