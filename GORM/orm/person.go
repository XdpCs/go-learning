package orm

type Person struct {
	ID       uint   `gorm:"column:user_id"`
	Username string `gorm:"size:260"`
	Sex      string `gorm:"size:260"`
	Email    string `gorm:"size:260"`
}
