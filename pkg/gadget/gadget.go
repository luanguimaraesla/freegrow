package gadget

import (
	"fmt"
)

type Gadget struct {
	UUID    string `json:"gadget_uuid"`
	UserID  int64  `json:"user_id"`
	Enabled bool   `json:"enabled"`
}

// New returns an empty gadget
func New() *Gadget {
	return &Gadget{}
}

// Create inserts the gadget in the database
func (g *Gadget) Create() error {
	if err := insertGadget(g); err != nil {
		return fmt.Errorf("unable to create gadget: %v", err)
	}

	return nil
}

// Update updates the gadget record in the database
func (g *Gadget) Update() error {
	if err := updateGadget(g); err != nil {
		return fmt.Errorf("unable to update gadget", err)
	}

	return nil
}

// Delete deletes the gadget record from the database
func (g *Gadget) Delete() error {
	if err := deleteGadget(g.UUID); err != nil {
		return fmt.Errorf("unable to delete gadget", err)
	}

	return nil
}
