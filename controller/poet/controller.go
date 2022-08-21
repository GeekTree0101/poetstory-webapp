package poet

import (
	context "context"
)

// Controller for poets
type Controller struct {
}

// Poet struct
type Poet struct {
	ID int `json:"id"`
}

// Index of poets
// GET /poet
func (c *Controller) Index(ctx context.Context) (poets []*Poet, err error) {
	return []*Poet{}, nil
}

// Show poet
// GET /poet/:id
func (c *Controller) Show(ctx context.Context, id int) (poet *Poet, err error) {
	return &Poet{
		ID: id,
	}, nil
}

// Delete poet
// DELETE /poet/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
