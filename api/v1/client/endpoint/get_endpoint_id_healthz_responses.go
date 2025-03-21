// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDHealthzReader is a Reader for the GetEndpointIDHealthz structure.
type GetEndpointIDHealthzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointIDHealthzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointIDHealthzOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEndpointIDHealthzInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointIDHealthzNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEndpointIDHealthzTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /endpoint/{id}/healthz] GetEndpointIDHealthz", response, response.Code())
	}
}

// NewGetEndpointIDHealthzOK creates a GetEndpointIDHealthzOK with default headers values
func NewGetEndpointIDHealthzOK() *GetEndpointIDHealthzOK {
	return &GetEndpointIDHealthzOK{}
}

/*
GetEndpointIDHealthzOK describes a response with status code 200, with default header values.

Success
*/
type GetEndpointIDHealthzOK struct {
	Payload *models.EndpointHealth
}

// IsSuccess returns true when this get endpoint Id healthz o k response has a 2xx status code
func (o *GetEndpointIDHealthzOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint Id healthz o k response has a 3xx status code
func (o *GetEndpointIDHealthzOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id healthz o k response has a 4xx status code
func (o *GetEndpointIDHealthzOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint Id healthz o k response has a 5xx status code
func (o *GetEndpointIDHealthzOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id healthz o k response a status code equal to that given
func (o *GetEndpointIDHealthzOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get endpoint Id healthz o k response
func (o *GetEndpointIDHealthzOK) Code() int {
	return 200
}

func (o *GetEndpointIDHealthzOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzOK %s", 200, payload)
}

func (o *GetEndpointIDHealthzOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzOK %s", 200, payload)
}

func (o *GetEndpointIDHealthzOK) GetPayload() *models.EndpointHealth {
	return o.Payload
}

func (o *GetEndpointIDHealthzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDHealthzInvalid creates a GetEndpointIDHealthzInvalid with default headers values
func NewGetEndpointIDHealthzInvalid() *GetEndpointIDHealthzInvalid {
	return &GetEndpointIDHealthzInvalid{}
}

/*
GetEndpointIDHealthzInvalid describes a response with status code 400, with default header values.

Invalid identity provided
*/
type GetEndpointIDHealthzInvalid struct {
}

// IsSuccess returns true when this get endpoint Id healthz invalid response has a 2xx status code
func (o *GetEndpointIDHealthzInvalid) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint Id healthz invalid response has a 3xx status code
func (o *GetEndpointIDHealthzInvalid) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id healthz invalid response has a 4xx status code
func (o *GetEndpointIDHealthzInvalid) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint Id healthz invalid response has a 5xx status code
func (o *GetEndpointIDHealthzInvalid) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id healthz invalid response a status code equal to that given
func (o *GetEndpointIDHealthzInvalid) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get endpoint Id healthz invalid response
func (o *GetEndpointIDHealthzInvalid) Code() int {
	return 400
}

func (o *GetEndpointIDHealthzInvalid) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzInvalid", 400)
}

func (o *GetEndpointIDHealthzInvalid) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzInvalid", 400)
}

func (o *GetEndpointIDHealthzInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDHealthzNotFound creates a GetEndpointIDHealthzNotFound with default headers values
func NewGetEndpointIDHealthzNotFound() *GetEndpointIDHealthzNotFound {
	return &GetEndpointIDHealthzNotFound{}
}

/*
GetEndpointIDHealthzNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type GetEndpointIDHealthzNotFound struct {
}

// IsSuccess returns true when this get endpoint Id healthz not found response has a 2xx status code
func (o *GetEndpointIDHealthzNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint Id healthz not found response has a 3xx status code
func (o *GetEndpointIDHealthzNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id healthz not found response has a 4xx status code
func (o *GetEndpointIDHealthzNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint Id healthz not found response has a 5xx status code
func (o *GetEndpointIDHealthzNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id healthz not found response a status code equal to that given
func (o *GetEndpointIDHealthzNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get endpoint Id healthz not found response
func (o *GetEndpointIDHealthzNotFound) Code() int {
	return 404
}

func (o *GetEndpointIDHealthzNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzNotFound", 404)
}

func (o *GetEndpointIDHealthzNotFound) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzNotFound", 404)
}

func (o *GetEndpointIDHealthzNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDHealthzTooManyRequests creates a GetEndpointIDHealthzTooManyRequests with default headers values
func NewGetEndpointIDHealthzTooManyRequests() *GetEndpointIDHealthzTooManyRequests {
	return &GetEndpointIDHealthzTooManyRequests{}
}

/*
GetEndpointIDHealthzTooManyRequests describes a response with status code 429, with default header values.

Rate-limiting too many requests in the given time frame
*/
type GetEndpointIDHealthzTooManyRequests struct {
}

// IsSuccess returns true when this get endpoint Id healthz too many requests response has a 2xx status code
func (o *GetEndpointIDHealthzTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint Id healthz too many requests response has a 3xx status code
func (o *GetEndpointIDHealthzTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id healthz too many requests response has a 4xx status code
func (o *GetEndpointIDHealthzTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint Id healthz too many requests response has a 5xx status code
func (o *GetEndpointIDHealthzTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id healthz too many requests response a status code equal to that given
func (o *GetEndpointIDHealthzTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get endpoint Id healthz too many requests response
func (o *GetEndpointIDHealthzTooManyRequests) Code() int {
	return 429
}

func (o *GetEndpointIDHealthzTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzTooManyRequests", 429)
}

func (o *GetEndpointIDHealthzTooManyRequests) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/healthz][%d] getEndpointIdHealthzTooManyRequests", 429)
}

func (o *GetEndpointIDHealthzTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
