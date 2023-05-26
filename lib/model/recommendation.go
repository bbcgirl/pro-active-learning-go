package model

import "fmt"

type RecommendationListType int

const (
	GENERAL RecommendationListType = 0
	ARTICLE RecommendationListType = 1
	GITHUB  RecommendationListType = 2
	SLIDE   RecommendationListType = 3
	ARXIV   RecommendationListType = 4
	VIDEO   RecommendationListType = 5
	EVENT   Recommendati