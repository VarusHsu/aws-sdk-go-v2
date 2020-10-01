// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Outpost.
func (c *Client) CreateOutpost(ctx context.Context, params *CreateOutpostInput, optFns ...func(*Options)) (*CreateOutpostOutput, error) {
	stack := middleware.NewStack("CreateOutpost", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateOutpostMiddlewares(stack)
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
	addOpCreateOutpostValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOutpost(options.Region), middleware.Before)
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
			OperationName: "CreateOutpost",
			Err:           err,
		}
	}
	out := result.(*CreateOutpostOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOutpostInput struct {

	// The name of the Outpost.
	Name *string

	// The Availability Zone. You must specify AvailabilityZone or AvailabilityZoneId.
	AvailabilityZone *string

	// The ID of the Availability Zone. You must specify AvailabilityZone or
	// AvailabilityZoneId.
	AvailabilityZoneId *string

	// The ID of the site.
	//
	// This member is required.
	SiteId *string

	// The Outpost description.
	Description *string
}

type CreateOutpostOutput struct {

	// Information about an Outpost.
	Outpost *types.Outpost

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateOutpostMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateOutpost{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOutpost{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateOutpost(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "CreateOutpost",
	}
}
