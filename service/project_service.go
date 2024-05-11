package service

import (
	"fmt"
	"gorm.io/gorm"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
	"time"
)

/*******************request dto start*******************************************/
type OtherLinksDto struct {
	OtherLinkUrl     string `json:"otherLinkUrl"`
	OtherLinkUrlType string `json:"otherLinkUrlType" default:"other"`
}

type ProjectCreateOrUpdateDto struct {
	Ticker      string          `json:"ticker"`
	ProjectName string          `json:"projectName"`
	Description string          `json:"description" `
	Twitter     string          `form:"twitter" json:"twitter" `
	OtherLinks  []OtherLinksDto `form:"otherLinks" json:"otherLinks" `
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

	otherLinks := param.OtherLinks
	if otherLinks != nil && len(otherLinks) != 0 {
		projectSocials := make([]model.MemooProjectSocial, 0)
		for _, item := range otherLinks {
			projectSocials = append(projectSocials, model.MemooProjectSocial{ProjectId: memooProject.ID,
				SocialUrl: item.OtherLinkUrl, SocialType: item.OtherLinkUrlType})
		}
		tx = dbTx.CreateInBatches(projectSocials, len(projectSocials))
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	return param, nil
}
