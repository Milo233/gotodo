package filemanager

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"gotodo/model"
	"net/http"
	"os"
	"strings"
)

// fixme 异常输出不统一
func UploadFileHandler(c *gin.Context) {
	//获取表单数据 参数为name值
	f, err := c.FormFile("uploadFile")
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	//将文件保存至本项目根目录中
	c.SaveUploadedFile(f, "files/" + f.Filename)
	if  strings.HasSuffix(f.Filename,".txt") {
		var book model.Book
		book.Lines = countFile(f.Filename) // fixme 要读取出来
		book.File = f.Filename
		book.CurrentPage = 1
		model.DB.Save(&book)
	}
	//保存成功返回正确的Json数据
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}


func countFile(filename string) int {
	file, _ := os.Open("files/" + filename)
	fileScanner := bufio.NewScanner(file)
	lines := 0
	for fileScanner.Scan() {
		fileScanner.Text() // todo 可以删除吗？
		lines++
	}
	defer file.Close()
	return lines
}