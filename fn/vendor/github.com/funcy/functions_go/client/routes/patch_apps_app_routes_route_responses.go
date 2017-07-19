package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// PatchAppsAppRoutesRouteReader is a Reader for the PatchAppsAppRoutesRoute structure.
type PatchAppsAppRoutesRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAppsAppRoutesRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAppsAppRoutesRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAppsAppRoutesRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAppsAppRoutesRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchAppsAppRoutesRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAppsAppRoutesRouteOK creates a PatchAppsAppRoutesRouteOK with default headers values
func NewPatchAppsAppRoutesRouteOK() *PatchAppsAppRoutesRouteOK {
	return &PatchAppsAppRoutesRouteOK{}
}

/*PatchAppsAppRoutesRouteOK handles this case with default header values.

Route updated
*/
type PatchAppsAppRoutesRouteOK struct {
	Payload *models.RouteWrapper
}

func (o *PatchAppsAppRoutesRouteOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app}/routes/{route}][%d] patchAppsAppRoutesRouteOK  %+v", 200, o.Payload)
}

func (o *PatchAppsAppRoutesRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppsAppRoutesRouteBadRequest creates a PatchAppsAppRoutesRouteBadRequest with default headers values
func NewPatchAppsAppRoutesRouteBadRequest() *PatchAppsAppRoutesRouteBadRequest {
	return &PatchAppsAppRoutesRouteBadRequest{}
}

/*PatchAppsAppRoutesRouteBadRequest handles this case with default header values.

Invalid route due to parameters being missing or invalid.
*/
type PatchAppsAppRoutesRouteBadRequest struct {
	Payload *models.Error
}

func (o *PatchAppsAppRoutesRouteBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app}/routes/{route}][%d] patchAppsAppRoutesRouteBadRequest  %+v", 400, o.Payload)
}

func (o *PatchAppsAppRoutesRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppsAppRoutesRouteNotFound creates a PatchAppsAppRoutesRouteNotFound with default headers values
func NewPatchAppsAppRoutesRouteNotFound() *PatchAppsAppRoutesRouteNotFound {
	return &PatchAppsAppRoutesRouteNotFound{}
}

/*PatchAppsAppRoutesRouteNotFound handles this case with default header values.

App / Route does not exist.
*/
type PatchAppsAppRoutesRouteNotFound struct {
	Payload *models.Error
}

func (o *PatchAppsAppRoutesRouteNotFound) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app}/routes/{route}][%d] patchAppsAppRoutesRouteNotFound  %+v", 404, o.Payload)
}

func (o *PatchAppsAppRoutesRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppsAppRoutesRouteDefault creates a PatchAppsAppRoutesRouteDefault with default headers values
func NewPatchAppsAppRoutesRouteDefault(code int) *PatchAppsAppRoutesRouteDefault {
	return &PatchAppsAppRoutesRouteDefault{
		_statusCode: code,
	}
}

/*PatchAppsAppRoutesRouteDefault handles this case with default header values.

Unexpected error
*/
type PatchAppsAppRoutesRouteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch apps app routes route default response
func (o *PatchAppsAppRoutesRouteDefault) Code() int {
	return o._statusCode
}

func (o *PatchAppsAppRoutesRouteDefault) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app}/routes/{route}][%d] PatchAppsAppRoutesRoute default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAppsAppRoutesRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
