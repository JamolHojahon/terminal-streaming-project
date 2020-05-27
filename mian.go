package main

import (
	"bufio"
	"log"
	"net"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// var (
	// 	configPath = flag.String("config", "./config.json", "path of the config file")
	// )
	// conf, err := config.FromFile(*configPath)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	listener, err := net.Listen("tcp", "+"+port)
	if err != nil {
		log.Fatalf("can't listen: %v", err)
	}
	log.Println("RUNNIG NOW", port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accepting connection failed:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader, writer := bufio.NewReader(conn), bufio.NewWriter(conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("can't read command %e", err)
		return
	}
	log.Println("CATCH:", str)
	_, err = writer.WriteString("THANOS GOING FOR U!")
	if err != nil {
		log.Println("Failed send response", err)
	}
}
