// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DepotServiceAssignShoppingListBody depot service assign shopping list body
//
// swagger:model DepotServiceAssignShoppingListBody
type DepotServiceAssignShoppingListBody struct {

	// bot Id
	BotID string `json:"botId,omitempty"`
}

// Validate validates this depot service assign shopping list body
func (m *DepotServiceAssignShoppingListBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this depot service assign shopping list body based on context it is used
func (m *DepotServiceAssignShoppingListBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DepotServiceAssignShoppingListBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DepotServiceAssignShoppingListBody) UnmarshalBinary(b []byte) error {
	var res DepotServiceAssignShoppingListBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}