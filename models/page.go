package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Title   string `gorm:"unique;not null"`
	Content []byte
	//AuthorId uint
	//Author   User `gorm:"foreignKey:AuthorID"`
}

func (p *Page) Save(db *gorm.DB) error {
	result := db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func LoadPage(title string, db *gorm.DB) (*Page, error) {
	var p Page
	result := db.Where("title = ?", title).First(&p)
	if result.Error != nil {
		return nil, result.Error
	}
	return &p, nil
}

func (p *Page) Update(content []byte, db *gorm.DB) error {
	result := db.Model(p).Update("Content", content)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Page) Delete(db *gorm.DB) error {
	result := db.Delete(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
