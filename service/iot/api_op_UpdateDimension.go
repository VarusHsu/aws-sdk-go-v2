// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the definition for a dimension. You cannot change the type of a
// dimension after it is created (you can delete it and re-create it).
func (c *Client) UpdateDimension(ctx context.Context, params *UpdateDimensionInput, optFns ...func(*Options)) (*UpdateDimensionOutput, error) {
	stack := middleware.NewStack("UpdateDimension", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDimensionMiddlewares(stack)
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
	addOpUpdateDimensionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDimension(options.Region), middleware.Before)
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
			OperationName: "UpdateDimension",
			Err:           err,
		}
	}
	out := result.(*UpdateDimensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDimensionInput struct {

	// A unique identifier for the dimension. Choose something that describes the type
	// and value to make it easy to remember what it does.
	//
	// This member is required.
	Name *string

	// Specifies the value or list of values for the dimension. For TOPIC_FILTER
	// dimensions, this is a pattern used to match the MQTT topic (for example,
	// "admin/#").
	//
	// This member is required.
	StringValues []*string
}

type UpdateDimensionOutput struct {

	// The value or list of values used to scope the dimension. For example, for topic
	// filters, this is the pattern used to match the MQTT topic name.
	StringValues []*string

	// The type of the dimension.
	Type types.DimensionType

	// The date and time, in milliseconds since epoch, when the dimension was initially
	// created.
	CreationDate *time.Time

	// A unique identifier for the dimension.
	Name *string

	// The ARN (Amazon resource name) of the created dimension.
	Arn *string

	// The date and time, in milliseconds since epoch, when the dimension was most
	// recently updated.
	LastModifiedDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDimensionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDimension{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDimension{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDimension(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateDimension",
	}
}
