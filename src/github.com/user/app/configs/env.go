package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var envs map[string]string

// LoadEnv function will be used to load environment variables
func LoadEnv() {
	jsonFile, error := os.Open("configs/env.json")

	defer jsonFile.Close()

	if error != nil {
		panic("Failed to open env json file: " + error.Error())
	}

	jsonByte, error := ioutil.ReadAll(jsonFile)

	if error != nil {
		panic("Something went wrong with env.json file : " + error.Error())
	}

	err := json.Unmarshal(jsonByte, &envs)
	if err != nil {
		panic("Something went wrong with env keys file : " + err.Error())
	}

	for envKey, envValue := range envs {
		os.Setenv(envKey, envValue)
	}

	log.Println("### ENVs setup done ####")
}
