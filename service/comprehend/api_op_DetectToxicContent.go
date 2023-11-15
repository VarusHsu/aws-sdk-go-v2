// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Performs toxicity analysis on the list of text strings that you provide as
// input. The analysis uses the order of strings in the list to determine context
// when predicting toxicity. The API response contains a results list that matches
// the size of the input list. For more information about toxicity detection, see
// Toxicity detection (https://docs.aws.amazon.com/comprehend/latest/dg/toxicity-detection.html)
// in the Amazon Comprehend Developer Guide
func (c *Client) DetectToxicContent(ctx context.Context, params *DetectToxicContentInput, optFns ...func(*Options)) (*DetectToxicContentOutput, error) {
	if params == nil {
		params = &DetectToxicContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectToxicContent", params, optFns, c.addOperationDetectToxicContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectToxicContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectToxicContentInput struct {

	// The language of the input text. Currently, English is the only supported
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A list of up to 10 text strings. The maximum size for the list is 10 KB.
	//
	// This member is required.
	TextSegments []types.TextSegment

	noSmithyDocumentSerde
}

type DetectToxicContentOutput struct {

	// Results of the content moderation analysis. Each entry in the results list
	// contains a list of toxic content types identified in the text, along with a
	// confidence score for each content type. The results list also includes a
	// toxicity score for each entry in the results list.
	ResultList []types.ToxicLabels

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectToxicContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectToxicContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectToxicContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DetectToxicContent"); err != nil {
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
	if err = addOpDetectToxicContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectToxicContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectToxicContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DetectToxicContent",
	}
}
