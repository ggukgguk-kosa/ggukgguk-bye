package step1

import (
	"fmt"
	"time"
)

type Author struct {
	ID        string
	Name      string
	Nickname  string
	Email     string
	CreatedAt time.Time
}

type Document struct {
	Author    string
	Content   string
	CreatedAt time.Time
	Media     DocumentMedia
	IsOpen    bool
	ShareTo   string
	Replies   []DocumentReply
}

type DocumentReply struct {
	Content   string
	CreatedAt time.Time
	Author    string
}

type DocumentMedia struct {
	ID   string
	Type string
}

func printDocument(member Author, documents []Document) {
	fmt.Println(member.Nickname + "님의 조각들입니다")
	for _, d := range documents {
		fmt.Println(d)
	}
	fmt.Println("")
}
