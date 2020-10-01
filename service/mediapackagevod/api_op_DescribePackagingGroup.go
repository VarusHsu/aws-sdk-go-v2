// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a description of a MediaPackage VOD PackagingGroup resource.
func (c *Client) DescribePackagingGroup(ctx context.Context, params *DescribePackagingGroupInput, optFns ...func(*Options)) (*DescribePackagingGroupOutput, error) {
	stack := middleware.NewStack("DescribePackagingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribePackagingGroupMiddlewares(stack)
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
	addOpDescribePackagingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackagingGroup(options.Region), middleware.Before)
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
			OperationName: "DescribePackagingGroup",
			Err:           err,
		}
	}
	out := result.(*DescribePackagingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackagingGroupInput struct {

	// The ID of a MediaPackage VOD PackagingGroup resource.
	//
	// This member is required.
	Id *string
}

type DescribePackagingGroupOutput struct {

	// CDN Authorization credentials
	Authorization *types.Authorization

	// A collection of tags associated with a resource
	Tags map[string]*string

	// The fully qualified domain name for Assets in the PackagingGroup.
	DomainName *string

	// The ARN of the PackagingGroup.
	Arn *string

	// The ID of the PackagingGroup.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribePackagingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackagingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackagingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePackagingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "DescribePackagingGroup",
	}
}
