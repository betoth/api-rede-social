package models

import (
	"errors"
	"strings"
	"time"
)

// Publication representar a post make by a user
type Publication struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"author_id,omitempty"`
	AuthorNickName string    `json:"author_nick_name,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

// Prepare prepara a publication
func (p *Publication) Prepare() error {

	if err := p.Validate(); err != nil {
		return err
	}

	p.Format()

	return nil
}

// Validate validate a publication
func (p *Publication) Validate() error {

	if p.Title == "" {
		return errors.New("Title is required")
	}
	if p.Content == "" {
		return errors.New("Content is required")
	}
	return nil
}

// Format remove spaces for publication
func (p *Publication) Format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)

}
