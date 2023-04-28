package searchEngine

import (
	"bufio"
	"searchEngine/models"
)

type ISearchEngine interface {
	AddData(txt *bufio.Scanner, docId int)
	Search(q string) []models.Posting
}
