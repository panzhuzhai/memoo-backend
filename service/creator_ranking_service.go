package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
)

/*******************response dto start*******************************************/
type TrendingCreatorsRespDto struct {
	CreatorsName              string  `form:"creatorsName"  json:"creatorsName" `
	CreatorsAddress           string  `form:"creatorsAddress" json:"creatorsAddress" `
	TokensCreatedNum          int64   `form:"tokensCreatedNum" json:"tokensCreatedNum" `
	AccumulativeMarketCap     float64 `form:"accumulativeMarketCap" json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap  float64 `form:"accumulativeATHMarketCap"  json:"accumulativeATHMarketCap" `
	AccumulativeHolders       int64   `form:"accumulativeHolders" json:"accumulativeHolders" `
	AccumulativeHoldersGrowth int64   `form:"accumulativeHoldersGrowth" json:"accumulativeHoldersGrowth" `
	TopTokenMarketCap         float64 `form:"topTokenMarketCap" json:"topTokenMarketCap"`
}

type TopCreatorsRespDto struct {
	CreatorsName             string  `form:"creatorsName"  json:"creatorsName" `
	TokensCreatedNum         int64   `form:"tokensCreatedNum" json:"tokensCreatedNum" `
	AccumulativeMarketCap    float64 `form:"accumulativeMarketCap" json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap float64 `form:"accumulativeATHMarketCap"  json:"accumulativeATHMarketCap" `
	AccumulativeHolders      int64   `form:"accumulativeHolders" json:"accumulativeHolders" `
	TopTokenMarketCap        float64 `form:"topTokenMarketCap" json:"topTokenMarketCap"`
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
