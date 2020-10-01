// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Redshift subnet group. You must provide a list of one or
// more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when
// creating Amazon Redshift subnet group. For information about subnet groups, go
// to Amazon Redshift Cluster Subnet Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateClusterSubnetGroup(ctx context.Context, params *CreateClusterSubnetGroupInput, optFns ...func(*Options)) (*CreateClusterSubnetGroupOutput, error) {
	stack := middleware.NewStack("CreateClusterSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateClusterSubnetGroupMiddlewares(stack)
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
	addOpCreateClusterSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClusterSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "CreateClusterSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*CreateClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateClusterSubnetGroupInput struct {

	// A list of tag instances.
	Tags []*types.Tag

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
	// request.
	//
	// This member is required.
	SubnetIds []*string

	// The name for the subnet group. Amazon Redshift stores the value as a lowercase
	// string. Constraints:
	//
	//     * Must contain no more than 255 alphanumeric
	// characters or hyphens.
	//
	//     * Must not be "Default".
	//
	//     * Must be unique for
	// all subnet groups that are created by your AWS account.
	//
	// Example:
	// examplesubnetgroup
	//
	// This member is required.
	ClusterSubnetGroupName *string

	// A description for the subnet group.
	//
	// This member is required.
	Description *string
}

type CreateClusterSubnetGroupOutput struct {

	// Describes a subnet group.
	ClusterSubnetGroup *types.ClusterSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateClusterSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateClusterSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateClusterSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateClusterSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateClusterSubnetGroup",
	}
}
