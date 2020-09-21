package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagId      int    `json:"tag_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (article []Article) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

//新增文章标签
func AddArticle(title string, tagId int, desc string, content string, state int, createdBy string) bool {
	db.Create(&Article{
		TagId:     tagId,
		Title:     title,
		Desc:      desc,
		Content:   content,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

//修改
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}
