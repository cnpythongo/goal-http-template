package controller

// 查询用户结构体
type GetUserListPayload struct {
	Page             int    `json:"page" binding:"required"`
	Size             int    `json:"size" binding:"required"`
	LastLoginAtStart string `json:"last_login_at_start"`
	LastLoginAtEnd   string `json:"last_login_at_end"`
}
