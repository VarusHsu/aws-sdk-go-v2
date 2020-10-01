// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified virtual tape from the virtual tape shelf (VTS). This
// operation is only supported in the tape gateway type.
func (c *Client) DeleteTapeArchive(ctx context.Context, params *DeleteTapeArchiveInput, optFns ...func(*Options)) (*DeleteTapeArchiveOutput, error) {
	stack := middleware.NewStack("DeleteTapeArchive", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteTapeArchiveMiddlewares(stack)
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
	addOpDeleteTapeArchiveValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTapeArchive(options.Region), middleware.Before)
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
			OperationName: "DeleteTapeArchive",
			Err:           err,
		}
	}
	out := result.(*DeleteTapeArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteTapeArchiveInput
type DeleteTapeArchiveInput struct {

	// The Amazon Resource Name (ARN) of the virtual tape to delete from the virtual
	// tape shelf (VTS).
	//
	// This member is required.
	TapeARN *string
}

// DeleteTapeArchiveOutput
type DeleteTapeArchiveOutput struct {

	// The Amazon Resource Name (ARN) of the virtual tape that was deleted from the
	// virtual tape shelf (VTS).
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteTapeArchiveMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTapeArchive{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTapeArchive{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTapeArchive(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteTapeArchive",
	}
}
