// Code generated by go-swagger; DO NOT EDIT.

package certificate_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteCertificateParams creates a new DeleteCertificateParams object
// with the default values initialized.
func NewDeleteCertificateParams() *DeleteCertificateParams {

	return &DeleteCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCertificateParamsWithTimeout creates a new DeleteCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCertificateParamsWithTimeout(timeout time.Duration) *DeleteCertificateParams {

	return &DeleteCertificateParams{

		timeout: timeout,
	}
}

// NewDeleteCertificateParamsWithContext creates a new DeleteCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCertificateParamsWithContext(ctx context.Context) *DeleteCertificateParams {

	return &DeleteCertificateParams{

		Context: ctx,
	}
}

// NewDeleteCertificateParamsWithHTTPClient creates a new DeleteCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCertificateParamsWithHTTPClient(client *http.Client) *DeleteCertificateParams {

	return &DeleteCertificateParams{
		HTTPClient: client,
	}
}

/*DeleteCertificateParams contains all the parameters to send to the API endpoint
for the delete certificate operation typically these are written to a http.Request
*/
type DeleteCertificateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete certificate params
func (o *DeleteCertificateParams) WithTimeout(timeout time.Duration) *DeleteCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete certificate params
func (o *DeleteCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete certificate params
func (o *DeleteCertificateParams) WithContext(ctx context.Context) *DeleteCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete certificate params
func (o *DeleteCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete certificate params
func (o *DeleteCertificateParams) WithHTTPClient(client *http.Client) *DeleteCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete certificate params
func (o *DeleteCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}