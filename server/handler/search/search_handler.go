package search

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/server"
)

var _ server.HTTPHandler = (*SearchHandler)(nil)

// 상품 검색 시 판매점 정보(물품, 판매점, 판매점 상세 정보, 위치, 주차여부 등)를 확인할 수 있다

type SearchHandler struct{}

func (h *SearchHandler) Method() string {
	return http.MethodPost
}

func (h *SearchHandler) Path() string {
	return "/search"
}

func (h *SearchHandler) HandleFunc() func(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
