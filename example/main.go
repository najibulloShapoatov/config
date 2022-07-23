package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/najibulloShapoatov/config"
	"time"
)

type MyConf struct {
	Debug         bool          `config:"APPDEBUG" default:"true"`
	Name          string        `config:"key.name"`
	Name1         string        `config:"key.name1" default:"name_f"`
	Multiline     string        `config:"key.multiline"`
	Int           int           `config:"test.int.value"`
	Float         float64       `config:"test.float.value"`
	NegativeFloat float64       `config:"test.negative.value"`
	Hex           int           `config:"test.hex.number"`
	Octal         int           `config:"test.octal.number"`
	Binary        int           `config:"test.binary.number"`
	Exp           float64       `config:"test.exponential.number"`
	NegExp        float64       `config:"test.negative.exponential.number"`
	B1            bool          `config:"test.bool.value1"`
	B2            bool          `config:"test.bool.value2"`
	B3            bool          `config:"test.bool.value3"`
	B4            bool          `config:"test.bool.value4"`
	B5            bool          `config:"test.bool.value5"`
	B6            bool          `config:"test.bool.value6"`
	B7            bool          `config:"test.bool.value7"`
	Duration1     time.Duration `config:"test.duration.value1" default:"1h"`
	Duration2     time.Duration `config:"test.duration.value2"`
}

func main() {

	godotenv.Load()

	//conf := config.Get()

	conf, err := config.Load(
		config.NewFileLoader("file.conf", true), // load config from this file
		config.NewEnvLoader(false, "app"),       // and also from environment variables
		config.NewStringLoader(`
test.int.value = 7
test.float.value = 3.17
test.negative.value = -1.7
`),
	)
	if err != nil {
		fmt.Println("failed to parse", err)
	}

	// unmarshall directly into a struct
	var cfg MyConf
	if err := conf.Unmarshal(&cfg); err != nil {
		fmt.Println("some error", err)
	}

	fmt.Printf("+%v", cfg)
	fmt.Println()

	// retrieve a value by name
	isDebug, exists := conf.GetBool("APPDEBUG")

	fmt.Println("isDebug: ", isDebug, " exsists: ", exists)

}
