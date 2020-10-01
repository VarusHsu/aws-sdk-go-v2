// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about an application revision.
func (c *Client) GetApplicationRevision(ctx context.Context, params *GetApplicationRevisionInput, optFns ...func(*Options)) (*GetApplicationRevisionOutput, error) {
	stack := middleware.NewStack("GetApplicationRevision", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetApplicationRevisionMiddlewares(stack)
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
	addOpGetApplicationRevisionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationRevision(options.Region), middleware.Before)
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
			OperationName: "GetApplicationRevision",
			Err:           err,
		}
	}
	out := result.(*GetApplicationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetApplicationRevision operation.
type GetApplicationRevisionInput struct {

	// The name of the application that corresponds to the revision.
	//
	// This member is required.
	ApplicationName *string

	// Information about the application revision to get, including type and location.
	//
	// This member is required.
	Revision *types.RevisionLocation
}

// Represents the output of a GetApplicationRevision operation.
type GetApplicationRevisionOutput struct {

	// General information about the revision.
	RevisionInfo *types.GenericRevisionInfo

	// Additional information about the revision, including type and location.
	Revision *types.RevisionLocation

	// The name of the application that corresponds to the revision.
	ApplicationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetApplicationRevisionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetApplicationRevision{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetApplicationRevision{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetApplicationRevision(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetApplicationRevision",
	}
}
