package config

// fileLoader loads and processes .conf files
type fileLoader struct {
	watchForChanges bool
	filename        string
}

// NewFileLoader returns a .conf file parser
func NewFileLoader(filename string, watchForChanges bool) Loader {
	loader := &fileLoader{
		watchForChanges: watchForChanges,
		filename:        filename,
	}
	return loader
}

// Parse the file and returns a map of key=value settings
func (f *fileLoader) Parse() (map[string]string, error) {
	return f.parseFile(f.filename)
}

func (f *fileLoader) IsWatchable() bool {
	return f.watchForChanges
}

func (f *fileLoader) parseFile(filename string) (map[string]string, error) {
	strLoader := &stringLoader{filename: filename}
	return strLoader.parseData()
}
