package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type SystemOrderMsgService struct {
}

func (s *SystemOrderMsgService) GetUnReadCount() (int64, error) {
	var count int64
	err := global.GVA_DB.Model(&order.SystemOrderMsg{}).
		Where("read_status", order.MsgReadStatusUnRead).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *SystemOrderMsgService) SetMessagesRead(msgIds []int64) error {
	return global.GVA_DB.Model(&order.SystemOrderMsg{}).
		Where("id in (?)", msgIds).
		Updates(map[string]interface{}{"read_status": order.MsgReadStatusRead}).Error
}

func (s *SystemOrderMsgService) PageQuery(req *request.SystemOrderMsgPageQueryReq) ([]*response.SystemOrderMsgPageData, int64, error) {
	db := global.GVA_DB.Model(&order.SystemOrderMsg{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	} else if total == 0 {
		return make([]*response.SystemOrderMsgPageData, 0), 0, nil
	}

	var msgList []*order.SystemOrderMsg
	err = req.Paginate()(db).Order("created_at desc").Find(&msgList).Error
	if err != nil {
		return nil, 0, err
	}

	userIds := make([]uint, len(msgList))
	for i, msg := range msgList {
		userIds[i] = msg.UserId
	}
	var users []user.FrontendUsers
	global.GVA_DB.Model(&user.FrontendUsers{}).Where("id in (?)", userIds).Find(&users)
	userMap := make(map[uint]user.FrontendUsers, len(msgList))
	for _, v := range users {
		userMap[v.ID] = v
	}

	resultList := make([]*response.SystemOrderMsgPageData, len(msgList))
	for i, v := range msgList {
		resultList[i] = &response.SystemOrderMsgPageData{
			Id:         v.ID,
			Type:       v.Type,
			OrderId:    v.OrderId,
			UserId:     v.UserId,
			ReadStatus: v.ReadStatus,
		}
		if u, ok := userMap[v.UserId]; ok {
			resultList[i].Phone = u.Phone
		}
	}

	return resultList, total, nil
}
