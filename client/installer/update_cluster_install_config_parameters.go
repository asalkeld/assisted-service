// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewUpdateClusterInstallConfigParams creates a new UpdateClusterInstallConfigParams object
// with the default values initialized.
func NewUpdateClusterInstallConfigParams() *UpdateClusterInstallConfigParams {
	var ()
	return &UpdateClusterInstallConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterInstallConfigParamsWithTimeout creates a new UpdateClusterInstallConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterInstallConfigParamsWithTimeout(timeout time.Duration) *UpdateClusterInstallConfigParams {
	var ()
	return &UpdateClusterInstallConfigParams{

		timeout: timeout,
	}
}

// NewUpdateClusterInstallConfigParamsWithContext creates a new UpdateClusterInstallConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterInstallConfigParamsWithContext(ctx context.Context) *UpdateClusterInstallConfigParams {
	var ()
	return &UpdateClusterInstallConfigParams{

		Context: ctx,
	}
}

// NewUpdateClusterInstallConfigParamsWithHTTPClient creates a new UpdateClusterInstallConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterInstallConfigParamsWithHTTPClient(client *http.Client) *UpdateClusterInstallConfigParams {
	var ()
	return &UpdateClusterInstallConfigParams{
		HTTPClient: client,
	}
}

/*UpdateClusterInstallConfigParams contains all the parameters to send to the API endpoint
for the update cluster install config operation typically these are written to a http.Request
*/
type UpdateClusterInstallConfigParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID
	/*InstallConfigParams*/
	InstallConfigParams string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) WithTimeout(timeout time.Duration) *UpdateClusterInstallConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) WithContext(ctx context.Context) *UpdateClusterInstallConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) WithHTTPClient(client *http.Client) *UpdateClusterInstallConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) WithClusterID(clusterID strfmt.UUID) *UpdateClusterInstallConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithInstallConfigParams adds the installConfigParams to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) WithInstallConfigParams(installConfigParams string) *UpdateClusterInstallConfigParams {
	o.SetInstallConfigParams(installConfigParams)
	return o
}

// SetInstallConfigParams adds the installConfigParams to the update cluster install config params
func (o *UpdateClusterInstallConfigParams) SetInstallConfigParams(installConfigParams string) {
	o.InstallConfigParams = installConfigParams
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterInstallConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.InstallConfigParams); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
