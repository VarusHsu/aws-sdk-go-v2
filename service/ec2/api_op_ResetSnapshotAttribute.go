// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets permission settings for the specified snapshot. For more information
// about modifying snapshot permissions, see Sharing Snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ResetSnapshotAttribute(ctx context.Context, params *ResetSnapshotAttributeInput, optFns ...func(*Options)) (*ResetSnapshotAttributeOutput, error) {
	stack := middleware.NewStack("ResetSnapshotAttribute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpResetSnapshotAttributeMiddlewares(stack)
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
	addOpResetSnapshotAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetSnapshotAttribute(options.Region), middleware.Before)
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
			OperationName: "ResetSnapshotAttribute",
			Err:           err,
		}
	}
	out := result.(*ResetSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetSnapshotAttributeInput struct {

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The attribute to reset. Currently, only the attribute for permission to create
	// volumes can be reset.
	//
	// This member is required.
	Attribute types.SnapshotAttributeName
}

type ResetSnapshotAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpResetSnapshotAttributeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpResetSnapshotAttribute{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpResetSnapshotAttribute{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetSnapshotAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ResetSnapshotAttribute",
	}
}
