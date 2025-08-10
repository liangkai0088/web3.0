package types

import "time"

type User struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
}

type Post struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	UserID    uint      `json:"user_id" db:"user_id"`
}

type Comment struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Content   string    `json:"content" db:"content"`
	UserID    uint      `json:"user_id" db:"user_id"`
	PostID    uint      `json:"post_id" db:"post_id"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserPageReq struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type PostCreateReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostEditReq struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostDeleteReq struct {
	PostID uint `form:"postId"`
}

type PostDetailReq struct {
	PostID uint `form:"postId"`
}

type PostPageReq struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type CommentCreateReq struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

type CommentListReq struct {
	PostID uint `form:"postId"`
}
