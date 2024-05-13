package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"memoo-backend/serializer"
)

/*******************response dto start*******************************************/
type TrendingCreatorsDto struct {
	UserProfileImageUrl       string  `json:"userProfileImageUrl"`
	CreatorsName              string  `json:"creatorsName" `
	CreatorsAddress           string  `json:"creatorsAddress" `
	TokensCreatedNum          int64   `json:"tokensCreatedNum" `
	AccumulativeMarketCap     float64 `json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap  float64 `json:"accumulativeATHMarketCap" `
	AccumulativeHolders       int64   `json:"accumulativeHolders" `
	AccumulativeHoldersGrowth int64   `json:"accumulativeHoldersGrowth" `
	TopTokenMarketCap         float64 `json:"topTokenMarketCap"`
}

type TopCreatorsDto struct {
	UserProfileImageUrl      string  `json:"userProfileImageUrl"`
	CreatorsName             string  `json:"creatorsName" `
	TokensCreatedNum         int64   `json:"tokensCreatedNum" `
	AccumulativeMarketCap    float64 `json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap float64 `json:"accumulativeATHMarketCap" `
	AccumulativeHolders      int64   `json:"accumulativeHolders" `
	TopTokenMarketCap        float64 `json:"topTokenMarketCap"`
}

/*******************response dto end*******************************************/
/*******************response swagger dto start*******************************************/

type TrendingCreatorsPaginator struct {
	database.PaginatorBase
	Records []TrendingCreatorsDto `json:"records"`
}

type TrendingCreatorsResp struct {
	serializer.ResponseNotWithData
	Data TrendingCreatorsPaginator `json:"data,omitempty"`
}

type TopCreatorsPaginator struct {
	database.PaginatorBase
	Records []TopCreatorsDto `json:"records"`
}

type TopCreatorsResp struct {
	serializer.ResponseNotWithData
	Data TopCreatorsPaginator `json:"data,omitempty"`
}

/*******************response swagger dto end*******************************************/

/*******************service start*******************************************/
func TrendingCreators(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func TopCreators(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

/*******************service end*******************************************/
