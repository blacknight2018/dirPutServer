package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func createFilePath(fileRelPath string, preDir string, fileData []byte) {
	if fileRelPath[0] == '/' {
		fileRelPath = fileRelPath[1:]
	}
	fileRelPath = preDir + "/" + fileRelPath
	spit := strings.Split(fileRelPath, "/")
	var prex string
	for i, v := range spit {
		if i == len(spit)-1 {
			break
		}
		os.Mkdir(prex+v, os.ModePerm)
		prex = prex + v + "/"
	}
	fs, err := os.Create(prex + spit[len(spit)-1])
	fmt.Println(fs, err)
	if err == nil {
		fs.Write(fileData)
		fs.Close()
	}

}

func main() {
	g := gin.Default()
	g.POST("", func(context *gin.Context) {
		fileRelPath := context.PostForm("file_path")
		fileData := context.PostForm("data")
		bytes, _ := base64.StdEncoding.DecodeString(fileData)
		saveDir := context.PostForm("dir")
		createFilePath(fileRelPath, saveDir, bytes)
		context.Writer.Write([]byte("ok"))
	})
	g.Run(":1424")
}
