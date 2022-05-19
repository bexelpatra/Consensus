package main

import "sync"

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}

var Blockchain []Block
var tempBlocks []Block

var candidateBlocks = make(chan Block)

var announcements = make(chan string)

var mutex = &sync.Mutex{}

var validators = make(map[string]int)
