package utils

import (
	"testing"
)

func TestGetSnapshot(t *testing.T) {
	err := GetVideoCover("../public/test.mp4", "E:\\workspace\\go\\src\\douyin-easy\\public\\test", 1)
	if err != nil {
		return
	}
}
