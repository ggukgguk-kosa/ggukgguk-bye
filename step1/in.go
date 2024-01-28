package step1

import (
	"context"
	"database/sql"
	"log"

	"github.com/ggukgguk-kosa/ggukgguk-bye/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
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
		log.Println("step1 - in - " + m.MemberEmail.String)

		meta := getMetaOf(*m)

		uploadedDocuments := getDocumentsOf(*m)
		sharedDocuments := getSharedDocumentsOf(*m)
		documents := append(uploadedDocuments, sharedDocuments...)

		MakePdf(Author{
			ID:        m.MemberID,
			Name:      m.MemberName,
			Nickname:  m.MemberNickname,
			Email:     m.MemberEmail.String,
			CreatedAt: m.MemberCreatedAt,
		}, meta, documents)
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
				Author:    loadReplyAuthorOf(*repl).MemberNickname,
			})
		}

		mediaFile := loadMediaFileOf(*r)
		if mediaFile.MediaFileID == "" || mediaFile.MediaFileBlocked.Bool {
			documents = append(documents, Document{
				Author:         r.MemberID,
				AuthorNickname: member.MemberNickname,
				Content:        r.RecordComment.String,
				CreatedAt:      r.RecordCreatedAt,
				IsOpen:         r.RecordIsOpen,
				ShareTo:        r.RecordShareTo.String,
				Replies:        documentReplies,
			})
		} else {
			documents = append(documents, Document{
				Author:         r.MemberID,
				AuthorNickname: member.MemberNickname,
				Content:        r.RecordComment.String,
				CreatedAt:      r.RecordCreatedAt,
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

func getSharedDocumentsOf(member models.Member) []Document {
	var documents []Document

	records := loadAllSharedRecordsOf(member)
	for _, r := range records {
		author := loadAuthorOf(*r)

		replies := loadAllReplyOf(*r)
		var documentReplies []DocumentReply
		for _, repl := range replies {
			documentReplies = append(documentReplies, DocumentReply{
				Content:   repl.ReplyContent,
				CreatedAt: repl.ReplyDate,
				Author:    loadReplyAuthorOf(*repl).MemberNickname,
			})
		}

		mediaFile := loadMediaFileOf(*r)
		if mediaFile.MediaFileID == "" || mediaFile.MediaFileBlocked.Bool {
			documents = append(documents, Document{
				Author:         r.MemberID,
				AuthorNickname: author.MemberNickname,
				Content:        r.RecordComment.String,
				CreatedAt:      r.RecordCreatedAt,
				IsOpen:         r.RecordIsOpen,
				ShareTo:        r.RecordShareTo.String,
				Replies:        documentReplies,
			})
		} else {
			documents = append(documents, Document{
				Author:         r.MemberID,
				AuthorNickname: author.MemberNickname,
				Content:        r.RecordComment.String,
				CreatedAt:      r.RecordCreatedAt,
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

func getMetaOf(member models.Member) []MetaOfDocuments {
	var meta []MetaOfDocuments

	diaries := loadAllDiaries(member)
	for _, d := range diaries {
		meta = append(meta, MetaOfDocuments{
			Year:    d.DiaryYear,
			Month:   d.DiaryMonth,
			Color:   d.MainColor,
			Keyword: d.MainKeyword,
		})
	}

	return meta
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
	}

	return records
}

func loadAllSharedRecordsOf(member models.Member) models.RecordSlice {
	ctx := context.Background()

	records, err := models.Records(
		Where("record_share_to = ? and record_share_accepted = True", member.MemberID),
	).AllG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.RecordSlice{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllRecordsOf")
		DBClose()
		log.Fatal(err.Error())
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

func loadAllDiaries(member models.Member) models.DiarySlice {
	ctx := context.Background()

	diaries, err := member.Diaries().AllG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.DiarySlice{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllReplyOf")
		DBClose()
		log.Fatal(err.Error())
	}

	return diaries
}

func loadAuthorOf(record models.Record) models.Member {
	ctx := context.Background()

	member, err := record.Member().OneG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.Member{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllReplyOf")
		DBClose()
		log.Fatal(err.Error())
	}

	return *member
}

func loadReplyAuthorOf(reply models.Reply) models.Member {
	ctx := context.Background()

	member, err := reply.Member().OneG(ctx)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.Member{}
		}

		log.Printf("DB 작업 중 오류 발생 - loadAllReplyOf")
		DBClose()
		log.Fatal(err.Error())
	}

	return *member
}
