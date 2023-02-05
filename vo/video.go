package vo

type VideoVo struct {
	Id            uint64 `json:"id"`            // 视频id
	Title         string `json:"title"`         // 标题
	Author        UserVo `json:"author"`        // 作者
	PlayUrl       string `json:"play_url"`      // 播放地址
	CoverUrl      string `json:"cover_url"`     // 视频封面地址
	FavoriteCount uint64 `json:"favoriteCount"` // 收到的喜欢数目
	CommentCount  uint64 `json:"commentCount"`  // 评论数
	IsFavorite    bool   `json:"is_favorite"`   // 是否喜欢
}
