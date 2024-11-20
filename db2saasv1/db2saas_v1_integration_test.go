//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the db2saasv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`Db2saasV1 Integration Tests`, func() {
	const externalConfigFile = "../db2saas_v1.env"

	var (
		err          error
		db2saasService *db2saasv1.Db2saasV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(db2saasv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			db2saasServiceOptions := &db2saasv1.Db2saasV1Options{}

			db2saasService, err = db2saasv1.NewDb2saasV1UsingExternalConfig(db2saasServiceOptions)
			Expect(err).To(BeNil())
			Expect(db2saasService).ToNot(BeNil())
			Expect(db2saasService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			db2saasService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetDb2SaasConnectionInfo - Get Db2 connection information`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions)`, func() {
			getDb2SaasConnectionInfoOptions := &db2saasv1.GetDb2SaasConnectionInfoOptions{
				DeploymentID: core.StringPtr("testString"),
				XDeploymentID: core.StringPtr("testString"),
			}

			successConnectionInfo, response, err := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successConnectionInfo).ToNot(BeNil())
		})
	})

	Describe(`PostDb2SaasWhitelist - Whitelisting of new IPs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDb2SaasWhitelist(postDb2SaasWhitelistOptions *PostDb2SaasWhitelistOptions)`, func() {
			ipAddressModel := &db2saasv1.IpAddress{
				Address: core.StringPtr("127.0.0.1"),
				Description: core.StringPtr("A sample IP address"),
			}

			postDb2SaasWhitelistOptions := &db2saasv1.PostDb2SaasWhitelistOptions{
				XDeploymentID: core.StringPtr("testString"),
				IpAddresses: []db2saasv1.IpAddress{*ipAddressModel},
			}

			successPostWhitelistIPs, response, err := db2saasService.PostDb2SaasWhitelist(postDb2SaasWhitelistOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successPostWhitelistIPs).ToNot(BeNil())
		})
	})

	Describe(`GetDb2SaasWhitelist - Get whitelisted IPs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDb2SaasWhitelist(getDb2SaasWhitelistOptions *GetDb2SaasWhitelistOptions)`, func() {
			getDb2SaasWhitelistOptions := &db2saasv1.GetDb2SaasWhitelistOptions{
				XDeploymentID: core.StringPtr("testString"),
			}

			successGetWhitelistIPs, response, err := db2saasService.GetDb2SaasWhitelist(getDb2SaasWhitelistOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successGetWhitelistIPs).ToNot(BeNil())
		})
	})

	Describe(`PostDb2SaasUser - Create new user ( available only for platform users)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions)`, func() {
			createUserAuthenticationModel := &db2saasv1.CreateUserAuthentication{
				Method: core.StringPtr("internal"),
				PolicyID: core.StringPtr("Default"),
			}

			postDb2SaasUserOptions := &db2saasv1.PostDb2SaasUserOptions{
				XDeploymentID: core.StringPtr("testString"),
				ID: core.StringPtr("test-user"),
				Iam: core.BoolPtr(false),
				Ibmid: core.StringPtr("test-ibm-id"),
				Name: core.StringPtr("test_user"),
				Password: core.StringPtr("dEkMc43@gfAPl!867^dSbu"),
				Role: core.StringPtr("bluuser"),
				Email: core.StringPtr("test_user@mycompany.com"),
				Locked: core.StringPtr("no"),
				Authentication: createUserAuthenticationModel,
			}

			successUserResponse, response, err := db2saasService.PostDb2SaasUser(postDb2SaasUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successUserResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDb2SaasUser - Get the list of Users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions)`, func() {
			getDb2SaasUserOptions := &db2saasv1.GetDb2SaasUserOptions{
				XDeploymentID: core.StringPtr("testString"),
			}

			successGetUserInfo, response, err := db2saasService.GetDb2SaasUser(getDb2SaasUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successGetUserInfo).ToNot(BeNil())
		})
	})

	Describe(`PutDb2SaasUser - Update the details of existing user`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutDb2SaasUser(putDb2SaasUserOptions *PutDb2SaasUserOptions)`, func() {
			updateUserAuthenticationModel := &db2saasv1.UpdateUserAuthentication{
				Method: core.StringPtr("internal"),
				PolicyID: core.StringPtr("Default"),
			}

			putDb2SaasUserOptions := &db2saasv1.PutDb2SaasUserOptions{
				XDeploymentID: core.StringPtr("testString"),
				ID: core.StringPtr("test-user"),
				NewID: core.StringPtr("test-user"),
				NewName: core.StringPtr("test_user"),
				NewOldPassword: core.StringPtr("dEkMc43@gfAPl!867^dSbu"),
				NewNewPassword: core.StringPtr("ihbgc26@gfAPl!1297^dFGy"),
				NewRole: core.StringPtr("bluuser"),
				NewEmail: core.StringPtr("test_user@mycompany.com"),
				NewLocked: core.StringPtr("no"),
				NewAuthentication: updateUserAuthenticationModel,
				NewIbmid: core.StringPtr("test-ibm-id"),
			}

			successUserResponse, response, err := db2saasService.PutDb2SaasUser(putDb2SaasUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successUserResponse).ToNot(BeNil())
		})
	})

	Describe(`GetbyidDb2SaasUser - Get specific user by Id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions)`, func() {
			getbyidDb2SaasUserOptions := &db2saasv1.GetbyidDb2SaasUserOptions{
				XDeploymentID: core.StringPtr("testString"),
			}

			successGetUserByID, response, err := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successGetUserByID).ToNot(BeNil())
		})
	})

	Describe(`PutDb2SaasAutoscale - Update auto scaling configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions)`, func() {
			putDb2SaasAutoscaleOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID: core.StringPtr("testString"),
				AutoScalingThreshold: core.Int64Ptr(int64(90)),
				AutoScalingPauseLimit: core.Int64Ptr(int64(70)),
			}

			successUpdateAutoScale, response, err := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successUpdateAutoScale).ToNot(BeNil())
		})
	})

	Describe(`GetDb2SaasAutoscale - Get auto scaling info`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions)`, func() {
			getDb2SaasAutoscaleOptions := &db2saasv1.GetDb2SaasAutoscaleOptions{
				XDeploymentID: core.StringPtr("testString"),
			}

			successAutoScaling, response, err := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(successAutoScaling).ToNot(BeNil())
		})
	})

	Describe(`DeleteDb2SaasUser - Delete a user (only platform admin)`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions)`, func() {
			deleteDb2SaasUserOptions := &db2saasv1.DeleteDb2SaasUserOptions{
				XDeploymentID: core.StringPtr("testString"),
				ID: core.StringPtr("test-user"),
			}

			result, response, err := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
