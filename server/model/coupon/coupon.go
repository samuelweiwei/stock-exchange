// 自动生成模板Coupon
package coupon

// Coupon coupon表 结构体  Coupon
type Coupon struct {
	Id           *uint64  `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键;size:20;"`                                         //主键
	Type         *Type    `json:"type" form:"type" gorm:"column:type;comment:优惠券类型：0-未知券 1-注册送券 2-实名送券 3-人工送券;"`                        //优惠券类型：0 未知券 1-注册送券 2-实例送券 3-手动送券
	Name         string   `json:"name" form:"name" gorm:"column:name;comment:优惠券名称;size:50;" validate:"required"`                       //优惠券名称
	Amount       *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:优惠金额;size:22;" validate:"required"`                  //优惠金额
	ActiveStatus *bool    `json:"activeStatus" form:"activeStatus" gorm:"column:active_status;comment:启用状态: 0-关闭;" validate:"required"` //启用状态: 0-关闭
	ValidDays    int      `json:"validDays" form:"validDays" gorm:"column:valid_days;comment:有效时长,单位天,只适用于注册券和实名券;size:19;"`            //有效时长,单位天,只适用于注册券和实名券
	ValidStart   int64    `json:"validStart" form:"validStart" gorm:"column:valid_start;comment:有效期开始时间，只适用于手动发券;size:19;"`             //有效期开始时间，只适用于手动发券
	ValidEnd     int64    `json:"validEnd" form:"validEnd" gorm:"column:valid_end;comment:有效期结束时间，只适用于手动发券;size:19;"`                   //有效期结束时间，只适用于手动发券
	Period       []int64  `json:"period" form:"period" gorm:"-"`                                                                        //券有效期时间段
	CreatedAt    int64    `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`                                    //创建时间
	UpdatedXAt   int64    `json:"UpdatedXAt" form:"UpdatedXAt" gorm:"column:updated_at;comment:更新时间;"`                                  //更新时间
}

// TableName coupon表 Coupon自定义表名 coupon
func (Coupon) TableName() string {
	return "coupon"
}
