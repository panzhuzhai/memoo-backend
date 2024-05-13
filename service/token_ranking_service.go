package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"memoo-backend/serializer"
)

/*******************response dto start*******************************************/

type TrendingTokensDto struct {
	TokenImageUrl  string  `json:"tokenImageUrl" `
	TokenName      string  `json:"tokenName" `
	Price          string  `json:"price" `
	Increase1H     float64 `json:"increase1H" `
	Increase24H    float64 `json:"increase24H" `
	Volume24H      float64 `json:"volume24H" `
	TokenMarketCap float64 `json:"topTokenMarketCap"`
	MeMooScore     string  `json:"meMooScore" `
}

type TopTokensDto struct {
	TokenImageUrl  string  `json:"tokenImageUrl" `
	TokenName      string  `json:"tokenName" `
	Price          string  `json:"price" `
	Increase1H     float64 `json:"increase1H" `
	Increase24H    float64 `json:"increase24H" `
	Volume24H      float64 `json:"volume24H" `
	TokenMarketCap float64 `json:"topTokenMarketCap"`
	MeMooScore     string  `json:"meMooScore" `
}

/*******************response dto end*******************************************/
/*******************response swagger dto start*******************************************/

type TrendingTokensPaginator struct {
	database.PaginatorBase
	Records []TrendingTokensDto `json:"records"`
}

type TrendingTokensResp struct {
	serializer.ResponseNotWithData
	Data TrendingTokensPaginator `json:"data,omitempty"`
}

type TopTokensPaginator struct {
	database.PaginatorBase
	Records []TopTokensDto `json:"records"`
}

type TopTokensResp struct {
	serializer.ResponseNotWithData
	Data TopTokensPaginator `json:"data,omitempty"`
}

/*******************response swagger dto end*******************************************/

/*******************service start*******************************************/

func TrendingTokens(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func TopTokens(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

/*******************service end*******************************************/
