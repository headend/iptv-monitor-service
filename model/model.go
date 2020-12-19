package model

import "time"

func (Monitor) TableName() string {
	return "monitor"
}

type Monitor struct {
	Id  int64 `gorm:"column:agent_id;AUTO_INCREMENT;primary_key"`
	Status_signal int64 `gorm:"column:status_signal;default:1"`
	Status_video  int64 `gorm:"column:status_video ;default:1"`
	Status_audio  int64 `gorm:"column:status_audio ;default:1"`
	Signal_monitor  bool  `gorm:"column:signal_monitor ;default:true"`
	Video_monitor bool  `gorm:"column:video_monitor;default:false"`
	Audio_monitor bool  `gorm:"column:audio_monitor;default:false"`
	Is_enable bool  `gorm:"column:is_enable;default:true"`
	Agent_id  int64 `gorm:"column:agent_id ;default:not null"`
	Profile_id  int64 `gorm:"column:profile_id ;default:not null"`
	Status_id int64 `gorm:"column:status_id;default:null"`
	Date_update time.Time `gorm:"column:date_update;default:null"`
}