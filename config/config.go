package config

import (
	"os"
	"strings"
)

var configs = map[string]string{}

func Start() {
	convertConfigs()
}

func Get(key string) string {
	return configs[key]
}

func readLinesFromFile() []string {
	bytes, err := os.ReadFile(".env")
	if err != nil {
		panic(err)
	}

	text := string(bytes)
	linesArray := strings.Split(text, "\n")
	return linesArray
}

func convertConfigs() {
	lines := readLinesFromFile()

	for _, line := range lines {
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 || equalIndex == len(line)-1 || equalIndex == 0 {
			continue
		}

		key := strings.TrimSpace(line[:equalIndex])
		value := strings.TrimSpace(line[equalIndex+1:])
		configs[key] = value
	}
}
