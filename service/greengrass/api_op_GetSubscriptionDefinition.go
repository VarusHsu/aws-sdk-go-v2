// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a subscription definition.
func (c *Client) GetSubscriptionDefinition(ctx context.Context, params *GetSubscriptionDefinitionInput, optFns ...func(*Options)) (*GetSubscriptionDefinitionOutput, error) {
	stack := middleware.NewStack("GetSubscriptionDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSubscriptionDefinitionMiddlewares(stack)
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
	addOpGetSubscriptionDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSubscriptionDefinition(options.Region), middleware.Before)
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
			OperationName: "GetSubscriptionDefinition",
			Err:           err,
		}
	}
	out := result.(*GetSubscriptionDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSubscriptionDefinitionInput struct {

	// The ID of the subscription definition.
	//
	// This member is required.
	SubscriptionDefinitionId *string
}

type GetSubscriptionDefinitionOutput struct {

	// The ID of the latest version associated with the definition.
	LatestVersion *string

	// The ID of the definition.
	Id *string

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string

	// The name of the definition.
	Name *string

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string

	// The ARN of the definition.
	Arn *string

	// Tag(s) attached to the resource arn.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSubscriptionDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSubscriptionDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSubscriptionDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSubscriptionDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetSubscriptionDefinition",
	}
}
