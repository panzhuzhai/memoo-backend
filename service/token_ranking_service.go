package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
)

/*******************response dto start*******************************************/

type TrendingTokensRespDto struct {
	TokenImageUrl  string  `json:"tokenImageUrl" `
	TokenName      string  `json:"tokenName" `
	Price          string  `json:"price" `
	Increase1H     float64 `json:"increase1H" `
	Increase24H    float64 `json:"increase24H" `
	Volume24H      float64 `json:"volume24H" `
	TokenMarketCap float64 `json:"topTokenMarketCap"`
	MeMooScore     string  `json:"meMooScore" `
}

type TopTokensRespDto struct {
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

/*******************service start*******************************************/

func TrendingTokens(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func TopTokens(param dto.PageDto) (database.Paginator, error) {
	return database.Paginator{}, nil
}

/*******************service end*******************************************/
