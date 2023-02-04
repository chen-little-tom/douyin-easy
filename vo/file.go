package vo

import "time"

// FileVo 文件模型
type FileVo struct {
	Id       uint64    `json:"id"`       // 文件id
	Suffix   string    `json:"suffix"`   // 文件后缀
	Prefix   string    `json:"prefix"`   // 文件前缀
	FileUrl  string    `json:"fileUrl"`  // 文件路径
	Size     uint64    `json:"size"`     // 文件大小 字节
	Author   uint64    `json:"author"`   // 文件作者
	CreateAt time.Time `json:"createAt"` // 创建时间
}
