package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"restApp/config"
)

/*
	Oracle DB bağlantı parametrelerini db.json dosyasından çeker ve
	geri döndürür

*/

func GetOracleConfig() config.OracleConfig {
	jsonFile, _ := os.Open("./.config/db.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	var oracle config.OracleConfig

	jsonString, _ := json.Marshal(result["oracle"])
	json.Unmarshal(jsonString, &oracle)

	return oracle
}
