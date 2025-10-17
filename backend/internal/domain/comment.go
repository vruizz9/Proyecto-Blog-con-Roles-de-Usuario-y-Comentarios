package domain

type Comment struct {
	ID      int64  `json:"id"`
	BlogID  int64  `json:"blog_id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}
