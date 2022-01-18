package util

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context, fileName, pathFolder string) (string, error) {
	exist, err := Exists(pathFolder)
	if !exist {
		err := os.Mkdir(pathFolder, 0777)
		if err != nil {
			return "can't create folder " + pathFolder, err
		}
	}

	file, err := c.FormFile(fileName)
	if err != nil {
		return "form file undefined", err
	}

	filename := pathFolder + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		return "error save file", err
	}

	return "file saved", err
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
