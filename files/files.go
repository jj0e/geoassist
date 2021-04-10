package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func New() *Manager {
	userInfo, _ := user.Current()
	homeDir := fmt.Sprintf("%s/GeoAssist", userInfo.HomeDir)

	fileManager := &Manager{
		Directory:      homeDir,
		ConfigJSONPath: fmt.Sprintf("%s/config.json", homeDir),
	}
	validate(fileManager)
	return fileManager
}

func fileExists(filename string) bool {
	exists := true
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		exists = false
	}
	return exists
}

func validate(fileManager *Manager) {
	if !fileExists(fileManager.Directory) {
		dir := os.Mkdir(fileManager.Directory, 0777)
		if dir != nil {
			log.Fatal(dir)
		}
	}

	if !fileExists(fileManager.ConfigJSONPath) {
		file, _ := os.Create(fileManager.ConfigJSONPath)
		file.Close()
		config := Config{
			Email:    "",
			Password: "",
		}

		f, _ := json.MarshalIndent(config, "", " ")
		err := ioutil.WriteFile(fileManager.ConfigJSONPath, f, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (fm *Manager) LoadConfig() Config {
	var config Config
	file, _ := os.Open(fm.ConfigJSONPath)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &config)
	return config
}

func (fm *Manager) UpdateConfig(conf *Config) {
	f, _ := json.MarshalIndent(conf, "", " ")
	err := ioutil.WriteFile(fm.ConfigJSONPath, f, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
