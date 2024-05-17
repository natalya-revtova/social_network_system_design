package api

import "time"

type RestError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Post struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Text       string    `json:"text"`
	Location   string    `json:"location"`
	ImagesURLS []string  `json:"images_urls"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Comment struct {
	PostID         int64     `json:"post_id"`
	UserID         int64     `json:"user_id"`
	Text           string    `json:"text"`
	ReplyCommentID int64     `json:"reply_comment_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type PaginatedPosts struct {
	Records []Post `json:"records"`
	Total   int    `json:"total"`
}

type Chat struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	AvatarURL      string  `json:"avatar_url"`
	UsersIDs       []int64 `json:"users_ids"`
	UnreadMessages bool    `json:"unread_messages"`
}

type Message struct {
	UserID     int64     `json:"user_id"`
	ChatID     int64     `json:"chat_id"`
	Text       string    `json:"text"`
	Seen       bool      `json:"seen"`
	ImagesURLS []string  `json:"images_urls"`
	CreatedAt  time.Time `json:"created_at"`
}

type PaginatedChats struct {
	Records []Chat `json:"records"`
	Total   int    `json:"total"`
}

type PaginatedMessages struct {
	Records []Message `json:"records"`
	Total   int       `json:"total"`
}

type PostsTop struct {
	Location string `json:"location"`
}

type ImageInfo struct {
	URL string `json:"url"`
}
