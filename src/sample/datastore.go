package sample

import (
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

// Item
type Item struct {
	Key       *datastore.Key `json:"-" datastore:"-"`
	KeyStr    string         `json:"key" datastore:"-"`
	Title     string         `json:"title" datastore:",noindex"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// Put
func Put(c context.Context, title string) (*Item, error) {
	now := time.Now()

	item := Item{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := datastore.RunInTransaction(c, func(c context.Context) error {
		key, err := datastore.Put(c, datastore.NewKey(c, "Item", title, 0, nil), &item)
		if err != nil {
			return err
		}
		item.Key = key
		item.KeyStr = key.Encode()

		return nil
	}, nil)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Get
func Get(c context.Context, key *datastore.Key) (*Item, error) {
	var item Item
	err := datastore.Get(c, key, &item)
	if err != nil {
		return nil, err
	}
	return &item, err
}
