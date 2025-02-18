// Code generated by go-swagger; DO NOT EDIT.

package cloud_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RefreshPricingReader is a Reader for the RefreshPricing structure.
type RefreshPricingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshPricingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewRefreshPricingDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewRefreshPricingDefault creates a RefreshPricingDefault with default headers values
func NewRefreshPricingDefault(code int) *RefreshPricingDefault {
	return &RefreshPricingDefault{
		_statusCode: code,
	}
}

/* RefreshPricingDefault describes a response with status code -1, with default header values.

successful operation
*/
type RefreshPricingDefault struct {
	_statusCode int
}

// Code gets the status code for the refresh pricing default response
func (o *RefreshPricingDefault) Code() int {
	return o._statusCode
}

func (o *RefreshPricingDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/customers/{cUUID}/providers/{pUUID}/refresh_pricing][%d] refreshPricing default ", o._statusCode)
}

func (o *RefreshPricingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
