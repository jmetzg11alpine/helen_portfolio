package models

type User struct {
	ID          uint           `gorm:"primaryKey"`
	Email       string         `json:"email" gorm:"unique"`
	Username    string         `json:"username"`
	Role        string         `json:"role"` // "admin" or "user"
	Comments    []BlogComment  `gorm:"foreignKey:UserID"`
	MerchOrders []MerchHistory `gorm:"foreignKey:UserID"`
}

type BlogPost struct {
	ID        uint          `gorm:"primaryKey"`
	Title     string        `json:"title"`
	SubTitle  string        `json:"sub_title"`
	Content   string        `json:"content"`
	Comments  []BlogComment `gorm:"foreignKey:BlogID"`
	CreatedAt int64         `json:"created_at"`
}

type BlogComment struct {
	ID        uint     `gorm:"primaryKey"`
	Content   string   `json:"content"`
	UserID    uint     `json:"user_id"`
	User      User     `gorm:"foreignKey:UserID"`
	BlogID    uint     `json:"blog_id"`
	BlogPost  BlogPost `gorm:"foreignKey:BlogID"`
	CreatedAt int64    `json:"created_at"`
}

type Merch struct {
	ID        uint    `gorm:"primaryKey"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	CreatedAt int64   `json:"created_at"`
}

type MerchHistory struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	ItemName  string `json:"item_name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	CreatedAt int64  `json:"created_at"`
}
