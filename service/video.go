package service

import (
	"douyin-easy/model"
	"douyin-easy/utils"
	"douyin-easy/utils/user_token"
	"douyin-easy/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type videoService struct{}

var VideoService videoService

func (vs videoService) Add(c *gin.Context) error {
	// 获取参数
	title := c.PostForm("title")
	token := c.PostForm("token")

	// 获取到文件
	file, err := c.FormFile("data")
	if err != nil {
		log.Printf("文件获取失败，%s\n", err)
		return errors.New("文件获取失败")
	}
	// 处理文件，储存文件
	filePath, err := FileService.saveFile(file, c)
	if err != nil {
		return errors.New("上传文件失败")
	}
	// 处理文件信息，存入数据库
	var f model.File
	user, err := UserService.GetLoginUser(token)
	if err != nil {
		return err
	}
	f.Author = user.Id
	f.Path = filePath
	f.Size = uint64(file.Size)
	f.Tag = 0 // 这里存储为本地
	index := strings.LastIndex(file.Filename, ".")
	f.Prefix = file.Filename[:index]
	f.Suffix = file.Filename[index+1:]
	f.CreateAt = time.Now()

	playId, err := model.FileModel.Add(f)
	if err != nil {
		log.Printf("文件存储失败,err->%s\n", err)
		return errors.New("文件存储失败")
	}
	var video model.Video
	video.CreateAt = time.Now()
	video.UpdateAt = time.Now()
	video.AuthorId = user.Id
	video.CommentCount = 0
	video.PlayId = playId
	video.FavoriteCount = 0
	video.Title = title

	// 生成封面 文件后缀 .png
	index = strings.LastIndex(filePath, ".")
	coverPath := filePath[:index] + ".png"
	err = utils.GetVideoCover(utils.StaticRoot+filePath, utils.StaticRoot+coverPath, 1)
	if err != nil {
		return err
	}
	// 读取封面并存储
	coverId, err := FileService.ReadAndSave(coverPath, user.Id)
	if err != nil {
		return err
	}
	video.CoverId = coverId
	err = model.VideoModel.Add(video)
	if err != nil {
		log.Printf("视频信息存储失败,err->%s\n", err)
		return errors.New("上传视频失败")
	}
	return nil
}

// Feed 获取视频流
func (vs videoService) Feed(lastTime time.Time, token string) ([]vo.VideoVo, time.Time, error) {
	var ret []vo.VideoVo
	nextTime := time.Now()
	videos, err := model.VideoModel.ListByTime(lastTime)
	if err != nil {
		log.Printf("获取视频流失败,err->%s\n", err)
		return ret, nextTime, errors.New("获取视频流失败")
	}
	// 设置下一次的lastTime
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreateAt
	}

	userIds := make([]uint64, 0)
	fileIds := make([]uint64, 0)

	for _, v := range videos {
		userIds = append(userIds, v.AuthorId)
		fileIds = append(fileIds, v.CoverId)
		fileIds = append(fileIds, v.PlayId)
	}

	fileMap, err := FileService.ListByIdsMap(fileIds)
	if err != nil {
		return ret, nextTime, err
	}
	loginUser, err := UserService.GetLoginUser(token)
	if err != nil {
		log.Printf("用户信息查询失败,err->%s\n", err)
	}
	userMap, err := UserService.DetailByIdsMap(userIds, loginUser)
	if err != nil {
		return ret, nextTime, err
	}
	for _, v := range videos {
		var video vo.VideoVo
		video.Title = v.Title
		video.Id = v.Id
		video.CommentCount = v.CommentCount
		video.FavoriteCount = v.FavoriteCount
		video.IsFavorite = false
		if fileInfo, ok := fileMap[v.CoverId]; ok {
			video.CoverUrl = fileInfo.FileUrl
		}
		if fileInfo, ok := fileMap[v.PlayId]; ok {
			video.PlayUrl = fileInfo.FileUrl
		}
		if userInfo, ok := userMap[v.AuthorId]; ok {
			video.Author = userInfo
		}
		ret = append(ret, video)
	}

	return ret, nextTime, nil
}

// List 查询用户所有投稿过的视频
func (vs videoService) List(token string) ([]vo.VideoVo, error) {
	ret := make([]vo.VideoVo, 0)
	userid, err := user_token.GetUserIdByToken(token)
	if err != nil {
		log.Printf("获取当前登录用户信息失败,err->%s\n", err)
		return ret, errors.New("获取当前登录用户失败")
	}
	videos, err := model.VideoModel.ListByUser(userid)
	if err != nil {
		log.Printf("查询用户投稿的视频失败,err->%s\n", err)
		return ret, errors.New("查询用户投稿的视频失败")
	}

	fileIds := make([]uint64, 0)
	for _, v := range videos {
		fileIds = append(fileIds, v.CoverId, v.PlayId)
	}
	fileMap, err := FileService.ListByIdsMap(fileIds)
	if err != nil {
		return ret, err
	}
	userVo, err := UserService.Detail(userid, token)
	if err != nil {
		return ret, err
	}
	for _, v := range videos {
		var video vo.VideoVo
		video.Title = v.Title
		video.Id = v.Id
		video.CommentCount = v.CommentCount
		video.Author = userVo
		video.IsFavorite = false
		video.FavoriteCount = v.FavoriteCount
		if fileInfo, ok := fileMap[v.PlayId]; ok {
			video.PlayUrl = fileInfo.FileUrl
		}
		if fileInfo, ok := fileMap[v.CoverId]; ok {
			video.CoverUrl = fileInfo.FileUrl
		}
		ret = append(ret, video)
	}
	return ret, nil
}
