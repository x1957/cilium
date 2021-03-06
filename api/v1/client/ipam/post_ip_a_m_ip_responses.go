package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIPAMIPReader is a Reader for the PostIPAMIP structure.
type PostIPAMIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIPAMIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIPAMIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostIPAMIPInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostIPAMIPExists()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostIPAMIPFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewPostIPAMIPDisabled()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostIPAMIPOK creates a PostIPAMIPOK with default headers values
func NewPostIPAMIPOK() *PostIPAMIPOK {
	return &PostIPAMIPOK{}
}

/*PostIPAMIPOK handles this case with default header values.

Success
*/
type PostIPAMIPOK struct {
}

func (o *PostIPAMIPOK) Error() string {
	return fmt.Sprintf("[POST /ipam/{ip}][%d] postIpAMIpOK ", 200)
}

func (o *PostIPAMIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMIPInvalid creates a PostIPAMIPInvalid with default headers values
func NewPostIPAMIPInvalid() *PostIPAMIPInvalid {
	return &PostIPAMIPInvalid{}
}

/*PostIPAMIPInvalid handles this case with default header values.

Invalid IP address
*/
type PostIPAMIPInvalid struct {
}

func (o *PostIPAMIPInvalid) Error() string {
	return fmt.Sprintf("[POST /ipam/{ip}][%d] postIpAMIpInvalid ", 400)
}

func (o *PostIPAMIPInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMIPExists creates a PostIPAMIPExists with default headers values
func NewPostIPAMIPExists() *PostIPAMIPExists {
	return &PostIPAMIPExists{}
}

/*PostIPAMIPExists handles this case with default header values.

IP already allocated
*/
type PostIPAMIPExists struct {
}

func (o *PostIPAMIPExists) Error() string {
	return fmt.Sprintf("[POST /ipam/{ip}][%d] postIpAMIpExists ", 409)
}

func (o *PostIPAMIPExists) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIPAMIPFailure creates a PostIPAMIPFailure with default headers values
func NewPostIPAMIPFailure() *PostIPAMIPFailure {
	return &PostIPAMIPFailure{}
}

/*PostIPAMIPFailure handles this case with default header values.

IP allocation failure. Details in message.
*/
type PostIPAMIPFailure struct {
	Payload models.Error
}

func (o *PostIPAMIPFailure) Error() string {
	return fmt.Sprintf("[POST /ipam/{ip}][%d] postIpAMIpFailure  %+v", 500, o.Payload)
}

func (o *PostIPAMIPFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIPAMIPDisabled creates a PostIPAMIPDisabled with default headers values
func NewPostIPAMIPDisabled() *PostIPAMIPDisabled {
	return &PostIPAMIPDisabled{}
}

/*PostIPAMIPDisabled handles this case with default header values.

Allocation for address family disabled
*/
type PostIPAMIPDisabled struct {
}

func (o *PostIPAMIPDisabled) Error() string {
	return fmt.Sprintf("[POST /ipam/{ip}][%d] postIpAMIpDisabled ", 501)
}

func (o *PostIPAMIPDisabled) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
