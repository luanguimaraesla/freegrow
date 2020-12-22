package gadget

import "fmt"

// Gadgets is the entity used for controlling
// interactions with many gadgets in the database
type Gadgets struct {
	userID int64
}

// NewGadgets returns a Gadgets controller
func NewGadgets() *Gadgets {
	return &Gadgets{}
}

// NewUserGadgets returns a Gadgets controller to a specific user
func NewUserGadgets(userID int64) *Gadgets {
	return &Gadgets{userID}
}

// All returns all the gadgets registered in the database
func (gs *Gadgets) All() ([]*Gadget, error) {
	// if no user was set
	if gs.userID == 0 {
		gadgets, err := getAllGadgets()
		if err != nil {
			return gadgets, fmt.Errorf("unable to get gadgets: %v", err)
		}

		return gadgets, nil
	}

	gadgets, err := getUserGadgets(gs.userID)
	if err != nil {
		return gadgets, fmt.Errorf("unable to get user %v gadgets: %v", gs.userID, err)
	}

	return gadgets, nil
}

// Get returns the gadget information according to its UUID
func (gs *Gadgets) Get(UUID string) (*Gadget, error) {
	gadget, err := getGadgetByUUID(UUID)
	if err != nil {
		return nil, fmt.Errorf("unable to load gadget: %v", err)
	}

	return gadget, nil
}

// Delete deletes an gadget based on its UUID
func (gs *Gadgets) Delete(UUID string) error {
	if err := deleteGadget(UUID); err != nil {
		return fmt.Errorf("unable to delete gadget: %v", err)
	}

	return nil
}

// Register registers a new gadget for an user
func (gs *Gadgets) Register(gadget *Gadget) error {
	if gs.userID == 0 {
		return fmt.Errorf("unable to register a new gadget: user_id not defined")
	}

	gadget.UserID = gs.userID

	if err := gadget.Create(); err != nil {
		return fmt.Errorf("unable to register a new gadget: %v", err)
	}

	return nil
}
