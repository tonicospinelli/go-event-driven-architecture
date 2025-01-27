// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BasketspbGetBasketResponse basketspb get basket response
//
// swagger:model basketspbGetBasketResponse
type BasketspbGetBasketResponse struct {

	// basket
	Basket *BasketspbBasket `json:"basket,omitempty"`
}

// Validate validates this basketspb get basket response
func (m *BasketspbGetBasketResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasketspbGetBasketResponse) validateBasket(formats strfmt.Registry) error {
	if swag.IsZero(m.Basket) { // not required
		return nil
	}

	if m.Basket != nil {
		if err := m.Basket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this basketspb get basket response based on the context it is used
func (m *BasketspbGetBasketResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasketspbGetBasketResponse) contextValidateBasket(ctx context.Context, formats strfmt.Registry) error {

	if m.Basket != nil {
		if err := m.Basket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BasketspbGetBasketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasketspbGetBasketResponse) UnmarshalBinary(b []byte) error {
	var res BasketspbGetBasketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
