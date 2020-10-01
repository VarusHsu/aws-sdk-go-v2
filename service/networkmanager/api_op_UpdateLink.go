// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the details for an existing link. To remove information for any of the
// parameters, specify an empty string.
func (c *Client) UpdateLink(ctx context.Context, params *UpdateLinkInput, optFns ...func(*Options)) (*UpdateLinkOutput, error) {
	stack := middleware.NewStack("UpdateLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateLinkMiddlewares(stack)
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
	addOpUpdateLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLink(options.Region), middleware.Before)
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
			OperationName: "UpdateLink",
			Err:           err,
		}
	}
	out := result.(*UpdateLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLinkInput struct {

	// The ID of the link.
	//
	// This member is required.
	LinkId *string

	// The upload and download speed in Mbps.
	Bandwidth *types.Bandwidth

	// The type of the link. Length Constraints: Maximum length of 128 characters.
	Type *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The provider of the link. Length Constraints: Maximum length of 128 characters.
	Provider *string

	// A description of the link. Length Constraints: Maximum length of 256 characters.
	Description *string
}

type UpdateLinkOutput struct {

	// Information about the link.
	Link *types.Link

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLink{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "UpdateLink",
	}
}
