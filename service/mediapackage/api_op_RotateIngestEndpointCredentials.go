// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rotate the IngestEndpoint's username and password, as specified by the
// IngestEndpoint's id.
func (c *Client) RotateIngestEndpointCredentials(ctx context.Context, params *RotateIngestEndpointCredentialsInput, optFns ...func(*Options)) (*RotateIngestEndpointCredentialsOutput, error) {
	stack := middleware.NewStack("RotateIngestEndpointCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRotateIngestEndpointCredentialsMiddlewares(stack)
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
	addOpRotateIngestEndpointCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRotateIngestEndpointCredentials(options.Region), middleware.Before)
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
			OperationName: "RotateIngestEndpointCredentials",
			Err:           err,
		}
	}
	out := result.(*RotateIngestEndpointCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateIngestEndpointCredentialsInput struct {

	// The ID of the channel the IngestEndpoint is on.
	//
	// This member is required.
	Id *string

	// The id of the IngestEndpoint whose credentials should be rotated
	//
	// This member is required.
	IngestEndpointId *string
}

type RotateIngestEndpointCredentialsOutput struct {

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest

	// A collection of tags associated with a resource
	Tags map[string]*string

	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// A short text description of the Channel.
	Description *string

	// The ID of the Channel.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRotateIngestEndpointCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRotateIngestEndpointCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRotateIngestEndpointCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opRotateIngestEndpointCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "RotateIngestEndpointCredentials",
	}
}
