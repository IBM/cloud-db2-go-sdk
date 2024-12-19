/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

// Package db2saasv1 : Operations and models for the Db2saasV1 service
package db2saasv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/cloud-db2-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// Db2saasV1 : Manage lifecycle of your Db2 on Cloud resources using the  APIs.
//
// API Version: 1.0.0
type Db2saasV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.db2.saas.ibm.com/dbapi/v4"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "db2saas"

const ParameterizedServiceURL = "https://{region}.db2.saas.ibm.com/dbapi/v4"

var defaultUrlVariables = map[string]string{
	"region": "us-south",
}

// Db2saasV1Options : Service options
type Db2saasV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDb2saasV1UsingExternalConfig : constructs an instance of Db2saasV1 with passed in options and external configuration.
func NewDb2saasV1UsingExternalConfig(options *Db2saasV1Options) (db2saas *Db2saasV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	db2saas, err = NewDb2saasV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = db2saas.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = db2saas.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDb2saasV1 : constructs an instance of Db2saasV1 with passed in options.
func NewDb2saasV1(options *Db2saasV1Options) (service *Db2saasV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &Db2saasV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "db2saas" suitable for processing requests.
func (db2saas *Db2saasV1) Clone() *Db2saasV1 {
	if core.IsNil(db2saas) {
		return nil
	}
	clone := *db2saas
	clone.Service = db2saas.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (db2saas *Db2saasV1) SetServiceURL(url string) error {
	err := db2saas.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (db2saas *Db2saasV1) GetServiceURL() string {
	return db2saas.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (db2saas *Db2saasV1) SetDefaultHeaders(headers http.Header) {
	db2saas.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (db2saas *Db2saasV1) SetEnableGzipCompression(enableGzip bool) {
	db2saas.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (db2saas *Db2saasV1) GetEnableGzipCompression() bool {
	return db2saas.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (db2saas *Db2saasV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	db2saas.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (db2saas *Db2saasV1) DisableRetries() {
	db2saas.Service.DisableRetries()
}

// GetDb2SaasConnectionInfo : Get Db2 connection information
func (db2saas *Db2saasV1) GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) (result *SuccessConnectionInfo, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasConnectionInfoWithContext(context.Background(), getDb2SaasConnectionInfoOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasConnectionInfoWithContext is an alternate form of the GetDb2SaasConnectionInfo method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasConnectionInfoWithContext(ctx context.Context, getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) (result *SuccessConnectionInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasConnectionInfoOptions, "getDb2SaasConnectionInfoOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasConnectionInfoOptions, "getDb2SaasConnectionInfoOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *getDb2SaasConnectionInfoOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/connectioninfo/{deployment_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasConnectionInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasConnectionInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasConnectionInfoOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasConnectionInfoOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_connection_info", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessConnectionInfo)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PostDb2SaasAllowlist : Allow listing of new IPs
func (db2saas *Db2saasV1) PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) (result *SuccessPostAllowedlistIPs, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasAllowlistWithContext(context.Background(), postDb2SaasAllowlistOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasAllowlistWithContext is an alternate form of the PostDb2SaasAllowlist method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasAllowlistWithContext(ctx context.Context, postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) (result *SuccessPostAllowedlistIPs, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasAllowlistOptions, "postDb2SaasAllowlistOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasAllowlistOptions, "postDb2SaasAllowlistOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/dbsettings/whitelistips`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if postDb2SaasAllowlistOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasAllowlistOptions.XDeploymentID))
	}

	body := make(map[string]interface{})
	if postDb2SaasAllowlistOptions.IpAddresses != nil {
		body["ip_addresses"] = postDb2SaasAllowlistOptions.IpAddresses
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_allowlist", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessPostAllowedlistIPs)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasAllowlist : Get allowed list of IPs
func (db2saas *Db2saasV1) GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) (result *SuccessGetAllowlistIPs, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasAllowlistWithContext(context.Background(), getDb2SaasAllowlistOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasAllowlistWithContext is an alternate form of the GetDb2SaasAllowlist method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasAllowlistWithContext(ctx context.Context, getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) (result *SuccessGetAllowlistIPs, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasAllowlistOptions, "getDb2SaasAllowlistOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasAllowlistOptions, "getDb2SaasAllowlistOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/dbsettings/whitelistips`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasAllowlistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasAllowlist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasAllowlistOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasAllowlistOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_allowlist", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetAllowlistIPs)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PostDb2SaasUser : Create new user ( available only for platform users)
func (db2saas *Db2saasV1) PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions) (result *SuccessUserResponse, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PostDb2SaasUserWithContext(context.Background(), postDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PostDb2SaasUserWithContext is an alternate form of the PostDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) PostDb2SaasUserWithContext(ctx context.Context, postDb2SaasUserOptions *PostDb2SaasUserOptions) (result *SuccessUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postDb2SaasUserOptions, "postDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(postDb2SaasUserOptions, "postDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range postDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PostDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if postDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*postDb2SaasUserOptions.XDeploymentID))
	}

	body := make(map[string]interface{})
	if postDb2SaasUserOptions.ID != nil {
		body["id"] = postDb2SaasUserOptions.ID
	}
	if postDb2SaasUserOptions.Iam != nil {
		body["iam"] = postDb2SaasUserOptions.Iam
	}
	if postDb2SaasUserOptions.Ibmid != nil {
		body["ibmid"] = postDb2SaasUserOptions.Ibmid
	}
	if postDb2SaasUserOptions.Name != nil {
		body["name"] = postDb2SaasUserOptions.Name
	}
	if postDb2SaasUserOptions.Password != nil {
		body["password"] = postDb2SaasUserOptions.Password
	}
	if postDb2SaasUserOptions.Role != nil {
		body["role"] = postDb2SaasUserOptions.Role
	}
	if postDb2SaasUserOptions.Email != nil {
		body["email"] = postDb2SaasUserOptions.Email
	}
	if postDb2SaasUserOptions.Locked != nil {
		body["locked"] = postDb2SaasUserOptions.Locked
	}
	if postDb2SaasUserOptions.Authentication != nil {
		body["authentication"] = postDb2SaasUserOptions.Authentication
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "post_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessUserResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasUser : Get the list of Users
func (db2saas *Db2saasV1) GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions) (result *SuccessGetUserInfo, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasUserWithContext(context.Background(), getDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasUserWithContext is an alternate form of the GetDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasUserWithContext(ctx context.Context, getDb2SaasUserOptions *GetDb2SaasUserOptions) (result *SuccessGetUserInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasUserOptions, "getDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasUserOptions, "getDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetUserInfo)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDb2SaasUser : Delete a user (only platform admin)
func (db2saas *Db2saasV1) DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions) (result map[string]interface{}, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.DeleteDb2SaasUserWithContext(context.Background(), deleteDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDb2SaasUserWithContext is an alternate form of the DeleteDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) DeleteDb2SaasUserWithContext(ctx context.Context, deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions) (result map[string]interface{}, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDb2SaasUserOptions, "deleteDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDb2SaasUserOptions, "deleteDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteDb2SaasUserOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "DeleteDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*deleteDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = db2saas.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetbyidDb2SaasUser : Get specific user by Id
func (db2saas *Db2saasV1) GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) (result *SuccessGetUserByID, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetbyidDb2SaasUserWithContext(context.Background(), getbyidDb2SaasUserOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetbyidDb2SaasUserWithContext is an alternate form of the GetbyidDb2SaasUser method which supports a Context parameter
func (db2saas *Db2saasV1) GetbyidDb2SaasUserWithContext(ctx context.Context, getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) (result *SuccessGetUserByID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getbyidDb2SaasUserOptions, "getbyidDb2SaasUserOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getbyidDb2SaasUserOptions, "getbyidDb2SaasUserOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/users/bluadmin`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getbyidDb2SaasUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetbyidDb2SaasUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getbyidDb2SaasUserOptions.XDeploymentID != nil {
		builder.AddHeader("x-deployment-id", fmt.Sprint(*getbyidDb2SaasUserOptions.XDeploymentID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getbyid_db2_saas_user", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessGetUserByID)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PutDb2SaasAutoscale : Update auto scaling configuration
func (db2saas *Db2saasV1) PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) (result *SuccessUpdateAutoScale, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.PutDb2SaasAutoscaleWithContext(context.Background(), putDb2SaasAutoscaleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PutDb2SaasAutoscaleWithContext is an alternate form of the PutDb2SaasAutoscale method which supports a Context parameter
func (db2saas *Db2saasV1) PutDb2SaasAutoscaleWithContext(ctx context.Context, putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) (result *SuccessUpdateAutoScale, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putDb2SaasAutoscaleOptions, "putDb2SaasAutoscaleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(putDb2SaasAutoscaleOptions, "putDb2SaasAutoscaleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/scaling/auto`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range putDb2SaasAutoscaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "PutDb2SaasAutoscale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if putDb2SaasAutoscaleOptions.XDbProfile != nil {
		builder.AddHeader("x-db-profile", fmt.Sprint(*putDb2SaasAutoscaleOptions.XDbProfile))
	}

	body := make(map[string]interface{})
	if putDb2SaasAutoscaleOptions.AutoScalingEnabled != nil {
		body["auto_scaling_enabled"] = putDb2SaasAutoscaleOptions.AutoScalingEnabled
	}
	if putDb2SaasAutoscaleOptions.AutoScalingThreshold != nil {
		body["auto_scaling_threshold"] = putDb2SaasAutoscaleOptions.AutoScalingThreshold
	}
	if putDb2SaasAutoscaleOptions.AutoScalingOverTimePeriod != nil {
		body["auto_scaling_over_time_period"] = putDb2SaasAutoscaleOptions.AutoScalingOverTimePeriod
	}
	if putDb2SaasAutoscaleOptions.AutoScalingPauseLimit != nil {
		body["auto_scaling_pause_limit"] = putDb2SaasAutoscaleOptions.AutoScalingPauseLimit
	}
	if putDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit != nil {
		body["auto_scaling_allow_plan_limit"] = putDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "put_db2_saas_autoscale", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessUpdateAutoScale)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDb2SaasAutoscale : Get auto scaling info
func (db2saas *Db2saasV1) GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) (result *SuccessAutoScaling, response *core.DetailedResponse, err error) {
	result, response, err = db2saas.GetDb2SaasAutoscaleWithContext(context.Background(), getDb2SaasAutoscaleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDb2SaasAutoscaleWithContext is an alternate form of the GetDb2SaasAutoscale method which supports a Context parameter
func (db2saas *Db2saasV1) GetDb2SaasAutoscaleWithContext(ctx context.Context, getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) (result *SuccessAutoScaling, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDb2SaasAutoscaleOptions, "getDb2SaasAutoscaleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDb2SaasAutoscaleOptions, "getDb2SaasAutoscaleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = db2saas.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(db2saas.Service.Options.URL, `/manage/scaling/auto`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getDb2SaasAutoscaleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("db2saas", "V1", "GetDb2SaasAutoscale")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDb2SaasAutoscaleOptions.XDbProfile != nil {
		builder.AddHeader("x-db-profile", fmt.Sprint(*getDb2SaasAutoscaleOptions.XDbProfile))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = db2saas.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_db2_saas_autoscale", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSuccessAutoScaling)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// CreateUserAuthentication : CreateUserAuthentication struct
type CreateUserAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Authentication policy ID.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// NewCreateUserAuthentication : Instantiate CreateUserAuthentication (Generic Model Constructor)
func (*Db2saasV1) NewCreateUserAuthentication(method string, policyID string) (_model *CreateUserAuthentication, err error) {
	_model = &CreateUserAuthentication{
		Method: core.StringPtr(method),
		PolicyID: core.StringPtr(policyID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalCreateUserAuthentication unmarshals an instance of CreateUserAuthentication from the specified map of raw messages.
func UnmarshalCreateUserAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateUserAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDb2SaasUserOptions : The DeleteDb2SaasUser options.
type DeleteDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// id of the user.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDb2SaasUserOptions : Instantiate DeleteDb2SaasUserOptions
func (*Db2saasV1) NewDeleteDb2SaasUserOptions(xDeploymentID string, id string) *DeleteDb2SaasUserOptions {
	return &DeleteDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		ID: core.StringPtr(id),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *DeleteDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *DeleteDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetID : Allow user to set ID
func (_options *DeleteDb2SaasUserOptions) SetID(id string) *DeleteDb2SaasUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDb2SaasUserOptions) SetHeaders(param map[string]string) *DeleteDb2SaasUserOptions {
	options.Headers = param
	return options
}

// GetDb2SaasAllowlistOptions : The GetDb2SaasAllowlist options.
type GetDb2SaasAllowlistOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasAllowlistOptions : Instantiate GetDb2SaasAllowlistOptions
func (*Db2saasV1) NewGetDb2SaasAllowlistOptions(xDeploymentID string) *GetDb2SaasAllowlistOptions {
	return &GetDb2SaasAllowlistOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasAllowlistOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasAllowlistOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasAllowlistOptions) SetHeaders(param map[string]string) *GetDb2SaasAllowlistOptions {
	options.Headers = param
	return options
}

// GetDb2SaasAutoscaleOptions : The GetDb2SaasAutoscale options.
type GetDb2SaasAutoscaleOptions struct {
	// CRN deployment id.
	XDbProfile *string `json:"x-db-profile" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasAutoscaleOptions : Instantiate GetDb2SaasAutoscaleOptions
func (*Db2saasV1) NewGetDb2SaasAutoscaleOptions(xDbProfile string) *GetDb2SaasAutoscaleOptions {
	return &GetDb2SaasAutoscaleOptions{
		XDbProfile: core.StringPtr(xDbProfile),
	}
}

// SetXDbProfile : Allow user to set XDbProfile
func (_options *GetDb2SaasAutoscaleOptions) SetXDbProfile(xDbProfile string) *GetDb2SaasAutoscaleOptions {
	_options.XDbProfile = core.StringPtr(xDbProfile)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasAutoscaleOptions) SetHeaders(param map[string]string) *GetDb2SaasAutoscaleOptions {
	options.Headers = param
	return options
}

// GetDb2SaasConnectionInfoOptions : The GetDb2SaasConnectionInfo options.
type GetDb2SaasConnectionInfoOptions struct {
	// Encoded CRN deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasConnectionInfoOptions : Instantiate GetDb2SaasConnectionInfoOptions
func (*Db2saasV1) NewGetDb2SaasConnectionInfoOptions(deploymentID string, xDeploymentID string) *GetDb2SaasConnectionInfoOptions {
	return &GetDb2SaasConnectionInfoOptions{
		DeploymentID: core.StringPtr(deploymentID),
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *GetDb2SaasConnectionInfoOptions) SetDeploymentID(deploymentID string) *GetDb2SaasConnectionInfoOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasConnectionInfoOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasConnectionInfoOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasConnectionInfoOptions) SetHeaders(param map[string]string) *GetDb2SaasConnectionInfoOptions {
	options.Headers = param
	return options
}

// GetDb2SaasUserOptions : The GetDb2SaasUser options.
type GetDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDb2SaasUserOptions : Instantiate GetDb2SaasUserOptions
func (*Db2saasV1) NewGetDb2SaasUserOptions(xDeploymentID string) *GetDb2SaasUserOptions {
	return &GetDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *GetDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDb2SaasUserOptions) SetHeaders(param map[string]string) *GetDb2SaasUserOptions {
	options.Headers = param
	return options
}

// GetbyidDb2SaasUserOptions : The GetbyidDb2SaasUser options.
type GetbyidDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetbyidDb2SaasUserOptions : Instantiate GetbyidDb2SaasUserOptions
func (*Db2saasV1) NewGetbyidDb2SaasUserOptions(xDeploymentID string) *GetbyidDb2SaasUserOptions {
	return &GetbyidDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *GetbyidDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *GetbyidDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetbyidDb2SaasUserOptions) SetHeaders(param map[string]string) *GetbyidDb2SaasUserOptions {
	options.Headers = param
	return options
}

// IpAddress : Details of an IP address.
type IpAddress struct {
	// The IP address, in IPv4/ipv6 format.
	Address *string `json:"address" validate:"required"`

	// Description of the IP address.
	Description *string `json:"description" validate:"required"`
}

// NewIpAddress : Instantiate IpAddress (Generic Model Constructor)
func (*Db2saasV1) NewIpAddress(address string, description string) (_model *IpAddress, err error) {
	_model = &IpAddress{
		Address: core.StringPtr(address),
		Description: core.StringPtr(description),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalIpAddress unmarshals an instance of IpAddress from the specified map of raw messages.
func UnmarshalIpAddress(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IpAddress)
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		err = core.SDKErrorf(err, "", "address-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostDb2SaasAllowlistOptions : The PostDb2SaasAllowlist options.
type PostDb2SaasAllowlistOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// List of IP addresses.
	IpAddresses []IpAddress `json:"ip_addresses" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPostDb2SaasAllowlistOptions : Instantiate PostDb2SaasAllowlistOptions
func (*Db2saasV1) NewPostDb2SaasAllowlistOptions(xDeploymentID string, ipAddresses []IpAddress) *PostDb2SaasAllowlistOptions {
	return &PostDb2SaasAllowlistOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		IpAddresses: ipAddresses,
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasAllowlistOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasAllowlistOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetIpAddresses : Allow user to set IpAddresses
func (_options *PostDb2SaasAllowlistOptions) SetIpAddresses(ipAddresses []IpAddress) *PostDb2SaasAllowlistOptions {
	_options.IpAddresses = ipAddresses
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasAllowlistOptions) SetHeaders(param map[string]string) *PostDb2SaasAllowlistOptions {
	options.Headers = param
	return options
}

// PostDb2SaasUserOptions : The PostDb2SaasUser options.
type PostDb2SaasUserOptions struct {
	// CRN deployment id.
	XDeploymentID *string `json:"x-deployment-id" validate:"required"`

	// The id of the User.
	ID *string `json:"id" validate:"required"`

	// Indicates if IAM is enabled.
	Iam *bool `json:"iam" validate:"required"`

	// IBM ID of the User.
	Ibmid *string `json:"ibmid" validate:"required"`

	// The name of the User.
	Name *string `json:"name" validate:"required"`

	// Password of the User.
	Password *string `json:"password" validate:"required"`

	// Role of the User.
	Role *string `json:"role" validate:"required"`

	// Email of the User.
	Email *string `json:"email" validate:"required"`

	// Indicates if the account is locked.
	Locked *string `json:"locked" validate:"required"`

	Authentication *CreateUserAuthentication `json:"authentication" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the PostDb2SaasUserOptions.Role property.
// Role of the User.
const (
	PostDb2SaasUserOptions_Role_Bluadmin = "bluadmin"
	PostDb2SaasUserOptions_Role_Bluuser = "bluuser"
)

// Constants associated with the PostDb2SaasUserOptions.Locked property.
// Indicates if the account is locked.
const (
	PostDb2SaasUserOptions_Locked_No = "no"
	PostDb2SaasUserOptions_Locked_Yes = "yes"
)

// NewPostDb2SaasUserOptions : Instantiate PostDb2SaasUserOptions
func (*Db2saasV1) NewPostDb2SaasUserOptions(xDeploymentID string, id string, iam bool, ibmid string, name string, password string, role string, email string, locked string, authentication *CreateUserAuthentication) *PostDb2SaasUserOptions {
	return &PostDb2SaasUserOptions{
		XDeploymentID: core.StringPtr(xDeploymentID),
		ID: core.StringPtr(id),
		Iam: core.BoolPtr(iam),
		Ibmid: core.StringPtr(ibmid),
		Name: core.StringPtr(name),
		Password: core.StringPtr(password),
		Role: core.StringPtr(role),
		Email: core.StringPtr(email),
		Locked: core.StringPtr(locked),
		Authentication: authentication,
	}
}

// SetXDeploymentID : Allow user to set XDeploymentID
func (_options *PostDb2SaasUserOptions) SetXDeploymentID(xDeploymentID string) *PostDb2SaasUserOptions {
	_options.XDeploymentID = core.StringPtr(xDeploymentID)
	return _options
}

// SetID : Allow user to set ID
func (_options *PostDb2SaasUserOptions) SetID(id string) *PostDb2SaasUserOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIam : Allow user to set Iam
func (_options *PostDb2SaasUserOptions) SetIam(iam bool) *PostDb2SaasUserOptions {
	_options.Iam = core.BoolPtr(iam)
	return _options
}

// SetIbmid : Allow user to set Ibmid
func (_options *PostDb2SaasUserOptions) SetIbmid(ibmid string) *PostDb2SaasUserOptions {
	_options.Ibmid = core.StringPtr(ibmid)
	return _options
}

// SetName : Allow user to set Name
func (_options *PostDb2SaasUserOptions) SetName(name string) *PostDb2SaasUserOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPassword : Allow user to set Password
func (_options *PostDb2SaasUserOptions) SetPassword(password string) *PostDb2SaasUserOptions {
	_options.Password = core.StringPtr(password)
	return _options
}

// SetRole : Allow user to set Role
func (_options *PostDb2SaasUserOptions) SetRole(role string) *PostDb2SaasUserOptions {
	_options.Role = core.StringPtr(role)
	return _options
}

// SetEmail : Allow user to set Email
func (_options *PostDb2SaasUserOptions) SetEmail(email string) *PostDb2SaasUserOptions {
	_options.Email = core.StringPtr(email)
	return _options
}

// SetLocked : Allow user to set Locked
func (_options *PostDb2SaasUserOptions) SetLocked(locked string) *PostDb2SaasUserOptions {
	_options.Locked = core.StringPtr(locked)
	return _options
}

// SetAuthentication : Allow user to set Authentication
func (_options *PostDb2SaasUserOptions) SetAuthentication(authentication *CreateUserAuthentication) *PostDb2SaasUserOptions {
	_options.Authentication = authentication
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostDb2SaasUserOptions) SetHeaders(param map[string]string) *PostDb2SaasUserOptions {
	options.Headers = param
	return options
}

// PutDb2SaasAutoscaleOptions : The PutDb2SaasAutoscale options.
type PutDb2SaasAutoscaleOptions struct {
	// CRN deployment id.
	XDbProfile *string `json:"x-db-profile" validate:"required"`

	// Indicates if automatic scaling is enabled or not.
	AutoScalingEnabled *string `json:"auto_scaling_enabled,omitempty"`

	// Specifies the resource utilization level that triggers an auto-scaling.
	AutoScalingThreshold *int64 `json:"auto_scaling_threshold,omitempty"`

	// Defines the time period over which auto-scaling adjustments are monitored and applied.
	AutoScalingOverTimePeriod *float64 `json:"auto_scaling_over_time_period,omitempty"`

	// Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
	AutoScalingPauseLimit *int64 `json:"auto_scaling_pause_limit,omitempty"`

	// Indicates the maximum number of scaling actions that are allowed within a specified time period.
	AutoScalingAllowPlanLimit *string `json:"auto_scaling_allow_plan_limit,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the PutDb2SaasAutoscaleOptions.AutoScalingEnabled property.
// Indicates if automatic scaling is enabled or not.
const (
	PutDb2SaasAutoscaleOptions_AutoScalingEnabled_False = "false"
	PutDb2SaasAutoscaleOptions_AutoScalingEnabled_True = "true"
)

// Constants associated with the PutDb2SaasAutoscaleOptions.AutoScalingAllowPlanLimit property.
// Indicates the maximum number of scaling actions that are allowed within a specified time period.
const (
	PutDb2SaasAutoscaleOptions_AutoScalingAllowPlanLimit_No = "NO"
	PutDb2SaasAutoscaleOptions_AutoScalingAllowPlanLimit_Yes = "YES"
)

// NewPutDb2SaasAutoscaleOptions : Instantiate PutDb2SaasAutoscaleOptions
func (*Db2saasV1) NewPutDb2SaasAutoscaleOptions(xDbProfile string) *PutDb2SaasAutoscaleOptions {
	return &PutDb2SaasAutoscaleOptions{
		XDbProfile: core.StringPtr(xDbProfile),
	}
}

// SetXDbProfile : Allow user to set XDbProfile
func (_options *PutDb2SaasAutoscaleOptions) SetXDbProfile(xDbProfile string) *PutDb2SaasAutoscaleOptions {
	_options.XDbProfile = core.StringPtr(xDbProfile)
	return _options
}

// SetAutoScalingEnabled : Allow user to set AutoScalingEnabled
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingEnabled(autoScalingEnabled string) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingEnabled = core.StringPtr(autoScalingEnabled)
	return _options
}

// SetAutoScalingThreshold : Allow user to set AutoScalingThreshold
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingThreshold(autoScalingThreshold int64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingThreshold = core.Int64Ptr(autoScalingThreshold)
	return _options
}

// SetAutoScalingOverTimePeriod : Allow user to set AutoScalingOverTimePeriod
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingOverTimePeriod(autoScalingOverTimePeriod float64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingOverTimePeriod = core.Float64Ptr(autoScalingOverTimePeriod)
	return _options
}

// SetAutoScalingPauseLimit : Allow user to set AutoScalingPauseLimit
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingPauseLimit(autoScalingPauseLimit int64) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingPauseLimit = core.Int64Ptr(autoScalingPauseLimit)
	return _options
}

// SetAutoScalingAllowPlanLimit : Allow user to set AutoScalingAllowPlanLimit
func (_options *PutDb2SaasAutoscaleOptions) SetAutoScalingAllowPlanLimit(autoScalingAllowPlanLimit string) *PutDb2SaasAutoscaleOptions {
	_options.AutoScalingAllowPlanLimit = core.StringPtr(autoScalingAllowPlanLimit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PutDb2SaasAutoscaleOptions) SetHeaders(param map[string]string) *PutDb2SaasAutoscaleOptions {
	options.Headers = param
	return options
}

// SuccessAutoScaling : The details of the autoscale.
type SuccessAutoScaling struct {
	// Indicates the maximum number of scaling actions that are allowed within a specified time period.
	AutoScalingAllowPlanLimit *bool `json:"auto_scaling_allow_plan_limit" validate:"required"`

	// Indicates if automatic scaling is enabled or not.
	AutoScalingEnabled *bool `json:"auto_scaling_enabled" validate:"required"`

	// The maximum limit for automatically increasing storage capacity to handle growing data needs.
	AutoScalingMaxStorage *int64 `json:"auto_scaling_max_storage" validate:"required"`

	// Defines the time period over which auto-scaling adjustments are monitored and applied.
	AutoScalingOverTimePeriod *int64 `json:"auto_scaling_over_time_period" validate:"required"`

	// Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
	AutoScalingPauseLimit *int64 `json:"auto_scaling_pause_limit" validate:"required"`

	// Specifies the resource utilization level that triggers an auto-scaling.
	AutoScalingThreshold *int64 `json:"auto_scaling_threshold" validate:"required"`

	// Specifies the unit of measurement for storage capacity.
	StorageUnit *string `json:"storage_unit" validate:"required"`

	// Represents the percentage of total storage capacity currently in use.
	StorageUtilizationPercentage *int64 `json:"storage_utilization_percentage" validate:"required"`

	// Indicates whether a system or service can automatically adjust resources based on demand.
	SupportAutoScaling *bool `json:"support_auto_scaling" validate:"required"`
}

// UnmarshalSuccessAutoScaling unmarshals an instance of SuccessAutoScaling from the specified map of raw messages.
func UnmarshalSuccessAutoScaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessAutoScaling)
	err = core.UnmarshalPrimitive(m, "auto_scaling_allow_plan_limit", &obj.AutoScalingAllowPlanLimit)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_allow_plan_limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_enabled", &obj.AutoScalingEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_max_storage", &obj.AutoScalingMaxStorage)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_max_storage-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_over_time_period", &obj.AutoScalingOverTimePeriod)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_over_time_period-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_pause_limit", &obj.AutoScalingPauseLimit)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_pause_limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "auto_scaling_threshold", &obj.AutoScalingThreshold)
	if err != nil {
		err = core.SDKErrorf(err, "", "auto_scaling_threshold-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_unit", &obj.StorageUnit)
	if err != nil {
		err = core.SDKErrorf(err, "", "storage_unit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_utilization_percentage", &obj.StorageUtilizationPercentage)
	if err != nil {
		err = core.SDKErrorf(err, "", "storage_utilization_percentage-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "support_auto_scaling", &obj.SupportAutoScaling)
	if err != nil {
		err = core.SDKErrorf(err, "", "support_auto_scaling-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfo : Responds with JSON of the connection information for the Db2 SaaS Instance.
type SuccessConnectionInfo struct {
	Public *SuccessConnectionInfoPublic `json:"public,omitempty"`

	Private *SuccessConnectionInfoPrivate `json:"private,omitempty"`
}

// UnmarshalSuccessConnectionInfo unmarshals an instance of SuccessConnectionInfo from the specified map of raw messages.
func UnmarshalSuccessConnectionInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfo)
	err = core.UnmarshalModel(m, "public", &obj.Public, UnmarshalSuccessConnectionInfoPublic)
	if err != nil {
		err = core.SDKErrorf(err, "", "public-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "private", &obj.Private, UnmarshalSuccessConnectionInfoPrivate)
	if err != nil {
		err = core.SDKErrorf(err, "", "private-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfoPrivate : SuccessConnectionInfoPrivate struct
type SuccessConnectionInfoPrivate struct {
	Hostname *string `json:"hostname,omitempty"`

	DatabaseName *string `json:"databaseName,omitempty"`

	SslPort *string `json:"sslPort,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	PrivateServiceName *string `json:"private_serviceName,omitempty"`

	CloudServiceOffering *string `json:"cloud_service_offering,omitempty"`

	VpeServiceCrn *string `json:"vpe_service_crn,omitempty"`

	DbVpcEndpointService *string `json:"db_vpc_endpoint_service,omitempty"`
}

// UnmarshalSuccessConnectionInfoPrivate unmarshals an instance of SuccessConnectionInfoPrivate from the specified map of raw messages.
func UnmarshalSuccessConnectionInfoPrivate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfoPrivate)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseName", &obj.DatabaseName)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sslPort", &obj.SslPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "sslPort-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		err = core.SDKErrorf(err, "", "ssl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseVersion", &obj.DatabaseVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseVersion-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "private_serviceName", &obj.PrivateServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "private_serviceName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cloud_service_offering", &obj.CloudServiceOffering)
	if err != nil {
		err = core.SDKErrorf(err, "", "cloud_service_offering-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "vpe_service_crn", &obj.VpeServiceCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "vpe_service_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "db_vpc_endpoint_service", &obj.DbVpcEndpointService)
	if err != nil {
		err = core.SDKErrorf(err, "", "db_vpc_endpoint_service-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessConnectionInfoPublic : SuccessConnectionInfoPublic struct
type SuccessConnectionInfoPublic struct {
	Hostname *string `json:"hostname,omitempty"`

	DatabaseName *string `json:"databaseName,omitempty"`

	SslPort *string `json:"sslPort,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	DatabaseVersion *string `json:"databaseVersion,omitempty"`
}

// UnmarshalSuccessConnectionInfoPublic unmarshals an instance of SuccessConnectionInfoPublic from the specified map of raw messages.
func UnmarshalSuccessConnectionInfoPublic(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessConnectionInfoPublic)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		err = core.SDKErrorf(err, "", "hostname-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseName", &obj.DatabaseName)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseName-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "sslPort", &obj.SslPort)
	if err != nil {
		err = core.SDKErrorf(err, "", "sslPort-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		err = core.SDKErrorf(err, "", "ssl-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "databaseVersion", &obj.DatabaseVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "databaseVersion-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetAllowlistIPs : Success response of get allowlist IPs.
type SuccessGetAllowlistIPs struct {
	// List of IP addresses.
	IpAddresses []IpAddress `json:"ip_addresses" validate:"required"`
}

// UnmarshalSuccessGetAllowlistIPs unmarshals an instance of SuccessGetAllowlistIPs from the specified map of raw messages.
func UnmarshalSuccessGetAllowlistIPs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetAllowlistIPs)
	err = core.UnmarshalModel(m, "ip_addresses", &obj.IpAddresses, UnmarshalIpAddress)
	if err != nil {
		err = core.SDKErrorf(err, "", "ip_addresses-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserByID : The details of the users.
type SuccessGetUserByID struct {
	// User's DV role.
	DvRole *string `json:"dvRole" validate:"required"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata" validate:"required"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid" validate:"required"`

	// Role assigned to the user.
	Role *string `json:"role" validate:"required"`

	// IAM ID for the user.
	Iamid *string `json:"iamid" validate:"required"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions" validate:"required"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean" validate:"required"`

	// User's password.
	Password *string `json:"password" validate:"required"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam" validate:"required"`

	// The display name of the user.
	Name *string `json:"name" validate:"required"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid" validate:"required"`

	// Unique identifier for the user.
	ID *string `json:"id" validate:"required"`

	// Account lock status for the user.
	Locked *string `json:"locked" validate:"required"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg" validate:"required"`

	// Email address of the user.
	Email *string `json:"email" validate:"required"`

	// Authentication details for the user.
	Authentication *SuccessGetUserByIDAuthentication `json:"authentication" validate:"required"`
}

// Constants associated with the SuccessGetUserByID.Role property.
// Role assigned to the user.
const (
	SuccessGetUserByID_Role_Bluadmin = "bluadmin"
	SuccessGetUserByID_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessGetUserByID.Locked property.
// Account lock status for the user.
const (
	SuccessGetUserByID_Locked_No = "no"
	SuccessGetUserByID_Locked_Yes = "yes"
)

// UnmarshalSuccessGetUserByID unmarshals an instance of SuccessGetUserByID from the specified map of raw messages.
func UnmarshalSuccessGetUserByID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserByID)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessGetUserByIDAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserByIDAuthentication : Authentication details for the user.
type SuccessGetUserByIDAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessGetUserByIDAuthentication unmarshals an instance of SuccessGetUserByIDAuthentication from the specified map of raw messages.
func UnmarshalSuccessGetUserByIDAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserByIDAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfo : Success response of get user.
type SuccessGetUserInfo struct {
	// The total number of resources.
	Count *int64 `json:"count" validate:"required"`

	// A list of user resource.
	Resources []SuccessGetUserInfoResourcesItem `json:"resources" validate:"required"`
}

// UnmarshalSuccessGetUserInfo unmarshals an instance of SuccessGetUserInfo from the specified map of raw messages.
func UnmarshalSuccessGetUserInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfo)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		err = core.SDKErrorf(err, "", "count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalSuccessGetUserInfoResourcesItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "resources-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfoResourcesItem : SuccessGetUserInfoResourcesItem struct
type SuccessGetUserInfoResourcesItem struct {
	// User's DV role.
	DvRole *string `json:"dvRole,omitempty"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid,omitempty"`

	// Role assigned to the user.
	Role *string `json:"role,omitempty"`

	// IAM ID for the user.
	Iamid *string `json:"iamid,omitempty"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions,omitempty"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean,omitempty"`

	// User's password.
	Password *string `json:"password,omitempty"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam,omitempty"`

	// The display name of the user.
	Name *string `json:"name,omitempty"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid,omitempty"`

	// Unique identifier for the user.
	ID *string `json:"id,omitempty"`

	// Account lock status for the user.
	Locked *string `json:"locked,omitempty"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg,omitempty"`

	// Email address of the user.
	Email *string `json:"email,omitempty"`

	// Authentication details for the user.
	Authentication *SuccessGetUserInfoResourcesItemAuthentication `json:"authentication,omitempty"`
}

// Constants associated with the SuccessGetUserInfoResourcesItem.Role property.
// Role assigned to the user.
const (
	SuccessGetUserInfoResourcesItem_Role_Bluadmin = "bluadmin"
	SuccessGetUserInfoResourcesItem_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessGetUserInfoResourcesItem.Locked property.
// Account lock status for the user.
const (
	SuccessGetUserInfoResourcesItem_Locked_No = "no"
	SuccessGetUserInfoResourcesItem_Locked_Yes = "yes"
)

// UnmarshalSuccessGetUserInfoResourcesItem unmarshals an instance of SuccessGetUserInfoResourcesItem from the specified map of raw messages.
func UnmarshalSuccessGetUserInfoResourcesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfoResourcesItem)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessGetUserInfoResourcesItemAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessGetUserInfoResourcesItemAuthentication : Authentication details for the user.
type SuccessGetUserInfoResourcesItemAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessGetUserInfoResourcesItemAuthentication unmarshals an instance of SuccessGetUserInfoResourcesItemAuthentication from the specified map of raw messages.
func UnmarshalSuccessGetUserInfoResourcesItemAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessGetUserInfoResourcesItemAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessPostAllowedlistIPs : Success response of post allowlist IPs.
type SuccessPostAllowedlistIPs struct {
	// status of the post allowlist IPs request.
	Status *string `json:"status" validate:"required"`
}

// UnmarshalSuccessPostAllowedlistIPs unmarshals an instance of SuccessPostAllowedlistIPs from the specified map of raw messages.
func UnmarshalSuccessPostAllowedlistIPs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessPostAllowedlistIPs)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUpdateAutoScale : Response of successful updation of scaling configurations.
type SuccessUpdateAutoScale struct {
	// Indicates the message of the updation.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalSuccessUpdateAutoScale unmarshals an instance of SuccessUpdateAutoScale from the specified map of raw messages.
func UnmarshalSuccessUpdateAutoScale(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUpdateAutoScale)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUserResponse : The details of the users.
type SuccessUserResponse struct {
	// User's DV role.
	DvRole *string `json:"dvRole" validate:"required"`

	// Metadata associated with the user.
	Metadata map[string]interface{} `json:"metadata" validate:"required"`

	// Formatted IBM ID.
	FormatedIbmid *string `json:"formatedIbmid" validate:"required"`

	// Role assigned to the user.
	Role *string `json:"role" validate:"required"`

	// IAM ID for the user.
	Iamid *string `json:"iamid" validate:"required"`

	// List of allowed actions of the user.
	PermittedActions []string `json:"permittedActions" validate:"required"`

	// Indicates if the user account has no issues.
	AllClean *bool `json:"allClean" validate:"required"`

	// User's password.
	Password *string `json:"password" validate:"required"`

	// Indicates if IAM is enabled or not.
	Iam *bool `json:"iam" validate:"required"`

	// The display name of the user.
	Name *string `json:"name" validate:"required"`

	// IBM ID of the user.
	Ibmid *string `json:"ibmid" validate:"required"`

	// Unique identifier for the user.
	ID *string `json:"id" validate:"required"`

	// Account lock status for the user.
	Locked *string `json:"locked" validate:"required"`

	// Initial error message.
	InitErrorMsg *string `json:"initErrorMsg" validate:"required"`

	// Email address of the user.
	Email *string `json:"email" validate:"required"`

	// Authentication details for the user.
	Authentication *SuccessUserResponseAuthentication `json:"authentication" validate:"required"`
}

// Constants associated with the SuccessUserResponse.Role property.
// Role assigned to the user.
const (
	SuccessUserResponse_Role_Bluadmin = "bluadmin"
	SuccessUserResponse_Role_Bluuser = "bluuser"
)

// Constants associated with the SuccessUserResponse.Locked property.
// Account lock status for the user.
const (
	SuccessUserResponse_Locked_No = "no"
	SuccessUserResponse_Locked_Yes = "yes"
)

// UnmarshalSuccessUserResponse unmarshals an instance of SuccessUserResponse from the specified map of raw messages.
func UnmarshalSuccessUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUserResponse)
	err = core.UnmarshalPrimitive(m, "dvRole", &obj.DvRole)
	if err != nil {
		err = core.SDKErrorf(err, "", "dvRole-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		err = core.SDKErrorf(err, "", "metadata-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "formatedIbmid", &obj.FormatedIbmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "formatedIbmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iamid", &obj.Iamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "permittedActions", &obj.PermittedActions)
	if err != nil {
		err = core.SDKErrorf(err, "", "permittedActions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "allClean", &obj.AllClean)
	if err != nil {
		err = core.SDKErrorf(err, "", "allClean-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		err = core.SDKErrorf(err, "", "password-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "iam", &obj.Iam)
	if err != nil {
		err = core.SDKErrorf(err, "", "iam-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmid", &obj.Ibmid)
	if err != nil {
		err = core.SDKErrorf(err, "", "ibmid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		err = core.SDKErrorf(err, "", "locked-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "initErrorMsg", &obj.InitErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "initErrorMsg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		err = core.SDKErrorf(err, "", "email-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalSuccessUserResponseAuthentication)
	if err != nil {
		err = core.SDKErrorf(err, "", "authentication-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SuccessUserResponseAuthentication : Authentication details for the user.
type SuccessUserResponseAuthentication struct {
	// Authentication method.
	Method *string `json:"method" validate:"required"`

	// Policy ID of authentication.
	PolicyID *string `json:"policy_id" validate:"required"`
}

// UnmarshalSuccessUserResponseAuthentication unmarshals an instance of SuccessUserResponseAuthentication from the specified map of raw messages.
func UnmarshalSuccessUserResponseAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SuccessUserResponseAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		err = core.SDKErrorf(err, "", "method-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "policy_id", &obj.PolicyID)
	if err != nil {
		err = core.SDKErrorf(err, "", "policy_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
