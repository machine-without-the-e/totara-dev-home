package config

import(
  "github.com/goccy/go-yaml"
  "os"
  "fmt"
)

type SiteConfig struct {
  ProjectRoot string `yaml:"projectroot"`
  ProjectName string `yaml:"projectname"`
}


func getSiteProjectDefaultConfig() (SiteConfig, error) {
  file, err := os.ReadFile("config/site/sites.yaml")
  
  if (err != nil) {
    return SiteConfig{}, err
  }

  config := SiteConfig{}

  err = yaml.Unmarshal(file, &config)

  if (err != nil) {
    return SiteConfig{}, err
  }

  return config, nil
}

func getSiteProjectConfig(site string) (SiteConfig, error) {
  file, err := os.ReadFile(fmt.Sprintf("config/site/sites/%s.yaml", site))
  
  if (err != nil) {
    return SiteConfig{}, err
  }

  config := SiteConfig{}

  err = yaml.Unmarshal(file, &config)

  if (err != nil) {
    return SiteConfig{}, err
  }

  return config, nil
}

func GetConfig() {

}
