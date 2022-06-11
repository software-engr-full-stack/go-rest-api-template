package app

import (
    "context"
    "time"
)

type Resource struct {
    ID           uint      `json:"-"`
    Name         string    `json:"name,omitempty"`

    // Could be type User, using string for now for simplicity
    Owner        string    `json:"owner,omitempty"`

    // Could be type Data, using string for now for simplicity
    Data         string    `json:"data,omitempty"`

    CreatedAt    time.Time `json:"-" db:"created_at"`
    UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

type ResourcePatch struct {
    Owner        *string `json:"owner"`
    Data         *string `json:"data"`
}

type ResourceFilter struct {
    ID       *uint
    Name     *string
}

type ResourceService interface {
    CreateResource(context.Context, *Resource) error

    ResourceByID(context.Context, uint) (*Resource, error)

    ResourceByName(context.Context, string) (*Resource, error)

    Resources(context.Context, ResourceFilter) ([]*Resource, error)

    UpdateResource(context.Context, *Resource, ResourcePatch) error

    DeleteResource(context.Context, uint) error
}
