package models

import (
	"fmt"
	"ielts/dao"

	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model
	UserID  int    `json:"userid" form:"userid"`
	Listen  string `json:"listen" form:"listen" `
	Speak   string `json:"speak"  form:"speak"`
	Read    string `json:"read"   form:"read"`
	Write   string `json:"write"  form:"write"`
	Total   string `json:"total"  form:"total"`
	Average string `json:"average"form:"average"`
	Note    string `json:"note"   form:"note"`
}

func CreatRecord(record *Record) (err error) {

	err = dao.DB.Debug().Create(&record).Error
	if err != nil {

		return err
	}
	return
}
func GetRecords(id int) (records []*Record, err error) {
	err = dao.DB.Debug().Where("user_id = ?", id).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteARecord(uid int, reid int) (err error) {
	var re Record
	err = dao.DB.Debug().Where("user_id = ? AND id = ? ", uid, reid).First(&re).Delete(&re).Error

	fmt.Println("要删除的数据", re)
	return
}
