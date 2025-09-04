/**
 * (C) Copyright IBM Corp. 2025.
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

package db2saasv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`Db2saasV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(db2saasService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(db2saasService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				URL: "https://db2saasv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(db2saasService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DB2SAAS_URL": "https://db2saasv1/api",
				"DB2SAAS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				})
				Expect(db2saasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := db2saasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != db2saasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
					URL: "https://testService/api",
				})
				Expect(db2saasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := db2saasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != db2saasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				})
				err := db2saasService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := db2saasService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != db2saasService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(db2saasService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(db2saasService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DB2SAAS_URL": "https://db2saasv1/api",
				"DB2SAAS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(db2saasService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DB2SAAS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			db2saasService, serviceErr := db2saasv1.NewDb2saasV1UsingExternalConfig(&db2saasv1.Db2saasV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(db2saasService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = db2saasv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := db2saasv1.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://us-south.db2.saas.ibm.com/dbapi/v4"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := db2saasv1.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions) - Operation response error`, func() {
		getDb2SaasConnectionInfoPath := "/connectioninfo/crn%253Av1%253Astaging%253Apublic%253Adashdb-for-transactions%253Aus-south%253Aa%252Fe7e3e87b512f474381c0684a5ecbba03%253A69db420f-33d5-4953-8bd8-1950abd356f6%253A%253A"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasConnectionInfo with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions)`, func() {
		getDb2SaasConnectionInfoPath := "/connectioninfo/crn%253Av1%253Astaging%253Apublic%253Adashdb-for-transactions%253Aus-south%253Aa%252Fe7e3e87b512f474381c0684a5ecbba03%253A69db420f-33d5-4953-8bd8-1950abd356f6%253A%253A"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0", "private_serviceName": "us-south-private.db2oc.test.saas.ibm.com:32764", "cloud_service_offering": "dashdb-for-transactions", "vpe_service_crn": "crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud", "db_vpc_endpoint_service": "feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679"}}`)
				}))
			})
			It(`Invoke GetDb2SaasConnectionInfo successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasConnectionInfoWithContext(ctx, getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasConnectionInfoWithContext(ctx, getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0", "private_serviceName": "us-south-private.db2oc.test.saas.ibm.com:32764", "cloud_service_offering": "dashdb-for-transactions", "vpe_service_crn": "crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud", "db_vpc_endpoint_service": "feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679"}}`)
				}))
			})
			It(`Invoke GetDb2SaasConnectionInfo successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasConnectionInfo with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasConnectionInfoOptions model with no property values
				getDb2SaasConnectionInfoOptionsModelNew := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasConnectionInfo successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				getDb2SaasConnectionInfoOptionsModel := new(db2saasv1.GetDb2SaasConnectionInfoOptions)
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions) - Operation response error`, func() {
		postDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDb2SaasAllowlist with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions)`, func() {
		postDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasAllowlist successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PostDb2SaasAllowlistWithContext(ctx, postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PostDb2SaasAllowlistWithContext(ctx, postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasAllowlist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PostDb2SaasAllowlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostDb2SaasAllowlist with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDb2SaasAllowlistOptions model with no property values
				postDb2SaasAllowlistOptionsModelNew := new(db2saasv1.PostDb2SaasAllowlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostDb2SaasAllowlist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				postDb2SaasAllowlistOptionsModel := new(db2saasv1.PostDb2SaasAllowlistOptions)
				postDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions) - Operation response error`, func() {
		getDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasAllowlist with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAllowlistOptions model
				getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions)`, func() {
		getDb2SaasAllowlistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				}))
			})
			It(`Invoke GetDb2SaasAllowlist successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasAllowlistOptions model
				getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasAllowlistWithContext(ctx, getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasAllowlistWithContext(ctx, getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAllowlistPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				}))
			})
			It(`Invoke GetDb2SaasAllowlist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasAllowlist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasAllowlistOptions model
				getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasAllowlist with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAllowlistOptions model
				getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasAllowlistOptions model with no property values
				getDb2SaasAllowlistOptionsModelNew := new(db2saasv1.GetDb2SaasAllowlistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasAllowlist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAllowlistOptions model
				getDb2SaasAllowlistOptionsModel := new(db2saasv1.GetDb2SaasAllowlistOptions)
				getDb2SaasAllowlistOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions) - Operation response error`, func() {
		postDb2SaasUserPath := "/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDb2SaasUser with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PostDb2SaasUserOptions model
				postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions)`, func() {
		postDb2SaasUserPath := "/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke PostDb2SaasUser successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PostDb2SaasUserOptions model
				postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PostDb2SaasUserWithContext(ctx, postDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PostDb2SaasUserWithContext(ctx, postDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasUserPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke PostDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PostDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PostDb2SaasUserOptions model
				postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostDb2SaasUser with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PostDb2SaasUserOptions model
				postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDb2SaasUserOptions model with no property values
				postDb2SaasUserOptionsModelNew := new(db2saasv1.PostDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PostDb2SaasUserOptions model
				postDb2SaasUserOptionsModel := new(db2saasv1.PostDb2SaasUserOptions)
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				postDb2SaasUserOptionsModel.Iam = core.BoolPtr(false)
				postDb2SaasUserOptionsModel.Ibmid = core.StringPtr("test-ibm-id")
				postDb2SaasUserOptionsModel.Name = core.StringPtr("test_user")
				postDb2SaasUserOptionsModel.Password = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.Role = core.StringPtr("bluuser")
				postDb2SaasUserOptionsModel.Email = core.StringPtr("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.Locked = core.StringPtr("no")
				postDb2SaasUserOptionsModel.Authentication = createUserAuthenticationModel
				postDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PostDb2SaasUser(postDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions) - Operation response error`, func() {
		getDb2SaasUserPath := "/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasUser with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasUserOptions model
				getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions)`, func() {
		getDb2SaasUserPath := "/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 1, "resources": [{"dvRole": "test-role", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "test-formated-ibm-id", "role": "bluadmin", "iamid": "test-iam-id", "permittedActions": ["PermittedActions"], "allClean": false, "password": "nd!@aegr63@989hcRFTcdcs63", "iam": false, "name": "admin", "ibmid": "test-ibm-id", "id": "admin", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "internal", "policy_id": "Default"}}]}`)
				}))
			})
			It(`Invoke GetDb2SaasUser successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasUserOptions model
				getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasUserWithContext(ctx, getDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasUserWithContext(ctx, getDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 1, "resources": [{"dvRole": "test-role", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "test-formated-ibm-id", "role": "bluadmin", "iamid": "test-iam-id", "permittedActions": ["PermittedActions"], "allClean": false, "password": "nd!@aegr63@989hcRFTcdcs63", "iam": false, "name": "admin", "ibmid": "test-ibm-id", "id": "admin", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "internal", "policy_id": "Default"}}]}`)
				}))
			})
			It(`Invoke GetDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasUserOptions model
				getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasUser with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasUserOptions model
				getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasUserOptions model with no property values
				getDb2SaasUserOptionsModelNew := new(db2saasv1.GetDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasUserOptions model
				getDb2SaasUserOptionsModel := new(db2saasv1.GetDb2SaasUserOptions)
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasUser(getDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDb2SaasUser(putDb2SaasUserOptions *PutDb2SaasUserOptions) - Operation response error`, func() {
		putDb2SaasUserPath := "/users/test-user"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasUserPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutDb2SaasUser with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PutDb2SaasUserOptions model
				putDb2SaasUserOptionsModel := new(db2saasv1.PutDb2SaasUserOptions)
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewIam = core.BoolPtr(false)
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDb2SaasUser(putDb2SaasUserOptions *PutDb2SaasUserOptions)`, func() {
		putDb2SaasUserPath := "/users/test-user"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasUserPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke PutDb2SaasUser successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PutDb2SaasUserOptions model
				putDb2SaasUserOptionsModel := new(db2saasv1.PutDb2SaasUserOptions)
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewIam = core.BoolPtr(false)
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PutDb2SaasUserWithContext(ctx, putDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PutDb2SaasUserWithContext(ctx, putDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasUserPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke PutDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PutDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PutDb2SaasUserOptions model
				putDb2SaasUserOptionsModel := new(db2saasv1.PutDb2SaasUserOptions)
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewIam = core.BoolPtr(false)
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutDb2SaasUser with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PutDb2SaasUserOptions model
				putDb2SaasUserOptionsModel := new(db2saasv1.PutDb2SaasUserOptions)
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewIam = core.BoolPtr(false)
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutDb2SaasUserOptions model with no property values
				putDb2SaasUserOptionsModelNew := new(db2saasv1.PutDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")

				// Construct an instance of the PutDb2SaasUserOptions model
				putDb2SaasUserOptionsModel := new(db2saasv1.PutDb2SaasUserOptions)
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewIam = core.BoolPtr(false)
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PutDb2SaasUser(putDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions)`, func() {
		deleteDb2SaasUserPath := "/users/test-user"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDb2SaasUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := db2saasService.DeleteDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDb2SaasUserOptions model
				deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDb2SaasUser with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the DeleteDb2SaasUserOptions model
				deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDb2SaasUserOptions model with no property values
				deleteDb2SaasUserOptionsModelNew := new(db2saasv1.DeleteDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) - Operation response error`, func() {
		getbyidDb2SaasUserPath := "/users/test-user"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetbyidDb2SaasUser with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetbyidDb2SaasUserOptions model
				getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions)`, func() {
		getbyidDb2SaasUserPath := "/users/test-user"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke GetbyidDb2SaasUser successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetbyidDb2SaasUserOptions model
				getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetbyidDb2SaasUserWithContext(ctx, getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetbyidDb2SaasUserWithContext(ctx, getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"dvRole": "DvRole", "metadata": {"anyKey": "anyValue"}, "formatedIbmid": "FormatedIbmid", "role": "bluadmin", "iamid": "Iamid", "permittedActions": ["PermittedActions"], "allClean": true, "password": "Password", "iam": false, "name": "Name", "ibmid": "Ibmid", "id": "ID", "locked": "no", "initErrorMsg": "InitErrorMsg", "email": "user@host.org", "authentication": {"method": "Method", "policy_id": "PolicyID"}}`)
				}))
			})
			It(`Invoke GetbyidDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetbyidDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetbyidDb2SaasUserOptions model
				getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetbyidDb2SaasUser with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetbyidDb2SaasUserOptions model
				getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetbyidDb2SaasUserOptions model with no property values
				getbyidDb2SaasUserOptionsModelNew := new(db2saasv1.GetbyidDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetbyidDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetbyidDb2SaasUserOptions model
				getbyidDb2SaasUserOptionsModel := new(db2saasv1.GetbyidDb2SaasUserOptions)
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				getbyidDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions) - Operation response error`, func() {
		putDb2SaasAutoscalePath := "/manage/scaling/auto"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PutDb2SaasAutoscale with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions)`, func() {
		putDb2SaasAutoscalePath := "/manage/scaling/auto"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke PutDb2SaasAutoscale successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PutDb2SaasAutoscaleWithContext(ctx, putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PutDb2SaasAutoscaleWithContext(ctx, putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Message"}`)
				}))
			})
			It(`Invoke PutDb2SaasAutoscale successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PutDb2SaasAutoscale(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PutDb2SaasAutoscale with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PutDb2SaasAutoscaleOptions model with no property values
				putDb2SaasAutoscaleOptionsModelNew := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutDb2SaasAutoscale successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				putDb2SaasAutoscaleOptionsModel := new(db2saasv1.PutDb2SaasAutoscaleOptions)
				putDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled = core.StringPtr("true")
				putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold = core.Int64Ptr(int64(90))
				putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod = core.Float64Ptr(float64(5))
				putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit = core.Int64Ptr(int64(70))
				putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit = core.StringPtr("YES")
				putDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions) - Operation response error`, func() {
		getDb2SaasAutoscalePath := "/manage/scaling/auto"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasAutoscale with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions)`, func() {
		getDb2SaasAutoscalePath := "/manage/scaling/auto"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auto_scaling_allow_plan_limit": false, "auto_scaling_enabled": true, "auto_scaling_max_storage": 21, "auto_scaling_over_time_period": 25, "auto_scaling_pause_limit": 21, "auto_scaling_threshold": 20, "storage_unit": "StorageUnit", "storage_utilization_percentage": 28, "support_auto_scaling": true}`)
				}))
			})
			It(`Invoke GetDb2SaasAutoscale successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasAutoscaleWithContext(ctx, getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasAutoscaleWithContext(ctx, getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"auto_scaling_allow_plan_limit": false, "auto_scaling_enabled": true, "auto_scaling_max_storage": 21, "auto_scaling_over_time_period": 25, "auto_scaling_pause_limit": 21, "auto_scaling_threshold": 20, "storage_unit": "StorageUnit", "storage_utilization_percentage": 28, "support_auto_scaling": true}`)
				}))
			})
			It(`Invoke GetDb2SaasAutoscale successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasAutoscale(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasAutoscale with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasAutoscaleOptions model with no property values
				getDb2SaasAutoscaleOptionsModelNew := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasAutoscale successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				getDb2SaasAutoscaleOptionsModel := new(db2saasv1.GetDb2SaasAutoscaleOptions)
				getDb2SaasAutoscaleOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions) - Operation response error`, func() {
		postDb2SaasDbConfigurationPath := "/manage/deployments/custom_setting"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDb2SaasDbConfiguration with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				postDb2SaasDbConfigurationOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				postDb2SaasDbConfigurationOptionsModel.Dbm = createCustomSettingsDbmModel
				postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions)`, func() {
		postDb2SaasDbConfigurationPath := "/manage/deployments/custom_setting"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasDbConfiguration successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				postDb2SaasDbConfigurationOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				postDb2SaasDbConfigurationOptionsModel.Dbm = createCustomSettingsDbmModel
				postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PostDb2SaasDbConfigurationWithContext(ctx, postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PostDb2SaasDbConfigurationWithContext(ctx, postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasDbConfigurationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasDbConfiguration successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				postDb2SaasDbConfigurationOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				postDb2SaasDbConfigurationOptionsModel.Dbm = createCustomSettingsDbmModel
				postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostDb2SaasDbConfiguration with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				postDb2SaasDbConfigurationOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				postDb2SaasDbConfigurationOptionsModel.Dbm = createCustomSettingsDbmModel
				postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDb2SaasDbConfigurationOptions model with no property values
				postDb2SaasDbConfigurationOptionsModelNew := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostDb2SaasDbConfiguration successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				postDb2SaasDbConfigurationOptionsModel := new(db2saasv1.PostDb2SaasDbConfigurationOptions)
				postDb2SaasDbConfigurationOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.Registry = createCustomSettingsRegistryModel
				postDb2SaasDbConfigurationOptionsModel.Db = createCustomSettingsDbModel
				postDb2SaasDbConfigurationOptionsModel.Dbm = createCustomSettingsDbmModel
				postDb2SaasDbConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions) - Operation response error`, func() {
		getDb2SaasTuneableParamPath := "/manage/tuneable_param"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasTuneableParam with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions)`, func() {
		getDb2SaasTuneableParamPath := "/manage/tuneable_param"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tuneable_param": {"db": {"ACT_SORTMEM_LIMIT": "'NONE', 'range(10, 100)'", "ALT_COLLATE": "'NULL', 'IDENTITY_16BIT'", "APPGROUP_MEM_SZ": "'range(1, 1000000)'", "APPLHEAPSZ": "'AUTOMATIC', 'range(16, 2147483647)'", "APPL_MEMORY": "'AUTOMATIC', 'range(128, 4294967295)'", "APP_CTL_HEAP_SZ": "'range(1, 64000)'", "ARCHRETRYDELAY": "'range(0, 65535)'", "AUTHN_CACHE_DURATION": "'range(1,10000)'", "AUTORESTART": "'ON', 'OFF'", "AUTO_CG_STATS": "'ON', 'OFF'", "AUTO_MAINT": "'ON', 'OFF'", "AUTO_REORG": "'ON', 'OFF'", "AUTO_REVAL": "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'", "AUTO_RUNSTATS": "'ON', 'OFF'", "AUTO_SAMPLING": "'ON', 'OFF'", "AUTO_STATS_VIEWS": "'ON', 'OFF'", "AUTO_STMT_STATS": "'ON', 'OFF'", "AUTO_TBL_MAINT": "'ON', 'OFF'", "AVG_APPLS": "'-'", "CATALOGCACHE_SZ": "'-'", "CHNGPGS_THRESH": "'range(5,99)'", "CUR_COMMIT": "'ON, AVAILABLE, DISABLED'", "DATABASE_MEMORY": "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'", "DBHEAP": "'AUTOMATIC', 'range(32, 2147483647)'", "DB_COLLNAME": "'-'", "DB_MEM_THRESH": "'range(0, 100)'", "DDL_COMPRESSION_DEF": "'YES', 'NO'", "DDL_CONSTRAINT_DEF": "'YES', 'NO'", "DECFLT_ROUNDING": "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'", "DEC_ARITHMETIC": "'-'", "DEC_TO_CHAR_FMT": "'NEW', 'V95'", "DFT_DEGREE": "'-1', 'ANY', 'range(1, 32767)'", "DFT_EXTENT_SZ": "'range(2, 256)'", "DFT_LOADREC_SES": "'range(1, 30000)'", "DFT_MTTB_TYPES": "'-'", "DFT_PREFETCH_SZ": "'range(0, 32767)', 'AUTOMATIC'", "DFT_QUERYOPT": "'range(0, 9)'", "DFT_REFRESH_AGE": "'-'", "DFT_SCHEMAS_DCC": "'YES', 'NO'", "DFT_SQLMATHWARN": "'YES', 'NO'", "DFT_TABLE_ORG": "'COLUMN', 'ROW'", "DLCHKTIME": "'range(1000, 600000)'", "ENABLE_XMLCHAR": "'YES', 'NO'", "EXTENDED_ROW_SZ": "'ENABLE', 'DISABLE'", "GROUPHEAP_RATIO": "'range(1, 99)'", "INDEXREC": "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'", "LARGE_AGGREGATION": "'YES', 'NO'", "LOCKLIST": "'AUTOMATIC', 'range(4, 134217728)'", "LOCKTIMEOUT": "'-1', 'range(0, 32767)'", "LOGINDEXBUILD": "'ON', 'OFF'", "LOG_APPL_INFO": "'YES', 'NO'", "LOG_DDL_STMTS": "'YES', 'NO'", "LOG_DISK_CAP": "'0', '-1', 'range(1, 2147483647)'", "MAXAPPLS": "'range(1, 60000)'", "MAXFILOP": "'range(64, 61440)'", "MAXLOCKS": "'AUTOMATIC', 'range(1, 100)'", "MIN_DEC_DIV_3": "'YES', 'NO'", "MON_ACT_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_DEADLOCK": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LCK_MSG_LVL": "'range(0, 3)'", "MON_LOCKTIMEOUT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LOCKWAIT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LW_THRESH": "'range(1000, 4294967295)'", "MON_OBJ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_PKGLIST_SZ": "'range(0, 1024)'", "MON_REQ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_RTN_DATA": "'NONE', 'BASE'", "MON_RTN_EXECLIST": "'OFF', 'ON'", "MON_UOW_DATA": "'NONE', 'BASE'", "MON_UOW_EXECLIST": "'ON', 'OFF'", "MON_UOW_PKGLIST": "'OFF', 'ON'", "NCHAR_MAPPING": "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'", "NUM_FREQVALUES": "'range(0, 32767)'", "NUM_IOCLEANERS": "'AUTOMATIC', 'range(0, 255)'", "NUM_IOSERVERS": "'AUTOMATIC', 'range(1, 255)'", "NUM_LOG_SPAN": "'range(0, 65535)'", "NUM_QUANTILES": "'range(0, 32767)'", "OPT_BUFFPAGE": "'-'", "OPT_DIRECT_WRKLD": "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'", "OPT_LOCKLIST": "'-'", "OPT_MAXLOCKS": "'-'", "OPT_SORTHEAP": "'-'", "PAGE_AGE_TRGT_GCR": "'range(1, 65535)'", "PAGE_AGE_TRGT_MCR": "'range(1, 65535)'", "PCKCACHESZ": "'AUTOMATIC', '-1', 'range(32, 2147483646)'", "PL_STACK_TRACE": "'NONE', 'ALL', 'UNHANDLED'", "SELF_TUNING_MEM": "'ON', 'OFF'", "SEQDETECT": "'YES', 'NO'", "SHEAPTHRES_SHR": "'AUTOMATIC', 'range(250, 2147483647)'", "SOFTMAX": "'-'", "SORTHEAP": "'AUTOMATIC', 'range(16, 4294967295)'", "SQL_CCFLAGS": "'-'", "STAT_HEAP_SZ": "'AUTOMATIC', 'range(1096, 2147483647)'", "STMTHEAP": "'AUTOMATIC', 'range(128, 2147483647)'", "STMT_CONC": "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'", "STRING_UNITS": "'SYSTEM', 'CODEUNITS32'", "SYSTIME_PERIOD_ADJ": "'NO', 'YES'", "TRACKMOD": "'YES', 'NO'", "UTIL_HEAP_SZ": "'AUTOMATIC', 'range(16, 2147483647)'", "WLM_ADMISSION_CTRL": "'YES', 'NO'", "WLM_AGENT_LOAD_TRGT": "'AUTOMATIC', 'range(1, 65535)'", "WLM_CPU_LIMIT": "'range(0, 100)'", "WLM_CPU_SHARES": "'range(1, 65535)'", "WLM_CPU_SHARE_MODE": "'HARD', 'SOFT'"}, "dbm": {"COMM_BANDWIDTH": "'range(0.1, 100000)', '-1'", "CPUSPEED": "'range(0.0000000001, 1)', '-1'", "DFT_MON_BUFPOOL": "'ON', 'OFF'", "DFT_MON_LOCK": "'ON', 'OFF'", "DFT_MON_SORT": "'ON', 'OFF'", "DFT_MON_STMT": "'ON', 'OFF'", "DFT_MON_TABLE": "'ON', 'OFF'", "DFT_MON_TIMESTAMP": "'ON', 'OFF'", "DFT_MON_UOW": "'ON', 'OFF'", "DIAGLEVEL": "'range(0, 4)'", "FEDERATED_ASYNC": "'range(0, 32767)', '-1', 'ANY'", "INDEXREC": "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'", "INTRA_PARALLEL": "'SYSTEM', 'NO', 'YES'", "KEEPFENCED": "'YES', 'NO'", "MAX_CONNRETRIES": "'range(0, 100)'", "MAX_QUERYDEGREE": "'range(1, 32767)', '-1', 'ANY'", "MON_HEAP_SZ": "'range(0, 2147483647)', 'AUTOMATIC'", "MULTIPARTSIZEMB": "'range(5, 5120)'", "NOTIFYLEVEL": "'range(0, 4)'", "NUM_INITAGENTS": "'range(0, 64000)'", "NUM_INITFENCED": "'range(0, 64000)'", "NUM_POOLAGENTS": "'-1', 'range(0, 64000)'", "RESYNC_INTERVAL": "'range(1, 60000)'", "RQRIOBLK": "'range(4096, 65535)'", "START_STOP_TIME": "'range(1, 1440)'", "UTIL_IMPACT_LIM": "'range(1, 100)'", "WLM_DISPATCHER": "'YES', 'NO'", "WLM_DISP_CONCUR": "'range(1, 32767)', 'COMPUTED'", "WLM_DISP_CPU_SHARES": "'NO', 'YES'", "WLM_DISP_MIN_UTIL": "'range(0, 100)'"}, "registry": {"DB2BIDI": "'YES', 'NO'", "DB2COMPOPT": "'-'", "DB2LOCK_TO_RB": "'STATEMENT'", "DB2STMM": "'NO', 'YES'", "DB2_ALTERNATE_AUTHZ_BEHAVIOUR": "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'", "DB2_ANTIJOIN": "'YES', 'NO', 'EXTEND'", "DB2_ATS_ENABLE": "'YES', 'NO'", "DB2_DEFERRED_PREPARE_SEMANTICS": "'NO', 'YES'", "DB2_EVALUNCOMMITTED": "'NO', 'YES'", "DB2_EXTENDED_OPTIMIZATION": "'-'", "DB2_INDEX_PCTFREE_DEFAULT": "'range(0, 99)'", "DB2_INLIST_TO_NLJN": "'NO', 'YES'", "DB2_MINIMIZE_LISTPREFETCH": "'NO', 'YES'", "DB2_OBJECT_TABLE_ENTRIES": "'range(0, 65532)'", "DB2_OPTPROFILE": "'NO', 'YES'", "DB2_OPTSTATS_LOG": "'-'", "DB2_OPT_MAX_TEMP_SIZE": "'-'", "DB2_PARALLEL_IO": "'-'", "DB2_REDUCED_OPTIMIZATION": "'-'", "DB2_SELECTIVITY": "'YES', 'NO', 'ALL'", "DB2_SKIPDELETED": "'NO', 'YES'", "DB2_SKIPINSERTED": "'NO', 'YES'", "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES": "'NO', 'YES'", "DB2_TRUNCATE_REUSESTORAGE": "'IMPORT', 'LOAD', 'TRUNCATE'", "DB2_USE_ALTERNATE_PAGE_CLEANING": "'ON', 'OFF'", "DB2_VIEW_REOPT_VALUES": "'NO', 'YES'", "DB2_WLM_SETTINGS": "'-'", "DB2_WORKLOAD": "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"}}}`)
				}))
			})
			It(`Invoke GetDb2SaasTuneableParam successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasTuneableParamWithContext(ctx, getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasTuneableParamWithContext(ctx, getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasTuneableParamPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tuneable_param": {"db": {"ACT_SORTMEM_LIMIT": "'NONE', 'range(10, 100)'", "ALT_COLLATE": "'NULL', 'IDENTITY_16BIT'", "APPGROUP_MEM_SZ": "'range(1, 1000000)'", "APPLHEAPSZ": "'AUTOMATIC', 'range(16, 2147483647)'", "APPL_MEMORY": "'AUTOMATIC', 'range(128, 4294967295)'", "APP_CTL_HEAP_SZ": "'range(1, 64000)'", "ARCHRETRYDELAY": "'range(0, 65535)'", "AUTHN_CACHE_DURATION": "'range(1,10000)'", "AUTORESTART": "'ON', 'OFF'", "AUTO_CG_STATS": "'ON', 'OFF'", "AUTO_MAINT": "'ON', 'OFF'", "AUTO_REORG": "'ON', 'OFF'", "AUTO_REVAL": "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'", "AUTO_RUNSTATS": "'ON', 'OFF'", "AUTO_SAMPLING": "'ON', 'OFF'", "AUTO_STATS_VIEWS": "'ON', 'OFF'", "AUTO_STMT_STATS": "'ON', 'OFF'", "AUTO_TBL_MAINT": "'ON', 'OFF'", "AVG_APPLS": "'-'", "CATALOGCACHE_SZ": "'-'", "CHNGPGS_THRESH": "'range(5,99)'", "CUR_COMMIT": "'ON, AVAILABLE, DISABLED'", "DATABASE_MEMORY": "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'", "DBHEAP": "'AUTOMATIC', 'range(32, 2147483647)'", "DB_COLLNAME": "'-'", "DB_MEM_THRESH": "'range(0, 100)'", "DDL_COMPRESSION_DEF": "'YES', 'NO'", "DDL_CONSTRAINT_DEF": "'YES', 'NO'", "DECFLT_ROUNDING": "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'", "DEC_ARITHMETIC": "'-'", "DEC_TO_CHAR_FMT": "'NEW', 'V95'", "DFT_DEGREE": "'-1', 'ANY', 'range(1, 32767)'", "DFT_EXTENT_SZ": "'range(2, 256)'", "DFT_LOADREC_SES": "'range(1, 30000)'", "DFT_MTTB_TYPES": "'-'", "DFT_PREFETCH_SZ": "'range(0, 32767)', 'AUTOMATIC'", "DFT_QUERYOPT": "'range(0, 9)'", "DFT_REFRESH_AGE": "'-'", "DFT_SCHEMAS_DCC": "'YES', 'NO'", "DFT_SQLMATHWARN": "'YES', 'NO'", "DFT_TABLE_ORG": "'COLUMN', 'ROW'", "DLCHKTIME": "'range(1000, 600000)'", "ENABLE_XMLCHAR": "'YES', 'NO'", "EXTENDED_ROW_SZ": "'ENABLE', 'DISABLE'", "GROUPHEAP_RATIO": "'range(1, 99)'", "INDEXREC": "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'", "LARGE_AGGREGATION": "'YES', 'NO'", "LOCKLIST": "'AUTOMATIC', 'range(4, 134217728)'", "LOCKTIMEOUT": "'-1', 'range(0, 32767)'", "LOGINDEXBUILD": "'ON', 'OFF'", "LOG_APPL_INFO": "'YES', 'NO'", "LOG_DDL_STMTS": "'YES', 'NO'", "LOG_DISK_CAP": "'0', '-1', 'range(1, 2147483647)'", "MAXAPPLS": "'range(1, 60000)'", "MAXFILOP": "'range(64, 61440)'", "MAXLOCKS": "'AUTOMATIC', 'range(1, 100)'", "MIN_DEC_DIV_3": "'YES', 'NO'", "MON_ACT_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_DEADLOCK": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LCK_MSG_LVL": "'range(0, 3)'", "MON_LOCKTIMEOUT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LOCKWAIT": "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'", "MON_LW_THRESH": "'range(1000, 4294967295)'", "MON_OBJ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_PKGLIST_SZ": "'range(0, 1024)'", "MON_REQ_METRICS": "'NONE', 'BASE', 'EXTENDED'", "MON_RTN_DATA": "'NONE', 'BASE'", "MON_RTN_EXECLIST": "'OFF', 'ON'", "MON_UOW_DATA": "'NONE', 'BASE'", "MON_UOW_EXECLIST": "'ON', 'OFF'", "MON_UOW_PKGLIST": "'OFF', 'ON'", "NCHAR_MAPPING": "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'", "NUM_FREQVALUES": "'range(0, 32767)'", "NUM_IOCLEANERS": "'AUTOMATIC', 'range(0, 255)'", "NUM_IOSERVERS": "'AUTOMATIC', 'range(1, 255)'", "NUM_LOG_SPAN": "'range(0, 65535)'", "NUM_QUANTILES": "'range(0, 32767)'", "OPT_BUFFPAGE": "'-'", "OPT_DIRECT_WRKLD": "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'", "OPT_LOCKLIST": "'-'", "OPT_MAXLOCKS": "'-'", "OPT_SORTHEAP": "'-'", "PAGE_AGE_TRGT_GCR": "'range(1, 65535)'", "PAGE_AGE_TRGT_MCR": "'range(1, 65535)'", "PCKCACHESZ": "'AUTOMATIC', '-1', 'range(32, 2147483646)'", "PL_STACK_TRACE": "'NONE', 'ALL', 'UNHANDLED'", "SELF_TUNING_MEM": "'ON', 'OFF'", "SEQDETECT": "'YES', 'NO'", "SHEAPTHRES_SHR": "'AUTOMATIC', 'range(250, 2147483647)'", "SOFTMAX": "'-'", "SORTHEAP": "'AUTOMATIC', 'range(16, 4294967295)'", "SQL_CCFLAGS": "'-'", "STAT_HEAP_SZ": "'AUTOMATIC', 'range(1096, 2147483647)'", "STMTHEAP": "'AUTOMATIC', 'range(128, 2147483647)'", "STMT_CONC": "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'", "STRING_UNITS": "'SYSTEM', 'CODEUNITS32'", "SYSTIME_PERIOD_ADJ": "'NO', 'YES'", "TRACKMOD": "'YES', 'NO'", "UTIL_HEAP_SZ": "'AUTOMATIC', 'range(16, 2147483647)'", "WLM_ADMISSION_CTRL": "'YES', 'NO'", "WLM_AGENT_LOAD_TRGT": "'AUTOMATIC', 'range(1, 65535)'", "WLM_CPU_LIMIT": "'range(0, 100)'", "WLM_CPU_SHARES": "'range(1, 65535)'", "WLM_CPU_SHARE_MODE": "'HARD', 'SOFT'"}, "dbm": {"COMM_BANDWIDTH": "'range(0.1, 100000)', '-1'", "CPUSPEED": "'range(0.0000000001, 1)', '-1'", "DFT_MON_BUFPOOL": "'ON', 'OFF'", "DFT_MON_LOCK": "'ON', 'OFF'", "DFT_MON_SORT": "'ON', 'OFF'", "DFT_MON_STMT": "'ON', 'OFF'", "DFT_MON_TABLE": "'ON', 'OFF'", "DFT_MON_TIMESTAMP": "'ON', 'OFF'", "DFT_MON_UOW": "'ON', 'OFF'", "DIAGLEVEL": "'range(0, 4)'", "FEDERATED_ASYNC": "'range(0, 32767)', '-1', 'ANY'", "INDEXREC": "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'", "INTRA_PARALLEL": "'SYSTEM', 'NO', 'YES'", "KEEPFENCED": "'YES', 'NO'", "MAX_CONNRETRIES": "'range(0, 100)'", "MAX_QUERYDEGREE": "'range(1, 32767)', '-1', 'ANY'", "MON_HEAP_SZ": "'range(0, 2147483647)', 'AUTOMATIC'", "MULTIPARTSIZEMB": "'range(5, 5120)'", "NOTIFYLEVEL": "'range(0, 4)'", "NUM_INITAGENTS": "'range(0, 64000)'", "NUM_INITFENCED": "'range(0, 64000)'", "NUM_POOLAGENTS": "'-1', 'range(0, 64000)'", "RESYNC_INTERVAL": "'range(1, 60000)'", "RQRIOBLK": "'range(4096, 65535)'", "START_STOP_TIME": "'range(1, 1440)'", "UTIL_IMPACT_LIM": "'range(1, 100)'", "WLM_DISPATCHER": "'YES', 'NO'", "WLM_DISP_CONCUR": "'range(1, 32767)', 'COMPUTED'", "WLM_DISP_CPU_SHARES": "'NO', 'YES'", "WLM_DISP_MIN_UTIL": "'range(0, 100)'"}, "registry": {"DB2BIDI": "'YES', 'NO'", "DB2COMPOPT": "'-'", "DB2LOCK_TO_RB": "'STATEMENT'", "DB2STMM": "'NO', 'YES'", "DB2_ALTERNATE_AUTHZ_BEHAVIOUR": "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'", "DB2_ANTIJOIN": "'YES', 'NO', 'EXTEND'", "DB2_ATS_ENABLE": "'YES', 'NO'", "DB2_DEFERRED_PREPARE_SEMANTICS": "'NO', 'YES'", "DB2_EVALUNCOMMITTED": "'NO', 'YES'", "DB2_EXTENDED_OPTIMIZATION": "'-'", "DB2_INDEX_PCTFREE_DEFAULT": "'range(0, 99)'", "DB2_INLIST_TO_NLJN": "'NO', 'YES'", "DB2_MINIMIZE_LISTPREFETCH": "'NO', 'YES'", "DB2_OBJECT_TABLE_ENTRIES": "'range(0, 65532)'", "DB2_OPTPROFILE": "'NO', 'YES'", "DB2_OPTSTATS_LOG": "'-'", "DB2_OPT_MAX_TEMP_SIZE": "'-'", "DB2_PARALLEL_IO": "'-'", "DB2_REDUCED_OPTIMIZATION": "'-'", "DB2_SELECTIVITY": "'YES', 'NO', 'ALL'", "DB2_SKIPDELETED": "'NO', 'YES'", "DB2_SKIPINSERTED": "'NO', 'YES'", "DB2_SYNC_RELEASE_LOCK_ATTRIBUTES": "'NO', 'YES'", "DB2_TRUNCATE_REUSESTORAGE": "'IMPORT', 'LOAD', 'TRUNCATE'", "DB2_USE_ALTERNATE_PAGE_CLEANING": "'ON', 'OFF'", "DB2_VIEW_REOPT_VALUES": "'NO', 'YES'", "DB2_WLM_SETTINGS": "'-'", "DB2_WORKLOAD": "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"}}}`)
				}))
			})
			It(`Invoke GetDb2SaasTuneableParam successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasTuneableParam with error: Operation request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasTuneableParam successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := new(db2saasv1.GetDb2SaasTuneableParamOptions)
				getDb2SaasTuneableParamOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions) - Operation response error`, func() {
		getDb2SaasBackupPath := "/manage/backups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasBackup with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasBackupOptions model
				getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				getDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions)`, func() {
		getDb2SaasBackupPath := "/manage/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "ID", "type": "Type", "status": "Status", "created_at": "CreatedAt", "size": 4, "duration": 8}]}`)
				}))
			})
			It(`Invoke GetDb2SaasBackup successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasBackupOptions model
				getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				getDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasBackupWithContext(ctx, getDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasBackupWithContext(ctx, getDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasBackupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "ID", "type": "Type", "status": "Status", "created_at": "CreatedAt", "size": 4, "duration": 8}]}`)
				}))
			})
			It(`Invoke GetDb2SaasBackup successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasBackupOptions model
				getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				getDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasBackup with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasBackupOptions model
				getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				getDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasBackupOptions model with no property values
				getDb2SaasBackupOptionsModelNew := new(db2saasv1.GetDb2SaasBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDb2SaasBackup successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasBackupOptions model
				getDb2SaasBackupOptionsModel := new(db2saasv1.GetDb2SaasBackupOptions)
				getDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions) - Operation response error`, func() {
		postDb2SaasBackupPath := "/manage/backups/backup"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDb2SaasBackup with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PostDb2SaasBackupOptions model
				postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				postDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions)`, func() {
		postDb2SaasBackupPath := "/manage/backups/backup"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:0c9c7889-54de-4ecc-8399-09a4d4ff228e:task:51ff2dc7-6cb9-41c0-9345-09e54550fb7b"}}`)
				}))
			})
			It(`Invoke PostDb2SaasBackup successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the PostDb2SaasBackupOptions model
				postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				postDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PostDb2SaasBackupWithContext(ctx, postDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PostDb2SaasBackupWithContext(ctx, postDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasBackupPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Db-Profile"]).ToNot(BeNil())
					Expect(req.Header["X-Db-Profile"][0]).To(Equal(fmt.Sprintf("%v", "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:0c9c7889-54de-4ecc-8399-09a4d4ff228e:task:51ff2dc7-6cb9-41c0-9345-09e54550fb7b"}}`)
				}))
			})
			It(`Invoke PostDb2SaasBackup successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PostDb2SaasBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PostDb2SaasBackupOptions model
				postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				postDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostDb2SaasBackup with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PostDb2SaasBackupOptions model
				postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				postDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDb2SaasBackupOptions model with no property values
				postDb2SaasBackupOptionsModelNew := new(db2saasv1.PostDb2SaasBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PostDb2SaasBackup successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the PostDb2SaasBackupOptions model
				postDb2SaasBackupOptionsModel := new(db2saasv1.PostDb2SaasBackupOptions)
				postDb2SaasBackupOptionsModel.XDbProfile = core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			db2saasService, _ := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
				URL:           "http://db2saasv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateUserAuthentication successfully`, func() {
				method := "internal"
				policyID := "Default"
				_model, err := db2saasService.NewCreateUserAuthentication(method, policyID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the DeleteDb2SaasUserOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				id := "test-user"
				deleteDb2SaasUserOptionsModel := db2saasService.NewDeleteDb2SaasUserOptions(xDeploymentID, id)
				deleteDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				deleteDb2SaasUserOptionsModel.SetID("test-user")
				deleteDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(deleteDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(deleteDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(deleteDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasAllowlistOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasAllowlistOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				getDb2SaasAllowlistOptionsModel := db2saasService.NewGetDb2SaasAllowlistOptions(xDeploymentID)
				getDb2SaasAllowlistOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasAllowlistOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasAllowlistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(getDb2SaasAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasAutoscaleOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				getDb2SaasAutoscaleOptionsModel := db2saasService.NewGetDb2SaasAutoscaleOptions(xDbProfile)
				getDb2SaasAutoscaleOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasAutoscaleOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				Expect(getDb2SaasAutoscaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasBackupOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasBackupOptions model
				xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				getDb2SaasBackupOptionsModel := db2saasService.NewGetDb2SaasBackupOptions(xDbProfile)
				getDb2SaasBackupOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				getDb2SaasBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasBackupOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasBackupOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				Expect(getDb2SaasBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasConnectionInfoOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				deploymentID := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A"
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				getDb2SaasConnectionInfoOptionsModel := db2saasService.NewGetDb2SaasConnectionInfoOptions(deploymentID, xDeploymentID)
				getDb2SaasConnectionInfoOptionsModel.SetDeploymentID("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")
				getDb2SaasConnectionInfoOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasConnectionInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasConnectionInfoOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasConnectionInfoOptionsModel.DeploymentID).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A")))
				Expect(getDb2SaasConnectionInfoOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(getDb2SaasConnectionInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasTuneableParamOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasTuneableParamOptions model
				getDb2SaasTuneableParamOptionsModel := db2saasService.NewGetDb2SaasTuneableParamOptions()
				getDb2SaasTuneableParamOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasTuneableParamOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasTuneableParamOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasUserOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				getDb2SaasUserOptionsModel := db2saasService.NewGetDb2SaasUserOptions(xDeploymentID)
				getDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(getDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetbyidDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the GetbyidDb2SaasUserOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				id := "test-user"
				getbyidDb2SaasUserOptionsModel := db2saasService.NewGetbyidDb2SaasUserOptions(xDeploymentID, id)
				getbyidDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				getbyidDb2SaasUserOptionsModel.SetID("test-user")
				getbyidDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getbyidDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(getbyidDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(getbyidDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(getbyidDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIpAddress successfully`, func() {
				address := "127.0.0.1"
				description := "A sample IP address"
				_model, err := db2saasService.NewIpAddress(address, description)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPostDb2SaasAllowlistOptions successfully`, func() {
				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				Expect(ipAddressModel).ToNot(BeNil())
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")
				Expect(ipAddressModel.Address).To(Equal(core.StringPtr("127.0.0.1")))
				Expect(ipAddressModel.Description).To(Equal(core.StringPtr("A sample IP address")))

				// Construct an instance of the PostDb2SaasAllowlistOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				postDb2SaasAllowlistOptionsIpAddresses := []db2saasv1.IpAddress{}
				postDb2SaasAllowlistOptionsModel := db2saasService.NewPostDb2SaasAllowlistOptions(xDeploymentID, postDb2SaasAllowlistOptionsIpAddresses)
				postDb2SaasAllowlistOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasAllowlistOptionsModel.SetIpAddresses([]db2saasv1.IpAddress{*ipAddressModel})
				postDb2SaasAllowlistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDb2SaasAllowlistOptionsModel).ToNot(BeNil())
				Expect(postDb2SaasAllowlistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(postDb2SaasAllowlistOptionsModel.IpAddresses).To(Equal([]db2saasv1.IpAddress{*ipAddressModel}))
				Expect(postDb2SaasAllowlistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDb2SaasBackupOptions successfully`, func() {
				// Construct an instance of the PostDb2SaasBackupOptions model
				xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				postDb2SaasBackupOptionsModel := db2saasService.NewPostDb2SaasBackupOptions(xDbProfile)
				postDb2SaasBackupOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDb2SaasBackupOptionsModel).ToNot(BeNil())
				Expect(postDb2SaasBackupOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				Expect(postDb2SaasBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDb2SaasDbConfigurationOptions successfully`, func() {
				// Construct an instance of the CreateCustomSettingsRegistry model
				createCustomSettingsRegistryModel := new(db2saasv1.CreateCustomSettingsRegistry)
				Expect(createCustomSettingsRegistryModel).ToNot(BeNil())
				createCustomSettingsRegistryModel.DB2BIDI = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2COMPOPT = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2LOCKTORB = core.StringPtr("STATEMENT")
				createCustomSettingsRegistryModel.DB2STMM = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
				createCustomSettingsRegistryModel.DB2ANTIJOIN = core.StringPtr("EXTEND")
				createCustomSettingsRegistryModel.DB2ATSENABLE = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
				createCustomSettingsRegistryModel.DB2INLISTTONLJN = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
				createCustomSettingsRegistryModel.DB2OPTPROFILE = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2OPTSTATSLOG = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2PARALLELIO = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2SELECTIVITY = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SKIPDELETED = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2SKIPINSERTED = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
				createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
				createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
				createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("NO")
				createCustomSettingsRegistryModel.DB2WLMSETTINGS = core.StringPtr("-")
				createCustomSettingsRegistryModel.DB2WORKLOAD = core.StringPtr("SAP")
				Expect(createCustomSettingsRegistryModel.DB2BIDI).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2COMPOPT).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2LOCKTORB).To(Equal(core.StringPtr("STATEMENT")))
				Expect(createCustomSettingsRegistryModel.DB2STMM).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR).To(Equal(core.StringPtr("EXTERNAL_ROUTINE_DBADM")))
				Expect(createCustomSettingsRegistryModel.DB2ANTIJOIN).To(Equal(core.StringPtr("EXTEND")))
				Expect(createCustomSettingsRegistryModel.DB2ATSENABLE).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2DEFERREDPREPARESEMANTICS).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2EVALUNCOMMITTED).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsRegistryModel.DB2EXTENDEDOPTIMIZATION).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2INDEXPCTFREEDEFAULT).To(Equal(core.StringPtr("10")))
				Expect(createCustomSettingsRegistryModel.DB2INLISTTONLJN).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2MINIMIZELISTPREFETCH).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsRegistryModel.DB2OBJECTTABLEENTRIES).To(Equal(core.StringPtr("5000")))
				Expect(createCustomSettingsRegistryModel.DB2OPTPROFILE).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsRegistryModel.DB2OPTSTATSLOG).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2OPTMAXTEMPSIZE).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2PARALLELIO).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2REDUCEDOPTIMIZATION).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2SELECTIVITY).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2SKIPDELETED).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsRegistryModel.DB2SKIPINSERTED).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsRegistryModel.DB2TRUNCATEREUSESTORAGE).To(Equal(core.StringPtr("IMPORT")))
				Expect(createCustomSettingsRegistryModel.DB2USEALTERNATEPAGECLEANING).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsRegistryModel.DB2VIEWREOPTVALUES).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsRegistryModel.DB2WLMSETTINGS).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsRegistryModel.DB2WORKLOAD).To(Equal(core.StringPtr("SAP")))

				// Construct an instance of the CreateCustomSettingsDb model
				createCustomSettingsDbModel := new(db2saasv1.CreateCustomSettingsDb)
				Expect(createCustomSettingsDbModel).ToNot(BeNil())
				createCustomSettingsDbModel.ACTSORTMEMLIMIT = core.StringPtr("NONE")
				createCustomSettingsDbModel.ALTCOLLATE = core.StringPtr("NULL")
				createCustomSettingsDbModel.APPGROUPMEMSZ = core.StringPtr("10")
				createCustomSettingsDbModel.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPLMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.APPCTLHEAPSZ = core.StringPtr("64000")
				createCustomSettingsDbModel.ARCHRETRYDELAY = core.StringPtr("65535")
				createCustomSettingsDbModel.AUTHNCACHEDURATION = core.StringPtr("10000")
				createCustomSettingsDbModel.AUTORESTART = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOCGSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOMAINT = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOREORG = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOREVAL = core.StringPtr("IMMEDIATE")
				createCustomSettingsDbModel.AUTORUNSTATS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSAMPLING = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOSTATSVIEWS = core.StringPtr("ON")
				createCustomSettingsDbModel.AUTOSTMTSTATS = core.StringPtr("OFF")
				createCustomSettingsDbModel.AUTOTBLMAINT = core.StringPtr("ON")
				createCustomSettingsDbModel.AVGAPPLS = core.StringPtr("-")
				createCustomSettingsDbModel.CATALOGCACHESZ = core.StringPtr("-")
				createCustomSettingsDbModel.CHNGPGSTHRESH = core.StringPtr("50")
				createCustomSettingsDbModel.CURCOMMIT = core.StringPtr("AVAILABLE")
				createCustomSettingsDbModel.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DBCOLLNAME = core.StringPtr("-")
				createCustomSettingsDbModel.DBMEMTHRESH = core.StringPtr("75")
				createCustomSettingsDbModel.DDLCOMPRESSIONDEF = core.StringPtr("YES")
				createCustomSettingsDbModel.DDLCONSTRAINTDEF = core.StringPtr("NO")
				createCustomSettingsDbModel.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
				createCustomSettingsDbModel.DECARITHMETIC = core.StringPtr("-")
				createCustomSettingsDbModel.DECTOCHARFMT = core.StringPtr("NEW")
				createCustomSettingsDbModel.DFTDEGREE = core.StringPtr("-1")
				createCustomSettingsDbModel.DFTEXTENTSZ = core.StringPtr("32")
				createCustomSettingsDbModel.DFTLOADRECSES = core.StringPtr("1000")
				createCustomSettingsDbModel.DFTMTTBTYPES = core.StringPtr("-")
				createCustomSettingsDbModel.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.DFTQUERYOPT = core.StringPtr("3")
				createCustomSettingsDbModel.DFTREFRESHAGE = core.StringPtr("-")
				createCustomSettingsDbModel.DFTSCHEMASDCC = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTSQLMATHWARN = core.StringPtr("YES")
				createCustomSettingsDbModel.DFTTABLEORG = core.StringPtr("COLUMN")
				createCustomSettingsDbModel.DLCHKTIME = core.StringPtr("10000")
				createCustomSettingsDbModel.ENABLEXMLCHAR = core.StringPtr("YES")
				createCustomSettingsDbModel.EXTENDEDROWSZ = core.StringPtr("ENABLE")
				createCustomSettingsDbModel.GROUPHEAPRATIO = core.StringPtr("50")
				createCustomSettingsDbModel.INDEXREC = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.LARGEAGGREGATION = core.StringPtr("YES")
				createCustomSettingsDbModel.LOCKLIST = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.LOCKTIMEOUT = core.StringPtr("-1")
				createCustomSettingsDbModel.LOGINDEXBUILD = core.StringPtr("ON")
				createCustomSettingsDbModel.LOGAPPLINFO = core.StringPtr("YES")
				createCustomSettingsDbModel.LOGDDLSTMTS = core.StringPtr("NO")
				createCustomSettingsDbModel.LOGDISKCAP = core.StringPtr("0")
				createCustomSettingsDbModel.MAXAPPLS = core.StringPtr("5000")
				createCustomSettingsDbModel.MAXFILOP = core.StringPtr("1024")
				createCustomSettingsDbModel.MAXLOCKS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.MINDECDIV3 = core.StringPtr("NO")
				createCustomSettingsDbModel.MONACTMETRICS = core.StringPtr("EXTENDED")
				createCustomSettingsDbModel.MONDEADLOCK = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLCKMSGLVL = core.StringPtr("2")
				createCustomSettingsDbModel.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
				createCustomSettingsDbModel.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
				createCustomSettingsDbModel.MONLWTHRESH = core.StringPtr("10000")
				createCustomSettingsDbModel.MONOBJMETRICS = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONPKGLISTSZ = core.StringPtr("512")
				createCustomSettingsDbModel.MONREQMETRICS = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONRTNDATA = core.StringPtr("BASE")
				createCustomSettingsDbModel.MONRTNEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWDATA = core.StringPtr("NONE")
				createCustomSettingsDbModel.MONUOWEXECLIST = core.StringPtr("ON")
				createCustomSettingsDbModel.MONUOWPKGLIST = core.StringPtr("OFF")
				createCustomSettingsDbModel.NCHARMAPPING = core.StringPtr("CHAR_CU32")
				createCustomSettingsDbModel.NUMFREQVALUES = core.StringPtr("50")
				createCustomSettingsDbModel.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.NUMLOGSPAN = core.StringPtr("10")
				createCustomSettingsDbModel.NUMQUANTILES = core.StringPtr("100")
				createCustomSettingsDbModel.OPTBUFFPAGE = core.StringPtr("-")
				createCustomSettingsDbModel.OPTDIRECTWRKLD = core.StringPtr("ON")
				createCustomSettingsDbModel.OPTLOCKLIST = core.StringPtr("-")
				createCustomSettingsDbModel.OPTMAXLOCKS = core.StringPtr("-")
				createCustomSettingsDbModel.OPTSORTHEAP = core.StringPtr("-")
				createCustomSettingsDbModel.PAGEAGETRGTGCR = core.StringPtr("5000")
				createCustomSettingsDbModel.PAGEAGETRGTMCR = core.StringPtr("3000")
				createCustomSettingsDbModel.PCKCACHESZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.PLSTACKTRACE = core.StringPtr("UNHANDLED")
				createCustomSettingsDbModel.SELFTUNINGMEM = core.StringPtr("ON")
				createCustomSettingsDbModel.SEQDETECT = core.StringPtr("YES")
				createCustomSettingsDbModel.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SOFTMAX = core.StringPtr("-")
				createCustomSettingsDbModel.SORTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.SQLCCFLAGS = core.StringPtr("-")
				createCustomSettingsDbModel.STATHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTHEAP = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.STMTCONC = core.StringPtr("LITERALS")
				createCustomSettingsDbModel.STRINGUNITS = core.StringPtr("SYSTEM")
				createCustomSettingsDbModel.SYSTIMEPERIODADJ = core.StringPtr("NO")
				createCustomSettingsDbModel.TRACKMOD = core.StringPtr("YES")
				createCustomSettingsDbModel.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbModel.WLMADMISSIONCTRL = core.StringPtr("YES")
				createCustomSettingsDbModel.WLMAGENTLOADTRGT = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPULIMIT = core.StringPtr("80")
				createCustomSettingsDbModel.WLMCPUSHARES = core.StringPtr("1000")
				createCustomSettingsDbModel.WLMCPUSHAREMODE = core.StringPtr("SOFT")
				Expect(createCustomSettingsDbModel.ACTSORTMEMLIMIT).To(Equal(core.StringPtr("NONE")))
				Expect(createCustomSettingsDbModel.ALTCOLLATE).To(Equal(core.StringPtr("NULL")))
				Expect(createCustomSettingsDbModel.APPGROUPMEMSZ).To(Equal(core.StringPtr("10")))
				Expect(createCustomSettingsDbModel.APPLHEAPSZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.APPLMEMORY).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.APPCTLHEAPSZ).To(Equal(core.StringPtr("64000")))
				Expect(createCustomSettingsDbModel.ARCHRETRYDELAY).To(Equal(core.StringPtr("65535")))
				Expect(createCustomSettingsDbModel.AUTHNCACHEDURATION).To(Equal(core.StringPtr("10000")))
				Expect(createCustomSettingsDbModel.AUTORESTART).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AUTOCGSTATS).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AUTOMAINT).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbModel.AUTOREORG).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AUTOREVAL).To(Equal(core.StringPtr("IMMEDIATE")))
				Expect(createCustomSettingsDbModel.AUTORUNSTATS).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AUTOSAMPLING).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbModel.AUTOSTATSVIEWS).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AUTOSTMTSTATS).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbModel.AUTOTBLMAINT).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.AVGAPPLS).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.CATALOGCACHESZ).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.CHNGPGSTHRESH).To(Equal(core.StringPtr("50")))
				Expect(createCustomSettingsDbModel.CURCOMMIT).To(Equal(core.StringPtr("AVAILABLE")))
				Expect(createCustomSettingsDbModel.DATABASEMEMORY).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.DBHEAP).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.DBCOLLNAME).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.DBMEMTHRESH).To(Equal(core.StringPtr("75")))
				Expect(createCustomSettingsDbModel.DDLCOMPRESSIONDEF).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.DDLCONSTRAINTDEF).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsDbModel.DECFLTROUNDING).To(Equal(core.StringPtr("ROUND_HALF_UP")))
				Expect(createCustomSettingsDbModel.DECARITHMETIC).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.DECTOCHARFMT).To(Equal(core.StringPtr("NEW")))
				Expect(createCustomSettingsDbModel.DFTDEGREE).To(Equal(core.StringPtr("-1")))
				Expect(createCustomSettingsDbModel.DFTEXTENTSZ).To(Equal(core.StringPtr("32")))
				Expect(createCustomSettingsDbModel.DFTLOADRECSES).To(Equal(core.StringPtr("1000")))
				Expect(createCustomSettingsDbModel.DFTMTTBTYPES).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.DFTPREFETCHSZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.DFTQUERYOPT).To(Equal(core.StringPtr("3")))
				Expect(createCustomSettingsDbModel.DFTREFRESHAGE).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.DFTSCHEMASDCC).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.DFTSQLMATHWARN).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.DFTTABLEORG).To(Equal(core.StringPtr("COLUMN")))
				Expect(createCustomSettingsDbModel.DLCHKTIME).To(Equal(core.StringPtr("10000")))
				Expect(createCustomSettingsDbModel.ENABLEXMLCHAR).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.EXTENDEDROWSZ).To(Equal(core.StringPtr("ENABLE")))
				Expect(createCustomSettingsDbModel.GROUPHEAPRATIO).To(Equal(core.StringPtr("50")))
				Expect(createCustomSettingsDbModel.INDEXREC).To(Equal(core.StringPtr("SYSTEM")))
				Expect(createCustomSettingsDbModel.LARGEAGGREGATION).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.LOCKLIST).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.LOCKTIMEOUT).To(Equal(core.StringPtr("-1")))
				Expect(createCustomSettingsDbModel.LOGINDEXBUILD).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.LOGAPPLINFO).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.LOGDDLSTMTS).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsDbModel.LOGDISKCAP).To(Equal(core.StringPtr("0")))
				Expect(createCustomSettingsDbModel.MAXAPPLS).To(Equal(core.StringPtr("5000")))
				Expect(createCustomSettingsDbModel.MAXFILOP).To(Equal(core.StringPtr("1024")))
				Expect(createCustomSettingsDbModel.MAXLOCKS).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.MINDECDIV3).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsDbModel.MONACTMETRICS).To(Equal(core.StringPtr("EXTENDED")))
				Expect(createCustomSettingsDbModel.MONDEADLOCK).To(Equal(core.StringPtr("HISTORY")))
				Expect(createCustomSettingsDbModel.MONLCKMSGLVL).To(Equal(core.StringPtr("2")))
				Expect(createCustomSettingsDbModel.MONLOCKTIMEOUT).To(Equal(core.StringPtr("HISTORY")))
				Expect(createCustomSettingsDbModel.MONLOCKWAIT).To(Equal(core.StringPtr("WITHOUT_HIST")))
				Expect(createCustomSettingsDbModel.MONLWTHRESH).To(Equal(core.StringPtr("10000")))
				Expect(createCustomSettingsDbModel.MONOBJMETRICS).To(Equal(core.StringPtr("BASE")))
				Expect(createCustomSettingsDbModel.MONPKGLISTSZ).To(Equal(core.StringPtr("512")))
				Expect(createCustomSettingsDbModel.MONREQMETRICS).To(Equal(core.StringPtr("NONE")))
				Expect(createCustomSettingsDbModel.MONRTNDATA).To(Equal(core.StringPtr("BASE")))
				Expect(createCustomSettingsDbModel.MONRTNEXECLIST).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.MONUOWDATA).To(Equal(core.StringPtr("NONE")))
				Expect(createCustomSettingsDbModel.MONUOWEXECLIST).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.MONUOWPKGLIST).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbModel.NCHARMAPPING).To(Equal(core.StringPtr("CHAR_CU32")))
				Expect(createCustomSettingsDbModel.NUMFREQVALUES).To(Equal(core.StringPtr("50")))
				Expect(createCustomSettingsDbModel.NUMIOCLEANERS).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.NUMIOSERVERS).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.NUMLOGSPAN).To(Equal(core.StringPtr("10")))
				Expect(createCustomSettingsDbModel.NUMQUANTILES).To(Equal(core.StringPtr("100")))
				Expect(createCustomSettingsDbModel.OPTBUFFPAGE).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.OPTDIRECTWRKLD).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.OPTLOCKLIST).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.OPTMAXLOCKS).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.OPTSORTHEAP).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.PAGEAGETRGTGCR).To(Equal(core.StringPtr("5000")))
				Expect(createCustomSettingsDbModel.PAGEAGETRGTMCR).To(Equal(core.StringPtr("3000")))
				Expect(createCustomSettingsDbModel.PCKCACHESZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.PLSTACKTRACE).To(Equal(core.StringPtr("UNHANDLED")))
				Expect(createCustomSettingsDbModel.SELFTUNINGMEM).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbModel.SEQDETECT).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.SHEAPTHRESSHR).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.SOFTMAX).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.SORTHEAP).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.SQLCCFLAGS).To(Equal(core.StringPtr("-")))
				Expect(createCustomSettingsDbModel.STATHEAPSZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.STMTHEAP).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.STMTCONC).To(Equal(core.StringPtr("LITERALS")))
				Expect(createCustomSettingsDbModel.STRINGUNITS).To(Equal(core.StringPtr("SYSTEM")))
				Expect(createCustomSettingsDbModel.SYSTIMEPERIODADJ).To(Equal(core.StringPtr("NO")))
				Expect(createCustomSettingsDbModel.TRACKMOD).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.UTILHEAPSZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbModel.WLMADMISSIONCTRL).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbModel.WLMAGENTLOADTRGT).To(Equal(core.StringPtr("1000")))
				Expect(createCustomSettingsDbModel.WLMCPULIMIT).To(Equal(core.StringPtr("80")))
				Expect(createCustomSettingsDbModel.WLMCPUSHARES).To(Equal(core.StringPtr("1000")))
				Expect(createCustomSettingsDbModel.WLMCPUSHAREMODE).To(Equal(core.StringPtr("SOFT")))

				// Construct an instance of the CreateCustomSettingsDbm model
				createCustomSettingsDbmModel := new(db2saasv1.CreateCustomSettingsDbm)
				Expect(createCustomSettingsDbmModel).ToNot(BeNil())
				createCustomSettingsDbmModel.COMMBANDWIDTH = core.StringPtr("1000")
				createCustomSettingsDbmModel.CPUSPEED = core.StringPtr("0.5")
				createCustomSettingsDbmModel.DFTMONBUFPOOL = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONLOCK = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONSORT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONSTMT = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONTABLE = core.StringPtr("OFF")
				createCustomSettingsDbmModel.DFTMONTIMESTAMP = core.StringPtr("ON")
				createCustomSettingsDbmModel.DFTMONUOW = core.StringPtr("ON")
				createCustomSettingsDbmModel.DIAGLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.FEDERATEDASYNC = core.StringPtr("32767")
				createCustomSettingsDbmModel.INDEXREC = core.StringPtr("RESTART")
				createCustomSettingsDbmModel.INTRAPARALLEL = core.StringPtr("YES")
				createCustomSettingsDbmModel.KEEPFENCED = core.StringPtr("YES")
				createCustomSettingsDbmModel.MAXCONNRETRIES = core.StringPtr("5")
				createCustomSettingsDbmModel.MAXQUERYDEGREE = core.StringPtr("4")
				createCustomSettingsDbmModel.MONHEAPSZ = core.StringPtr("AUTOMATIC")
				createCustomSettingsDbmModel.MULTIPARTSIZEMB = core.StringPtr("100")
				createCustomSettingsDbmModel.NOTIFYLEVEL = core.StringPtr("2")
				createCustomSettingsDbmModel.NUMINITAGENTS = core.StringPtr("100")
				createCustomSettingsDbmModel.NUMINITFENCED = core.StringPtr("20")
				createCustomSettingsDbmModel.NUMPOOLAGENTS = core.StringPtr("10")
				createCustomSettingsDbmModel.RESYNCINTERVAL = core.StringPtr("1000")
				createCustomSettingsDbmModel.RQRIOBLK = core.StringPtr("8192")
				createCustomSettingsDbmModel.STARTSTOPTIME = core.StringPtr("10")
				createCustomSettingsDbmModel.UTILIMPACTLIM = core.StringPtr("50")
				createCustomSettingsDbmModel.WLMDISPATCHER = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPCONCUR = core.StringPtr("16")
				createCustomSettingsDbmModel.WLMDISPCPUSHARES = core.StringPtr("YES")
				createCustomSettingsDbmModel.WLMDISPMINUTIL = core.StringPtr("10")
				Expect(createCustomSettingsDbmModel.COMMBANDWIDTH).To(Equal(core.StringPtr("1000")))
				Expect(createCustomSettingsDbmModel.CPUSPEED).To(Equal(core.StringPtr("0.5")))
				Expect(createCustomSettingsDbmModel.DFTMONBUFPOOL).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbmModel.DFTMONLOCK).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbmModel.DFTMONSORT).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbmModel.DFTMONSTMT).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbmModel.DFTMONTABLE).To(Equal(core.StringPtr("OFF")))
				Expect(createCustomSettingsDbmModel.DFTMONTIMESTAMP).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbmModel.DFTMONUOW).To(Equal(core.StringPtr("ON")))
				Expect(createCustomSettingsDbmModel.DIAGLEVEL).To(Equal(core.StringPtr("2")))
				Expect(createCustomSettingsDbmModel.FEDERATEDASYNC).To(Equal(core.StringPtr("32767")))
				Expect(createCustomSettingsDbmModel.INDEXREC).To(Equal(core.StringPtr("RESTART")))
				Expect(createCustomSettingsDbmModel.INTRAPARALLEL).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbmModel.KEEPFENCED).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbmModel.MAXCONNRETRIES).To(Equal(core.StringPtr("5")))
				Expect(createCustomSettingsDbmModel.MAXQUERYDEGREE).To(Equal(core.StringPtr("4")))
				Expect(createCustomSettingsDbmModel.MONHEAPSZ).To(Equal(core.StringPtr("AUTOMATIC")))
				Expect(createCustomSettingsDbmModel.MULTIPARTSIZEMB).To(Equal(core.StringPtr("100")))
				Expect(createCustomSettingsDbmModel.NOTIFYLEVEL).To(Equal(core.StringPtr("2")))
				Expect(createCustomSettingsDbmModel.NUMINITAGENTS).To(Equal(core.StringPtr("100")))
				Expect(createCustomSettingsDbmModel.NUMINITFENCED).To(Equal(core.StringPtr("20")))
				Expect(createCustomSettingsDbmModel.NUMPOOLAGENTS).To(Equal(core.StringPtr("10")))
				Expect(createCustomSettingsDbmModel.RESYNCINTERVAL).To(Equal(core.StringPtr("1000")))
				Expect(createCustomSettingsDbmModel.RQRIOBLK).To(Equal(core.StringPtr("8192")))
				Expect(createCustomSettingsDbmModel.STARTSTOPTIME).To(Equal(core.StringPtr("10")))
				Expect(createCustomSettingsDbmModel.UTILIMPACTLIM).To(Equal(core.StringPtr("50")))
				Expect(createCustomSettingsDbmModel.WLMDISPATCHER).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbmModel.WLMDISPCONCUR).To(Equal(core.StringPtr("16")))
				Expect(createCustomSettingsDbmModel.WLMDISPCPUSHARES).To(Equal(core.StringPtr("YES")))
				Expect(createCustomSettingsDbmModel.WLMDISPMINUTIL).To(Equal(core.StringPtr("10")))

				// Construct an instance of the PostDb2SaasDbConfigurationOptions model
				xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				postDb2SaasDbConfigurationOptionsModel := db2saasService.NewPostDb2SaasDbConfigurationOptions(xDbProfile)
				postDb2SaasDbConfigurationOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				postDb2SaasDbConfigurationOptionsModel.SetRegistry(createCustomSettingsRegistryModel)
				postDb2SaasDbConfigurationOptionsModel.SetDb(createCustomSettingsDbModel)
				postDb2SaasDbConfigurationOptionsModel.SetDbm(createCustomSettingsDbmModel)
				postDb2SaasDbConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDb2SaasDbConfigurationOptionsModel).ToNot(BeNil())
				Expect(postDb2SaasDbConfigurationOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				Expect(postDb2SaasDbConfigurationOptionsModel.Registry).To(Equal(createCustomSettingsRegistryModel))
				Expect(postDb2SaasDbConfigurationOptionsModel.Db).To(Equal(createCustomSettingsDbModel))
				Expect(postDb2SaasDbConfigurationOptionsModel.Dbm).To(Equal(createCustomSettingsDbmModel))
				Expect(postDb2SaasDbConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPostDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the CreateUserAuthentication model
				createUserAuthenticationModel := new(db2saasv1.CreateUserAuthentication)
				Expect(createUserAuthenticationModel).ToNot(BeNil())
				createUserAuthenticationModel.Method = core.StringPtr("internal")
				createUserAuthenticationModel.PolicyID = core.StringPtr("Default")
				Expect(createUserAuthenticationModel.Method).To(Equal(core.StringPtr("internal")))
				Expect(createUserAuthenticationModel.PolicyID).To(Equal(core.StringPtr("Default")))

				// Construct an instance of the PostDb2SaasUserOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				postDb2SaasUserOptionsID := "test-user"
				postDb2SaasUserOptionsIam := false
				postDb2SaasUserOptionsIbmid := "test-ibm-id"
				postDb2SaasUserOptionsName := "test_user"
				postDb2SaasUserOptionsPassword := "dEkMc43@gfAPl!867^dSbu"
				postDb2SaasUserOptionsRole := "bluuser"
				postDb2SaasUserOptionsEmail := "test_user@mycompany.com"
				postDb2SaasUserOptionsLocked := "no"
				var postDb2SaasUserOptionsAuthentication *db2saasv1.CreateUserAuthentication = nil
				postDb2SaasUserOptionsModel := db2saasService.NewPostDb2SaasUserOptions(xDeploymentID, postDb2SaasUserOptionsID, postDb2SaasUserOptionsIam, postDb2SaasUserOptionsIbmid, postDb2SaasUserOptionsName, postDb2SaasUserOptionsPassword, postDb2SaasUserOptionsRole, postDb2SaasUserOptionsEmail, postDb2SaasUserOptionsLocked, postDb2SaasUserOptionsAuthentication)
				postDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				postDb2SaasUserOptionsModel.SetID("test-user")
				postDb2SaasUserOptionsModel.SetIam(false)
				postDb2SaasUserOptionsModel.SetIbmid("test-ibm-id")
				postDb2SaasUserOptionsModel.SetName("test_user")
				postDb2SaasUserOptionsModel.SetPassword("dEkMc43@gfAPl!867^dSbu")
				postDb2SaasUserOptionsModel.SetRole("bluuser")
				postDb2SaasUserOptionsModel.SetEmail("test_user@mycompany.com")
				postDb2SaasUserOptionsModel.SetLocked("no")
				postDb2SaasUserOptionsModel.SetAuthentication(createUserAuthenticationModel)
				postDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(postDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(postDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(postDb2SaasUserOptionsModel.Iam).To(Equal(core.BoolPtr(false)))
				Expect(postDb2SaasUserOptionsModel.Ibmid).To(Equal(core.StringPtr("test-ibm-id")))
				Expect(postDb2SaasUserOptionsModel.Name).To(Equal(core.StringPtr("test_user")))
				Expect(postDb2SaasUserOptionsModel.Password).To(Equal(core.StringPtr("dEkMc43@gfAPl!867^dSbu")))
				Expect(postDb2SaasUserOptionsModel.Role).To(Equal(core.StringPtr("bluuser")))
				Expect(postDb2SaasUserOptionsModel.Email).To(Equal(core.StringPtr("test_user@mycompany.com")))
				Expect(postDb2SaasUserOptionsModel.Locked).To(Equal(core.StringPtr("no")))
				Expect(postDb2SaasUserOptionsModel.Authentication).To(Equal(createUserAuthenticationModel))
				Expect(postDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDb2SaasAutoscaleOptions successfully`, func() {
				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
				putDb2SaasAutoscaleOptionsModel := db2saasService.NewPutDb2SaasAutoscaleOptions(xDbProfile)
				putDb2SaasAutoscaleOptionsModel.SetXDbProfile("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingEnabled("true")
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingThreshold(int64(90))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingOverTimePeriod(float64(5))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingPauseLimit(int64(70))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingAllowPlanLimit("YES")
				putDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				Expect(putDb2SaasAutoscaleOptionsModel.XDbProfile).To(Equal(core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A")))
				Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingEnabled).To(Equal(core.StringPtr("true")))
				Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingThreshold).To(Equal(core.Int64Ptr(int64(90))))
				Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingOverTimePeriod).To(Equal(core.Float64Ptr(float64(5))))
				Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingPauseLimit).To(Equal(core.Int64Ptr(int64(70))))
				Expect(putDb2SaasAutoscaleOptionsModel.AutoScalingAllowPlanLimit).To(Equal(core.StringPtr("YES")))
				Expect(putDb2SaasAutoscaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the UpdateUserAuthentication model
				updateUserAuthenticationModel := new(db2saasv1.UpdateUserAuthentication)
				Expect(updateUserAuthenticationModel).ToNot(BeNil())
				updateUserAuthenticationModel.Method = core.StringPtr("internal")
				updateUserAuthenticationModel.PolicyID = core.StringPtr("Default")
				Expect(updateUserAuthenticationModel.Method).To(Equal(core.StringPtr("internal")))
				Expect(updateUserAuthenticationModel.PolicyID).To(Equal(core.StringPtr("Default")))

				// Construct an instance of the PutDb2SaasUserOptions model
				xDeploymentID := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
				id := "test-user"
				putDb2SaasUserOptionsNewID := "test-user"
				putDb2SaasUserOptionsNewIam := false
				putDb2SaasUserOptionsNewIbmid := "test-ibm-id"
				putDb2SaasUserOptionsNewName := "test_user"
				putDb2SaasUserOptionsNewPassword := "dEkMc43@gfAPl!867^dSbu"
				putDb2SaasUserOptionsNewRole := "bluuser"
				putDb2SaasUserOptionsNewEmail := "test_user@mycompany.com"
				putDb2SaasUserOptionsNewLocked := "no"
				var putDb2SaasUserOptionsNewAuthentication *db2saasv1.UpdateUserAuthentication = nil
				putDb2SaasUserOptionsModel := db2saasService.NewPutDb2SaasUserOptions(xDeploymentID, id, putDb2SaasUserOptionsNewID, putDb2SaasUserOptionsNewIam, putDb2SaasUserOptionsNewIbmid, putDb2SaasUserOptionsNewName, putDb2SaasUserOptionsNewPassword, putDb2SaasUserOptionsNewRole, putDb2SaasUserOptionsNewEmail, putDb2SaasUserOptionsNewLocked, putDb2SaasUserOptionsNewAuthentication)
				putDb2SaasUserOptionsModel.SetXDeploymentID("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")
				putDb2SaasUserOptionsModel.SetID("test-user")
				putDb2SaasUserOptionsModel.SetNewID("test-user")
				putDb2SaasUserOptionsModel.SetNewIam(false)
				putDb2SaasUserOptionsModel.SetNewIbmid("test-ibm-id")
				putDb2SaasUserOptionsModel.SetNewName("test_user")
				putDb2SaasUserOptionsModel.SetNewPassword("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.SetNewRole("bluuser")
				putDb2SaasUserOptionsModel.SetNewEmail("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.SetNewLocked("no")
				putDb2SaasUserOptionsModel.SetNewAuthentication(updateUserAuthenticationModel)
				putDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(putDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::")))
				Expect(putDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(putDb2SaasUserOptionsModel.NewID).To(Equal(core.StringPtr("test-user")))
				Expect(putDb2SaasUserOptionsModel.NewIam).To(Equal(core.BoolPtr(false)))
				Expect(putDb2SaasUserOptionsModel.NewIbmid).To(Equal(core.StringPtr("test-ibm-id")))
				Expect(putDb2SaasUserOptionsModel.NewName).To(Equal(core.StringPtr("test_user")))
				Expect(putDb2SaasUserOptionsModel.NewPassword).To(Equal(core.StringPtr("dEkMc43@gfAPl!867^dSbu")))
				Expect(putDb2SaasUserOptionsModel.NewRole).To(Equal(core.StringPtr("bluuser")))
				Expect(putDb2SaasUserOptionsModel.NewEmail).To(Equal(core.StringPtr("test_user@mycompany.com")))
				Expect(putDb2SaasUserOptionsModel.NewLocked).To(Equal(core.StringPtr("no")))
				Expect(putDb2SaasUserOptionsModel.NewAuthentication).To(Equal(updateUserAuthenticationModel))
				Expect(putDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserAuthentication successfully`, func() {
				method := "internal"
				policyID := "Default"
				_model, err := db2saasService.NewUpdateUserAuthentication(method, policyID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCreateCustomSettingsDb successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.CreateCustomSettingsDb)
			model.ACTSORTMEMLIMIT = core.StringPtr("NONE")
			model.ALTCOLLATE = core.StringPtr("NULL")
			model.APPGROUPMEMSZ = core.StringPtr("10")
			model.APPLHEAPSZ = core.StringPtr("AUTOMATIC")
			model.APPLMEMORY = core.StringPtr("AUTOMATIC")
			model.APPCTLHEAPSZ = core.StringPtr("64000")
			model.ARCHRETRYDELAY = core.StringPtr("65535")
			model.AUTHNCACHEDURATION = core.StringPtr("10000")
			model.AUTORESTART = core.StringPtr("ON")
			model.AUTOCGSTATS = core.StringPtr("ON")
			model.AUTOMAINT = core.StringPtr("OFF")
			model.AUTOREORG = core.StringPtr("ON")
			model.AUTOREVAL = core.StringPtr("IMMEDIATE")
			model.AUTORUNSTATS = core.StringPtr("ON")
			model.AUTOSAMPLING = core.StringPtr("OFF")
			model.AUTOSTATSVIEWS = core.StringPtr("ON")
			model.AUTOSTMTSTATS = core.StringPtr("OFF")
			model.AUTOTBLMAINT = core.StringPtr("ON")
			model.AVGAPPLS = core.StringPtr("-")
			model.CATALOGCACHESZ = core.StringPtr("-")
			model.CHNGPGSTHRESH = core.StringPtr("50")
			model.CURCOMMIT = core.StringPtr("AVAILABLE")
			model.DATABASEMEMORY = core.StringPtr("AUTOMATIC")
			model.DBHEAP = core.StringPtr("AUTOMATIC")
			model.DBCOLLNAME = core.StringPtr("-")
			model.DBMEMTHRESH = core.StringPtr("75")
			model.DDLCOMPRESSIONDEF = core.StringPtr("YES")
			model.DDLCONSTRAINTDEF = core.StringPtr("NO")
			model.DECFLTROUNDING = core.StringPtr("ROUND_HALF_UP")
			model.DECARITHMETIC = core.StringPtr("-")
			model.DECTOCHARFMT = core.StringPtr("NEW")
			model.DFTDEGREE = core.StringPtr("-1")
			model.DFTEXTENTSZ = core.StringPtr("32")
			model.DFTLOADRECSES = core.StringPtr("1000")
			model.DFTMTTBTYPES = core.StringPtr("-")
			model.DFTPREFETCHSZ = core.StringPtr("AUTOMATIC")
			model.DFTQUERYOPT = core.StringPtr("3")
			model.DFTREFRESHAGE = core.StringPtr("-")
			model.DFTSCHEMASDCC = core.StringPtr("YES")
			model.DFTSQLMATHWARN = core.StringPtr("YES")
			model.DFTTABLEORG = core.StringPtr("COLUMN")
			model.DLCHKTIME = core.StringPtr("10000")
			model.ENABLEXMLCHAR = core.StringPtr("YES")
			model.EXTENDEDROWSZ = core.StringPtr("ENABLE")
			model.GROUPHEAPRATIO = core.StringPtr("50")
			model.INDEXREC = core.StringPtr("SYSTEM")
			model.LARGEAGGREGATION = core.StringPtr("YES")
			model.LOCKLIST = core.StringPtr("AUTOMATIC")
			model.LOCKTIMEOUT = core.StringPtr("-1")
			model.LOGINDEXBUILD = core.StringPtr("ON")
			model.LOGAPPLINFO = core.StringPtr("YES")
			model.LOGDDLSTMTS = core.StringPtr("NO")
			model.LOGDISKCAP = core.StringPtr("0")
			model.MAXAPPLS = core.StringPtr("5000")
			model.MAXFILOP = core.StringPtr("1024")
			model.MAXLOCKS = core.StringPtr("AUTOMATIC")
			model.MINDECDIV3 = core.StringPtr("NO")
			model.MONACTMETRICS = core.StringPtr("EXTENDED")
			model.MONDEADLOCK = core.StringPtr("HISTORY")
			model.MONLCKMSGLVL = core.StringPtr("2")
			model.MONLOCKTIMEOUT = core.StringPtr("HISTORY")
			model.MONLOCKWAIT = core.StringPtr("WITHOUT_HIST")
			model.MONLWTHRESH = core.StringPtr("10000")
			model.MONOBJMETRICS = core.StringPtr("BASE")
			model.MONPKGLISTSZ = core.StringPtr("512")
			model.MONREQMETRICS = core.StringPtr("NONE")
			model.MONRTNDATA = core.StringPtr("BASE")
			model.MONRTNEXECLIST = core.StringPtr("ON")
			model.MONUOWDATA = core.StringPtr("NONE")
			model.MONUOWEXECLIST = core.StringPtr("ON")
			model.MONUOWPKGLIST = core.StringPtr("OFF")
			model.NCHARMAPPING = core.StringPtr("CHAR_CU32")
			model.NUMFREQVALUES = core.StringPtr("50")
			model.NUMIOCLEANERS = core.StringPtr("AUTOMATIC")
			model.NUMIOSERVERS = core.StringPtr("AUTOMATIC")
			model.NUMLOGSPAN = core.StringPtr("10")
			model.NUMQUANTILES = core.StringPtr("100")
			model.OPTBUFFPAGE = core.StringPtr("-")
			model.OPTDIRECTWRKLD = core.StringPtr("ON")
			model.OPTLOCKLIST = core.StringPtr("-")
			model.OPTMAXLOCKS = core.StringPtr("-")
			model.OPTSORTHEAP = core.StringPtr("-")
			model.PAGEAGETRGTGCR = core.StringPtr("5000")
			model.PAGEAGETRGTMCR = core.StringPtr("3000")
			model.PCKCACHESZ = core.StringPtr("AUTOMATIC")
			model.PLSTACKTRACE = core.StringPtr("UNHANDLED")
			model.SELFTUNINGMEM = core.StringPtr("ON")
			model.SEQDETECT = core.StringPtr("YES")
			model.SHEAPTHRESSHR = core.StringPtr("AUTOMATIC")
			model.SOFTMAX = core.StringPtr("-")
			model.SORTHEAP = core.StringPtr("AUTOMATIC")
			model.SQLCCFLAGS = core.StringPtr("-")
			model.STATHEAPSZ = core.StringPtr("AUTOMATIC")
			model.STMTHEAP = core.StringPtr("AUTOMATIC")
			model.STMTCONC = core.StringPtr("LITERALS")
			model.STRINGUNITS = core.StringPtr("SYSTEM")
			model.SYSTIMEPERIODADJ = core.StringPtr("NO")
			model.TRACKMOD = core.StringPtr("YES")
			model.UTILHEAPSZ = core.StringPtr("AUTOMATIC")
			model.WLMADMISSIONCTRL = core.StringPtr("YES")
			model.WLMAGENTLOADTRGT = core.StringPtr("1000")
			model.WLMCPULIMIT = core.StringPtr("80")
			model.WLMCPUSHARES = core.StringPtr("1000")
			model.WLMCPUSHAREMODE = core.StringPtr("SOFT")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.CreateCustomSettingsDb
			err = db2saasv1.UnmarshalCreateCustomSettingsDb(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateCustomSettingsDbm successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.CreateCustomSettingsDbm)
			model.COMMBANDWIDTH = core.StringPtr("1000")
			model.CPUSPEED = core.StringPtr("0.5")
			model.DFTMONBUFPOOL = core.StringPtr("ON")
			model.DFTMONLOCK = core.StringPtr("OFF")
			model.DFTMONSORT = core.StringPtr("ON")
			model.DFTMONSTMT = core.StringPtr("ON")
			model.DFTMONTABLE = core.StringPtr("OFF")
			model.DFTMONTIMESTAMP = core.StringPtr("ON")
			model.DFTMONUOW = core.StringPtr("ON")
			model.DIAGLEVEL = core.StringPtr("2")
			model.FEDERATEDASYNC = core.StringPtr("32767")
			model.INDEXREC = core.StringPtr("RESTART")
			model.INTRAPARALLEL = core.StringPtr("YES")
			model.KEEPFENCED = core.StringPtr("YES")
			model.MAXCONNRETRIES = core.StringPtr("5")
			model.MAXQUERYDEGREE = core.StringPtr("4")
			model.MONHEAPSZ = core.StringPtr("AUTOMATIC")
			model.MULTIPARTSIZEMB = core.StringPtr("100")
			model.NOTIFYLEVEL = core.StringPtr("2")
			model.NUMINITAGENTS = core.StringPtr("100")
			model.NUMINITFENCED = core.StringPtr("20")
			model.NUMPOOLAGENTS = core.StringPtr("10")
			model.RESYNCINTERVAL = core.StringPtr("1000")
			model.RQRIOBLK = core.StringPtr("8192")
			model.STARTSTOPTIME = core.StringPtr("10")
			model.UTILIMPACTLIM = core.StringPtr("50")
			model.WLMDISPATCHER = core.StringPtr("YES")
			model.WLMDISPCONCUR = core.StringPtr("16")
			model.WLMDISPCPUSHARES = core.StringPtr("YES")
			model.WLMDISPMINUTIL = core.StringPtr("10")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.CreateCustomSettingsDbm
			err = db2saasv1.UnmarshalCreateCustomSettingsDbm(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateCustomSettingsRegistry successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.CreateCustomSettingsRegistry)
			model.DB2BIDI = core.StringPtr("YES")
			model.DB2COMPOPT = core.StringPtr("-")
			model.DB2LOCKTORB = core.StringPtr("STATEMENT")
			model.DB2STMM = core.StringPtr("YES")
			model.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("EXTERNAL_ROUTINE_DBADM")
			model.DB2ANTIJOIN = core.StringPtr("EXTEND")
			model.DB2ATSENABLE = core.StringPtr("YES")
			model.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("YES")
			model.DB2EVALUNCOMMITTED = core.StringPtr("NO")
			model.DB2EXTENDEDOPTIMIZATION = core.StringPtr("-")
			model.DB2INDEXPCTFREEDEFAULT = core.StringPtr("10")
			model.DB2INLISTTONLJN = core.StringPtr("YES")
			model.DB2MINIMIZELISTPREFETCH = core.StringPtr("NO")
			model.DB2OBJECTTABLEENTRIES = core.StringPtr("5000")
			model.DB2OPTPROFILE = core.StringPtr("NO")
			model.DB2OPTSTATSLOG = core.StringPtr("-")
			model.DB2OPTMAXTEMPSIZE = core.StringPtr("-")
			model.DB2PARALLELIO = core.StringPtr("-")
			model.DB2REDUCEDOPTIMIZATION = core.StringPtr("-")
			model.DB2SELECTIVITY = core.StringPtr("YES")
			model.DB2SKIPDELETED = core.StringPtr("NO")
			model.DB2SKIPINSERTED = core.StringPtr("YES")
			model.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("YES")
			model.DB2TRUNCATEREUSESTORAGE = core.StringPtr("IMPORT")
			model.DB2USEALTERNATEPAGECLEANING = core.StringPtr("ON")
			model.DB2VIEWREOPTVALUES = core.StringPtr("NO")
			model.DB2WLMSETTINGS = core.StringPtr("-")
			model.DB2WORKLOAD = core.StringPtr("SAP")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.CreateCustomSettingsRegistry
			err = db2saasv1.UnmarshalCreateCustomSettingsRegistry(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalCreateUserAuthentication successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.CreateUserAuthentication)
			model.Method = core.StringPtr("internal")
			model.PolicyID = core.StringPtr("Default")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.CreateUserAuthentication
			err = db2saasv1.UnmarshalCreateUserAuthentication(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalIpAddress successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.IpAddress)
			model.Address = core.StringPtr("127.0.0.1")
			model.Description = core.StringPtr("A sample IP address")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.IpAddress
			err = db2saasv1.UnmarshalIpAddress(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUpdateUserAuthentication successfully`, func() {
			// Construct an instance of the model.
			model := new(db2saasv1.UpdateUserAuthentication)
			model.Method = core.StringPtr("internal")
			model.PolicyID = core.StringPtr("Default")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *db2saasv1.UpdateUserAuthentication
			err = db2saasv1.UnmarshalUpdateUserAuthentication(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
