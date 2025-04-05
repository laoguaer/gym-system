package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	Model
	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`

	Articles []*Article `gorm:"many2many:article_tag;" json:"articles,omitempty"`
}

type TagVO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
}

func GetTagList(db *gorm.DB, page, size int, keyword string) (list []TagVO, total int64, err error) {
	db = db.Table("tag t").
		Joins("LEFT JOIN article_tag at ON t.id = at.tag_id").
		Select("t.id", "t.name", "COUNT(at.article_id) AS article_count", "t.created_at", "t.updated_at")

	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	result := db.
		Group("t.id").Order("t.updated_at DESC").
		Count(&total).
		Scopes(Paginate(page, size)).
		Find(&list)

	return list, total, result.Error
}

func GetTagOption(db *gorm.DB) ([]OptionVO, error) {
	list := make([]OptionVO, 0)
	result := db.Model(&Tag{}).Select("id", "name").Find(&list)
	return list, result.Error
}

// 根据 [文章id] 获取 [标签名称列表]
func GetTagNamesByArticleId(db *gorm.DB, id int) ([]string, error) {
	list := make([]string, 0)
	result := db.Table("tag").
		Joins("LEFT JOIN article_tag ON tag.id = article_tag.tag_id").
		Where("article_id", id).
		Pluck("name", &list)
	return list, result.Error
}

func SaveOrUpdateTag(db *gorm.DB, id int, name string) (*Tag, error) {
	tag := Tag{
		Model: Model{ID: id},
		Name:  name,
	}

	var result *gorm.DB
	if id > 0 {
		result = db.Updates(tag)
	} else {
		result = db.Create(&tag)
	}

	return &tag, result.Error
}

/*
CREATE TABLE `tag` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `component` varchar(255) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC
*/

type TagModel struct {
	Id int `gorm:"column:id"`
	createdAt time.Time
	updatedAt time.Time
	Name string `gorm:"column:name"`
	Component string `gorm:"column:component"`
}

func GetTagsByComponent(db *gorm.DB, component string) ([]TagModel, error) {
	list := make([]TagModel, 0)
	result := db.Table("tag").Where("component = ?", component).Find(&list)
	return list, result.Error	
}
