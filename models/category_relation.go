package models

type CategoryRelation struct {
	ParentID uint `gorm:"primaryKey"`
	ChildID  uint `gorm:"primaryKey"`

	Parent Category `gorm:"foreignKey:ParentID"`
	Child  Category `gorm:"foreignKey:ChildID"`
}
