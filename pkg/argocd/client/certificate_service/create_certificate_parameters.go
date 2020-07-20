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

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// NewCreateCertificateParams creates a new CreateCertificateParams object
// with the default values initialized.
func NewCreateCertificateParams() *CreateCertificateParams {
	var ()
	return &CreateCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCertificateParamsWithTimeout creates a new CreateCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCertificateParamsWithTimeout(timeout time.Duration) *CreateCertificateParams {
	var ()
	return &CreateCertificateParams{

		timeout: timeout,
	}
}

// NewCreateCertificateParamsWithContext creates a new CreateCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCertificateParamsWithContext(ctx context.Context) *CreateCertificateParams {
	var ()
	return &CreateCertificateParams{

		Context: ctx,
	}
}

// NewCreateCertificateParamsWithHTTPClient creates a new CreateCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCertificateParamsWithHTTPClient(client *http.Client) *CreateCertificateParams {
	var ()
	return &CreateCertificateParams{
		HTTPClient: client,
	}
}

/*CreateCertificateParams contains all the parameters to send to the API endpoint
for the create certificate operation typically these are written to a http.Request
*/
type CreateCertificateParams struct {

	/*Body*/
	Body *models.V1alpha1RepositoryCertificateList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create certificate params
func (o *CreateCertificateParams) WithTimeout(timeout time.Duration) *CreateCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create certificate params
func (o *CreateCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create certificate params
func (o *CreateCertificateParams) WithContext(ctx context.Context) *CreateCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create certificate params
func (o *CreateCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create certificate params
func (o *CreateCertificateParams) WithHTTPClient(client *http.Client) *CreateCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create certificate params
func (o *CreateCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create certificate params
func (o *CreateCertificateParams) WithBody(body *models.V1alpha1RepositoryCertificateList) *CreateCertificateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create certificate params
func (o *CreateCertificateParams) SetBody(body *models.V1alpha1RepositoryCertificateList) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}