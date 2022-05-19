package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

// "bufio"
// "crypto/sha256"
// "encoding/hex"
// "encoding/json"
// "fmt"
// "io"
// "log"
// "math/rand"
// "net"
// "os"
// "strconv"
// "sync"
// "time"

// "github.com/davecgh/go-spew/spew"
// "github.com/joho/godotenv"

//https://mycoralhealth.medium.com/code-your-own-proof-of-stake-blockchain-in-go-610cd99aa658
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
