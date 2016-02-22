// Package conf is an helper package for managing configuration files.
package conf

import (
	"gopkg.in/ini.v1"
)

// Conf is a struct containing the instance information on ripple.
type Conf struct {
	SQLInfo struct {
		User   string
		Pass   string
		Name   string
		Host   string
		DBType string
	}
	Ports struct {
		HTTP  string
		HTTPS string
	}
}

// GetConf creates a new Conf instance with the passed bytes
func GetConf(b []byte) (Conf, error) {
	c := Conf{}
	cf, err := ini.Load(b)
	if err != nil {
		return c, err
	}

	sqlinfo := cf.Section("sqlinfo")
	c.SQLInfo.User = sqlinfo.Key("user").Value()
	c.SQLInfo.Pass = sqlinfo.Key("password").Value()
	c.SQLInfo.Name = sqlinfo.Key("name").Value()
	c.SQLInfo.Host = sqlinfo.Key("host").Value()
	c.SQLInfo.DBType = sqlinfo.Key("type").Value()

	http := cf.Section("http")
	c.Ports.HTTP = http.Key("http_port").Value()
	c.Ports.HTTPS = http.Key("https_port").Value()

	return c, nil
}
