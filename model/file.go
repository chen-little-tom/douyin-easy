package model

import (
	"gorm.io/gorm"
	"time"
)

// File 文件模型
type File struct {
	Id       uint64    `json:"id" gorm:"column:id;primaryKey"`   // 文件id
	Tag      uint8     `json:"tag" gorm:"column:tag;"`           // 文件标志 0 本地文件 1 网络文件
	Suffix   string    `json:"suffix" gorm:"column:suffix"`      // 文件后缀
	Prefix   string    `json:"prefix" gorm:"column:prefix"`      // 文件前缀
	Path     string    `json:"path" gorm:"column:path"`          // 文件路径
	Size     uint64    `json:"size" gorm:"column:size"`          // 文件大小 字节
	Author   uint64    `json:"author" gorm:"column:author"`      // 文件作者
	CreateAt time.Time `json:"createAt" gorm:"column:create_at"` // 创建时间
}

type fileModel struct{}

var FileModel fileModel

// FileDB 获取文件DB
func (fm fileModel) FileDB() *gorm.DB {
	return DB.Model(&File{})
}

// Add 添加文件
// file 文件信息
func (fm fileModel) Add(file File) (uint64, error) {
	tx := fm.FileDB().Create(&file)
	return file.Id, tx.Error
}

// Del 删除文件
// 文件id
func (fm fileModel) Del(id uint64) error {
	tx := DB.Delete(&File{}, id)
	return tx.Error
}

// Detail 文件信息
// id 文件id
func (fm fileModel) Detail(id uint64) (File, error) {
	var file File
	tx := fm.FileDB().Where("id = ?", id).First(&file)
	return file, tx.Error
}

// ListByIds 根据id列表查询对应的信息
func (fm fileModel) ListByIds(ids []uint64) ([]File, error) {
	var files []File
	tx := DB.Model(&File{}).Where("id IN ?", ids).Find(&files)
	return files, tx.Error
}
