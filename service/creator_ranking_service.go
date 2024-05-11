package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
)

/*******************response dto start*******************************************/
type TrendingCreatorsRespDto struct {
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

type TopCreatorsRespDto struct {
	UserProfileImageUrl      string  `json:"userProfileImageUrl"`
	CreatorsName             string  `json:"creatorsName" `
	TokensCreatedNum         int64   `json:"tokensCreatedNum" `
	AccumulativeMarketCap    float64 `json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap float64 `json:"accumulativeATHMarketCap" `
	AccumulativeHolders      int64   `json:"accumulativeHolders" `
	TopTokenMarketCap        float64 `json:"topTokenMarketCap"`
}

/*******************response dto end*******************************************/

/*******************service start*******************************************/
func TrendingCreators(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func TopCreators(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

/*******************service end*******************************************/
