package config

import (
	"os"
	"strings"
)

// envLoader loads and processes
type envLoader struct {
	watchForChanges bool
	prefix          string
}

// NewEnvLoader returns a env vars parser
func NewEnvLoader(watchForChanges bool, prefix string) Loader {
	loader := &envLoader{
		watchForChanges: watchForChanges,
		prefix:          prefix,
	}
	return loader
}

func (f *envLoader) IsWatchable() bool {
	return f.watchForChanges
}

func (f envLoader) Parse() (map[string]string, error) {
	// Collect the environment variable
	// If prefix is provided returning only variables with prefix
	var res = map[string]string{}
	if f.prefix == "" {
		for _, v := range os.Environ() {
			kv := strings.SplitN(v, "=", 2)
			res[kv[0]] = kv[1]
		}
	} else {
		for _, v := range os.Environ() {
			if strings.HasPrefix(strings.ToLower(v), strings.ToLower(f.prefix)) {
				kv := strings.SplitN(v, "=", 2)
				res[kv[0]] = kv[1]
			}
		}
	}

	return res, nil
}
