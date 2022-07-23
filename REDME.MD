# Simple Configuration library

Allows the application to load its configuration from `.config` files or environment variables and automatic watches to changes

## Install

To install the package

```
$ go get github.com/najibulloShapoatov/config
```

Methods for Config struct

| Method                                      |                                                                                  Description                                                                                  |
|---------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| Load(loaders ...Loader) error               |                Load runs the given loaders in order to load and parse the configuration values. The first loader that returns an error stops the load process                 |
| Has(key string) bool                        |                                                                   Has returns true if the given key exists                                                                    |
| SetWatchInterval(duration time.Duration)    |                                   SetWatchInterval sets interval of to automatic load configurations every duration, default interval is 5s                                   |
| StopWatch()                                 |                                                                        StopWatch stopping all watches                                                                         |
| Unmarshal(destinationPtr interface{}) error |                                          Unmarshal decodes the configuration in a structure based on the `config` and `default` tags                                          |
| GetString(key string) (string, bool)        |                    GetString returns the value at the given key as a string and true if the key exists or empty string and false if the key doesn't exist                     |
| GetFloat(key string) (float64, bool)        |      GetFloat returns the value at the given key parsed as a float and true if the key exists or 0.0 and false if the key doesn't exist or failed to parse as a float64       |
| GetInt(key string) (int, bool)              |           GetInt returns the value at the given key parsed as a int and true if the key exists or 0 and false if the key doesn't exist or failed to parse as a int            |
| GetBool(key string) (bool, bool)            | GetDuration returns the value at the given key parsed as Duration and true if the key exists or Duration(0) and false if the key doesn't exist or failed to parse as Duration |
| GetKeys() (res []string)                    |                                                                        GetKeys return all loaded keys                                                                         |


You can add custom loader like `json` or `yaml`  or `xml` etc. you need implement `Loader` interface and add your loader in load function.
```go
// Loader is used to load and parse configuration values from various formats and location
type Loader interface {
	// Parse method is called
	Parse() (map[string]string, error)
	// IsWatchable method should return true if you need watch this loader
	IsWatchable() bool
}
```



## Usage example

```go
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

	err := godotenv.Load()
	if err != nil {
		return 
	}

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

	fmt.Println("isDebug: ", isDebug, " exists: ", exists)
}
```

## Configuration files

Configuration files are mostly `key=value` files but with few additions. For example, numbers are evaluated by the parser and booleans can be all truthy values besides true or false. Other files can be included using the `include` directive

```bash
# String values
key.name = "value" # this is inline comment
key.multiline = "multi \
line \
string"

# Number values
test.int.value = 5
test.float.value = 3.14
test.negative.value = -1.2
test.hex.number = 0x1234 # will parse to 4660
test.octal.number = 0o123 # will parse to 83
test.binary.number = 0b1010101 # will parse to 85
test.exponential.number = 1e3 # will parse to 1000
test.negative.exponential.number = 2e-2 # will parse to 0.02

# Boolean values
test.bool.value1 = yes       # or no
test.bool.value2 = on        # or off
test.bool.value3 = set       # or unset
test.bool.value4 = active    # or inactive
test.bool.value5 = enabled   # or disabled
test.bool.value6 = true      # or false
test.bool.value6 = 1         # or 0

# Duration
test.duration.value1 = "1h5m"
test.duration.value2 = "3s"

# Include other file
include "sub-config-file.conf"
```
