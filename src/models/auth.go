package models

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "Other"}[g]
}

type Designation int

const (
	Manager Designation = iota
	Developer
	Designer
	Tester
)

func (d Designation) String() string {
	return [...]string{"Manager", "Developer", "Designer", "Tester"}[d]
}

type User struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Email            string `gorm:"unique;not null" json:"email"`
	Username         string `gorm:"unique;not null" json:"username"`
	Password         string `gorm:"not null" json:"-"`
	Address          string `gorm:"not null" json:"address"`
	Age              int    `gorm:"not null" json:"age"`
	Gender           string `gorm:"not null" json:"gender"`
	Designation      string `gorm:"not null" json:"designation"`
	NationalIDNumber string `gorm:"not null" json:"national_id_number"`
}
