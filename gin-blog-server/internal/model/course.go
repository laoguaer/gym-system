package model

import (
	"time"

	"gorm.io/gorm"
)

// Course 课程信息模型
type Course struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Tags        string    `json:"tags" gorm:"type:varchar(255);not null"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	CoachID     int       `json:"coach_id" gorm:"not null"`
	MaxCapacity int       `json:"max_capacity" gorm:"not null;default:30"`
	IsSingle    int       `json:"is_single" gorm:"type:bit(1);not null"`

	// // 关联教练信息
	// Coach *Coach `json:"coach,omitempty" gorm:"foreignKey:CoachID"`
}

// CourseQuery 课程查询参数
type CourseQuery struct {
	Page     int    `form:"page_num" binding:"required,min=1"`
	Size     int    `form:"page_size" binding:"required,min=1,max=50"`
	Title    string `form:"title"`
	CoachID  int    `form:"coach_id"`
	TagID    int    `form:"tag_id"`
	IsSingle *bool  `form:"is_single"`
}

// CourseVO 课程视图对象
type CourseVO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CoachID     int       `json:"coach_id"`
	MaxCapacity int       `json:"max_capacity"`
	IsSingle    int       `json:"is_single"`
	Coach       *Coach    `json:"coach,omitempty"`
	TagList     []string  `json:"tag_list,omitempty"`
}

// GetCourseList 获取课程列表
func GetCourseList(db *gorm.DB, page, size int, title string, coachID int, tagID int, isSingle *bool) (list []CourseVO, total int64, err error) {
	// 构建基础查询
	query := db.Model(&Course{})

	// 应用过滤条件
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if coachID > 0 {
		query = query.Where("coach_id = ?", coachID)
	}

	if isSingle != nil {
		query = query.Where("is_single = ?", *isSingle)
	}

	// 如果有标签ID过滤，需要联表查询
	if tagID > 0 {
		query = query.Joins("JOIN course_tag ON course.id = course_tag.course_id").Where("course_tag.tag_id = ?", tagID)
	}

	// 获取总数
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	var courses []Course
	err = query.Limit(size).Offset((page - 1) * size).Find(&courses).Error
	if err != nil {
		return nil, 0, err
	}

	// 转换为VO对象
	list = make([]CourseVO, len(courses))
	for i, course := range courses {
		list[i] = CourseVO{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Tags:        course.Tags,
			StartTime:   course.StartTime,
			EndTime:     course.EndTime,
			CoachID:     course.CoachID,
			MaxCapacity: course.MaxCapacity,
			IsSingle:    course.IsSingle, // 将uint8转换为bool
		}

		// 如果需要，可以在这里加载教练信息和标签列表
	}

	return list, total, nil
}

// GetCourseById 根据ID获取课程信息
func GetCourseById(db *gorm.DB, id int) (*CourseVO, error) {
	var course Course
	result := db.Model(&course).Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}

	// 转换为VO对象
	courseVO := CourseVO{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		Tags:        course.Tags,
		StartTime:   course.StartTime,
		EndTime:     course.EndTime,
		CoachID:     course.CoachID,
		MaxCapacity: course.MaxCapacity,
		IsSingle:    course.IsSingle, // 将uint8转换为bool
	}

	// 加载教练信息
	if course.CoachID > 0 {
		coach, err := GetCoachById(db, course.CoachID)
		if err == nil {
			courseVO.Coach = coach
		}
	}

	// 加载标签列表
	tagList, err := GetCourseTagNames(db, id)
	if err == nil {
		courseVO.TagList = tagList
	}

	return &courseVO, nil
}

// CreateCourse 创建课程
func CreateCourse(db *gorm.DB, course *Course) (*Course, error) {
	result := db.Create(course)
	return course, result.Error
}

// UpdateCourse 更新课程信息
func UpdateCourse(db *gorm.DB, course *Course) error {
	result := db.Model(&Course{}).Where("id = ?", course.ID).Updates(course)
	return result.Error
}

// DeleteCourse 删除课程
func DeleteCourse(db *gorm.DB, id int) error {
	// 先删除关联的标签
	if err := db.Where("course_id = ?", id).Delete(&CourseTag{}).Error; err != nil {
		return err
	}

	// 再删除课程
	result := db.Delete(&Course{}, id)
	return result.Error
}

// GetCoursesByCoachId 根据教练ID获取课程列表
func GetCoursesByCoachId(db *gorm.DB, coachID int) ([]Course, error) {
	var courses []Course
	result := db.Where("coach_id = ?", coachID).Find(&courses)
	return courses, result.Error
}

// GetCourseList 获取课程列表
func GetCourseListForAi(db *gorm.DB, title string, coachIdList []int, isSingle *bool) (list []CourseVO, err error) {
	// 构建基础查询
	query := db.Model(&Course{})

	var courses []Course

	// 应用过滤条件
	if title != "" {
		query = query.Where("title LIKE?", "%"+title+"%")
	}
	if len(coachIdList) > 0 {
		query = query.Where("coach_id IN ?", coachIdList)
	}
	if isSingle != nil {
		query = query.Where("is_single =?", *isSingle)
	}
	// 获取数据
	err = query.Find(&courses).Error
	if err != nil {
		return nil, err
	}

	// 转换为VO对象
	list = make([]CourseVO, len(courses))
	for i, course := range courses {
		list[i] = CourseVO{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Tags:        course.Tags,
			StartTime:   course.StartTime,
			EndTime:     course.EndTime,
			CoachID:     course.CoachID,
			MaxCapacity: course.MaxCapacity,
			IsSingle:    course.IsSingle, // 将uint8转换为bool
		}

		// 如果需要，可以在这里加载教练信息和标签列表
	}

	return list, nil
}
