package model

import "time"

func (Monitor) TableName() string {
	return "monitor"
}

type Monitor struct {
	id  int64 `gorm:"column:id ;default:null"`
	status_signal int64 `gorm:"column:status_signal;default:1"`
	status_video  int64 `gorm:"column:status_video ;default:1"`
	status_audio  int64 `gorm:"column:status_audio ;default:1"`
	signal_monitor  bool  `gorm:"column:signal_monitor ;default:true"`
	video_monitor bool  `gorm:"column:video_monitor;default:false"`
	audio_monitor bool  `gorm:"column:audio_monitor;default:false"`
	is_enable bool  `gorm:"column:is_enable;default:true"`
	agent_id  int64 `gorm:"column:agent_id ;default:not null"`
	profile_id  int64 `gorm:"column:profile_id ;default:not null"`
	status_id int64 `gorm:"column:status_id;default:null"`
	date_update time.Time `gorm:"column:date_update;default:null"`
}

