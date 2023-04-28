package LinearFastSearchEngine

import (
	"bufio"
	"searchEngine/models"
	"strings"
)

type TermInfo struct {
	Term      string
	DocId     int
	Frequency int
}

type LinearEngine struct {
	data []TermInfo
}

func stopWorld(text string) bool {
	return text != "a" && text != "an" && text != "the" && text != "of" && text != "off" && text != "is"
}

func CreateLinearEngine(capacity int) *LinearEngine {
	return &LinearEngine{data: make([]TermInfo, 0, capacity)}
}

func (de *LinearEngine) AddData(txt *bufio.Scanner, DocId int) {
	txt.Split(bufio.ScanWords)
	for txt.Scan() {
		text := strings.ToLower(txt.Text())
		if stopWorld(text) && !(strings.Contains(text, ".")) && text != ";" && text != "!" && text != "?" {
			counter := 0
			if len(de.data) > 0 {
				for _, vData := range de.data {
					if vData.DocId == DocId && vData.Term == text {
						vData.Frequency++
						counter++
						break
					}
				}
			}
			if counter == 0 {
				newTermInfo := TermInfo{
					Term:      text,
					DocId:     DocId,
					Frequency: 1,
				}
				de.data = append(de.data, newTermInfo)
			}
		}
	}
}

func (de *LinearEngine) Search(q string) []models.Posting {
	out := make([]models.Posting, 0, len(q))
	for _, vData := range de.data {
		if vData.Term == q {
			newPosting := models.Posting{Id: vData.DocId, Frequency: vData.Frequency}
			out = append(out, newPosting)
		}
	}

	return out
}
