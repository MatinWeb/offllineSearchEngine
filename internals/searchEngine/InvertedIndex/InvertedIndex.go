package invertedindex

import (
	"bufio"
	"searchEngine/models"
	"strings"
)

type InvertedIndex struct {
	data map[string][]models.Posting
}

func stopWords(text string) bool {
	return text != "a" && text != "an" && text != "the" && text != "of" && text != "off" && text != "is"
}

func removePunctuation(text string) bool {
	return !(strings.Contains(text, ".")) && text != ";" && text != "!" && text != "?"
}

func CreateInvertedIndex(capacity int) *InvertedIndex {
	return &InvertedIndex{data: make(map[string][]models.Posting, capacity)}
}

func (in *InvertedIndex) AddData(text *bufio.Scanner, docId int) {
	text.Split(bufio.ScanWords)
	for text.Scan() {
		text := strings.ToLower(text.Text())
		if stopWords(text) && removePunctuation(text) {
			_, ok := in.data[text]
			if !ok {
				in.data[text] = []models.Posting{{Id: docId, Frequency: 1}}
			} else {
				lastIndex := len(in.data[text]) - 1
				if in.data[text][lastIndex].Id == docId {
					in.data[text][lastIndex].Frequency++
				} else {
					in.data[text] = append(in.data[text], models.Posting{Id: docId, Frequency: 1})
				}
			}
		}
	}
}

func (in *InvertedIndex) Search(q string) []models.Posting {
	res, ok := in.data[q]
	if !ok {
		return nil
	}
	return res
}
