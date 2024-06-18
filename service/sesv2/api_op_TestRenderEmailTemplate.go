// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a preview of the MIME content of an email when provided with a template
// and a set of replacement data.
//
// You can execute this operation no more than once per second.
func (c *Client) TestRenderEmailTemplate(ctx context.Context, params *TestRenderEmailTemplateInput, optFns ...func(*Options)) (*TestRenderEmailTemplateOutput, error) {
	if params == nil {
		params = &TestRenderEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestRenderEmailTemplate", params, optFns, c.addOperationTestRenderEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestRenderEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// >Represents a request to create a preview of the MIME content of an email when
// provided with a template and a set of replacement data.
type TestRenderEmailTemplateInput struct {

	// A list of replacement values to apply to the template. This parameter is a JSON
	// object, typically consisting of key-value pairs in which the keys correspond to
	// replacement tags in the email template.
	//
	// This member is required.
	TemplateData *string

	// The name of the template.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

// The following element is returned by the service.
type TestRenderEmailTemplateOutput struct {

	// The complete MIME message rendered by applying the data in the TemplateData
	// parameter to the template specified in the TemplateName parameter.
	//
	// This member is required.
	RenderedTemplate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestRenderEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestRenderEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestRenderEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TestRenderEmailTemplate"); err != nil {
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
	if err = addOpTestRenderEmailTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestRenderEmailTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestRenderEmailTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestRenderEmailTemplate",
	}
}
