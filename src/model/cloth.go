package model

import "time"

type Cloth struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:255;not null"`
    Color     string    `gorm:"size:100;not null"`
    Size      string    `gorm:"size:5;not null"`
    Price     float64   `gorm:"not null"`
    Stock     *int      
    CreatedAt time.Time
    UpdatedAt time.Time
}
