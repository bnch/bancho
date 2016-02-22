package conf

import (
	"io/ioutil"
)

// SampleConf is a sample configuration file.
const SampleConf = `
[sqlinfo]
user = root
password =
name = bancho
; leave blank for localhost:3306.
; If you want to set it up better, have a look at https://github.com/go-sql-driver/mysql#dsn-data-source-name
; host is the protocol(address) part.
host =
; At the moment only mysql is built in, so either use that or nothing.
type = mysql

[http]
http_port = :3000
; https will only work if cert.pem and key.pem are found in the current directory.
https_port = :10443
`

// WriteSampleConf writes SampleConf to bancho.ini
func WriteSampleConf() error {
	return ioutil.WriteFile("bancho.ini", []byte(SampleConf), 0644)
}
