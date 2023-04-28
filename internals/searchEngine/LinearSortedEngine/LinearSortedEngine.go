package LinearSortedEngine

import (
	"bufio"
	"searchEngine/models"
	"sort"
	"strings"
)

type TermInfo struct {
	term      string
	docId     int
	frequency int
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

func (de *LinearEngine) AddData(txt *bufio.Scanner, docId int) {
	txt.Split(bufio.ScanWords)
	for txt.Scan() {
		text := strings.ToLower(txt.Text())
		if stopWorld(text) && !(strings.Contains(text, ".")) && text != ";" && text != "!" && text != "?" {
			counter := 0
			if len(de.data) > 0 {
				for _, vData := range de.data {
					if vData.docId == docId && vData.term == text {
						vData.frequency++
						counter++
						break
					}
				}
			}
			if counter == 0 {
				newTermInfo := TermInfo{
					term:      text,
					docId:     docId,
					frequency: 1,
				}
				de.data = append(de.data, newTermInfo)
			}

		}

	}
	sort.Slice(de.data, func(i, j int) bool {
		return de.data[i].term < de.data[j].term
	})
}

func binarySearch(input []TermInfo, search string) int {
	mid := len(input) / 2
	switch {
	case len(input) == 0:
		return -1
	case input[mid].term < search:
		return binarySearch(input[:mid], search)
	case input[mid].term > search:
		return binarySearch(input[mid+1:], search)
	default:
		return mid
	}
}

func (de *LinearEngine) Search(q string) []models.Posting {
	out := make([]models.Posting, 0, len(q))
	searchResult := binarySearch(de.data, q)
	if searchResult != -1 {
		out = append(out, models.Posting{Id: de.data[searchResult].docId, Frequency: de.data[searchResult].frequency})
		for i := searchResult - 1; i >= 0; i-- {
			if de.data[i].term == de.data[searchResult].term {
				out = append(out, models.Posting{Id: de.data[i].docId, Frequency: de.data[i].frequency})
			} else {
				break
			}
		}
		for i := searchResult + 1; i < len(de.data); i++ {
			if de.data[i].term == de.data[searchResult].term {
				out = append(out, models.Posting{Id: de.data[i].docId, Frequency: de.data[i].frequency})
			} else {
				break
			}
		}
	}

	return out

}
