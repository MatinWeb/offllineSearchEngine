package searchEngine

import (
	invertedindex "searchEngine/internals/searchEngine/InvertedIndex"
	"searchEngine/internals/searchEngine/LinearFastAddEngine"

	//"searchEngine/internals/searchEngine/LinearFastAddEngine"
	"searchEngine/internals/searchEngine/LinearFastSearchEngine"
	"searchEngine/internals/searchEngine/LinearSortedEngine"
	"searchEngine/internals/searchEngine/LinearSortedEngineWithPosting"
	//"searchEngine/internals/searchEngine/LinearFastSearchEngine"
)

func NewSearchEngine(s string, cap int) ISearchEngine {
	switch s {
	case "LinearFastAddEngine":
		return LinearFastAddEngine.CreateLinearEngine(cap)
	case "LinearFastSearchEngine":
		return LinearFastSearchEngine.CreateLinearEngine(cap)
	case "LinearSortedEngine":
		return LinearSortedEngine.CreateLinearEngine(cap)
	case "LinearSortedEngineWithPosting":
		return LinearSortedEngineWithPosting.CreateLinearEngine(cap)
	case "Invertedindex":
		return invertedindex.CreateInvertedIndex(cap)
	default:
		return nil
	}
}
