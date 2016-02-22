package conf

import (
	"io/ioutil"
)

var cached *Conf

// Get retrieves data from bancho.ini, and if already set, it gets data from cached.
func Get() (Conf, error) {
	if cached != nil {
		return *cached, nil
	}
	c, err := GetFromFile("bancho.ini")
	if err != nil {
		return c, err
	}
	cached = &c
	return c, nil
}

// GetFromFile retrieves the data from a file and uses it to make a new Conf.
func GetFromFile(confFileName string) (Conf, error) {
	data, err := ioutil.ReadFile(confFileName)
	if err != nil {
		return Conf{}, err
	}
	return GetConf(data)
}
