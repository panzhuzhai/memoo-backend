package database

import (
	"gorm.io/gorm"
	"math"
)

// Param 分页参数
type Param struct {
	DB       *gorm.DB
	Page     int
	RowStart int64
	Limit    int
	OrderBy  []string
	ShowSQL  bool
}

// Paginator 分页返回
type PaginatorBase struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int   `json:"total_page"`
	RowStart    int64 `json:"rowStart"`
	// 使用 oneOf 注解标签指定可能的类型
	// swagger:allOf
	// swagger:discriminator Type
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PrevPage int `json:"prev_page"`
	NextPage int `json:"next_page"`
}

// Paginator 分页返回
type Paginator struct {
	PaginatorBase
	Records interface{} `json:"records"`
}

// Paginator 分页返回
type PaginatorById struct {
	TotalRecord int64 `json:"totalRecord"`
	TotalPage   int   `json:"totalPage"`
	RowStart    int64 `json:"rowStart"`
	// 使用 x-anyOf 扩展指定可能的类型
	// 示例：Type1 或 Type2
	// anyOf: [dto.AirdropDto]
	Records interface{} `json:"records"`
	Limit   int         `json:"limit"`
}

type IDType struct {
	ID uint64
}

// Paging 分页
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int64
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}

// Paging 分页
func PagingById(p *Param, result interface{}, outTable string) *PaginatorById {
	db := p.DB
	if p.ShowSQL {
		db = db.Debug()
	}
	if p.RowStart < 0 {
		p.RowStart = 0
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	done := make(chan bool, 1)
	var paginator PaginatorById
	var count int64
	go countRecords(db, result, done, &count)
	if outTable != "" {
		db.Where(outTable+".id>=?", p.RowStart).Order(outTable + ".id asc").Limit(p.Limit).Scan(result)
		<-done
	} else {
		db.Where("id>=?", p.RowStart).Order("id asc").Limit(p.Limit).Scan(result)
		<-done
	}
	paginator.TotalRecord = count
	paginator.Records = result

	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))
	return &paginator
}
