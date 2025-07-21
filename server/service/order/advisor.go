package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/errorx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	orderResp "github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"gorm.io/gorm"
)

type AdvisorService struct{}

func (advisorService *AdvisorService) Create(req *orderReq.AdvisorCreateReq) error {
	if req.Exp != nil {
		if *req.Exp < 0 || *req.Exp > 127 {
			return errorx.NewWithCode(errorx.IllegalAdvisorExpr)
		}
	}

	advisor := &order.Advisor{
		NickName:           req.NickName,
		Duty:               req.Duty,
		Intro:              req.Intro,
		AvatarUrl:          req.AvatarUrl,
		Exp:                req.Exp,
		SevenDayReturn:     req.SevenDayReturn,
		SevenDayReturnRate: req.SevenDayReturnRate,
		ActiveStatus:       order.Inactive,
	}
	return global.GVA_DB.Create(advisor).Error
}

func (advisorService *AdvisorService) Update(req *orderReq.AdvisorUpdateReq) error {
	if req.Exp != nil {
		if *req.Exp < 0 || *req.Exp > 127 {
			return errorx.NewWithCode(errorx.IllegalAdvisorExpr)
		}
	}

	var existAdvisor order.Advisor
	err := global.GVA_DB.First(&existAdvisor, req.Id).Error
	if err != nil {
		return err
	}

	existAdvisor.NickName = req.NickName
	existAdvisor.Duty = req.Duty
	existAdvisor.Intro = req.Intro
	existAdvisor.AvatarUrl = req.AvatarUrl
	existAdvisor.Exp = req.Exp
	existAdvisor.SevenDayReturn = req.SevenDayReturn
	existAdvisor.SevenDayReturnRate = req.SevenDayReturnRate

	return global.GVA_DB.Save(&existAdvisor).Error
}

func (advisorService *AdvisorService) UpdateStatus(req *orderReq.AdvisorStatusUpdateReq) error {
	var existAdvisor order.Advisor
	err := global.GVA_DB.First(&existAdvisor, req.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewWithCode(errorx.AdvisorNotFound)
	} else if err != nil {
		return err
	}

	existAdvisor.ActiveStatus = req.ActiveStatus
	return global.GVA_DB.Save(&existAdvisor).Error
}

func (advisorService *AdvisorService) Get(id int) (*orderResp.AdvisorDetail, error) {
	var advisor order.Advisor
	err := global.GVA_DB.First(&advisor, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewWithCode(errorx.AdvisorNotFound)
	} else if err != nil {
		return nil, err
	}

	resp := &orderResp.AdvisorDetail{
		Id:                 advisor.ID,
		NickName:           advisor.NickName,
		Duty:               advisor.Duty,
		Intro:              advisor.Intro,
		Exp:                advisor.Exp,
		AvatarUrl:          advisor.AvatarUrl,
		SevenDayReturn:     advisor.SevenDayReturn,
		SevenDayReturnRate: advisor.SevenDayReturnRate,
		ActiveStatus:       advisor.ActiveStatus,
	}
	return resp, nil
}

func (advisorService *AdvisorService) PageQuery(req *orderReq.AdvisorPageQueryReq) ([]*orderResp.AdvisorPageData, int64, error) {
	db := global.GVA_DB.Model(&order.Advisor{})
	if req.ActiveStatus != nil {
		db = db.Where("active_status = ?", *req.ActiveStatus)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil || total == 0 {
		return nil, 0, err
	}

	var advisors []*order.Advisor
	err = db.Scopes(req.Paginate()).Order("updated_at desc").Find(&advisors).Error

	list := make([]*orderResp.AdvisorPageData, len(advisors))
	for i, advisor := range advisors {
		list[i] = &orderResp.AdvisorPageData{
			AdvisorDetail: orderResp.AdvisorDetail{
				Id:                 advisor.ID,
				NickName:           advisor.NickName,
				Duty:               advisor.Duty,
				Intro:              advisor.Intro,
				Exp:                advisor.Exp,
				AvatarUrl:          advisor.AvatarUrl,
				SevenDayReturn:     advisor.SevenDayReturn,
				SevenDayReturnRate: advisor.SevenDayReturnRate,
				ActiveStatus:       advisor.ActiveStatus,
			},
		}
	}
	return list, total, nil
}

func (advisorService *AdvisorService) ListAdvisors(req *orderReq.AdvisorListQueryReq) ([]*orderResp.AdvisorOption, error) {
	var advisors []*order.Advisor
	db := global.GVA_DB
	if req.ActiveStatus != nil {
		db = db.Where("active_status = ?", *req.ActiveStatus)
	}

	if err := db.Find(&advisors).Error; err != nil {
		return nil, err
	}

	options := make([]*orderResp.AdvisorOption, len(advisors))
	for i, v := range advisors {
		options[i] = &orderResp.AdvisorOption{
			Id:       v.ID,
			NickName: v.NickName,
		}
	}

	return options, nil
}
