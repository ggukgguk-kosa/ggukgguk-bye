package step1

import (
	"context"
	"database/sql"
	"log"

	"github.com/ggukgguk-kosa/ggukgguk-bye/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var dbPt *sql.DB

func DBInit(dbUri string) {
	log.Println("init db")

	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		log.Printf("DB 작업 중 오류 발생 - DBInit")
		DBClose()
		log.Fatal(err)
	}

	dbPt = db

	boil.SetDB(db)
}

func ReadAndWrite(members models.MemberSlice) {
	log.Println("operate start")
	for _, m := range members {
		log.Println("start - " + m.MemberEmail.String)
		documents := getDocumentsOf(*m)
		printDocument(Author{
			ID:        m.MemberID,
			Name:      m.MemberName,
			Nickname:  m.MemberNickname,
			Email:     m.MemberEmail.String,
			CreatedAt: m.MemberCreatedAt,
		}, documents)
	}
}

func getDocumentsOf(member models.Member) []Document {
	var documents []Document

	records := loadAllRecordsOf(member)
	for _, r := range records {
		replies := loadAllReplyOf(*r)
		var documentReplies []DocumentReply
		for _, repl := range replies {
			documentReplies = append(documentReplies, DocumentReply{
				Content:   repl.ReplyContent,
				CreatedAt: repl.ReplyDate,
				Author:    repl.MemberID,
			})
		}

		mediaFile := loadMediaFileOf(*r)
		if mediaFile.MediaFileID == "" || mediaFile.MediaFileBlocked.Bool {
			documents = append(documents, Document{
				Author:    r.MemberID,
				Content:   r.RecordComment.String,
				CreatedAt: r.RecordCreatedAt,
				IsOpen:    r.RecordIsOpen,
				ShareTo:   r.RecordShareTo.String,
				Replies:   documentReplies,
			})
		} else {
			documents = append(documents, Document{
				Author:    r.MemberID,
				Content:   r.RecordComment.String,
				CreatedAt: r.RecordCreatedAt,
				Media: DocumentMedia{
					ID:   mediaFile.MediaFileID,
					Type: mediaFile.MediaTypeID,
				},
				IsOpen:  r.RecordIsOpen,
				ShareTo: r.RecordShareTo.String,
				Replies: documentReplies,
			})
		}
	}

	return documents
}

func LoadAllMembers() models.MemberSlice {
	log.Println("load all members")
	ctx := context.Background()

	members, err := models.Members().AllG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.MemberSlice{}
		}

		log.Printf("DB 작업 중 오류 발생 - LoadAllMembers")
		DBClose()
		log.Fatal(err.Error())
		return models.MemberSlice{}
	}

	return members
}

func DBClose() {
	log.Println("close db")
	dbPt.Close()
}

func loadAllRecordsOf(member models.Member) models.RecordSlice {
	ctx := context.Background()

	records, err := member.Records().AllG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.RecordSlice{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllRecordsOf")
		DBClose()
		log.Fatal(err.Error())
		return models.RecordSlice{}
	}

	return records
}

func loadAllReplyOf(record models.Record) models.ReplySlice {
	ctx := context.Background()

	reply, err := record.Replies().AllG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.ReplySlice{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllReplyOf")
		DBClose()
		log.Fatal(err.Error())
		return models.ReplySlice{}
	}

	return reply
}

func loadMediaFileOf(record models.Record) models.MediaFile {
	ctx := context.Background()

	mediaFile, err := record.MediaFile().OneG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.MediaFile{
				MediaFileID: "",
			}
		}

		log.Printf("DB 작업 중 오류 발생 - loadMediaFileOf")
		DBClose()
		log.Fatal(err.Error())
	}

	return *mediaFile
}
