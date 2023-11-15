// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests whether the specified event pattern matches the provided event. Most
// services in Amazon Web Services treat : or / as the same character in Amazon
// Resource Names (ARNs). However, EventBridge uses an exact match in event
// patterns and rules. Be sure to use the correct ARN characters when creating
// event patterns so that they match the ARN syntax in the event you want to match.
func (c *Client) TestEventPattern(ctx context.Context, params *TestEventPatternInput, optFns ...func(*Options)) (*TestEventPatternOutput, error) {
	if params == nil {
		params = &TestEventPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestEventPattern", params, optFns, c.addOperationTestEventPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestEventPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestEventPatternInput struct {

	// The event, in JSON format, to test against the event pattern. The JSON must
	// follow the format specified in Amazon Web Services Events (https://docs.aws.amazon.com/eventbridge/latest/userguide/aws-events.html)
	// , and the following fields are mandatory:
	//   - id
	//   - account
	//   - source
	//   - time
	//   - region
	//   - resources
	//   - detail-type
	//
	// This member is required.
	Event *string

	// The event pattern. For more information, see Events and Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	//
	// This member is required.
	EventPattern *string

	noSmithyDocumentSerde
}

type TestEventPatternOutput struct {

	// Indicates whether the event matches the event pattern.
	Result bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestEventPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestEventPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestEventPattern{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TestEventPattern"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpTestEventPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestEventPattern(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTestEventPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestEventPattern",
	}
}
