package service

import (
	"fmt"
	"memoo-backend/dto"
	"memoo-backend/middleware/database"
	"memoo-backend/model"
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

type TokenCreateOrUpdateDto struct {
	TokenName         string `json:"tokenName"`
	ContractAddress   string `json:"contractAddress"`
	LPContractAddress string `json:"lpContractAddress"`
	TokenDescription  string `json:"tokenDescription" `
	ProjectCreateOrUpdateDto
}

/*******************request dto end*******************************************/

/*******************response dto start*******************************************/
type TokenListRespDto struct {
	TokenImageUrl string `json:"tokenImageUrl" `
	TokenName     string `json:"tokenName" `
	TotalRaised   string `json:"totalRaised"`
	LaunchDate    int64  `json:"launchDate" `
	MeMooScore    string `json:"meMooScore" `
	Status        string `json:"status" `
}

/*******************response dto end*******************************************/

/*******************service start*******************************************/

func TokenList(param TokenListReqDto, address string) (database.Paginator, error) {
	return database.Paginator{}, nil
}

func TokenNewOrEdit(param *TokenCreateOrUpdateDto, address string, tokenIconUrls []string, bannerUrls []string) (*TokenCreateOrUpdateDto, error) {
	dbTx := database.DB.Begin()
	var memooToken model.MemooToken
	dbTx.Table("memoo_tokens").Where("ticker=?", param.Ticker).First(&memooToken)
	imageUrl := ""
	if tokenIconUrls != nil && len(tokenIconUrls) != 0 {
		imageUrl = tokenIconUrls[0]
	} else {
		imageUrl = memooToken.ImageUrl
	}
	status := DRAFT
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
		Description: param.Description, Twitter: param.Twitter, OtherLinks: param.OtherLinks}
	_, err := HandleProjectNewOrEdit(dbTx, &paramProject, address, bannerUrls)
	if err != nil {
		dbTx.Rollback()
		return nil, err
	}
	dbTx.Commit()
	return param, nil
}

/*******************service end*******************************************/
