package search

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/team-four-fingers/kakao/local"
	"github.com/team-four-fingers/kakao/local/keyword"
	"math/rand"
	"net/http"
	"server/pkg/grouper"
	"server/server"
	"server/server/handler/common"
	"strconv"
)

var _ server.HTTPHandler = (*Handler)(nil)

var (
	productByCategoryGroupCode = map[string]Product{
		local.CategoryGroupCode.Cafe: {
			Name:     "아메리카노",
			Price:    5000,
			ImageUrl: "https://cdn.pixabay.com/photo/2017/05/07/08/56/pancakes-2291908_960_720.jpg",
		},
		local.CategoryGroupCode.Restaurant: {
			Name:     "불고기",
			Price:    10000,
			ImageUrl: "https://cdn.pixabay.com/photo/2017/05/07/08/56/pancakes-2291908_960_720.jpg",
		},
		local.CategoryGroupCode.Convenience: {
			Name:     "삼각김밥",
			Price:    2000,
			ImageUrl: "https://cdn.pixabay.com/photo/2017/05/07/08/56/pancakes-2291908_960_720.jpg",
		},
		local.CategoryGroupCode.DepartmentStore: {
			Name:     "삼겹살",
			Price:    12000,
			ImageUrl: "https://cdn.pixabay.com/photo/2017/05/07/08/56/pancakes-2291908_960_720.jpg",
		},
	}
)

type Handler struct {
	localCli local.Client
}

func NewHandler(localCli local.Client) *Handler {
	return &Handler{localCli: localCli}
}

func (h *Handler) Method() string {
	return http.MethodPost
}

func (h *Handler) Path() string {
	return "/search"
}

func (h *Handler) HandleFunc() func(c echo.Context) error {
	return func(c echo.Context) error {
		body := &RequestBody{}
		if err := c.Bind(body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var (
			origin   = body.Origin
			eatType  = body.EatType
			whenType = body.WhenType
			query    = body.Query
			radius   = body.Radius
		)

		categoryGroupCodeByEatType := map[string][]string{
			EatType.PickUp:         {local.CategoryGroupCode.Convenience},
			EatType.AtStore:        {local.CategoryGroupCode.Restaurant, local.CategoryGroupCode.Cafe},
			EatType.OnlyIngredient: {local.CategoryGroupCode.DepartmentStore, local.CategoryGroupCode.Convenience},
			EatType.SelfCooking:    {local.CategoryGroupCode.DepartmentStore, local.CategoryGroupCode.Convenience},
		}

		categoryGroupCodeByWhenType := map[string]string{
			WhenType.Brunch: local.CategoryGroupCode.Cafe,
			WhenType.Lunch:  local.CategoryGroupCode.Restaurant,
			WhenType.Dinner: local.CategoryGroupCode.Restaurant,
		}

		categoryGroupCodesToUse := make([]string, 0, 20)

		codesFromEatType, ok := categoryGroupCodeByEatType[eatType]
		if ok {
			categoryGroupCodesToUse = append(categoryGroupCodesToUse, codesFromEatType...)
		}

		codesFromWhenType, ok := categoryGroupCodeByWhenType[whenType]
		if ok {
			categoryGroupCodesToUse = append(categoryGroupCodesToUse, codesFromWhenType)
		}

		categoryGroupCodesToUse = lo.Uniq(categoryGroupCodesToUse)

		g := grouper.NewPanicSafeGroup()

		resultsSlices := make([][]Result, 4)
		for i, code := range categoryGroupCodesToUse {
			g.Go(func() error {
				results := make([]Result, 0, 15)

				keywordResp, err := h.localCli.SearchByKeyword(query, code, origin.X, origin.Y, radius)
				if err != nil {
					if errors.Is(err, local.ErrNoResult) {
						return nil
					}
					return err
				}

				for _, document := range keywordResp.Documents {
					store, err := documentToStore(document)
					if err != nil {
						return err
					}

					results = append(results, Result{
						ResultId: i,
						WhenTypes: []string{
							WhenType.Brunch,
						},
						Product: productByCategoryGroupCode[code],
						Store:   *store,
					})
				}

				resultsSlices[i] = results
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		results := lo.Flatten(resultsSlices)
		results = lo.UniqBy(results, func(result Result) string {
			return result.Store.Name
		})
		for i, _ := range results {
			results[i].ResultId = i
		}

		return c.JSON(http.StatusOK, &Response{
			Results: results,
		})
	}
}

func documentToStore(document keyword.Document) (*common.Store, error) {
	x, err := strconv.ParseFloat(document.X, 64)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(document.Y, 64)
	if err != nil {
		return nil, err
	}

	distance, err := strconv.ParseInt(document.Distance, 10, 64)
	if err != nil {
		return nil, err
	}

	return &common.Store{
		Name: document.PlaceName,
		Coordinate: common.Coordinate{
			X: x,
			Y: y,
		},
		OperationHours:     "9:00 ~ 18:00",
		HasParkingLot:      rand.Intn(2) == 1,
		DistanceFromOrigin: int(distance),
	}, nil
}

type keywordSearchFunc func(query string, origin common.Coordinate, radius int) ([]keyword.Document, error)

func (h *Handler) searchCafe(query string, origin common.Coordinate, radius int) ([]keyword.Document, error) {
	keywordResp, err := h.localCli.SearchByKeyword(query, local.CategoryGroupCode.Cafe, origin.X, origin.Y, radius)
	if err != nil {
		return nil, err
	}

	return keywordResp.Documents, nil
}

func (h *Handler) searchRestaurant(query string, origin common.Coordinate, radius int) ([]keyword.Document, error) {
	keywordResp, err := h.localCli.SearchByKeyword(query, local.CategoryGroupCode.Restaurant, origin.X, origin.Y, radius)
	if err != nil {
		return nil, err
	}

	return keywordResp.Documents, nil
}

func (h *Handler) searchConvenience(query string, origin common.Coordinate, radius int) ([]keyword.Document, error) {
	keywordResp, err := h.localCli.SearchByKeyword(query, local.CategoryGroupCode.Convenience, origin.X, origin.Y, radius)
	if err != nil {
		return nil, err
	}

	return keywordResp.Documents, nil
}
