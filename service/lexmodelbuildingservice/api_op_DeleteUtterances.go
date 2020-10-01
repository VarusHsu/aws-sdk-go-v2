// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes stored utterances. Amazon Lex stores the utterances that users send to
// your bot. Utterances are stored for 15 days for use with the GetUtterancesView
// () operation, and then stored indefinitely for use in improving the ability of
// your bot to respond to user input. Use the DeleteUtterances operation to
// manually delete stored utterances for a specific user. When you use the
// DeleteUtterances operation, utterances stored for improving your bot's ability
// to respond to user input are deleted immediately. Utterances stored for use with
// the GetUtterancesView operation are deleted after 15 days. This operation
// requires permissions for the lex:DeleteUtterances action.
func (c *Client) DeleteUtterances(ctx context.Context, params *DeleteUtterancesInput, optFns ...func(*Options)) (*DeleteUtterancesOutput, error) {
	stack := middleware.NewStack("DeleteUtterances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteUtterancesMiddlewares(stack)
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
	addOpDeleteUtterancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUtterances(options.Region), middleware.Before)
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
			OperationName: "DeleteUtterances",
			Err:           err,
		}
	}
	out := result.(*DeleteUtterancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUtterancesInput struct {

	// The name of the bot that stored the utterances.
	//
	// This member is required.
	BotName *string

	// The unique identifier for the user that made the utterances. This is the user ID
	// that was sent in the PostContent
	// (http://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html) or
	// PostText (http://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html)
	// operation request that contained the utterance.
	//
	// This member is required.
	UserId *string
}

type DeleteUtterancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteUtterancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUtterances{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUtterances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUtterances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteUtterances",
	}
}
