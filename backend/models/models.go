package models

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
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	BlogID    uint     `json:"blog_id"`
	BlogPost  BlogPost `gorm:"foreignKey:BlogID"`
	CreatedAt int64    `json:"created_at"`
	Approved  bool     `json:"approved"`
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
	ItemName  string `json:"item_name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	CreatedAt int64  `json:"created_at"`
}
