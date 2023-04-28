package LinearFastAddEngine

import (
	"bufio"
	"searchEngine/models"
	"strings"
)

type TermInfo struct {
	term  string
	docId int
}

type LinearEngine struct {
	data []TermInfo
}

func stopWords(text string) bool {
	return text != "a" && text != "an" && text != "the" && text != "of" && text != "off" && text != "is"
}

func removePunctuation(text string) bool {
	return !(strings.Contains(text, ".")) && text != ";" && text != "!" && text != "?"
}

func CreateLinearEngine(capacity int) *LinearEngine {
	return &LinearEngine{data: make([]TermInfo, 0, capacity)}
}

func (de *LinearEngine) AddData(txt *bufio.Scanner, docId int) {
	txt.Split(bufio.ScanWords)
	for txt.Scan() {
		text := strings.ToLower(txt.Text())
		if stopWords(text) && removePunctuation(text) {
			newTermInfo := TermInfo{
				term:  text,
				docId: docId,
			}
			de.data = append(de.data, newTermInfo)
		}

	}
}

type PostingSlice struct {
	postings []models.Posting
}

func (p *PostingSlice) Output(docId int) {
	for _, v := range p.postings {
		if v.Id == docId {
			v.Frequency++
		} else {
			newPosting := models.Posting{Id: docId, Frequency: 1}
			p.postings = append(p.postings, newPosting)
		}
	}
}

func find(id int, res []models.Posting) int {
	for i, v := range res {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (de *LinearEngine) Search(q string) []models.Posting {
	q = strings.ToLower(q)
	var out []models.Posting
	for _, vData := range de.data {
		if q == vData.term {
			v := find(vData.docId, out)
			if v == -1 {
				out = append(out, models.Posting{Id: vData.docId, Frequency: 1})
			} else {
				out[v].Frequency++
			}
		}
	}
	return out
}
