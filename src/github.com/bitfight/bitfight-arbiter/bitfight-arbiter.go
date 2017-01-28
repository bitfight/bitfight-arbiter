package main

// import global things
import "github.com/op/go-logging"

// import out things
import daemon "bitfight-arbiter-daemon"
//import engine "bitfight-arbiter-engine"

// do some global setup
var log = logging.MustGetLogger("main")

func main() {
	log.Info("Welcome to the bitfight arbiter server")

	daemon_done := make(chan bool)
	go daemon.Start(daemon_done)

	// wait for daemon done
	<-daemon_done
}
