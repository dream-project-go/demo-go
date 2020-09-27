package sample

import (
	"gormdemo/db"
	"log"
	"time"
)

// CREATE TABLE `sample` (
// 	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
// 	`title` varchar(100) DEFAULT '' COMMENT '标题',
// 	`mtime` int(11) unsigned NOT NULL COMMENT '最后修改时间',
// 	PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试样例表';

type Sample struct {
	Id    int64  `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"` // 标题
	Mtime int64  `gorm:"column:mtime" json:"mtime"` // 最后修改时间
}

func NewSample() *Sample {
	return &Sample{}
}

func (s *Sample) Model() *db.Model {
	return &db.Model{
		DB: "test",
	}
}

func (s *Sample) AddInfo(d *Sample) error {
	s.Mtime = time.Now().Unix()
	err := s.Model().Master().Table("sample").Create(&s).Error
	if s.Model().CheckError(err) {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Sample) GetInfo(Id int) (Sample, error) {
	var info Sample
	err := s.Model().Master().Table("sample").Where("id = ?", Id).First(&info).Error
	if s.Model().CheckError(err) {
		log.Println(err)
	}
	return info
}
