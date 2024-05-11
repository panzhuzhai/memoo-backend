package service

import (
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"time"
)

/*******************respnose dto start*******************************************/

type IdoActiveRespDto struct {
	TokenImageUrl string `json:"tokenImageUrl" `
	TokenName     string `json:"tokenName"`
	EndsIn        int64  `json:"endsIn"`
	MeMooScore    string `json:"meMooScore" `
	TotalRaised   string `json:"totalRaised"`
	Status        string `json:"status"`
}

type IdoCompletedRespDto struct {
	TokenImageUrl string `json:"tokenImageUrl" `
	TokenName     string `json:"tokenName"`
	Ticker        string `json:"ticker"`
	AthRoi        string `json:"athRoi"`
	MeMooScore    string `json:"meMooScore" `
	TotalRaised   string `json:"totalRaised"`
}

type IdoUpcomingRespDto struct {
	TokenImageUrl string `json:"tokenImageUrl" `
	TokenName     string `json:"tokenName"`
	IdoDate       int64  `json:"idoDate"`
	MeMooScore    string `json:"meMooScore" `
	TotalRaised   string `json:"totalRaised"`
	Status        string `json:"status"`
}

type MeMooScoreBreakdownDto struct {
	SocialInfo       string `json:"socialInfo" `
	CommunitySize    string `json:"communitySize" `
	CommunityActivit string `json:"communityActivit" `
	Commitment       string `json:"commitment" `
	CreatorActivity  string `json:"creatorActivity" `
	TotalRaised      string `json:"totalRaised" `
	MarketCap        string `json:"marketCap" `
	Liquidity        string `json:"liquidity" `
	Holders          string `json:"holders" `
}

type ProjectCreatorDto struct {
	CreatorAddress string `json:"creatorAddress" `
	UserName       string `json:"userName" `
	Twitter        string `json:"twitter" `
}

type SocialInfoDto struct {
	ProjectWebsite string              `json:"projectWebsite" `
	ProjectSocial  []ProjectSocialDto  `json:"projectSocial"`
	ProjectCreator []ProjectCreatorDto `json:"projectCreator"`
}

type TokenInfoDto struct {
	TokenImageUrl     string    `json:"tokenImageUrl" `
	Description       string    `json:"description"`
	TokenName         string    `json:"tokenName"`
	Ticker            string    `json:"ticker"`
	ContractAddress   string    `json:"contractAddress" `
	LPContractAddress string    `json:"lpContractAddress" `
	CreatedAt         time.Time `json:"createdAt" `
}

type IdoUpcomingDetailRespDto struct {
	MeMooScore string `json:"meMooScore" `
	EndsIn     int64  `json:"endsIn"`
	IdoDate    int64  `json:"idoDate"`
	MaxSupply  int64  `json:"maxSupply"`
	FDV        int64  `json:"fdv"`
	TokenInfoDto
	MeMooScoreBreakdownDto
	SocialInfoDto
}

type IdoActiveDetailRespDto struct {
	MeMooScore  string `json:"meMooScore" `
	EndsIn      int64  `json:"endsIn"`
	Price       string `json:"price" `
	TotalRaised string `json:"totalRaised" `
	Contributed string `json:"contributed" `
	IdoDate     int64  `json:"idoDate"`
	MaxSupply   int64  `json:"maxSupply"`
	FDV         int64  `json:"fdv"`
	TokenInfoDto
	MeMooScoreBreakdownDto
	SocialInfoDto
}

type IdoDateDto struct {
	IdoDate     int64   `json:"idoDate"`
	LpLock      bool    `json:"lpLock"`
	Liquidity   bool    `json:"liquidity"`
	FDV         int64   `json:"fdv"`
	Volume24H   float64 `json:"volume24H" `
	Increase1H  float64 `json:"increase1H" `
	Increase24H float64 `json:"increase24H" `
	MaxSupply   int64   `json:"maxSupply"`
	AllTimeHigh string  `json:"allTimeHigh" `
	AllTimeLow  string  `json:"allTimeLow" `
	Holders     int64   `json:"holders"`
}
type IdoLaunchedDetailRespDto struct {
	MeMooScore  string `json:"meMooScore" `
	EndsIn      int64  `json:"endsIn"`
	MarketCap   string `json:"marketCap" `
	Price       string `json:"price" `
	TotalRaised string `json:"totalRaised" `
	Contributed string `json:"contributed" `
	Count       int64  `json:"count" `
	IdoDateDto
	TokenInfoDto
	MeMooScoreBreakdownDto
	SocialInfoDto
}

type IdoLaunchedDetailTop10ListDto struct {
	Address    string `json:"address" `
	Proportion string `json:"proportion" `
}

/*******************respnose dto end*******************************************/

/*******************service start*******************************************/
func IdoActive(param dto.PageDto, address string) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func IdoCompleted(param dto.PageDto, address string) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func IdoUpcoming(param dto.PageDto, address string) (database.Paginator, error) {
	return database.Paginator{}, nil
}
func IdoUpcomingDetail(address string) (IdoUpcomingDetailRespDto, error) {
	return IdoUpcomingDetailRespDto{}, nil
}

func IdoActiveDetail(address string) (IdoActiveDetailRespDto, error) {
	return IdoActiveDetailRespDto{}, nil
}

func IdoLaunchedDetail(address string) (IdoLaunchedDetailRespDto, error) {
	return IdoLaunchedDetailRespDto{}, nil
}

func IdoLaunchedDetailTop10List(param dto.PageDto, address string) (database.Paginator, error) {
	return database.Paginator{}, nil
}

/*******************service end*******************************************/
