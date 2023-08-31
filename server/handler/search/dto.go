package search

import "server/server/handler/common"

type RequestBody struct {
	Query    string            `json:"query"`
	Origin   common.Coordinate `json:"origin"`
	Radius   int               `json:"radius"`
	WhenType string            `json:"when_type"`
	EatType  string            `json:"eat_type"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
}

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	ResultId  int          `json:"result_id"`
	WhenTypes []string     `json:"when_types"`
	Product   Product      `json:"product"`
	Store     common.Store `json:"store"`
}

var WhenType = struct {
	Breakfast              string
	Lunch                  string
	Dinner                 string
	Brunch                 string
	BeforeAndAfterExercise string
	AfterExercise          string
	Light                  string
	RomanticDate           string
}{
	Breakfast:              "아침",
	Lunch:                  "점심",
	Dinner:                 "저녁",
	Brunch:                 "브런치",
	BeforeAndAfterExercise: "운동 전/후",
	Light:                  "가벼운",
	RomanticDate:           "로맨틱 데이트",
}

var EatType = struct {
	SelfCooking    string
	PickUp         string
	AtStore        string
	OnlyIngredient string
}{
	SelfCooking:    "직접조리",
	PickUp:         "픽업",
	AtStore:        "매장식사",
	OnlyIngredient: "재료만",
}
