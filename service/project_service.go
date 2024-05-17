package service

import (
	"fmt"
	"gorm.io/gorm"
	"memoo-backend/constants"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
	"mime/multipart"
	"time"
)

/*******************request dto start*******************************************/
type LinksDto struct {
	LinkUrl     string `form:"linkUrl" `
	LinkUrlType string `form:"linkUrlType"`
}

type ProjectCreateOrUpdateDto struct {
	Ticker               string                  `form:"ticker" binding:"required"`
	ProjectName          string                  `form:"projectName"`
	Website              string                  `form:"website"`
	Description          string                  `form:"description" binding:"required"`
	Twitter              string                  `form:"twitter"  binding:"required"`
	OtherLinkStr         string                  `form:"otherLinkStr"  `         //[{"linkUrl": "https://twitter.com/W8sFgv45Jt16576","linkUrlType":"twitter"}, {"linkUrl": "https://t.me/layercraftofficialchat","linkUrlType":"telegram"}]
	PinnedTwitterLinkStr string                  `form:"pinnedTwitterLinkStr"  ` //[{"linkUrl": "https://twitter.com/W8sFgv45Jt16576","linkUrlType":"twitter"}, {"linkUrl": "https://t.me/layercraftofficialchat","linkUrlType":"telegram"}]
	Banners              []*multipart.FileHeader `form:"banners" swaggerignore:"true"`
}

/*******************request dto end*******************************************/

/*******************service start*******************************************/
func ProjectNewOrEdit(param *ProjectCreateOrUpdateDto, address string, bannerUrls []string,
	otherLinks []LinksDto, pinnedTwitterLinks []LinksDto) (*ProjectCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	response, err := HandleProjectNewOrEdit(dbTx, param, address, bannerUrls, otherLinks, pinnedTwitterLinks)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	dbTx.Commit()
	return response, nil
}

func HandleProjectNewOrEdit(dbTx *gorm.DB, param *ProjectCreateOrUpdateDto, address string, bannerUrls []string,
	otherLinks []LinksDto, pinnedTwitterLinks []LinksDto) (*ProjectCreateOrUpdateDto, error) {
	updateStmt := fmt.Sprintf("insert into memoo_projects(created_at,updated_at,ticker,project_name,description,twitter) " +
		"values (?, ?, ?, ?, ?, ?) " +
		"on conflict (project_name) do  update set updated_at=?,description=?,twitter=?")
	tx := dbTx.Exec(updateStmt, time.Now(), time.Now(), param.Ticker, param.ProjectName, param.Description, param.Twitter,
		time.Now(), param.Description, param.Twitter)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var memooProject model.MemooProject
	dbTx.Table("memoo_projects").Where("project_name=?", param.ProjectName).First(&memooProject)

	projectBanners := make([]model.MemooProjectBanner, 0)
	for _, item := range bannerUrls {
		projectBanner := model.MemooProjectBanner{ProjectId: memooProject.ID, BannerUrl: item}
		projectBanners = append(projectBanners, projectBanner)
	}
	if projectBanners != nil && len(projectBanners) != 0 {
		tx = dbTx.CreateInBatches(projectBanners, len(projectBanners))
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	projectCreators := make([]model.MemooProjectCreator, 0)
	projectCreators = append(projectCreators, model.MemooProjectCreator{ProjectId: memooProject.ID, CreatorAddress: address})
	tx = dbTx.CreateInBatches(projectCreators, len(projectCreators))
	if tx.Error != nil {
		return nil, tx.Error
	}

	_, err := onHandleProjectSocials(otherLinks, dbTx, memooProject, constants.OTHER_LINK)
	if err != nil {
		return nil, err
	}
	_, err = onHandleProjectSocials(pinnedTwitterLinks, dbTx, memooProject, constants.PINNED_TWITTER_LINK)
	if err != nil {
		return nil, err
	}
	return param, nil
}

func onHandleProjectSocials(links []LinksDto, dbTx *gorm.DB, memooProject model.MemooProject, majorCategories string) (interface{}, error) {
	if links == nil || len(links) == 0 {
		return nil, nil
	}
	projectSocials := make([]model.MemooProjectSocial, 0)
	for _, item := range links {
		projectSocials = append(projectSocials, model.MemooProjectSocial{ProjectId: memooProject.ID,
			SocialUrl: item.LinkUrl, SocialType: item.LinkUrlType, MajorCategories: majorCategories})
	}
	tx := dbTx.CreateInBatches(projectSocials, len(projectSocials))
	if tx.Error != nil {
		return nil, tx.Error
	}
	return nil, nil
}

/*******************service end*******************************************/
