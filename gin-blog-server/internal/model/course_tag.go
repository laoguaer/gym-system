package model

import (
	"time"

	"gorm.io/gorm"
)

// CourseTag 课程标签关联模型
type CourseTag struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CourseID  int       `json:"course_id" gorm:"not null"`
	TagID     int       `json:"tag_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}


// GetCourseTagNames 根据课程ID获取标签名称列表
func GetCourseTagNames(db *gorm.DB, courseID int) ([]string, error) {
	var tagNames []string
	result := db.Table("tag").Select("tag.name").
		Joins("JOIN course_tag ON tag.id = course_tag.tag_id").
		Where("course_tag.course_id = ?", courseID).
		Pluck("name", &tagNames)
	return tagNames, result.Error
}

// GetCourseTagIDs 根据课程ID获取标签ID列表
func GetCourseTagIDs(db *gorm.DB, courseID int) ([]int, error) {
	var tagIDs []int
	result := db.Model(&CourseTag{}).
		Where("course_id = ?", courseID).
		Pluck("tag_id", &tagIDs)
	return tagIDs, result.Error
}

// AddCourseTag 添加课程标签关联
func AddCourseTag(db *gorm.DB, courseID int, tagID int) error {
	// 检查关联是否已存在
	var count int64
	db.Model(&CourseTag{}).
		Where("course_id = ? AND tag_id = ?", courseID, tagID).
		Count(&count)

	if count > 0 {
		return nil // 关联已存在，无需重复添加
	}

	// 添加新关联
	courseTag := CourseTag{
		CourseID: courseID,
		TagID:    tagID,
	}
	return db.Create(&courseTag).Error
}

// RemoveCourseTag 移除课程标签关联
func RemoveCourseTag(db *gorm.DB, courseID int, tagID int) error {
	return db.Where("course_id = ? AND tag_id = ?", courseID, tagID).Delete(&CourseTag{}).Error
}

// UpdateCourseTags 更新课程的标签关联（替换所有标签）
func UpdateCourseTags(db *gorm.DB, courseID int, tagIDs []int) error {
	// 开启事务
	return db.Transaction(func(tx *gorm.DB) error {
		// 删除所有现有关联
		if err := tx.Where("course_id = ?", courseID).Delete(&CourseTag{}).Error; err != nil {
			return err
		}

		// 添加新关联
		for _, tagID := range tagIDs {
			courseTag := CourseTag{
				CourseID: courseID,
				TagID:    tagID,
			}
			if err := tx.Create(&courseTag).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetCoursesByTagID 根据标签ID获取课程列表
func GetCoursesByTagID(db *gorm.DB, tagID int, page, size int) ([]Course, int64, error) {
	var courses []Course
	var total int64

	// 构建查询
	query := db.Table("course").
		Joins("JOIN course_tag ON course.id = course_tag.course_id").
		Where("course_tag.tag_id = ?", tagID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	if err := query.Limit(size).Offset((page - 1) * size).Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}
