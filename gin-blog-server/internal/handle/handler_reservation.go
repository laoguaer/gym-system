package handle

type GetUserReversationListQuery struct {
	UserID  int    `form:"user_id" binding:"required,min=1"`
	DayTime string `form:"day_time"`
}

