// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elasticbeanstalkiface provides an interface to enable mocking the AWS Elastic Beanstalk service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticbeanstalkiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

// ElasticBeanstalkAPI provides an interface to enable mocking the
// elasticbeanstalk.ElasticBeanstalk service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Elastic Beanstalk.
//    func myFunc(svc elasticbeanstalkiface.ElasticBeanstalkAPI) bool {
//        // Make svc.AbortEnvironmentUpdate request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elasticbeanstalk.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockElasticBeanstalkClient struct {
//        elasticbeanstalkiface.ElasticBeanstalkAPI
//    }
//    func (m *mockElasticBeanstalkClient) AbortEnvironmentUpdate(input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockElasticBeanstalkClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ElasticBeanstalkAPI interface {
	AbortEnvironmentUpdateRequest(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*request.Request, *elasticbeanstalk.AbortEnvironmentUpdateOutput)

	AbortEnvironmentUpdate(*elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)

	ApplyEnvironmentManagedActionRequest(*elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*request.Request, *elasticbeanstalk.ApplyEnvironmentManagedActionOutput)

	ApplyEnvironmentManagedAction(*elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error)

	CheckDNSAvailabilityRequest(*elasticbeanstalk.CheckDNSAvailabilityInput) (*request.Request, *elasticbeanstalk.CheckDNSAvailabilityOutput)

	CheckDNSAvailability(*elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)

	ComposeEnvironmentsRequest(*elasticbeanstalk.ComposeEnvironmentsInput) (*request.Request, *elasticbeanstalk.EnvironmentDescriptionsMessage)

	ComposeEnvironments(*elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)

	CreateApplicationRequest(*elasticbeanstalk.CreateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	CreateApplication(*elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	CreateApplicationVersionRequest(*elasticbeanstalk.CreateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	CreateApplicationVersion(*elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	CreateConfigurationTemplateRequest(*elasticbeanstalk.CreateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	CreateConfigurationTemplate(*elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	CreateEnvironmentRequest(*elasticbeanstalk.CreateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	CreateEnvironment(*elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	CreatePlatformVersionRequest(*elasticbeanstalk.CreatePlatformVersionInput) (*request.Request, *elasticbeanstalk.CreatePlatformVersionOutput)

	CreatePlatformVersion(*elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error)

	CreateStorageLocationRequest(*elasticbeanstalk.CreateStorageLocationInput) (*request.Request, *elasticbeanstalk.CreateStorageLocationOutput)

	CreateStorageLocation(*elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)

	DeleteApplicationRequest(*elasticbeanstalk.DeleteApplicationInput) (*request.Request, *elasticbeanstalk.DeleteApplicationOutput)

	DeleteApplication(*elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)

	DeleteApplicationVersionRequest(*elasticbeanstalk.DeleteApplicationVersionInput) (*request.Request, *elasticbeanstalk.DeleteApplicationVersionOutput)

	DeleteApplicationVersion(*elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)

	DeleteConfigurationTemplateRequest(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.DeleteConfigurationTemplateOutput)

	DeleteConfigurationTemplate(*elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)

	DeleteEnvironmentConfigurationRequest(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*request.Request, *elasticbeanstalk.DeleteEnvironmentConfigurationOutput)

	DeleteEnvironmentConfiguration(*elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)

	DeletePlatformVersionRequest(*elasticbeanstalk.DeletePlatformVersionInput) (*request.Request, *elasticbeanstalk.DeletePlatformVersionOutput)

	DeletePlatformVersion(*elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error)

	DescribeApplicationVersionsRequest(*elasticbeanstalk.DescribeApplicationVersionsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationVersionsOutput)

	DescribeApplicationVersions(*elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)

	DescribeApplicationsRequest(*elasticbeanstalk.DescribeApplicationsInput) (*request.Request, *elasticbeanstalk.DescribeApplicationsOutput)

	DescribeApplications(*elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)

	DescribeConfigurationOptionsRequest(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationOptionsOutput)

	DescribeConfigurationOptions(*elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)

	DescribeConfigurationSettingsRequest(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.DescribeConfigurationSettingsOutput)

	DescribeConfigurationSettings(*elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)

	DescribeEnvironmentHealthRequest(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentHealthOutput)

	DescribeEnvironmentHealth(*elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)

	DescribeEnvironmentManagedActionHistoryRequest(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput)

	DescribeEnvironmentManagedActionHistory(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)

	DescribeEnvironmentManagedActionsRequest(*elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentManagedActionsOutput)

	DescribeEnvironmentManagedActions(*elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)

	DescribeEnvironmentResourcesRequest(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*request.Request, *elasticbeanstalk.DescribeEnvironmentResourcesOutput)

	DescribeEnvironmentResources(*elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)

	DescribeEnvironmentsRequest(*elasticbeanstalk.DescribeEnvironmentsInput) (*request.Request, *elasticbeanstalk.EnvironmentDescriptionsMessage)

	DescribeEnvironments(*elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)

	DescribeEventsRequest(*elasticbeanstalk.DescribeEventsInput) (*request.Request, *elasticbeanstalk.DescribeEventsOutput)

	DescribeEvents(*elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)

	DescribeEventsPages(*elasticbeanstalk.DescribeEventsInput, func(*elasticbeanstalk.DescribeEventsOutput, bool) bool) error

	DescribeInstancesHealthRequest(*elasticbeanstalk.DescribeInstancesHealthInput) (*request.Request, *elasticbeanstalk.DescribeInstancesHealthOutput)

	DescribeInstancesHealth(*elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)

	DescribePlatformVersionRequest(*elasticbeanstalk.DescribePlatformVersionInput) (*request.Request, *elasticbeanstalk.DescribePlatformVersionOutput)

	DescribePlatformVersion(*elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error)

	ListAvailableSolutionStacksRequest(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*request.Request, *elasticbeanstalk.ListAvailableSolutionStacksOutput)

	ListAvailableSolutionStacks(*elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)

	ListPlatformVersionsRequest(*elasticbeanstalk.ListPlatformVersionsInput) (*request.Request, *elasticbeanstalk.ListPlatformVersionsOutput)

	ListPlatformVersions(*elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error)

	RebuildEnvironmentRequest(*elasticbeanstalk.RebuildEnvironmentInput) (*request.Request, *elasticbeanstalk.RebuildEnvironmentOutput)

	RebuildEnvironment(*elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)

	RequestEnvironmentInfoRequest(*elasticbeanstalk.RequestEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RequestEnvironmentInfoOutput)

	RequestEnvironmentInfo(*elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)

	RestartAppServerRequest(*elasticbeanstalk.RestartAppServerInput) (*request.Request, *elasticbeanstalk.RestartAppServerOutput)

	RestartAppServer(*elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)

	RetrieveEnvironmentInfoRequest(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*request.Request, *elasticbeanstalk.RetrieveEnvironmentInfoOutput)

	RetrieveEnvironmentInfo(*elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)

	SwapEnvironmentCNAMEsRequest(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*request.Request, *elasticbeanstalk.SwapEnvironmentCNAMEsOutput)

	SwapEnvironmentCNAMEs(*elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)

	TerminateEnvironmentRequest(*elasticbeanstalk.TerminateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	TerminateEnvironment(*elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	UpdateApplicationRequest(*elasticbeanstalk.UpdateApplicationInput) (*request.Request, *elasticbeanstalk.ApplicationDescriptionMessage)

	UpdateApplication(*elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)

	UpdateApplicationResourceLifecycleRequest(*elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*request.Request, *elasticbeanstalk.UpdateApplicationResourceLifecycleOutput)

	UpdateApplicationResourceLifecycle(*elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error)

	UpdateApplicationVersionRequest(*elasticbeanstalk.UpdateApplicationVersionInput) (*request.Request, *elasticbeanstalk.ApplicationVersionDescriptionMessage)

	UpdateApplicationVersion(*elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)

	UpdateConfigurationTemplateRequest(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*request.Request, *elasticbeanstalk.ConfigurationSettingsDescription)

	UpdateConfigurationTemplate(*elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)

	UpdateEnvironmentRequest(*elasticbeanstalk.UpdateEnvironmentInput) (*request.Request, *elasticbeanstalk.EnvironmentDescription)

	UpdateEnvironment(*elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)

	ValidateConfigurationSettingsRequest(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*request.Request, *elasticbeanstalk.ValidateConfigurationSettingsOutput)

	ValidateConfigurationSettings(*elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
}

var _ ElasticBeanstalkAPI = (*elasticbeanstalk.ElasticBeanstalk)(nil)
