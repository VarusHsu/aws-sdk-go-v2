// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a snapshot of the application's state data.
func (c *Client) CreateApplicationSnapshot(ctx context.Context, params *CreateApplicationSnapshotInput, optFns ...func(*Options)) (*CreateApplicationSnapshotOutput, error) {
	stack := middleware.NewStack("CreateApplicationSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateApplicationSnapshotMiddlewares(stack)
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
	addOpCreateApplicationSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationSnapshot(options.Region), middleware.Before)
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
			OperationName: "CreateApplicationSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateApplicationSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationSnapshotInput struct {

	// An identifier for the application snapshot.
	//
	// This member is required.
	SnapshotName *string

	// The name of an existing application
	//
	// This member is required.
	ApplicationName *string
}

type CreateApplicationSnapshotOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateApplicationSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApplicationSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApplicationSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApplicationSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "CreateApplicationSnapshot",
	}
}
