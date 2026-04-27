package entity


import(
	"time"
)

type PostEntity struct{
	ID uint
	UserID uint
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}



func NewPostEntity(userID uint,title, content string) *PostEntity{
	return &PostEntity{
		UserID: userID,
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

