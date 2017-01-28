package bitfight_arbiter_daemon

import (
	"os"
	"net"
	"strconv"
)

import (
	"github.com/op/go-logging"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "31313"
	CONN_TYPE = "tcp"
)

var log = logging.MustGetLogger("daemon")

func Start(done chan bool) {
	log.Info("Daemon Starting")

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Errorf( "Error listening:", err.Error() )
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Infof( "Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Errorf("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new go routine.
		go handleRequest(conn)
	}

	done <- true
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 5000)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Errorf("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received of length " + strconv.Itoa(reqLen) + "\n"))

	// Close the connection when you're done with it.
	conn.Close()
}