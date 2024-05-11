package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
)

/*******************response dto start*******************************************/

type TrendingTokensRespDto struct {
	TokenImageUrl  string  `json:"tokenImageUrl" `
	TokenName      string  `form:"tokenName"  json:"tokenName" `
	Price          string  `form:"price" json:"price" `
	Increase1H     float64 `form:"increase1H" json:"increase1H" `
	Increase24H    float64 `form:"increase24H" json:"increase24H" `
	Volume24H      float64 `form:"volume24H"  json:"volume24H" `
	TokenMarketCap float64 `form:"topTokenMarketCap" json:"topTokenMarketCap"`
	MeMooScore     string  `json:"meMooScore" `
}

type TopTokensRespDto struct {
	TokenImageUrl  string  `json:"tokenImageUrl" `
	TokenName      string  `form:"tokenName"  json:"tokenName" `
	Price          string  `form:"price" json:"price" `
	Increase1H     float64 `form:"increase1H" json:"increase1H" `
	Increase24H    float64 `form:"increase24H" json:"increase24H" `
	Volume24H      float64 `form:"volume24H"  json:"volume24H" `
	TokenMarketCap float64 `form:"topTokenMarketCap" json:"topTokenMarketCap"`
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
