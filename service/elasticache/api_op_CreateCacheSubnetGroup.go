// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new cache subnet group. Use this parameter only when you are creating
// a cluster in an Amazon Virtual Private Cloud (Amazon VPC).
func (c *Client) CreateCacheSubnetGroup(ctx context.Context, params *CreateCacheSubnetGroupInput, optFns ...func(*Options)) (*CreateCacheSubnetGroupOutput, error) {
	stack := middleware.NewStack("CreateCacheSubnetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateCacheSubnetGroupMiddlewares(stack)
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
	addOpCreateCacheSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheSubnetGroup(options.Region), middleware.Before)
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
			OperationName: "CreateCacheSubnetGroup",
			Err:           err,
		}
	}
	out := result.(*CreateCacheSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheSubnetGroup operation.
type CreateCacheSubnetGroupInput struct {

	// A list of VPC subnet IDs for the cache subnet group.
	//
	// This member is required.
	SubnetIds []*string

	// A description for the cache subnet group.
	//
	// This member is required.
	CacheSubnetGroupDescription *string

	// A name for the cache subnet group. This value is stored as a lowercase string.
	// Constraints: Must contain no more than 255 alphanumeric characters or hyphens.
	// Example: mysubnetgroup
	//
	// This member is required.
	CacheSubnetGroupName *string
}

type CreateCacheSubnetGroupOutput struct {

	// Represents the output of one of the following operations:
	//
	//     *
	// CreateCacheSubnetGroup
	//
	//     * ModifyCacheSubnetGroup
	CacheSubnetGroup *types.CacheSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateCacheSubnetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheSubnetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheSubnetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCacheSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateCacheSubnetGroup",
	}
}
