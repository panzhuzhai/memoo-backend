package service

import (
	"fmt"
	"gorm.io/gorm"
	"memoo-backend/constants"
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
	"memoo-backend/serializer"
	"mime/multipart"
	"time"
)

const (
	DRAFT    = "Draft"
	QUEUE    = "Queue"
	IDO      = "IDO"
	LAUNCHED = "Launched"
)

/*******************request dto start*******************************************/

type TokenListReqDto struct {
	status string `json:"status"`
	dto.PageDto
}

type TokenDetailReqDto struct {
	Ticker string `json:"ticker"`
}

// TokenCreateOrUpdateDto represents the request body that includes an array of files.
type TokenCreateOrUpdateDto struct {
	TokenName                string                `form:"tokenName" binding:"required"` //
	ContractAddress          string                `form:"contractAddress"`
	LPContractAddress        string                `form:"lPContractAddress"`
	TokenDescription         string                `form:"tokenDescription" `
	PreLaunchDuration        string                `form:"preLaunchDuration" binding:"required"` //IMMEDIATE、1DAY、3DAY
	PreMarketAcquisition     string                `form:"preMarketAcquisition" binding:"required"`
	TokenIcon                *multipart.FileHeader `form:"tokenIcon" binding:"required" swaggerignore:"true"`
	ProjectCreateOrUpdateDto                       //tokenIcon tokenIcon files list banner banner iles list
}

/*******************request dto end*******************************************/

/*******************response dto start*******************************************/
type TokenListDto struct {
	Ticker      string `json:"ticker" `
	TokenIcon   string `json:"tokenIcon" `
	TokenName   string `json:"tokenName" `
	TotalRaised string `json:"totalRaised"`
	LaunchDate  int64  `json:"launchDate" `
	MeMooScore  string `json:"meMooScore" `
	Status      string `json:"status" `
}

type TokenDetailDto struct {
	TokenName            string `json:"tokenName" ` //
	ContractAddress      string `json:"contractAddress"`
	LPContractAddress    string `json:"lPContractAddress"`
	TokenDescription     string `json:"tokenDescription" `
	PreLaunchDuration    string `json:"preLaunchDuration" ` //IMMEDIATE、1DAY、3DAY
	PreMarketAcquisition string `json:"preMarketAcquisition" `
	TokenIcon            string `json:"tokenIcon" `

	Ticker            string   `json:"ticker" `
	ProjectName       string   `json:"projectName"`
	Website           string   `json:"website"`
	Description       string   `json:"description" `
	Twitter           string   `json:"twitter"  `
	OtherLink         []string `json:"otherLinkStr"  `         //[{"linkUrl": "https://twitter.com/W8sFgv45Jt16576","linkUrlType":"twitter"}, {"linkUrl": "https://t.me/layercraftofficialchat","linkUrlType":"telegram"}]
	PinnedTwitterLink []string `json:"pinnedTwitterLinkStr"  ` //[{"linkUrl": "https://twitter.com/W8sFgv45Jt16576","linkUrlType":"twitter"}, {"linkUrl": "https://t.me/layercraftofficialchat","linkUrlType":"telegram"}]
	Banners           []string `json:"banners"`
}

/*******************response dto end*******************************************/

/*******************swagger use response start*******************************************/
type TokenListPaginator struct {
	database.PaginatorBase
	Records []TokenListDto `json:"records"`
}

type TokenListResp struct {
	serializer.ResponseNotWithData
	Data TokenListPaginator `json:"data,omitempty"`
}

type TokenDetailResp struct {
	serializer.ResponseNotWithData
	Data TokenDetailDto `json:"data,omitempty"`
}

/*******************swagger use response end*******************************************/

/*******************service start*******************************************/

func TokenList(param TokenListReqDto, address string) (*database.Paginator, error) {
	db := database.DB.Table("memoo_tokens")
	if param.status != "" && len(param.status) != 0 {
		db = db.Where("status=?", param.status)
	}
	db = db.Order("id asc")
	var records []TokenListDto
	paginator := database.Paging(&database.Param{
		DB:      db,
		Page:    int(param.PageNumber),
		Limit:   param.PageSize,
		ShowSQL: true,
	}, &records)
	return paginator, nil
}

func TokenDetail(param TokenDetailReqDto, address string) (*TokenDetailDto, error) {
	var returnDto TokenDetailDto
	db := database.DB.Table("memoo_tokens").Where("address=? and ticker=?", address, param.Ticker).First(&returnDto)
	if db.Error != nil {
		return nil, db.Error
	}
	if returnDto.Ticker == "" {
		return nil, nil
	}
	var project model.MemooProject
	db = database.DB.Table("memoo_projects").Where("ticker=?", param.Ticker).First(&project)
	if db.Error != nil {
		return nil, db.Error
	}
	if project.ID == 0 {
		return nil, nil
	}
	returnDto.ProjectName = project.ProjectName
	returnDto.Website = project.Website
	returnDto.Description = project.Description
	returnDto.Twitter = project.Twitter

	var projectBanners []model.MemooProjectBanner
	db = database.DB.Table("memoo_project_banners").Where("project_id=?", project.ID).Find(&projectBanners)
	if db.Error != nil {
		return nil, db.Error
	}
	if project.ID == 0 {
		return nil, nil
	}
	banners := make([]string, 0)
	for _, item := range projectBanners {
		banners = append(banners, item.BannerUrl)
	}
	returnDto.Banners = banners

	var projectCreators []model.MemooProjectCreator
	db = database.DB.Table("memoo_project_creators").Where("project_id=?", project.ID).Find(&projectCreators)
	if db.Error != nil {
		return nil, db.Error
	}
	if project.ID == 0 {
		return nil, nil
	}

	var projectSocials []model.MemooProjectSocial
	db = database.DB.Table("memoo_project_socials").Where("project_id=?", project.ID).Find(&projectSocials)
	if db.Error != nil {
		return nil, db.Error
	}
	if project.ID == 0 {
		return nil, nil
	}
	otherLink := make([]string, 0)
	pinnedTwitterLink := make([]string, 0)
	for _, item := range projectSocials {
		if item.MajorCategories == constants.OTHER_LINK {
			otherLink = append(otherLink, item.SocialUrl)
		}
		if item.MajorCategories == constants.PINNED_TWITTER_LINK {
			pinnedTwitterLink = append(pinnedTwitterLink, item.SocialUrl)
		}
	}
	returnDto.OtherLink = otherLink
	returnDto.PinnedTwitterLink = pinnedTwitterLink
	return &returnDto, nil
}

