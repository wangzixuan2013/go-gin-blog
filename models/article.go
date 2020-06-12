package models

import (
	//"github.com/chromedp/cdproto/database"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate (scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())

	return nil
}

func (article *Article) BeforeSave(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn",time.Now().Unix())

	return nil
}

func ExistArticleByID(id int) bool  {

	var article Article
	if err := db.Select("id").Where("id = ? ",id).First(&article).Error;err == nil{
		if article.ID > 0{
			return true
		}
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int)  {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
	//db.Model(&Article{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article)  {
	db.Where("id = ? ",id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

//
func EditArticle(id int, data map[string]interface{}) bool {
	if err := db.Model(&Article{}).Where("id = ?",id).Update(data).Error;err == nil{
		return true
	}
	return false
}

func AddArticle(data map[string]interface{}) bool {

	if err := db.Create(&Article{
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	}).Error;err == nil{
		return true
	}

	return false
}

func DeleteArticle(id int) bool  {
	if err:= db.Where("id = ?",id).Delete(&Article{}).Error;err == nil{
		return true
	}

	return false
}





