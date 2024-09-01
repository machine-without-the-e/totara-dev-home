package config

import (
  "testing"
)

func Test_getSiteProjectDefaultConfig(t *testing.T) {
  _, err := getSiteProjectDefaultConfig()
  
  if (err == nil) {
    return;
  }

  t.Error("Error retrieving site default configuration")
}
