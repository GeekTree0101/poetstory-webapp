package feed

import (
	context "context"
)

// Controller for feeds
type Controller struct {
}

// Feed struct
type Feed struct {
	ID int `json:"id"`
}

// Index of feeds
// GET /feed
func (c *Controller) Index(ctx context.Context) (feeds []*Feed, err error) {
	return []*Feed{}, nil
}

// Show feed
// GET /feed/:id
func (c *Controller) Show(ctx context.Context, id int) (feed *Feed, err error) {
	return &Feed{
		ID: id,
	}, nil
}

// Delete feed
// DELETE /feed/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
