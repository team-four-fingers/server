package search

import "server/server/handler/common"

// 상품 검색 시 판매점 정보(물품, 판매점, 판매점 상세 정보, 위치, 주차여부 등)를 확인할 수 있다

type SearchRequest struct {
	Query          string             `json:"query"`
	Origin         *common.Coordinate `json:"origin"`
	RadiusInMeters int64              `json:"radius_in_meters"`
}

type SearchResponse struct {
}
