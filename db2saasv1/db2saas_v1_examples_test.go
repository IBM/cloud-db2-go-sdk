//go:build examples

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
	 "encoding/json"
	 "fmt"
	 "os"
 
	 "github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	 "github.com/IBM/go-sdk-core/v5/core"
	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
 )
 
 //
 // This file provides an example of how to use the db2saas service.
 //
 // The following configuration properties are assumed to be defined:
 // DB2SAAS_URL=<service base url>
 // DB2SAAS_AUTH_TYPE=iam
 // DB2SAAS_APIKEY=<IAM apikey>
 // DB2SAAS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
 //
 // These configuration properties can be exported as environment variables, or stored
 // in a configuration file and then:
 // export IBM_CREDENTIALS_FILE=<name of configuration file>
 //
 var _ = Describe(`Db2saasV1 Examples Tests`, func() {
 
	 const externalConfigFile = "../db2saas_v1.env"
 
	 var (
		 db2saasService *db2saasv1.Db2saasV1
		 config       map[string]string
	 )
 
	 var shouldSkipTest = func() {
		 Skip("External configuration is not available, skipping examples...")
	 }
 
	 Describe(`External configuration`, func() {
		 It("Successfully load the configuration", func() {
			 var err error
			 _, err = os.Stat(externalConfigFile)
			 if err != nil {
				 Skip("External configuration file not found, skipping examples: " + err.Error())
			 }
 
			 os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			 config, err = core.GetServiceProperties(db2saasv1.DefaultServiceName)
			 if err != nil {
				 Skip("Error loading service properties, skipping examples: " + err.Error())
			 } else if len(config) == 0 {
				 Skip("Unable to load service properties, skipping examples")
			 }
 
			 shouldSkipTest = func() {}
		 })
	 })
 
	 Describe(`Client initialization`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It("Successfully construct the service client instance", func() {
			 var err error
 
			 // begin-common
 
			 db2saasServiceOptions := &db2saasv1.Db2saasV1Options{}
 
			 db2saasService, err = db2saasv1.NewDb2saasV1UsingExternalConfig(db2saasServiceOptions)
 
			 if err != nil {
				 panic(err)
			 }
 
			 // end-common
 
			 Expect(db2saasService).ToNot(BeNil())
		 })
	 })
 
	 Describe(`Db2saasV1 request examples`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasConnectionInfo request example`, func() {
			 fmt.Println("\nGetDb2SaasConnectionInfo() result:")
			 // begin-get_db2_saas_connection_info
 
			 getDb2SaasConnectionInfoOptions := db2saasService.NewGetDb2SaasConnectionInfoOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A",
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successConnectionInfo, response, err := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successConnectionInfo, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_connection_info
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successConnectionInfo).ToNot(BeNil())
		 })
		 It(`PostDb2SaasAllowlist request example`, func() {
			 fmt.Println("\nPostDb2SaasAllowlist() result:")
			 // begin-post_db2_saas_allowlist
 
			 ipAddressModel := &db2saasv1.IpAddress{
				 Address: core.StringPtr("127.0.0.1"),
				 Description: core.StringPtr("A sample IP address"),
			 }
 
			 postDb2SaasAllowlistOptions := db2saasService.NewPostDb2SaasAllowlistOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
				 []db2saasv1.IpAddress{*ipAddressModel},
			 )
 
			 successPostAllowedlistIPs, response, err := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successPostAllowedlistIPs, "", "  ")
			 fmt.Println(string(b))
 
			 // end-post_db2_saas_allowlist
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successPostAllowedlistIPs).ToNot(BeNil())
		 })
		 It(`GetDb2SaasAllowlist request example`, func() {
			 fmt.Println("\nGetDb2SaasAllowlist() result:")
			 // begin-get_db2_saas_allowlist
 
			 getDb2SaasAllowlistOptions := db2saasService.NewGetDb2SaasAllowlistOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successGetAllowlistIPs, response, err := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successGetAllowlistIPs, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_allowlist
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetAllowlistIPs).ToNot(BeNil())
		 })
		 It(`PostDb2SaasUser request example`, func() {
			 fmt.Println("\nPostDb2SaasUser() result:")
			 // begin-post_db2_saas_user
 
			 createUserAuthenticationModel := &db2saasv1.CreateUserAuthentication{
				 Method: core.StringPtr("internal"),
				 PolicyID: core.StringPtr("Default"),
			 }
 
			 postDb2SaasUserOptions := db2saasService.NewPostDb2SaasUserOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
				 "test-user",
				 false,
				 "test-ibm-id",
				 "test_user",
				 "dEkMc43@gfAPl!867^dSbu",
				 "bluuser",
				 "test_user@mycompany.com",
				 "no",
				 createUserAuthenticationModel,
			 )
 
			 successUserResponse, response, err := db2saasService.PostDb2SaasUser(postDb2SaasUserOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successUserResponse, "", "  ")
			 fmt.Println(string(b))
 
			 // end-post_db2_saas_user
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successUserResponse).ToNot(BeNil())
		 })
		 It(`GetDb2SaasUser request example`, func() {
			 fmt.Println("\nGetDb2SaasUser() result:")
			 // begin-get_db2_saas_user
 
			 getDb2SaasUserOptions := db2saasService.NewGetDb2SaasUserOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successGetUserInfo, response, err := db2saasService.GetDb2SaasUser(getDb2SaasUserOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successGetUserInfo, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_user
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetUserInfo).ToNot(BeNil())
		 })
		 It(`GetbyidDb2SaasUser request example`, func() {
			 fmt.Println("\nGetbyidDb2SaasUser() result:")
			 // begin-getbyid_db2_saas_user
 
			 getbyidDb2SaasUserOptions := db2saasService.NewGetbyidDb2SaasUserOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successGetUserByID, response, err := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successGetUserByID, "", "  ")
			 fmt.Println(string(b))
 
			 // end-getbyid_db2_saas_user
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetUserByID).ToNot(BeNil())
		 })
		 It(`PutDb2SaasAutoscale request example`, func() {
			 fmt.Println("\nPutDb2SaasAutoscale() result:")
			 // begin-put_db2_saas_autoscale
 
			 putDb2SaasAutoscaleOptions := db2saasService.NewPutDb2SaasAutoscaleOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successUpdateAutoScale, response, err := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successUpdateAutoScale, "", "  ")
			 fmt.Println(string(b))
 
			 // end-put_db2_saas_autoscale
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successUpdateAutoScale).ToNot(BeNil())
		 })
		 It(`GetDb2SaasAutoscale request example`, func() {
			 fmt.Println("\nGetDb2SaasAutoscale() result:")
			 // begin-get_db2_saas_autoscale
 
			 getDb2SaasAutoscaleOptions := db2saasService.NewGetDb2SaasAutoscaleOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successAutoScaling, response, err := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successAutoScaling, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_autoscale
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successAutoScaling).ToNot(BeNil())
		 })
		 It(`DeleteDb2SaasUser request example`, func() {
			 fmt.Println("\nDeleteDb2SaasUser() result:")
			 // begin-delete_db2_saas_user
 
			 deleteDb2SaasUserOptions := db2saasService.NewDeleteDb2SaasUserOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
				 "test-user",
			 )
 
			 result, response, err := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(result, "", "  ")
			 fmt.Println(string(b))
 
			 // end-delete_db2_saas_user
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(result).ToNot(BeNil())
		 })
	 })
 })
 