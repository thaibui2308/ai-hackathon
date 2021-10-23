package config

import (
	"strconv"

	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/thaibui2308/ai-hackathon/models"
)

func GetConfiguration(filename string) (config models.Config, err error) {
	Config := models.Config{}
	data, err := yaml.ReadFile(filename)
	if err != nil {
		return models.Config{}, err
	}
	Config.Username, _ = data.Get("Username")
	Config.Repository, _ = data.Get("Repository")
	ID, _ := data.GetInt("CommitID")
	Config.CommitID = strconv.Itoa(int(ID))
	return Config, nil
}
