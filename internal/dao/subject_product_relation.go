package dao

import "github.com/jinzhu/gorm"

type SubjectProductRelationDao struct {
	*Dao
}

func NewSubjectProductRelation(db *gorm.DB) *SubjectProductRelationDao {
	return &SubjectProductRelationDao{
		NewDao(db),
	}
}
