package service

import (
	"fmt"
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
	"memoo-backend/oss"
	"memoo-backend/serializer"
	"time"
)

/*******************request dto start*******************************************/
type UserCreateOrUpdateDto struct {
	UserName   string `form:"userName"  `
	UserBio    string `form:"userBio"  `
	Twitter    string `form:"twitter" `
	WebsiteUrl string `form:"websiteUrl"  `
}

/*******************request dto end*******************************************/

/*******************response dto start*******************************************/
type ProjectSocialDto struct {
	SocialUrl  string `json:"socialUrl" `
	SocialType string `json:"socialType" `
}

type UserViewOthersDto struct {
	ImageUrl                 string             `json:"imageUrl" `
	TokensCreatedNum         int64              `json:"tokensCreatedNum" `
	AccumulativeMarketCap    float64            `json:"accumulativeMarketCap" `
	AccumulativeATHMarketCap float64            `json:"accumulativeATHMarketCap" `
	TopTokenMarketCap        float64            `json:"accumulativeMarketCap" `
	TopTokenATHMarketCap     float64            `json:"accumulativeATHMarketCap" `
	TopTokenHolders          int64              `json:"accumulativeHolders" `
	TopTokenHoldersGrowth    int64              `json:"accumulativeHoldersGrowth" `
	TokenName                string             `json:"tokenName" `
	Description              string             `json:"description"`
	CreatorAddress           string             `json:"creatorAddress"`
	CreatorSocialsAccount    string             `json:"creatorAddress"`
	CreatorWebsiteUrl        string             `json:"creatorWebsiteUrl"`
	CreatorTwitter           string             `json:"CreatorTwitter"`
	ProjectSocial            []ProjectSocialDto `json:"projectSocial"`
}

type UserViewOthersListDto struct {
	ImageUrl    string `json:"imageUrl" `
	TokenName   string `json:"tokenName"`
	Status      string `json:"status"`
	TotalRaised string `json:"totalRaised"`
	LaunchDate  int64  `json:"launchDate" `
	MeMooScore  string `json:"meMooScore" `
	Emotion     string `json:"emotion" ` //WIF、FOMO、LEASH
}

/*******************response dto end*******************************************/

/*******************swagger use response start*******************************************/
type UserViewOthersResp struct {
	serializer.ResponseNotWithData
	Data UserViewOthersDto `json:"data,omitempty"`
}

type UserViewOthersListPaginator struct {
	database.PaginatorBase
	Records []UserViewOthersListDto `json:"records"`
}

type UserViewOthersListResp struct {
	serializer.ResponseNotWithData
	Data UserViewOthersListPaginator `json:"data,omitempty"`
}

/*******************swagger use response end*******************************************/

/*******************service start*******************************************/
func UserViewOthers(address string) (UserViewOthersDto, error) {
	return UserViewOthersDto{}, nil
}

func UserViewOthersList(param dto.PageDto, address string) (*database.Paginator, error) {
	db := database.DB.Table("memoo_tokens").Order("id asc")
	var records []UserViewOthersListDto
	paginator := database.Paging(&database.Param{
		DB:      db,
		Page:    int(param.PageNumber),
		Limit:   param.PageSize,
		ShowSQL: true,
	}, &records)
	recordNews := make([]UserViewOthersListDto, 0)
	for _, item := range records {
		item.ImageUrl = oss.GetS3ObjectURL(item.ImageUrl)
		recordNews = append(recordNews, item)
	}
	return paginator, nil
}

func UserNewOrEdit(param *UserCreateOrUpdateDto, address string, bannerUrls []string, profileImageUrls []string) (*UserCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	var memooUser model.MemooUser
	dbTx.Table("memoo_users").Where("address=?", address).First(&memooUser)
	updateStmt := fmt.Sprintf("insert into memoo_users(created_at,updated_at,user_name,user_bio,address,twitter,website_url,profile_image_url) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?) " +
		"on conflict (address) do  update set updated_at=?,user_name=?,user_bio=?,twitter=?,website_url=?,profile_image_url=?")
	profileImageUrl := ""
	if profileImageUrls != nil && len(profileImageUrls) != 0 {
		profileImageUrl = profileImageUrls[0]
	} else {
		profileImageUrl = memooUser.ProfileImageUrl
	}
	tx := dbTx.Exec(updateStmt, time.Now(), time.Now(), param.UserName, param.UserBio, address, param.Twitter, param.WebsiteUrl, profileImageUrl,
		time.Now(), param.UserName, param.UserBio, param.Twitter, param.WebsiteUrl, profileImageUrl)
	if tx.Error != nil {
		dbTx.Rollback()
		return nil, tx.Error
	}
	memooUserBanners := make([]model.MemooUserBanner, 0)
	for _, item := range bannerUrls {
		memooUserBanner := model.MemooUserBanner{UserAddress: address, ProfileBannerUrl: item}
		memooUserBanners = append(memooUserBanners, memooUserBanner)
	}
	if memooUserBanners == nil || len(memooUserBanners) == 0 {
		dbTx.Commit()
		return param, nil
	}
	tx = dbTx.CreateInBatches(memooUserBanners, len(memooUserBanners))
	if tx.Error != nil {
		dbTx.Rollback()
		return nil, tx.Error
	}
	dbTx.Commit()
	return param, nil
}

/*******************service end*******************************************/
