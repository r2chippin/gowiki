package model

import (
	"os"
)

// Page as a struct with two fields representing the title and body.
type Page struct {
	Title string
	Body  []byte
}

// Save the Page.Body to a text file, using the Page.Title as the file name.
func (p *Page) Save(td string) error {
	targetDir := "./Data"
	if td != "" {
		targetDir = td
	}
	fileName := p.Title + ".txt"
	filePath := targetDir + "/" + fileName
	return os.WriteFile(filePath, p.Body, 0600)
}

/*
LoadPage constructs the file name from the title parameter, reads the file's
contents into a new variable body, and returns a pointer to a Page literal
constructed with the proper Page.Title and Page.Body values.
*/
func LoadPage(title string, td string) (*Page, error) {
	targetDir := "./Data"
	if td != "" {
		targetDir = td
	}
	fileName := title + ".txt"
	filePath := targetDir + "/" + fileName
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
