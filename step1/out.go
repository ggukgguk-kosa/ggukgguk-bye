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

type MetaOfDocuments struct {
	Year    string
	Month   string
	Color   string
	Keyword string
}

type Document struct {
	Author         string
	AuthorNickname string
	Content        string
	CreatedAt      time.Time
	Media          DocumentMedia
	IsOpen         bool
	ShareTo        string
	Replies        []DocumentReply
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

func printDocument(member Author, diaries []MetaOfDocuments, documents []Document) {
	fmt.Println(member.Nickname + "님의 다이어리입니다")
	for _, m := range diaries {
		fmt.Println(m)
	}
	fmt.Println("")

	fmt.Println(member.Nickname + "님의 조각들입니다")
	for _, d := range documents {
		if d.ShareTo == member.ID {
			fmt.Println(d.AuthorNickname + "님에게 공유받은 조각")
		}
		fmt.Println(d)
	}
	fmt.Println("")
}
