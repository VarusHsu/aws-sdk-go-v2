// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a notebook instance.
func (c *Client) DescribeNotebookInstance(ctx context.Context, params *DescribeNotebookInstanceInput, optFns ...func(*Options)) (*DescribeNotebookInstanceOutput, error) {
	stack := middleware.NewStack("DescribeNotebookInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeNotebookInstanceMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeNotebookInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotebookInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeNotebookInstance",
			Err:           err,
		}
	}
	out := result.(*DescribeNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotebookInstanceInput struct {

	// The name of the notebook instance that you want information about.
	//
	// This member is required.
	NotebookInstanceName *string
}

type DescribeNotebookInstanceOutput struct {

	// Whether root access is enabled or disabled for users of the notebook instance.
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for users.
	RootAccess types.RootAccess

	// The name of the Amazon SageMaker notebook instance.
	NotebookInstanceName *string

	// A timestamp. Use this parameter to return the time when the notebook instance
	// was created
	CreationTime *time.Time

	// A timestamp. Use this parameter to retrieve the time when the notebook instance
	// was last modified.
	LastModifiedTime *time.Time

	// The URL that you use to connect to the Jupyter notebook that is running in your
	// notebook instance.
	Url *string

	// The network interface IDs that Amazon SageMaker created at the time of creating
	// the instance.
	NetworkInterfaceId *string

	// A list of the Elastic Inference (EI) instance types associated with this
	// notebook instance. Currently only one EI instance type can be associated with a
	// notebook instance. For more information, see Using Elastic Inference in Amazon
	// SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).
	AcceleratorTypes []types.NotebookInstanceAcceleratorType

	// The Amazon Resource Name (ARN) of the IAM role associated with the instance.
	RoleArn *string

	// The size, in GB, of the ML storage volume attached to the notebook instance.
	VolumeSizeInGB *int32

	// Returns the name of a notebook instance lifecycle configuration. For information
	// about notebook instance lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html)
	NotebookInstanceLifecycleConfigName *string

	// Describes whether Amazon SageMaker provides internet access to the notebook
	// instance. If this value is set to Disabled, the notebook instance does not have
	// internet access, and cannot connect to Amazon SageMaker training and endpoint
	// services. For more information, see Notebook Instances Are Internet-Enabled by
	// Default
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access).
	DirectInternetAccess types.DirectInternetAccess

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string

	// The AWS KMS key ID Amazon SageMaker uses to encrypt data when storing it on the
	// ML storage volume attached to the instance.
	KmsKeyId *string

	// The ID of the VPC subnet.
	SubnetId *string

	// An array of up to three Git repositories associated with the notebook instance.
	// These can be either the names of Git repositories stored as resources in your
	// account, or the URL of Git repositories in AWS CodeCommit
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any
	// other Git repository. These repositories are cloned at the same level as the
	// default repository of your notebook instance. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []*string

	// The type of ML compute instance running on the notebook instance.
	InstanceType types.InstanceType

	// The IDs of the VPC security groups.
	SecurityGroups []*string

	// The Git repository associated with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in AWS CodeCommit
	// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any
	// other Git repository. When you open a notebook instance, it opens in the
	// directory that contains this repository. For more information, see Associating
	// Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string

	// The status of the notebook instance.
	NotebookInstanceStatus types.NotebookInstanceStatus

	// If status is Failed, the reason it failed.
	FailureReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeNotebookInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNotebookInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNotebookInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNotebookInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeNotebookInstance",
	}
}
