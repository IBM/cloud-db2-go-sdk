package db2saasv1_test

import (
	"log"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var _ = Describe(`DB2 saas v1 integration tests`, func() {
	var (
		err              error
		serviceURL       string
		deploymentID     string
		encodedCRN       string
		db2saasV1Service *db2saasv1.Db2saasV1
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`Client Initialization`, func() {
		It(`Successfully construct the service client instance`, func() {
			db2saasV1ServiceOptions := &db2saasv1.Db2saasV1Options{}

			db2saasV1Service, err = db2saasv1.NewDb2saasV1(db2saasV1ServiceOptions)

			Expect(err).To(BeNil())
			Expect(db2saasV1Service).ToNot(BeNil())
			Expect(db2saasV1Service.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			db2saasV1Service.EnableRetries(4, 30*time.Second)
		})
	})

	// ------------------ CONNECTIONINFO -----------------
	Describe(`Get connection info`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
	})

	// ------------------ CONNECTIONINFO - SUCCESS -----------------
	It(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions)`, func() {
		getDb2saasConnectionInfoOptions := &db2saasv1.GetDb2SaasConnectionInfoOptions{
			DeploymentID:  &encodedCRN,
			XDeploymentID: &deploymentID,
		}

		successResult, detailedResponse, err := db2saasV1Service.GetDb2SaasConnectionInfo(getDb2saasConnectionInfoOptions)

		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(http.StatusOK))
		Expect(successResult).ToNot(BeNil())

	})

	// ------------------ AUTOSCALING -----------------
	Describe(`Update Autoscale configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
	})

	// ------------------ AUTOSCALING- SUCCESS WITH ALL THE POSSIBLE UPDATABLE FIELDS -----------------
	Context(`Update all the possible autoscale configurations`, func() {
		It(`Success case for all the possible updation of autoscale configurations`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID:             &deploymentID,
				AutoScalingEnabled:        core.StringPtr("YES"),
				AutoScalingThreshold:      core.Int64Ptr(90),
				AutoScalingOverTimePeriod: core.Float64Ptr(5),
				AutoScalingPauseLimit:     core.Int64Ptr(70),
				AutoScalingAllowPlanLimit: core.StringPtr("YES"),
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(successResult).ToNot(BeNil())
		})
	})

	// ------------------ AUTOSCALING - SUCCESS WITH EMPTY PAYLOAD -----------------
	Context(`Update autoscale configuration with empty payload`, func() {
		It(`Success case with empty payload`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID: &deploymentID,
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusOK))
			Expect(successResult).ToNot(BeNil())
		})
	})

	// ------------------ AUTOSCALING - FAILURE WITH INVALID VALUE FOR AutoScalingEnabled - IT HAS TO BE ONE OF[YES NO] -----------------
	Context(`Invalid value for AutoScalingEnabled`, func() {
		It(`Failure case for invalid value of AutoScaleEnabled`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID:      &deploymentID,
				AutoScalingEnabled: core.StringPtr("YESs"),
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusUnprocessableEntity))
			Expect(successResult).To(BeNil())
		})
	})

	// ------------------ AUTOSCALING - FAILURE WITH INVALID VALUE FOR AutoScalingThreshold - IT HAS TO BE GREATER THAN 0 -----------------
	Context(`Invalid value for AutoScalingThreshold`, func() {
		It(`Failure case for invalid value of AutoScalingThreshold`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID:        &deploymentID,
				AutoScalingEnabled:   core.StringPtr("YES"),
				AutoScalingThreshold: core.Int64Ptr(0),
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusUnprocessableEntity))
			Expect(successResult).To(BeNil())
		})
	})

	// ------------------ AUTOSCALING - FAILURE WITH INVALID VALUE FOR AutoScalingOverTimePeriod - IT HAS TO BE ONE OF[5 15 30 45 60] -----------------
	Context(`Invalid value for AutoScalingOverTimePeriod`, func() {
		It(`Failure case for invalid value of AutoScalingOverTimePeriod`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID:             &deploymentID,
				AutoScalingEnabled:        core.StringPtr("YES"),
				AutoScalingThreshold:      core.Int64Ptr(70),
				AutoScalingOverTimePeriod: core.Float64Ptr(51),
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusUnprocessableEntity))
			Expect(successResult).To(BeNil())
		})
	})

	// ------------------ AUTOSCALING - FAILURE WITH INVALID VALUE FOR AutoScalingAllowPlanLimit - IT HAS TO BE ONE OF[YES NO] -----------------
	Context(`Invalid value for AutoScalingAllowPlanLimit`, func() {
		It(`Failure case for invalid value of AutoScalingAllowPlanLimit`, func() {
			updateScalingOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDeploymentID:             &deploymentID,
				AutoScalingEnabled:        core.StringPtr("YES"),
				AutoScalingThreshold:      core.Int64Ptr(70),
				AutoScalingOverTimePeriod: core.Float64Ptr(5),
				AutoScalingAllowPlanLimit: core.StringPtr("Noo"),
			}

			successResult, detailedResponse, err := db2saasV1Service.PutDb2SaasAutoscale(updateScalingOptions)

			Expect(err).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(http.StatusUnprocessableEntity))
			Expect(successResult).To(BeNil())
		})
	})

})
