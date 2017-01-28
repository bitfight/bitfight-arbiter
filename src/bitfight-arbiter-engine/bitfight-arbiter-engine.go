package bitfight_arbiter_engine

import "strconv"

import "github.com/op/go-logging"

var log = logging.MustGetLogger("engine")

func Lint(code string, connectionId int)(valid bool) {
	log.Info("Linting some code for connection: " + strconv.Itoa(connectionId) )
	// @TODO Mark to Complete since he likes stuff like this
	return true
}