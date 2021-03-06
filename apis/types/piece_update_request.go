// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PieceUpdateRequest request used to update piece attributes.
// swagger:model PieceUpdateRequest
type PieceUpdateRequest struct {

	// the downloader clientID
	//
	ClientID string `json:"clientID,omitempty"`

	// the uploader peerID
	//
	DstPID string `json:"dstPID,omitempty"`

	// pieceStatus indicates whether the peer task successfully download the piece.
	//
	// Enum: [FAILED SUCCESS INVALID SEMISUC]
	PieceStatus string `json:"pieceStatus,omitempty"`
}

// Validate validates this piece update request
func (m *PieceUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePieceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pieceUpdateRequestTypePieceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FAILED","SUCCESS","INVALID","SEMISUC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pieceUpdateRequestTypePieceStatusPropEnum = append(pieceUpdateRequestTypePieceStatusPropEnum, v)
	}
}

const (

	// PieceUpdateRequestPieceStatusFAILED captures enum value "FAILED"
	PieceUpdateRequestPieceStatusFAILED string = "FAILED"

	// PieceUpdateRequestPieceStatusSUCCESS captures enum value "SUCCESS"
	PieceUpdateRequestPieceStatusSUCCESS string = "SUCCESS"

	// PieceUpdateRequestPieceStatusINVALID captures enum value "INVALID"
	PieceUpdateRequestPieceStatusINVALID string = "INVALID"

	// PieceUpdateRequestPieceStatusSEMISUC captures enum value "SEMISUC"
	PieceUpdateRequestPieceStatusSEMISUC string = "SEMISUC"
)

// prop value enum
func (m *PieceUpdateRequest) validatePieceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pieceUpdateRequestTypePieceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PieceUpdateRequest) validatePieceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PieceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePieceStatusEnum("pieceStatus", "body", m.PieceStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PieceUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieceUpdateRequest) UnmarshalBinary(b []byte) error {
	var res PieceUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
