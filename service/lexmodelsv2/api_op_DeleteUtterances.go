// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes stored utterances.
//
// Amazon Lex stores the utterances that users send to your bot. Utterances are
// stored for 15 days for use with the [ListAggregatedUtterances]operation, and then stored indefinitely for
// use in improving the ability of your bot to respond to user input..
//
// Use the DeleteUtterances operation to manually delete utterances for a specific
// session. When you use the DeleteUtterances operation, utterances stored for
// improving your bot's ability to respond to user input are deleted immediately.
// Utterances stored for use with the ListAggregatedUtterances operation are
// deleted after 15 days.
//
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
func (c *Client) DeleteUtterances(ctx context.Context, params *DeleteUtterancesInput, optFns ...func(*Options)) (*DeleteUtterancesOutput, error) {
	if params == nil {
		params = &DeleteUtterancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUtterances", params, optFns, c.addOperationDeleteUtterancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUtterancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUtterancesInput struct {

	// The unique identifier of the bot that contains the utterances.
	//
	// This member is required.
	BotId *string

	// The identifier of the language and locale where the utterances were collected.
	// The string must match one of the supported locales. For more information, see [Supported languages].
	//
	// [Supported languages]: https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html
	LocaleId *string

	// The unique identifier of the session with the user. The ID is returned in the
	// response from the [RecognizeText]and [RecognizeUtterance] operations.
	//
	// [RecognizeUtterance]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_runtime_RecognizeUtterance.html
	// [RecognizeText]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_runtime_RecognizeText.html
	SessionId *string

	noSmithyDocumentSerde
}

type DeleteUtterancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteUtterancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUtterances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUtterances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteUtterances"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteUtterancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUtterances(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUtterances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteUtterances",
	}
}
