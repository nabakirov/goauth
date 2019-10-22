package auth

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTokens(tokenFilePath string, tokens *map[string]bool) {
	data, err := ioutil.ReadFile(tokenFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &tokens)
}
