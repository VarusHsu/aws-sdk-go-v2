// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an empty cluster. Each cluster supports five nodes. You use the
// CreateJob () action separately to create the jobs for each of these nodes. The
// cluster does not ship until these five node jobs have been created.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	stack := middleware.NewStack("CreateCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateClusterMiddlewares(stack)
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
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
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
			OperationName: "CreateCluster",
			Err:           err,
		}
	}
	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The ID for the address that you want the cluster shipped to.
	//
	// This member is required.
	AddressId *string

	// An optional description of this specific cluster, for example Environmental Data
	// Cluster-01.
	Description *string

	// The RoleARN that you want to associate with this cluster. RoleArn values are
	// created by using the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) API
	// action in AWS Identity and Access Management (IAM).
	//
	// This member is required.
	RoleARN *string

	// The tax documents required in your AWS Region.
	TaxDocuments *types.TaxDocuments

	// The forwarding address ID for a cluster. This field is not supported in most
	// regions.
	ForwardingAddressId *string

	// The shipping speed for each node in this cluster. This speed doesn't dictate how
	// soon you'll get each Snowball Edge device, rather it represents how quickly each
	// device moves to its destination while in transit. Regional shipping speeds are
	// as follows:
	//
	//     * In Australia, you have access to express shipping. Typically,
	// Snowballs shipped express are delivered in about a day.
	//
	//     * In the European
	// Union (EU), you have access to express shipping. Typically, Snowballs shipped
	// express are delivered in about a day. In addition, most countries in the EU have
	// access to standard shipping, which typically takes less than a week, one way.
	//
	//
	// * In India, Snowballs are delivered in one to seven days.
	//
	//     * In the United
	// States of America (US), you have access to one-day shipping and two-day
	// shipping.
	//
	//     <ul> <li> <p>In Australia, you have access to express shipping.
	// Typically, devices shipped express are delivered in about a day.</p> </li> <li>
	// <p>In the European Union (EU), you have access to express shipping. Typically,
	// Snowball Edges shipped express are delivered in about a day. In addition, most
	// countries in the EU have access to standard shipping, which typically takes less
	// than a week, one way.</p> </li> <li> <p>In India, Snowball Edges are delivered
	// in one to seven days.</p> </li> <li> <p>In the US, you have access to one-day
	// shipping and two-day shipping.</p> </li> </ul>
	//
	// This member is required.
	ShippingOption types.ShippingOption

	// The type of AWS Snowball device to use for this cluster. </p> <note> <p>For
	// cluster jobs, AWS Snowball currently supports only the <code>EDGE</code> device
	// type.</p> </note>
	SnowballType types.SnowballType

	// The Amazon Simple Notification Service (Amazon SNS) notification settings for
	// this cluster.
	Notification *types.Notification

	// The KmsKeyARN value that you want to associate with this cluster. KmsKeyARN
	// values are created by using the CreateKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html) API
	// action in AWS Key Management Service (AWS KMS).
	KmsKeyARN *string

	// The type of job for this cluster. Currently, the only job type supported for
	// clusters is LOCAL_USE.
	//
	// This member is required.
	JobType types.JobType

	// The resources associated with the cluster job. These resources include Amazon S3
	// buckets and optional AWS Lambda functions written in the Python language.
	//
	// This member is required.
	Resources *types.JobResource
}

type CreateClusterOutput struct {

	// The automatically generated ID for a cluster.
	ClusterId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateCluster",
	}
}
