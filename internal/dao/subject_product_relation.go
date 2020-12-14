package dao

import "gorm.io/gorm"

type SubjectProductRelationDao struct {
	*Dao
}

func NewSubjectProductRelation(db *gorm.DB) *SubjectProductRelationDao {
	return &SubjectProductRelationDao{
		NewDao(db),
	}
}
