package config

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	var (
		c = `
  mongo: 
    port: 27017
    host: localhost
    database: sakila
    username: root
    password: root
    options: x
  `
		conf Config
	)

	err := parseConfig(&conf, []byte(c))
	if err != nil {
		t.Errorf("could not parse yaml: %s", err.Error())
		t.Fail()
	}

	if conf == (Config{}) {
		t.Errorf("config is null")
		t.Fail()
	}

	if conf.Mongo.Port != 27017 {
		t.Errorf("got different port, expecting 27017 found: %d", conf.Mongo.Port)
	}

	if conf.Mongo.Password != "root" {
		t.Errorf("got different password, expecting 'root' found: %s", conf.Mongo.Password)
	}
}
