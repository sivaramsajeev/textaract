package main

import (
  "fmt"
  "io/ioutil"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "strings"
  "github.com/aws/aws-sdk-go/service/textract"
)

var textractSession *textract.Textract

func init() {
  textractSession = textract.New(session.Must(session.NewSession(&aws.Config{
    Region: aws.String("us-east-1"),
    })))
}

func main() {
  file, err := ioutil.ReadFile("I'm_Still_Worthy.jpg")
  Must(err)

  resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
    Document: &textract.Document {
      Bytes: file,
    },
  })
  Must(err)

  fmt.Printf("The following words have been identified from the picture \n%s\n", strings.Repeat("-", 100))
  for i:=1; i<len(resp.Blocks);i++ {
    if *resp.Blocks[i].BlockType == "LINE" {  //either a LINE/WORD
      fmt.Println(*resp.Blocks[i].Text)
    }
  }
}


func Must(err error) {
  if err != nil {
    panic(err)
  }
} 
