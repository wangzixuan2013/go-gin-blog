package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag * Tag) BeforeUpdate (scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeSave(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn",time.Now().Unix())

	return nil
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface {}) (count int){
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?",name).First(&tag)
	if tag.ID > 0{
		return true
	}
	return false
}

func AddTag(name string,state int,createdBy string) bool {
	if err := db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createdBy,
	}).Error; err != nil {
		return false
	}
	return true
}

func ExistTagByID(id int)  bool {

	var tag Tag
	db.Select("id").Where("id = ?",id).First(&tag)
	if tag.ID > 0{
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	if err := db.Where("id = ? ",id).Delete(&Tag{}).Error;err !=nil{
		return false
	}
	return true
}

func EditTag(id int,data interface{}) bool {

	//db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	if err := db.Model(&Tag{}).Where("id = ? ",id).Updates(data).Error;err != nil{
		return false
	}
	return true
}

func CleanAllTag() bool {
	if err := db.Unscoped().Where("deleted_on > ? ",0).Delete(&Tag{}).Error;err != nil{
		return false
	}
	return true
}
