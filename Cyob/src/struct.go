package main

type Block struct {
	Index     int    // 블록체인에 기록된 위치
	Timestamp string // 데이터가 작성된 시간
	BPM       int
	Hash      string // 데이터 id
	PrevHash  string // 이전 데이터의 id
}

type Message struct {
	BPM int
}
