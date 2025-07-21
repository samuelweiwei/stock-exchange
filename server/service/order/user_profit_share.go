package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type UserProfitShareService struct{}

func (s *UserProfitShareService) QueryMyProfitShareRecord(req *request.MyUserProfitShareRecordPageQueryReq) ([]response.MyUserProfitShareRecord, int64, error) {
	db := global.GVA_DB.Model(&order.UserProfitShare{}).Where("to_user_id = ?", req.UserId)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	var shares []order.UserProfitShare
	if err := db.Scopes(req.Paginate()).Order("created_at desc").Find(&shares).Error; err != nil {
		return nil, 0, err
	}

	fromUserIds := make([]uint, 0)
	for _, share := range shares {
		fromUserIds = append(fromUserIds, share.FromUserId)
	}
	fromUserMap := queryFrontUserMap(fromUserIds)

	records := make([]response.MyUserProfitShareRecord, len(shares))
	for i, v := range shares {
		records[i] = response.MyUserProfitShareRecord{
			Amount: v.Amount.InexactFloat64(),
			Date:   v.CreatedAt.UnixMilli(),
		}
		if u, ok := fromUserMap[v.FromUserId]; ok {
			records[i].FromUserName = u.NickName
		}
	}

	return records, total, nil
}

// 查询前台用户信息
func queryFrontUserMap(userIds []uint) map[uint]user.FrontendUsers {
	var fromUsers []user.FrontendUsers
	fromUserMap := make(map[uint]user.FrontendUsers)
	if err := global.GVA_DB.Find(&fromUsers, userIds).Error; err == nil {
		for _, u := range fromUsers {
			fromUserMap[u.ID] = u
		}
	}
	return fromUserMap
}
