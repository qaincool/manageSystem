package model

type Teacher struct {
	// 外键
	UserId string `gorm:"column:user_id"`

	TeacherID      int    `gorm:"column:teacher_id;primaryKey;autoIncrement"`
	TeacherPicture string `gorm:"column:teacher_picture"`
	TeacherDetail  string `gorm:"column:teacher_detail"`

	Videos []*Video `gorm:"foreignKey:create_user"`
}
