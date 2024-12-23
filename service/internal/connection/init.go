package connection

import (
	"flag"
)

func init() {
	db := flag.Bool("db", false, "")

	flag.Parse()

	// connect
	// initJwt()
	loadYml()
	makeVariable()
	makeFolder()
	connectPostgresql(*db)
	connectRabbitmq()
	connectRedis()
	initSmptAuth()
	initLogger()

	print("connect ok")
}
