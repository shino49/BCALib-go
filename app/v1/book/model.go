package book

import (
	"gorm.io/gorm"
)

type BookMeta struct {
	gorm.Model
	Title    string
	Author   string
	WordNum  int64
	CoverPic string
	FilePath string
}