func TokenNewOrEdit(param *TokenCreateOrUpdateDto, address string, tokenIconUrls []string, bannerUrls []string, otherLinks []LinksDto, pinnedTwitterLinks []LinksDto, status string) (*TokenCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	var memooToken model.MemooToken
	dbTx.Table("memoo_tokens").Where("ticker=?", param.Ticker).First(&memooToken)
	imageUrl := ""
	if tokenIconUrls != nil && len(tokenIconUrls) != 0 {
		imageUrl = tokenIconUrls[0]
	} else {
		imageUrl = memooToken.ImageUrl
	}
	updateStmt := fmt.Sprintf("insert into memoo_tokens(created_at,updated_at,ticker,token_name,contract_address,lp_contract_address,image_url,description,status,address) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) " +
		"on conflict (ticker) do  update set updated_at=?,token_name=?,image_url=?,description=?")
	tx := dbTx.Exec(updateStmt, time.Now(), time.Now(), param.Ticker, param.TokenName, param.ContractAddress, param.LPContractAddress, imageUrl, param.TokenDescription, status, address,
		time.Now(), param.TokenName, imageUrl, param.TokenDescription)
	if tx.Error != nil {
		dbTx.Rollback()
		return nil, tx.Error
	}
	paramProject := ProjectCreateOrUpdateDto{Ticker: param.Ticker, ProjectName: param.ProjectName,
		Description: param.Description, Twitter: param.Twitter}
	_, err := HandleProjectNewOrEdit(dbTx, &paramProject, address, bannerUrls, otherLinks, pinnedTwitterLinks)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	dbTx.Commit()
	return param, nil
}

func TokenConfirm(param *TokenCreateOrUpdateDto, address string, tokenIconUrls []string, bannerUrls []string, otherLinks []LinksDto, pinnedTwitterLinks []LinksDto, status string) (*TokenCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	err := removeToken(dbTx, param.Ticker)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	var memooToken model.MemooToken
	dbTx.Table("memoo_tokens").Where("ticker=?", param.Ticker).First(&memooToken)
	imageUrl := ""
	if tokenIconUrls != nil && len(tokenIconUrls) != 0 {
		imageUrl = tokenIconUrls[0]
	} else {
		imageUrl = memooToken.ImageUrl
	}
	updateStmt := fmt.Sprintf("insert into memoo_tokens(created_at,updated_at,ticker,token_name,contract_address,lp_contract_address,image_url,description,status) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?) " +
		"on conflict (ticker) do  update set updated_at=?,token_name=?,image_url=?,description=?")
	tx := dbTx.Exec(updateStmt, time.Now(), time.Now(), param.Ticker, param.TokenName, param.ContractAddress, param.LPContractAddress, imageUrl, param.TokenDescription, status,
		time.Now(), param.TokenName, imageUrl, param.TokenDescription)
	if tx.Error != nil {
		dbTx.Rollback()
		return nil, tx.Error
	}
	paramProject := ProjectCreateOrUpdateDto{Ticker: param.Ticker, ProjectName: param.ProjectName,
		Description: param.Description, Twitter: param.Twitter}
	_, err = HandleProjectNewOrEdit(dbTx, &paramProject, address, bannerUrls, otherLinks, pinnedTwitterLinks)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	dbTx.Commit()
	return param, nil
}

func removeToken(dbTx *gorm.DB, ticker string) error {
	tx := dbTx.Exec("delete from memoo_tokens where ticker=? and status=?", ticker, DRAFT)
	if tx.Error != nil {
		return tx.Error
	}
	var project model.MemooProject
	dbTx.Table("memoo_projects").Where("ticker=?", ticker).First(&project)
	if project.ID == 0 {
		return nil
	}
	err := removeProject(dbTx, project.ID)
	return err
}

func removeProject(dbTx *gorm.DB, projectId uint) error {
	tx := dbTx.Exec("delete from memoo_projects where id=? ", projectId)
	if tx.Error != nil {
		return tx.Error
	}
	tx = dbTx.Exec("delete from memoo_project_banners where project_id=?", projectId)
	if tx.Error != nil {
		return tx.Error
	}
	tx = dbTx.Exec("delete from memoo_project_creator where project_id=?", projectId)
	if tx.Error != nil {
		return tx.Error
	}
	tx = dbTx.Exec("delete from memoo_project_social where project_id=?", projectId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

/*******************service end*******************************************/
