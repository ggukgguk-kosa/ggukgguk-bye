package step1

import (
	"fmt"
	"image"
	_ "image/jpeg" // JPEG 포맷 지원
	_ "image/png"  // PNG 포맷 지원
	"log"
	"os"
	"regexp"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-pdf/fpdf"
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

var filePath string

func PdfInit(config string) {
	filePath = config
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

func makeGreetingContent() string {
	greetingString := "안녕하세요, 꾹꾹입니다.\n"
	greetingString += "혜화에서의 반 년이 벌써 아득해지고, 꾹꾹도 이제 서버를 닫습니다.\n"
	greetingString += "대신 꾹꾹에서 함께해주신 추억 조각들을 정리해봤어요.\n\n\n"
	return greetingString
}

func makeMetaContent(member Author, diaries []MetaOfDocuments) string {
	var metaString string
	for _, m := range diaries {
		metaString += "- " + m.Year + "년 " + m.Month + "월: [" + m.Keyword + "]\n"
	}
	return metaString
}

func makeDocumentContent(member Author, document Document) (string, string) {
	headContnet := document.CreatedAt.Format("2006년 01월 02일 15시 04분")

	var mainContent string

	if document.ShareTo != "" && document.ShareTo == member.ID {
		mainContent += "(" + document.AuthorNickname + "님에게 받은 조각)\n"
	}

	if document.ShareTo != "" && document.ShareTo != member.ID {
		mainContent += "(" + document.ShareTo + "님에게 보낸 조각)\n"
	}

	mainContent += document.Content + "\n\n"
	if len(document.Replies) > 0 {
		mainContent += "댓글들 \n"
	}
	for _, r := range document.Replies {
		mainContent += r.Author + "\n" + r.CreatedAt.Format("2006년 01월 02일 15시 04분") + "\n" + r.Content + "\n\n"
	}

	emojiPattern := regexp.MustCompile(`[\p{So}\p{Sk}\p{Sm}\p{Sc}\p{S}]`)
	mainContent = emojiPattern.ReplaceAllString(mainContent, "") // prevent error - "character outside the supported range"

	return headContnet, mainContent
}

func MakePdf(member Author, diaries []MetaOfDocuments, documents []Document) {
	log.Println("step1 - out - " + member.Email)

	if member.Email == "" || len(documents) == 0 {
		log.Println("step1 - skip - " + member.Email)
		return
	}

	pdf := fpdf.New("P", "mm", "A4", "")

	// first page
	pdf.AddPage()
	pdf.ImageOptions("./static/logo.png", 40, 20, 128, 0, false, fpdf.ImageOptions{}, 0, "")
	pdf.ImageOptions("./static/phrase.png", 40, 80, 128, 0, false, fpdf.ImageOptions{}, 0, "")

	pdf.AddUTF8Font("Pretendard", "", "./static/PretendardVariable.ttf")
	pdf.SetFont("Pretendard", "", 24)
	pdf.SetFont("Pretendard", "", 14)
	pdf.MultiCell(0, 10, "\n\n\n\n\n\n\n\n\n\n\n"+makeGreetingContent(), "", "", false)

	pdf.SetFont("Pretendard", "", 14)
	if len(diaries) > 0 {
		pdf.MultiCell(0, 10, "\n"+member.Nickname+"님의 다이어리\n"+makeMetaContent(member, diaries), "", "", false)
	}

	// next pages
	for _, d := range documents {
		pdf.AddPage()

		header, content := makeDocumentContent(member, d)
		pdf.SetFont("Pretendard", "", 18)
		pdf.MultiCell(0, 10, header, "", "", false)

		isMediaInserted := false
		var mediaPath string

		if d.Media.ID != "" && d.Media.Type == "image" {
			mediaPath = filePath + "/image/" + d.Media.ID
		}

		if d.Media.ID != "" && d.Media.Type == "video" {
			mediaPath = filePath + "/thumbnail/" + d.Media.ID + "/EXTRACT_-0001.jpg"
		}

		if _, err := os.Stat(mediaPath); err == nil {
			mtype, err := mimetype.DetectFile(mediaPath)
			if err != nil {
				log.Println("It is not valid file: ")
				log.Println(err)
			}

			var imageOptions fpdf.ImageOptions
			if mtype.String() == "image/jpeg" {
				imageOptions = fpdf.ImageOptions{
					ImageType:             "JPEG",
					ReadDpi:               true,
					AllowNegativePosition: true,
				}
			} else if mtype.String() == "image/png" {
				imageOptions = fpdf.ImageOptions{
					ImageType:             "PNG",
					ReadDpi:               true,
					AllowNegativePosition: true,
				}
			}

			pdf.ImageOptions(mediaPath, 40, 30, 128, 0, false, imageOptions, 0, "")
			isMediaInserted = true
		}

		if isMediaInserted {
			file, err := os.Open(mediaPath)
			if err != nil {
				log.Println("이미지를 여는 중 오류 발생: ")
				log.Println(mediaPath)
			}
			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				log.Println("이미지를 디코딩하는 중 오류 발생: ")
				log.Println(err)

				content = "\n\n\n\n\n\n\n\n\n\n\n" + content
			} else {
				var padding string

				if img.Bounds().Dx() > img.Bounds().Dy() {
					padding = "\n\n\n\n\n\n\n\n\n\n\n" // 가로 이미지
				} else {
					padding = "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n" // 세로 이미지
				}
				content = padding + content
			}
		}

		pdf.SetFont("Pretendard", "", 12)
		pdf.MultiCell(0, 10, content, "", "", false)
	}

	if err := os.MkdirAll("./workdir/"+member.Email+"/", os.ModePerm); err != nil {
		fmt.Println(err)
	}

	err := pdf.OutputFileAndClose("./workdir/" + member.Email + "/ggukgguk-archive.pdf")

	if err != nil {
		log.Printf("PDF 작업 중 오류 발생 - MakePdf")
		log.Printf(err.Error())
	}

	log.Println("step1 - end - " + member.Email)
}
