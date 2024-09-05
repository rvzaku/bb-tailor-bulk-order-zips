package models

type Gender struct {
	Name         string        `gorm:"type:varchar(255);primaryKey"`
	UserProfiles []UserProfile `gorm:"foreignKey:Gender"`
}
