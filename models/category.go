package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`

	ParentRelations []CategoryRelation `gorm:"foreignKey:ParentID"`
	ChildRelations  []CategoryRelation `gorm:"foreignKey:ChildID"`
}
