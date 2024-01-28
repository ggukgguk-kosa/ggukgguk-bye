package step2

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var mailSrv *gmail.Service
var basePath string

// see: https://developers.google.com/gmail/api/quickstart/go
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// see: https://developers.google.com/gmail/api/quickstart/go
// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// see: https://developers.google.com/gmail/api/quickstart/go
// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// see: https://developers.google.com/gmail/api/quickstart/go
// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// see: https://developers.google.com/gmail/api/quickstart/go
func MailInit() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	mailSrv = srv
	basePath = "./workdir"
}

func sendAllPdf() {
	mailList := getAllDirs()
	for _, mail := range mailList {
		log.Println("step2 - in - " + mail)

		mailContent := "안녕하세요, 꾹꾹입니다. \n\n"
		mailContent += "벌써 새해가 되고, 2월을 앞두고 있는데, 잘 지내고 계신가요?\n"
		mailContent += "다름이 아니라, KOSA 교육과정이 종료된 이후 거의 반 년 이상이 지났고, 이제는 꾹꾹 서버를 완전히 닫으려고 합니다.\n"
		mailContent += "다만 여러분들께서 올려주셨던 조각들을 그대로 지우기 아쉬워서, 유저별 조각을 PDF 파일로 한데 모아 공유드리려고 해요.\n"
		mailContent += "열정적으로 함께 공부했던 그 시절을 다시 한 번 떠올려보시기 바라며, 앞으로 우리 모두에게 좋은 일만 가득했으면 좋겠습니다.\n\n"
		mailContent += "꾹꾹 일동 올림\n(김진웅, 이나현, 임영택, 한재성)\n\n\n"
		mailContent += "* PDF 생성 및 메일 전송은 모두 자동화된 과정으로 진행했기 때문에, 누구도 아카이브를 열어보지 않았으니 안심하셔도 됩니다."

		sendEmailWithAttachment(mail, "[꾹꾹] 서버 종료 및 아카이브 공유", mailContent, "/"+mail+"/ggukgguk-archive.pdf")

		log.Println("step2 - end - " + mail)
	}
}

func getAllDirs() []string {
	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Println("파일을 읽는 중 오류 발생")
		log.Fatal(err)
	}

	var dirList []string

	for _, file := range files {
		dirList = append(dirList, file.Name())
	}

	return dirList
}

func sendEmailWithAttachment(to, subject, body, attachmentPath string) {
	_, err := mailSrv.Users.Messages.Send("me", createMessageWithAttachment(to, subject, body, attachmentPath)).Do()

	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	} else {
		log.Println("step2 - out - " + to)
	}
}

// original author: https://socketloop.com/tutorials/golang-send-email-with-attachment-rfc2822-using-gmail-api-example
func createMessageWithAttachment(to string, subject string, content string, filePath string) *gmail.Message {
	var message gmail.Message

	fileBytes, err := os.ReadFile(basePath + filePath)
	if err != nil {
		log.Fatalf("Unable to read file for attachment: %v", err)
	}

	fileMIMEType := "application/pdf"

	fileData := base64.StdEncoding.EncodeToString(fileBytes)

	boundary := randStr(32, "alphanum")

	encodedSubject := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="

	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"to: " + to + "\n" +
		"subject: " + encodedSubject + "\n\n" +

		"--" + boundary + "\n" +
		"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		content + "\n\n" +
		"--" + boundary + "\n" +

		"Content-Type: " + fileMIMEType + "; name=" + string('"') + filepath.Base(filePath) + string('"') + " \n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; filename=" + string('"') + filepath.Base(filePath) + string('"') + " \n\n" +
		chunkSplit(fileData, 76, "\n") +
		"--" + boundary + "--")

	message.Raw = base64.URLEncoding.EncodeToString(messageBody)

	return &message
}

// original author: https://socketloop.com/tutorials/golang-send-email-with-attachment-rfc2822-using-gmail-api-example
func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

// original author: https://socketloop.com/tutorials/golang-send-email-with-attachment-rfc2822-using-gmail-api-example
func chunkSplit(body string, limit int, end string) string {

	var charSlice []rune

	for _, char := range body {
		charSlice = append(charSlice, char)
	}

	var result string = ""

	for len(charSlice) >= 1 {
		result = result + string(charSlice[:limit]) + end

		charSlice = charSlice[limit:]

		if len(charSlice) < limit {
			limit = len(charSlice)
		}

	}

	return result

}
