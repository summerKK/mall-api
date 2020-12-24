package dao

import "gorm.io/gorm"

type SubjectProductRelation struct {
	*Dao
}

func NewSubjectProductRelation(db *gorm.DB) *SubjectProductRelation {
	return &SubjectProductRelation{
		NewDao(db),
	}
}
