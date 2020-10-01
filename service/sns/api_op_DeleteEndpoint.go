// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the endpoint for a device and mobile app from Amazon SNS. This action is
// idempotent. For more information, see Using Amazon SNS Mobile Push Notifications
// (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html). When you delete
// an endpoint that is also subscribed to a topic, then you must also unsubscribe
// the endpoint from the topic.
func (c *Client) DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) {
	stack := middleware.NewStack("DeleteEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteEndpointMiddlewares(stack)
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
	addOpDeleteEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEndpoint(options.Region), middleware.Before)
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
			OperationName: "DeleteEndpoint",
			Err:           err,
		}
	}
	out := result.(*DeleteEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for DeleteEndpoint action.
type DeleteEndpointInput struct {

	// EndpointArn of endpoint to delete.
	//
	// This member is required.
	EndpointArn *string
}

type DeleteEndpointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "DeleteEndpoint",
	}
}
