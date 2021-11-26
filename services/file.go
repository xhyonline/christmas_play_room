package services

import (
	"embed"
	"github.com/xhyonline/xutil/logger"
	"io/fs"
)

//go:embed template/*
var StaticFS embed.FS

// GetStaticHTML
func GetStaticHTML(filename string) string {
	body, err := fs.ReadFile(StaticFS, "template/html/"+filename)
	if err != nil {
		logger.Error("找不到文件,出现错误 :", err.Error())
		return ""
	}
	return string(body)
}
