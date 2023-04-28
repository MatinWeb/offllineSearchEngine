package LinearFastAddEngine

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
	"searchEngine/models"
	"strings"
	"testing"
)

func TestCreateLinearEngine(t *testing.T) {

	data := []struct {
		name     string
		cap      int
		expected int
	}{
		{
			"first", 1, 1,
		},
		{
			"second", 13, 13,
		},
		{
			"third", 24, 24,
		},
		{
			"fourth", 30, 30,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			got := d.cap
			result := CreateLinearEngine(got)
			if d.expected != cap(result.data) {
				t.Errorf("expected %d but got %d ", d.expected, cap(result.data))
			}
		})
	}
}

func TestAddData(t *testing.T) {

	var linear LinearEngine

	data := []struct {
		name     string
		text     string
		id       int
		expected int
	}{
		{
			"first", "hello matin ", 1, 2,
		},
		{
			"second", "how are you ? ", 2, 5,
		},
		{
			"third", "matin you are adult ? ", 3, 9,
		},
		{
			"fourth", "how do you design your plan ", 4, 15,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			scan := bufio.NewScanner(strings.NewReader(d.text))
			linear.AddData(scan, d.id)

			if len(linear.data) != d.expected {
				t.Errorf("expected %d but got %d ", d.expected, len(linear.data))
			}
		})
	}
}

func TestSearch(t *testing.T) {
	var linear LinearEngine
	addData := []struct {
		name     string
		text     string
		id       int
		expected int
	}{
		{
			"first", "hello matin ", 1, 2,
		},
		{
			"second", "how are you ? ", 2, 5,
		},
		{
			"third", "matin you are adult ? ", 3, 9,
		},
		{
			"fourth", "how do you design your plan ? you know what happened ? i dont think so ! you dont know ", 4, 15,
		},
	}

	for _, d := range addData {
		t.Run(d.name, func(t *testing.T) {
			scan := bufio.NewScanner(strings.NewReader(d.text))
			linear.AddData(scan, d.id)

		})
	}

	searchData := []struct {
		name     string
		text     string
		expected []models.Posting
	}{
		{
			name: "first", text: "matin ", expected: []models.Posting{
				{1, 1}, {3, 2},
			},
		},
		{
			name: "second", text: "you ", expected: []models.Posting{
				{2, 1}, {3, 1}, {4, 4},
			},
		},
	}
	for _, d := range searchData {
		t.Run(d.name, func(t *testing.T) {
			result := linear.Search(d.text)

			for iResult, _ := range result {
				if diff := cmp.Diff(d.expected[iResult], result[iResult]); diff != "" {
					t.Error(diff)
				}
			}

		})
	}

}
