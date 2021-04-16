package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/v3/validators"
)
// Journal is used by pop to map your journals database table to your go code.
type Journal struct {
    ID uuid.UUID `json:"id" db:"id"`
    Title string `json:"title" db:"title"`
    Body string `json:"body" db:"body"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (j Journal) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Journals is not required by pop and may be deleted
type Journals []Journal

// String is not required by pop and may be deleted
func (j Journals) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (j *Journal) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: j.Title, Name: "Title"},
		&validators.StringIsPresent{Field: j.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (j *Journal) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (j *Journal) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
