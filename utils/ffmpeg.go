package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func GetVideoCover(videoPath, coverPath string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Printf("生成视频封面失败,err->%s\n", err)
		return errors.New("生成视频封面失败")
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Printf("图片数据解码失败,err->%s\n", err)
		return errors.New("生成视频封面失败")
	}
	err = imaging.Save(img, coverPath)
	if err != nil {
		log.Printf("视频封面存储失败,err->%s\n", err)
		return err
	}
	return nil
}
