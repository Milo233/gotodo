package service

import (
	"bufio"
	"gotodo/model"
	"os"
	"strconv"
)

// ReadBookService
type ReadBookService struct {
	BookName string `form:"bookname"`
	Index int `form:"index"`
	PageSize int `form:"pagesize"`
}

//func (service *ReadBookService) ReadBook() serializer.Response {
func (service *ReadBookService) ReadBook()  map[string]string{
	//todos := []model.Todo{}
	//total := 0
	var book model.Book
	model.DB.Where("file = ?",service.BookName).Find(&book)
	if service.Index <= 0 {
		// 查询数据库 读取文件上次读到第几页了
		service.Index = book.CurrentPage
	} else {
		// update current_page
		book.CurrentPage = service.Index
		model.DB.Model(&book).Where("file = ?", service.BookName).Update("current_page",service.Index)
	}

	var resultMap = make(map[string]string)
	resultMap["total"] = strconv.Itoa(book.Lines / service.PageSize)
	resultMap["index"] = strconv.Itoa(service.Index)
	lastPageContent, current, nextPageContent := ReadLine(service.BookName, service.Index, service.PageSize)
	resultMap["lastPageContent"] = lastPageContent
	resultMap["current"] = current
	resultMap["nextPageContent"] = nextPageContent
	return resultMap
}

// 太慢可以预加载,
func ReadLine(filename string,index int,pagesize int) (string,string,string) {
	lineNumber := pagesize * (index - 1) + 1
	file, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(file)
	currentLine := 1
	lastPageContent := ""
	current := ""
	nextPageContent := ""
	for fileScanner.Scan() {
		if currentLine >= lineNumber - pagesize && currentLine < lineNumber {
			lastPageContent += fileScanner.Text() + "<br/>"
		} else if currentLine >= lineNumber && currentLine < lineNumber + pagesize {
			current += fileScanner.Text() + "<br/>"
		} else if currentLine >= lineNumber + pagesize && currentLine < lineNumber + 2 * pagesize {
			nextPageContent += fileScanner.Text() + "<br/>"
		} else if currentLine >= lineNumber + 2 * pagesize {
			break
		}
		currentLine++
	}
	defer file.Close()
	return lastPageContent, current, nextPageContent
}