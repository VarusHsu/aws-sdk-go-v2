// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a directory with an Amazon SNS topic. This establishes the directory
// as a publisher to the specified Amazon SNS topic. You can then receive email or
// text (SMS) messages when the status of your directory changes. You get notified
// if your directory goes from an Active status to an Impaired or Inoperable
// status. You also receive a notification when the directory returns to an Active
// status.
func (c *Client) RegisterEventTopic(ctx context.Context, params *RegisterEventTopicInput, optFns ...func(*Options)) (*RegisterEventTopicOutput, error) {
	if params == nil {
		params = &RegisterEventTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterEventTopic", params, optFns, c.addOperationRegisterEventTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterEventTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Registers a new event topic.
type RegisterEventTopicInput struct {

	// The Directory ID that will publish status messages to the Amazon SNS topic.
	//
	// This member is required.
	DirectoryId *string

	// The Amazon SNS topic name to which the directory will publish status messages.
	// This Amazon SNS topic must be in the same region as the specified Directory ID.
	//
	// This member is required.
	TopicName *string

	noSmithyDocumentSerde
}

// The result of a RegisterEventTopic request.
type RegisterEventTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterEventTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterEventTopic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterEventTopic{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterEventTopic"); err != nil {
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
	if err = addOpRegisterEventTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterEventTopic(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterEventTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterEventTopic",
	}
}
