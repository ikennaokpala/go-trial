package dao

import (
	"io/ioutil"

	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"
)

// Classic sets up basic data access default from
// configuration YAML and returns DB object
func Classic() gorm.DB {

	return New()
}

//New same as Classic
func New() gorm.DB {
	var err error

	db := DB{}

	path := db.ConfigPath()

	data, err := ioutil.ReadFile(path)
	err = yaml.Unmarshal([]byte(data), &db)

	if err != nil {
		panic(err.Error())
	}

	return db.Connect()
}
