package service

import (
	"douyin-easy/model"
	"douyin-easy/utils"
	"douyin-easy/vo"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type fileService struct{}

var FileService fileService

// Add 添加文件
func (fs fileService) Add(c *gin.Context) (uint64, error) {
	// 获取到文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("文件获取失败，%s\n", err)
		return 0, errors.New("文件获取失败")
	}
	// 处理文件，储存文件
	filePath, err := fs.saveFile(file, c)
	if err != nil {
		return 0, errors.New("上传文件失败")
	}
	// 处理文件信息，存入数据库
	var f model.File
	//loginUser := utils.GetLoginUser(c)
	//f.Author = loginUser.UId
	f.Path = filePath
	f.Size = uint64(file.Size)
	f.Tag = 0 // 这里存储为本地
	index := strings.LastIndex(file.Filename, ".")
	f.Prefix = file.Filename[:index]
	f.Suffix = file.Filename[index+1:]
	f.CreateAt = time.Now()

	fileId, err := model.FileModel.Add(f)
	if err != nil {
		log.Printf("文件存储失败,err->%s\n", err)
		return 0, errors.New("文件存储失败")
	}
	return fileId, nil
}

// Del 删除文件
// id 文件id
func (fs fileService) Del(id uint64) error {
	// 获取文件信息
	file, err := model.FileModel.Detail(id)
	if err != nil {
		log.Printf("文件获取失败,err->%s\n", err)
		return errors.New("获取文件信息失败")
	}
	err = os.Remove(utils.StaticRoot + file.Path)
	if err != nil {
		log.Printf("删除文件失败，err->%s\n", err)
		return errors.New("删除文件失败")
	}
	err = model.FileModel.Del(id)
	if err != nil {
		log.Printf("删除文件信息失败，err->%s\n", err)
		return errors.New("删除文件失败")
	}
	return nil
}

// Detail 文件详情
// id 文件id
func (fs fileService) Detail(id uint64) (vo.FileVo, error) {
	// 获取文件源信息
	var fileVo vo.FileVo
	file, err := model.FileModel.Detail(id)
	if err != nil {
		log.Printf("文件获取失败,err->%s\n", err)
		return fileVo, errors.New("获取文件信息失败")
	}
	// 处理文件信息
	fileVo.Id = file.Id
	fileVo.Suffix = file.Suffix
	fileVo.Prefix = file.Prefix
	fileVo.Author = file.Author
	fileVo.Size = file.Size
	fileVo.CreateAt = file.CreateAt
	fileVo.FileUrl = fs.getFileUrl(file.Path)
	return fileVo, nil
}

// saveFile 保存文件
func (fs fileService) saveFile(file *multipart.FileHeader, c *gin.Context) (string, error) {
	path := fmt.Sprintf(utils.UploadFileRoot, time.Now().Year(), time.Now().Month(), time.Now().Day())

	realPath := utils.StaticRoot + path
	if _, err := os.Stat(realPath); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(realPath, os.ModePerm)
			if err != nil {
				log.Printf("创建文件夹失败，%s\n", err)
				return "", errors.New("创建文件夹失败")
			}
		}
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
	realPath = realPath + fileName
	filePath := path + fileName
	err := c.SaveUploadedFile(file, realPath)
	if err != nil {
		log.Printf("文件存储失败，%s\n", err)
		return "", errors.New("文件存储失败")
	}
	return filePath, nil
}

// getFileUrl 获取文件的网络地址
func (fs fileService) getFileUrl(filePath string) string {
	return utils.DeployStaticRoot + filePath
}

// ListByIds 根据指定的id 批量查询数据
func (fs fileService) ListByIds(ids []uint64) ([]vo.FileVo, error) {
	var ret []vo.FileVo
	files, err := model.FileModel.ListByIds(ids)
	if err != nil {
		log.Printf("批量查询文件信息失败,err->%s\n", err)
		return ret, errors.New("批量查询文件信息失败")
	}
	for _, f := range files {
		var item vo.FileVo
		item.Id = f.Id
		item.Size = f.Size
		item.Suffix = f.Suffix
		item.Prefix = f.Prefix
		item.Author = f.Author
		item.FileUrl = fs.getFileUrl(f.Path)
		ret = append(ret, item)
	}
	return ret, nil
}

// ListByIdsMap 根据指定的id 批量查询数据 封装为map
func (fs fileService) ListByIdsMap(ids []uint64) (map[uint64]vo.FileVo, error) {
	ret := make(map[uint64]vo.FileVo, len(ids))
	fileVos, err := fs.ListByIds(ids)
	if err != nil {
		return ret, err
	}
	for _, f := range fileVos {
		ret[f.Id] = f
	}
	return ret, nil
}

// ReadAndSave 读取文件并存储信息到数据库
// Return 返回数据库文件id
func (fs fileService) ReadAndSave(filePath string, userId uint64) (uint64, error) {
	localPath := utils.StaticRoot + filePath
	fileInfo, err := os.Stat(localPath)
	if err != nil {
		log.Printf("读取文件信息失败,err->%s\n", err)
		return 0, errors.New("读取文件信息失败")
	}
	var file model.File
	file.Size = uint64(fileInfo.Size())
	file.Tag = 0
	file.Path = filePath
	file.Author = userId
	strings.ReplaceAll(filePath, "\\", "/")
	index := strings.LastIndex(filePath, "/")
	filename := filePath[index+1:]

	index = strings.LastIndex(filename, ".")
	file.Prefix = filename[:index]
	file.Suffix = filename[index+1:]
	file.CreateAt = time.Now()

	// 存储文件
	return model.FileModel.Add(file)
}
