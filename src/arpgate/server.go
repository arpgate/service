package arpgate

// export GOPATH=/Users/mat/arpgate/manage  // adjust for your username
// go get github.com/go-martini/martini
// arpgate/devoops needs to be checked out
// go get github.com/mattn/go-sqlite3

import (
	"fmt"
	"github.com/go-martini/martini"
)

func Start() {

	isInit := GetConf(DB_KEY_SYS_INITALIZED)
	fmt.Println("is initalized: " + isInit)

	if isInit == "false" {
		DbInit()
	}

	m := martini.Classic()

	//test := arpgate.GetMD5Hash("some")
	//fmt.Println(test)

	// static files  -temp to see the bootstrap template
	m.Use(martini.Static("../devoops"))

	// version
	m.Get("/v1", func() string {
		return "0.0.1"
	})

	go ListenMqtt()

	fmt.Println("Listening on http://localhost:3000")
	m.Run()
}
