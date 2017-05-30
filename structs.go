package main

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
)

type (
	Admin struct {
		Login          string `json:"login"`
		HashedPassword uint32
	}

	Like struct {
		Ip string
	}

	Tag struct {
		ID   string `json:"-"`
		Name string `json:"name"`
	}

	Image struct {
		ID  string `json:"-"`
		Url string `json:"url"`
	}

	Post struct {
		ID        string   `json:"-"`
		Name      string   `json:"name"`
		Content   string   `json:"text"`
		Permalink string   `json:"permalink"`
		Tags      []Tag    `json:"-"`
		TagsIDs   []string `json:"-"`
		Likes     uint16   `json:"likes"`
		CreatedAt string   `json:"created-at"`
	}

	Message struct {
		Phone string `json:"phone"`
		Email string `json:"e-mail"`
		Text  string `json:"text"`
		Theme string `json:"theme"`
	}
)

func (p Post) GetID() string {
	return p.ID
}
func (t Tag) GetID() string {
	return t.ID
}
func (i Image) GetID() string {
	return i.ID
}
func (p Post) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "tags",
			Name: "tags",
		},
	}
}
func (p Post) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range p.Tags {
		result = append(result, p.Tags[key])
	}

	return result
}

func (p Post) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, TagID := range p.TagsIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   TagID,
			Type: "tags",
			Name: "tags",
		})
	}

	return result
}
func (p *Post) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "tags" {
		p.TagsIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
func (p *Post) AddToManyIDs(name string, IDs []string) error {
	if name == "tags" {
		p.TagsIDs = append(p.TagsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
