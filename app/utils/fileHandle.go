package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// FileHandler 移动文件到合适的位置
func FileHandler(file *multipart.FileHeader) (string, error) {

	fileByte, _ := file.Open()
	defer fileByte.Close()

	now := strconv.Itoa(int(time.Now().Unix()))

	basePath := "/go/src/github.com/AnnatarHe/exam-online-be"
	bgPath := "/public/img/" + now + file.Filename

	f, err := os.Create(basePath + bgPath)
	if err != nil {
		return "", err
	}

	io.Copy(f, fileByte)
	return bgPath, nil
}
