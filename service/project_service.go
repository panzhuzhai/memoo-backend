package service

import (
	"fmt"
	"gorm.io/gorm"
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
	Ticker        string                  `form:"ticker" binding:"required"`
	ProjectName   string                  `form:"projectName"`
	Website       string                  `form:"website"`
	Introduction  string                  `form:"introduction"`
	Description   string                  `form:"description" binding:"required"`
	Twitter       string                  `form:"twitter"  binding:"required"`
	Telegram      string                  `form:"telegram"  `
	PinnedTwitter string                  `form:"pinnedTwitter"  `
	Banners       []*multipart.FileHeader `form:"banners" swaggerignore:"true"`
}

/*******************request dto end*******************************************/

/*******************service start*******************************************/
func ProjectNewOrEdit(param *ProjectCreateOrUpdateDto, address string, bannerUrls []string) (*ProjectCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	response, err := HandleProjectNewOrEdit(dbTx, param, address, bannerUrls)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	dbTx.Commit()
	return response, nil
}

func HandleProjectNewOrEdit(dbTx *gorm.DB, param *ProjectCreateOrUpdateDto, address string, bannerUrls []string) (*ProjectCreateOrUpdateDto, error) {
	updateStmt := fmt.Sprintf("insert into memoo_projects(created_at,updated_at,ticker,project_name,website,introduction,description,twitter,creator_address,telegram,pinned_twitter " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) " +
		"on conflict (ticker) do  update set updated_at=?,project_name=?,website=?,introduction=?,description=?,twitter=?,creator_address=?,telegram=?,pinned_twitter=?")
	tx := dbTx.Exec(updateStmt, time.Now(), time.Now(), param.Ticker, param.ProjectName, param.Website, param.Introduction, param.Description, param.Twitter, address, param.Telegram, param.PinnedTwitter,
		time.Now(), param.ProjectName, param.Website, param.Introduction, param.Description, param.Twitter, address, param.Telegram, param.PinnedTwitter)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var memooProject model.MemooProject
	dbTx.Table("memoo_projects").Where("project_name=?", param.ProjectName).First(&memooProject)
	var countBanner int64
	dbTx.Table("memoo_project_banners").Where("project_id=?", memooProject.ID).First(&countBanner)
	if countBanner != 0 {
		return param, nil
	}
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
	return param, nil
}

/*******************service end*******************************************/
