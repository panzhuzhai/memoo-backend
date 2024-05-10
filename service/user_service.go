package service

import (
	"fmt"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
	"time"
)

type UserCreateOrUpdateDto struct {
	UserName   string `form:"userName"  json:"userName" `
	UserBio    string `form:"userBio" json:"userBio" `
	Twitter    string `form:"twitter" json:"twitter" `
	WebsiteUrl string `form:"websiteUrl" json:"websiteUrl" `
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
