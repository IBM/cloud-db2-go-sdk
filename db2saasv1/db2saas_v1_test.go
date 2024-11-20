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
		getDb2SaasConnectionInfoPath := "/connectioninfo/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("testString")
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
		getDb2SaasConnectionInfoPath := "/connectioninfo/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasConnectionInfoPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "host_ros": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud:30515", "certificateBase64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIVENDQWdXZ0F3SUJBZ0lVTGRpR1U2QzdZajMwcS9VUVB3ek5ka2YyakJjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0hqRWNNQm9HQTFVRUF3d1RTVUpOSUVOc2IzVmtJRVJoZEdGaVlYTmxjekFlRncweU1EQTRNRFl3T1RReQpNVEZhRncwek1EQTRNRFF3T1RReU1URmFNQjR4SERBYUJnTlZCQU1NRTBsQ1RTQkRiRzkxWkNCRVlYUmhZbUZ6ClpYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDb1NIdS9TWkd5NHc0bHB0elQKbFVRQTV6Q0krTldhblQ0czAvTXFkQmJwRW9FWjYxLy9VQVFPaHVTUG85M05obG1NQWZUWENpUi9jVkxuQmxBMQpuZnEzcC9pWm1VMnJwSUxnUmdLeTdsNEZSMVVPaGlRa3RnN3d6Q0J1M2k0bTRJQkZ0NVVvRng5djl6eFkrK0tSCnNnYXhmK28yMEoxLzZBSHFwem5GaWJuTDdLcGlZMUs1c25BdHFwTUVsNHMyR3dlZXQ0dEFjZ3hRSlRVR3hvamsKUDMvUmtxSUI1RFBNSXJ0ZFMrWWpBdlM0alBpREVRT0FvZDg5aDBOays3bkpldllJT0lRVTN0OC81YlNYRDFFVwp3bmRqdHlkeC95Qlo5YlZ4bms4eWI1S1NCVUNpaHJsL1AxVmdNdStLb2w2M0ZZMmdSbndwb3FEOVRNWkJkeTRYCk5PRUZBZ01CQUFHalV6QlJNQjBHQTFVZERnUVdCQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBZkJnTlYKSFNNRUdEQVdnQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQkFRQk1XRHV2Z0JKVk5JYUp2NkFzL3FybWZKbVJObU80clhVcXhiTXdJdEZ1Ciswb1RIOTZMWU1OSjgyS1hXOFc1K0ZUOXJ2TjdzQzhRQzBYVzFIWkM4dlgvdE96dmluL1lqVW5nUlducUFBQXEKL0U4TnRtMFpuMEs4cnRzanJtaklLKzlwNjRObE1ENWJjcUpDMGZFSkpBQVpBSUozejRNSHhsTDhnc0plS0JyOApvcWhhejJOaXJtSEZ3Z3RDc0htVlI4UCt5TUtrN24xVlhlcmpHYWhORkQ2MzhGRnRoSHNvNmV0NGQ5NEpLTXFPCmt1cWhFOCtMcTZlalRWUTRYdldaUG4wVWlZQkVpTjFsT1JaZ0h5d3JvNjJ5Z2dFekhCaXF5dEI2SEN6TllyYXoKVElQUTNGanhGQXNYU3NhVzZPL2VteERNSDN4ZUZ5WmRZWWw5bGMxSnFVWW4KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "host_ros": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud:30515", "certificateBase64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIVENDQWdXZ0F3SUJBZ0lVTGRpR1U2QzdZajMwcS9VUVB3ek5ka2YyakJjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0hqRWNNQm9HQTFVRUF3d1RTVUpOSUVOc2IzVmtJRVJoZEdGaVlYTmxjekFlRncweU1EQTRNRFl3T1RReQpNVEZhRncwek1EQTRNRFF3T1RReU1URmFNQjR4SERBYUJnTlZCQU1NRTBsQ1RTQkRiRzkxWkNCRVlYUmhZbUZ6ClpYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDb1NIdS9TWkd5NHc0bHB0elQKbFVRQTV6Q0krTldhblQ0czAvTXFkQmJwRW9FWjYxLy9VQVFPaHVTUG85M05obG1NQWZUWENpUi9jVkxuQmxBMQpuZnEzcC9pWm1VMnJwSUxnUmdLeTdsNEZSMVVPaGlRa3RnN3d6Q0J1M2k0bTRJQkZ0NVVvRng5djl6eFkrK0tSCnNnYXhmK28yMEoxLzZBSHFwem5GaWJuTDdLcGlZMUs1c25BdHFwTUVsNHMyR3dlZXQ0dEFjZ3hRSlRVR3hvamsKUDMvUmtxSUI1RFBNSXJ0ZFMrWWpBdlM0alBpREVRT0FvZDg5aDBOays3bkpldllJT0lRVTN0OC81YlNYRDFFVwp3bmRqdHlkeC95Qlo5YlZ4bms4eWI1S1NCVUNpaHJsL1AxVmdNdStLb2w2M0ZZMmdSbndwb3FEOVRNWkJkeTRYCk5PRUZBZ01CQUFHalV6QlJNQjBHQTFVZERnUVdCQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBZkJnTlYKSFNNRUdEQVdnQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQkFRQk1XRHV2Z0JKVk5JYUp2NkFzL3FybWZKbVJObU80clhVcXhiTXdJdEZ1Ciswb1RIOTZMWU1OSjgyS1hXOFc1K0ZUOXJ2TjdzQzhRQzBYVzFIWkM4dlgvdE96dmluL1lqVW5nUlducUFBQXEKL0U4TnRtMFpuMEs4cnRzanJtaklLKzlwNjRObE1ENWJjcUpDMGZFSkpBQVpBSUozejRNSHhsTDhnc0plS0JyOApvcWhhejJOaXJtSEZ3Z3RDc0htVlI4UCt5TUtrN24xVlhlcmpHYWhORkQ2MzhGRnRoSHNvNmV0NGQ5NEpLTXFPCmt1cWhFOCtMcTZlalRWUTRYdldaUG4wVWlZQkVpTjFsT1JaZ0h5d3JvNjJ5Z2dFekhCaXF5dEI2SEN6TllyYXoKVElQUTNGanhGQXNYU3NhVzZPL2VteERNSDN4ZUZ5WmRZWWw5bGMxSnFVWW4KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}}`)
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"public": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "host_ros": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud:30515", "certificateBase64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIVENDQWdXZ0F3SUJBZ0lVTGRpR1U2QzdZajMwcS9VUVB3ek5ka2YyakJjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0hqRWNNQm9HQTFVRUF3d1RTVUpOSUVOc2IzVmtJRVJoZEdGaVlYTmxjekFlRncweU1EQTRNRFl3T1RReQpNVEZhRncwek1EQTRNRFF3T1RReU1URmFNQjR4SERBYUJnTlZCQU1NRTBsQ1RTQkRiRzkxWkNCRVlYUmhZbUZ6ClpYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDb1NIdS9TWkd5NHc0bHB0elQKbFVRQTV6Q0krTldhblQ0czAvTXFkQmJwRW9FWjYxLy9VQVFPaHVTUG85M05obG1NQWZUWENpUi9jVkxuQmxBMQpuZnEzcC9pWm1VMnJwSUxnUmdLeTdsNEZSMVVPaGlRa3RnN3d6Q0J1M2k0bTRJQkZ0NVVvRng5djl6eFkrK0tSCnNnYXhmK28yMEoxLzZBSHFwem5GaWJuTDdLcGlZMUs1c25BdHFwTUVsNHMyR3dlZXQ0dEFjZ3hRSlRVR3hvamsKUDMvUmtxSUI1RFBNSXJ0ZFMrWWpBdlM0alBpREVRT0FvZDg5aDBOays3bkpldllJT0lRVTN0OC81YlNYRDFFVwp3bmRqdHlkeC95Qlo5YlZ4bms4eWI1S1NCVUNpaHJsL1AxVmdNdStLb2w2M0ZZMmdSbndwb3FEOVRNWkJkeTRYCk5PRUZBZ01CQUFHalV6QlJNQjBHQTFVZERnUVdCQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBZkJnTlYKSFNNRUdEQVdnQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQkFRQk1XRHV2Z0JKVk5JYUp2NkFzL3FybWZKbVJObU80clhVcXhiTXdJdEZ1Ciswb1RIOTZMWU1OSjgyS1hXOFc1K0ZUOXJ2TjdzQzhRQzBYVzFIWkM4dlgvdE96dmluL1lqVW5nUlducUFBQXEKL0U4TnRtMFpuMEs4cnRzanJtaklLKzlwNjRObE1ENWJjcUpDMGZFSkpBQVpBSUozejRNSHhsTDhnc0plS0JyOApvcWhhejJOaXJtSEZ3Z3RDc0htVlI4UCt5TUtrN24xVlhlcmpHYWhORkQ2MzhGRnRoSHNvNmV0NGQ5NEpLTXFPCmt1cWhFOCtMcTZlalRWUTRYdldaUG4wVWlZQkVpTjFsT1JaZ0h5d3JvNjJ5Z2dFekhCaXF5dEI2SEN6TllyYXoKVElQUTNGanhGQXNYU3NhVzZPL2VteERNSDN4ZUZ5WmRZWWw5bGMxSnFVWW4KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}, "private": {"hostname": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud", "databaseName": "bluedb", "host_ros": "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud:30515", "certificateBase64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIVENDQWdXZ0F3SUJBZ0lVTGRpR1U2QzdZajMwcS9VUVB3ek5ka2YyakJjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0hqRWNNQm9HQTFVRUF3d1RTVUpOSUVOc2IzVmtJRVJoZEdGaVlYTmxjekFlRncweU1EQTRNRFl3T1RReQpNVEZhRncwek1EQTRNRFF3T1RReU1URmFNQjR4SERBYUJnTlZCQU1NRTBsQ1RTQkRiRzkxWkNCRVlYUmhZbUZ6ClpYTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDb1NIdS9TWkd5NHc0bHB0elQKbFVRQTV6Q0krTldhblQ0czAvTXFkQmJwRW9FWjYxLy9VQVFPaHVTUG85M05obG1NQWZUWENpUi9jVkxuQmxBMQpuZnEzcC9pWm1VMnJwSUxnUmdLeTdsNEZSMVVPaGlRa3RnN3d6Q0J1M2k0bTRJQkZ0NVVvRng5djl6eFkrK0tSCnNnYXhmK28yMEoxLzZBSHFwem5GaWJuTDdLcGlZMUs1c25BdHFwTUVsNHMyR3dlZXQ0dEFjZ3hRSlRVR3hvamsKUDMvUmtxSUI1RFBNSXJ0ZFMrWWpBdlM0alBpREVRT0FvZDg5aDBOays3bkpldllJT0lRVTN0OC81YlNYRDFFVwp3bmRqdHlkeC95Qlo5YlZ4bms4eWI1S1NCVUNpaHJsL1AxVmdNdStLb2w2M0ZZMmdSbndwb3FEOVRNWkJkeTRYCk5PRUZBZ01CQUFHalV6QlJNQjBHQTFVZERnUVdCQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBZkJnTlYKSFNNRUdEQVdnQlNldmNpblMvR0VwdmlmZkQ0ZUtPU0FNSGljUmpBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQkFRQk1XRHV2Z0JKVk5JYUp2NkFzL3FybWZKbVJObU80clhVcXhiTXdJdEZ1Ciswb1RIOTZMWU1OSjgyS1hXOFc1K0ZUOXJ2TjdzQzhRQzBYVzFIWkM4dlgvdE96dmluL1lqVW5nUlducUFBQXEKL0U4TnRtMFpuMEs4cnRzanJtaklLKzlwNjRObE1ENWJjcUpDMGZFSkpBQVpBSUozejRNSHhsTDhnc0plS0JyOApvcWhhejJOaXJtSEZ3Z3RDc0htVlI4UCt5TUtrN24xVlhlcmpHYWhORkQ2MzhGRnRoSHNvNmV0NGQ5NEpLTXFPCmt1cWhFOCtMcTZlalRWUTRYdldaUG4wVWlZQkVpTjFsT1JaZ0h5d3JvNjJ5Z2dFekhCaXF5dEI2SEN6TllyYXoKVElQUTNGanhGQXNYU3NhVzZPL2VteERNSDN4ZUZ5WmRZWWw5bGMxSnFVWW4KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "sslPort": "30450", "ssl": true, "databaseVersion": "11.5.0"}}`)
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasConnectionInfoOptionsModel.DeploymentID = core.StringPtr("testString")
				getDb2SaasConnectionInfoOptionsModel.XDeploymentID = core.StringPtr("testString")
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
	Describe(`PostDb2SaasWhitelist(postDb2SaasWhitelistOptions *PostDb2SaasWhitelistOptions) - Operation response error`, func() {
		postDb2SaasWhitelistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasWhitelistPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PostDb2SaasWhitelist with error: Operation response processing error`, func() {
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

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				postDb2SaasWhitelistOptionsModel := new(db2saasv1.PostDb2SaasWhitelistOptions)
				postDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				postDb2SaasWhitelistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PostDb2SaasWhitelist(postDb2SaasWhitelistOptions *PostDb2SaasWhitelistOptions)`, func() {
		postDb2SaasWhitelistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasWhitelistPath))
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasWhitelist successfully with retries`, func() {
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

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				postDb2SaasWhitelistOptionsModel := new(db2saasv1.PostDb2SaasWhitelistOptions)
				postDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				postDb2SaasWhitelistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.PostDb2SaasWhitelistWithContext(ctx, postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.PostDb2SaasWhitelistWithContext(ctx, postDb2SaasWhitelistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(postDb2SaasWhitelistPath))
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": "Status"}`)
				}))
			})
			It(`Invoke PostDb2SaasWhitelist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.PostDb2SaasWhitelist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				postDb2SaasWhitelistOptionsModel := new(db2saasv1.PostDb2SaasWhitelistOptions)
				postDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				postDb2SaasWhitelistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PostDb2SaasWhitelist with error: Operation validation and request error`, func() {
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

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				postDb2SaasWhitelistOptionsModel := new(db2saasv1.PostDb2SaasWhitelistOptions)
				postDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				postDb2SaasWhitelistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PostDb2SaasWhitelistOptions model with no property values
				postDb2SaasWhitelistOptionsModelNew := new(db2saasv1.PostDb2SaasWhitelistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModelNew)
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
			It(`Invoke PostDb2SaasWhitelist successfully`, func() {
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

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				postDb2SaasWhitelistOptionsModel := new(db2saasv1.PostDb2SaasWhitelistOptions)
				postDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				postDb2SaasWhitelistOptionsModel.IpAddresses = []db2saasv1.IpAddress{*ipAddressModel}
				postDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptionsModel)
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
	Describe(`GetDb2SaasWhitelist(getDb2SaasWhitelistOptions *GetDb2SaasWhitelistOptions) - Operation response error`, func() {
		getDb2SaasWhitelistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasWhitelistPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDb2SaasWhitelist with error: Operation response processing error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasWhitelistOptions model
				getDb2SaasWhitelistOptionsModel := new(db2saasv1.GetDb2SaasWhitelistOptions)
				getDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				getDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				db2saasService.EnableRetries(0, 0)
				result, response, operationErr = db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDb2SaasWhitelist(getDb2SaasWhitelistOptions *GetDb2SaasWhitelistOptions)`, func() {
		getDb2SaasWhitelistPath := "/dbsettings/whitelistips"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasWhitelistPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				}))
			})
			It(`Invoke GetDb2SaasWhitelist successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the GetDb2SaasWhitelistOptions model
				getDb2SaasWhitelistOptionsModel := new(db2saasv1.GetDb2SaasWhitelistOptions)
				getDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				getDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.GetDb2SaasWhitelistWithContext(ctx, getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.GetDb2SaasWhitelistWithContext(ctx, getDb2SaasWhitelistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasWhitelistPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "127.0.0.1", "description": "A sample IP address"}]}`)
				}))
			})
			It(`Invoke GetDb2SaasWhitelist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := db2saasService.GetDb2SaasWhitelist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDb2SaasWhitelistOptions model
				getDb2SaasWhitelistOptionsModel := new(db2saasv1.GetDb2SaasWhitelistOptions)
				getDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				getDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDb2SaasWhitelist with error: Operation validation and request error`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasWhitelistOptions model
				getDb2SaasWhitelistOptionsModel := new(db2saasv1.GetDb2SaasWhitelistOptions)
				getDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				getDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDb2SaasWhitelistOptions model with no property values
				getDb2SaasWhitelistOptionsModelNew := new(db2saasv1.GetDb2SaasWhitelistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModelNew)
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
			It(`Invoke GetDb2SaasWhitelist successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the GetDb2SaasWhitelistOptions model
				getDb2SaasWhitelistOptionsModel := new(db2saasv1.GetDb2SaasWhitelistOptions)
				getDb2SaasWhitelistOptionsModel.XDeploymentID = core.StringPtr("testString")
				getDb2SaasWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptionsModel)
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				postDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewOldPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewNewPassword = core.StringPtr("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewOldPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewNewPassword = core.StringPtr("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewOldPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewNewPassword = core.StringPtr("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
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
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewOldPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewNewPassword = core.StringPtr("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
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
				putDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				putDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewID = core.StringPtr("test-user")
				putDb2SaasUserOptionsModel.NewName = core.StringPtr("test_user")
				putDb2SaasUserOptionsModel.NewOldPassword = core.StringPtr("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.NewNewPassword = core.StringPtr("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.NewRole = core.StringPtr("bluuser")
				putDb2SaasUserOptionsModel.NewEmail = core.StringPtr("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.NewLocked = core.StringPtr("no")
				putDb2SaasUserOptionsModel.NewAuthentication = updateUserAuthenticationModel
				putDb2SaasUserOptionsModel.NewIbmid = core.StringPtr("test-ibm-id")
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
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDb2SaasUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
				}))
			})
			It(`Invoke DeleteDb2SaasUser successfully with retries`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())
				db2saasService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDb2SaasUserOptions model
				deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := db2saasService.DeleteDb2SaasUserWithContext(ctx, deleteDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				db2saasService.DisableRetries()
				result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = db2saasService.DeleteDb2SaasUserWithContext(ctx, deleteDb2SaasUserOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteDb2SaasUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"anyKey": "anyValue"}`)
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
				result, response, operationErr := db2saasService.DeleteDb2SaasUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDb2SaasUserOptions model
				deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := db2saasService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDb2SaasUserOptions model with no property values
				deleteDb2SaasUserOptionsModelNew := new(db2saasv1.DeleteDb2SaasUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModelNew)
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
			It(`Invoke DeleteDb2SaasUser successfully`, func() {
				db2saasService, serviceErr := db2saasv1.NewDb2saasV1(&db2saasv1.Db2saasV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(db2saasService).ToNot(BeNil())

				// Construct an instance of the DeleteDb2SaasUserOptions model
				deleteDb2SaasUserOptionsModel := new(db2saasv1.DeleteDb2SaasUserOptions)
				deleteDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
				deleteDb2SaasUserOptionsModel.ID = core.StringPtr("test-user")
				deleteDb2SaasUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptionsModel)
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
	Describe(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions) - Operation response error`, func() {
		getbyidDb2SaasUserPath := "/users/bluadmin"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
		getbyidDb2SaasUserPath := "/users/bluadmin"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getbyidDb2SaasUserPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getbyidDb2SaasUserOptionsModel.XDeploymentID = core.StringPtr("testString")
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
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				putDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				putDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				putDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
		getDb2SaasAutoscalePath := "/scaling/auto"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
		getDb2SaasAutoscalePath := "/scaling/auto"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDb2SaasAutoscalePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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

					Expect(req.Header["X-Deployment-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Deployment-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				getDb2SaasAutoscaleOptionsModel.XDeploymentID = core.StringPtr("testString")
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
				xDeploymentID := "testString"
				id := "test-user"
				deleteDb2SaasUserOptionsModel := db2saasService.NewDeleteDb2SaasUserOptions(xDeploymentID, id)
				deleteDb2SaasUserOptionsModel.SetXDeploymentID("testString")
				deleteDb2SaasUserOptionsModel.SetID("test-user")
				deleteDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(deleteDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(deleteDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasAutoscaleOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasAutoscaleOptions model
				xDeploymentID := "testString"
				getDb2SaasAutoscaleOptionsModel := db2saasService.NewGetDb2SaasAutoscaleOptions(xDeploymentID)
				getDb2SaasAutoscaleOptionsModel.SetXDeploymentID("testString")
				getDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasAutoscaleOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasAutoscaleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasConnectionInfoOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasConnectionInfoOptions model
				deploymentID := "testString"
				xDeploymentID := "testString"
				getDb2SaasConnectionInfoOptionsModel := db2saasService.NewGetDb2SaasConnectionInfoOptions(deploymentID, xDeploymentID)
				getDb2SaasConnectionInfoOptionsModel.SetDeploymentID("testString")
				getDb2SaasConnectionInfoOptionsModel.SetXDeploymentID("testString")
				getDb2SaasConnectionInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasConnectionInfoOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasConnectionInfoOptionsModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasConnectionInfoOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasConnectionInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasUserOptions model
				xDeploymentID := "testString"
				getDb2SaasUserOptionsModel := db2saasService.NewGetDb2SaasUserOptions(xDeploymentID)
				getDb2SaasUserOptionsModel.SetXDeploymentID("testString")
				getDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDb2SaasWhitelistOptions successfully`, func() {
				// Construct an instance of the GetDb2SaasWhitelistOptions model
				xDeploymentID := "testString"
				getDb2SaasWhitelistOptionsModel := db2saasService.NewGetDb2SaasWhitelistOptions(xDeploymentID)
				getDb2SaasWhitelistOptionsModel.SetXDeploymentID("testString")
				getDb2SaasWhitelistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDb2SaasWhitelistOptionsModel).ToNot(BeNil())
				Expect(getDb2SaasWhitelistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getDb2SaasWhitelistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetbyidDb2SaasUserOptions successfully`, func() {
				// Construct an instance of the GetbyidDb2SaasUserOptions model
				xDeploymentID := "testString"
				getbyidDb2SaasUserOptionsModel := db2saasService.NewGetbyidDb2SaasUserOptions(xDeploymentID)
				getbyidDb2SaasUserOptionsModel.SetXDeploymentID("testString")
				getbyidDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getbyidDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(getbyidDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(getbyidDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIpAddress successfully`, func() {
				address := "127.0.0.1"
				description := "A sample IP address"
				_model, err := db2saasService.NewIpAddress(address, description)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
				xDeploymentID := "testString"
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
				postDb2SaasUserOptionsModel.SetXDeploymentID("testString")
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
				Expect(postDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
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
			It(`Invoke NewPostDb2SaasWhitelistOptions successfully`, func() {
				// Construct an instance of the IpAddress model
				ipAddressModel := new(db2saasv1.IpAddress)
				Expect(ipAddressModel).ToNot(BeNil())
				ipAddressModel.Address = core.StringPtr("127.0.0.1")
				ipAddressModel.Description = core.StringPtr("A sample IP address")
				Expect(ipAddressModel.Address).To(Equal(core.StringPtr("127.0.0.1")))
				Expect(ipAddressModel.Description).To(Equal(core.StringPtr("A sample IP address")))

				// Construct an instance of the PostDb2SaasWhitelistOptions model
				xDeploymentID := "testString"
				postDb2SaasWhitelistOptionsIpAddresses := []db2saasv1.IpAddress{}
				postDb2SaasWhitelistOptionsModel := db2saasService.NewPostDb2SaasWhitelistOptions(xDeploymentID, postDb2SaasWhitelistOptionsIpAddresses)
				postDb2SaasWhitelistOptionsModel.SetXDeploymentID("testString")
				postDb2SaasWhitelistOptionsModel.SetIpAddresses([]db2saasv1.IpAddress{*ipAddressModel})
				postDb2SaasWhitelistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(postDb2SaasWhitelistOptionsModel).ToNot(BeNil())
				Expect(postDb2SaasWhitelistOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(postDb2SaasWhitelistOptionsModel.IpAddresses).To(Equal([]db2saasv1.IpAddress{*ipAddressModel}))
				Expect(postDb2SaasWhitelistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutDb2SaasAutoscaleOptions successfully`, func() {
				// Construct an instance of the PutDb2SaasAutoscaleOptions model
				xDeploymentID := "testString"
				putDb2SaasAutoscaleOptionsModel := db2saasService.NewPutDb2SaasAutoscaleOptions(xDeploymentID)
				putDb2SaasAutoscaleOptionsModel.SetXDeploymentID("testString")
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingEnabled("true")
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingThreshold(int64(90))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingOverTimePeriod(float64(5))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingPauseLimit(int64(70))
				putDb2SaasAutoscaleOptionsModel.SetAutoScalingAllowPlanLimit("YES")
				putDb2SaasAutoscaleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDb2SaasAutoscaleOptionsModel).ToNot(BeNil())
				Expect(putDb2SaasAutoscaleOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
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
				xDeploymentID := "testString"
				id := "test-user"
				putDb2SaasUserOptionsNewID := "test-user"
				putDb2SaasUserOptionsNewName := "test_user"
				putDb2SaasUserOptionsNewOldPassword := "dEkMc43@gfAPl!867^dSbu"
				putDb2SaasUserOptionsNewNewPassword := "ihbgc26@gfAPl!1297^dFGy"
				putDb2SaasUserOptionsNewRole := "bluuser"
				putDb2SaasUserOptionsNewEmail := "test_user@mycompany.com"
				putDb2SaasUserOptionsNewLocked := "no"
				var putDb2SaasUserOptionsNewAuthentication *db2saasv1.UpdateUserAuthentication = nil
				putDb2SaasUserOptionsModel := db2saasService.NewPutDb2SaasUserOptions(xDeploymentID, id, putDb2SaasUserOptionsNewID, putDb2SaasUserOptionsNewName, putDb2SaasUserOptionsNewOldPassword, putDb2SaasUserOptionsNewNewPassword, putDb2SaasUserOptionsNewRole, putDb2SaasUserOptionsNewEmail, putDb2SaasUserOptionsNewLocked, putDb2SaasUserOptionsNewAuthentication)
				putDb2SaasUserOptionsModel.SetXDeploymentID("testString")
				putDb2SaasUserOptionsModel.SetID("test-user")
				putDb2SaasUserOptionsModel.SetNewID("test-user")
				putDb2SaasUserOptionsModel.SetNewName("test_user")
				putDb2SaasUserOptionsModel.SetNewOldPassword("dEkMc43@gfAPl!867^dSbu")
				putDb2SaasUserOptionsModel.SetNewNewPassword("ihbgc26@gfAPl!1297^dFGy")
				putDb2SaasUserOptionsModel.SetNewRole("bluuser")
				putDb2SaasUserOptionsModel.SetNewEmail("test_user@mycompany.com")
				putDb2SaasUserOptionsModel.SetNewLocked("no")
				putDb2SaasUserOptionsModel.SetNewAuthentication(updateUserAuthenticationModel)
				putDb2SaasUserOptionsModel.SetNewIbmid("test-ibm-id")
				putDb2SaasUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putDb2SaasUserOptionsModel).ToNot(BeNil())
				Expect(putDb2SaasUserOptionsModel.XDeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(putDb2SaasUserOptionsModel.ID).To(Equal(core.StringPtr("test-user")))
				Expect(putDb2SaasUserOptionsModel.NewID).To(Equal(core.StringPtr("test-user")))
				Expect(putDb2SaasUserOptionsModel.NewName).To(Equal(core.StringPtr("test_user")))
				Expect(putDb2SaasUserOptionsModel.NewOldPassword).To(Equal(core.StringPtr("dEkMc43@gfAPl!867^dSbu")))
				Expect(putDb2SaasUserOptionsModel.NewNewPassword).To(Equal(core.StringPtr("ihbgc26@gfAPl!1297^dFGy")))
				Expect(putDb2SaasUserOptionsModel.NewRole).To(Equal(core.StringPtr("bluuser")))
				Expect(putDb2SaasUserOptionsModel.NewEmail).To(Equal(core.StringPtr("test_user@mycompany.com")))
				Expect(putDb2SaasUserOptionsModel.NewLocked).To(Equal(core.StringPtr("no")))
				Expect(putDb2SaasUserOptionsModel.NewAuthentication).To(Equal(updateUserAuthenticationModel))
				Expect(putDb2SaasUserOptionsModel.NewIbmid).To(Equal(core.StringPtr("test-ibm-id")))
				Expect(putDb2SaasUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
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
