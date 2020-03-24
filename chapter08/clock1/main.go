package main

import(
	"io"
	"net"
	"time"
	"log"
)

func main(){

	listener,err := net.Listen("tcp","0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close() // As destructor, closes the connection
	for {
		_, err := io.WriteString(c , time.Now().Format("15:04:05.000\n"))
		if err != nil {
			return 
		}
		time.Sleep(1 * time.Second)
	}
}

//$ nc localhost 8000
//11:57:31.997
//11:57:33.000
//11:57:34.003
//11:57:35.004
