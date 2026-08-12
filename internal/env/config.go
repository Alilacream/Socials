package env

import (
	"log"
	"strconv"
)

// i'll call every variable within the vars and set them here
var err error

type EnvConf struct {
	Port         string
	DB_url       string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
	Secret       string
}

// ReturnAll rather than assingning every var and checking, let us just assign it and a func
// i just assigned the data structure here to lazy to do it in model
// NOTE: i should make a test func for reading env var
func ReturnAll() *EnvConf {
	var maxopen int
	var maxidle int
	maxopen, err = strconv.Atoi(GetVar("DB_MAX_OPEN_CONNS"))
	maxidle, err = strconv.Atoi(GetVar("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("Couldn't Load max open connection env var %s", err.Error())
		return nil
	}
	return &EnvConf{
		Port:         GetVar("PORT"),
		DB_url:       GetVar("DB_URL"),
		MaxOpenConns: maxopen,
		MaxIdleConns: maxidle,
		MaxIdleTime:  GetVar("DB_MAX_IDLE_TIME"),
		Secret:       GetVar("SECRET_KEY"),
	}
}
