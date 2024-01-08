package marshaller

import (
	"log"
	"testing"
	"time"
)

type SearchByTagsReq struct {
	Device           *Device           `json:"device,omitempty" es:"term,device"`
	VipLeftDuration  *VIPLeftDuration  `json:"vip_left_duration,omitempty" es:"range,vip_left_duration"`
	RegTime          *RegTime          `json:"register_time,omitempty" es:"range,reg_time"`
	RegRangeDuration *RegRangeDuration `json:"register_range_time,omitempty" es:"range,reg_range_duration"`
	Active           *Active           `json:"active,omitempty" es:"range,active"`
	RegSource        *RegSource        `json:"reg_source,omitempty" es:"terms,reg_source"`
}

type Device struct {
	DeviceType string `json:"device_type" es:"value,device_type"` // 设备类型  ios | android
}

// VIPLeftDuration vip剩余时长
type VIPLeftDuration struct {
	LeftDurationMin int64 `json:"left_duration_min" es:"lte,duration_min"` // 剩余时长 时间段(小值)
	LeftDurationMax int64 `json:"left_duration_max" es:"gte,duration_max"` // 剩余时长时间段 (大值)
}

// RegSource 注册渠道
type RegSource struct {
	RegSource string `json:"reg_source" es:"values,reg_source"` // -1不限制 多个渠道,号间隔
}

// RegTime 注册时间筛选,自定义时间段
type RegTime struct {
	Start int64 `json:"start" es:"lte,start"` // 注册时间段开始日期
	End   int64 `json:"end" es:"gte,end"`     // 注册时间段截止日期
}

// RegRangeDuration 自定义时间间隔
type RegRangeDuration struct {
	RegRangeMin int64 `json:"reg_range_min"` // 注册距今时间段(小值)
	RegRangeMax int64 `json:"reg_range_max"` // 注册距今时间段(大值)
}

// Active 活跃时间
type Active struct {
	LastActiveMin int64 `json:"last_active_min " es:"lte,last_active_min"` // 最后活跃时间段 (大值)
	LastActiveMax int64 `json:"last_active_max" es:"gte,last_active_max"`  // 最后活跃时间段 (小值)
}

func Test(t *testing.T) {
	now := time.Now()
	req := &SearchByTagsReq{
		Device: &Device{
			DeviceType: "ios",
		},
		VipLeftDuration: &VIPLeftDuration{
			LeftDurationMin: now.Add(-time.Hour * 24).Unix(),
			LeftDurationMax: now.Add(time.Hour * 48).Unix(),
		},
		RegTime: &RegTime{
			Start: now.Add(-time.Hour * 24).Unix(),
			End:   now.Add(time.Hour * 48).Unix(),
		},
		RegSource: &RegSource{
			RegSource: "1,2,3",
		},
	}

	sql, err := Marshal(req)
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println(sql)
}
