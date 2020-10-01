// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Indicates whether the specified access point currently has a policy that allows
// public access. For more information about public access through access points,
// see Managing Data Access with Amazon S3 Access Points
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points.html) in the
// Amazon Simple Storage Service Developer Guide.
func (c *Client) GetAccessPointPolicyStatus(ctx context.Context, params *GetAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusOutput, error) {
	stack := middleware.NewStack("GetAccessPointPolicyStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetAccessPointPolicyStatusMiddlewares(stack)
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
	addOpGetAccessPointPolicyStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "GetAccessPointPolicyStatus",
			Err:           err,
		}
	}
	out := result.(*GetAccessPointPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyStatusInput struct {

	// The name of the access point whose policy status you want to retrieve.
	//
	// This member is required.
	Name *string

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string
}

type GetAccessPointPolicyStatusOutput struct {

	// Indicates the current policy status of the specified access point.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetAccessPointPolicyStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicyStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicyStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointPolicyStatus",
	}
}
