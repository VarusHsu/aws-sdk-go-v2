// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a calculated attribute for a customer profile.
func (c *Client) GetCalculatedAttributeForProfile(ctx context.Context, params *GetCalculatedAttributeForProfileInput, optFns ...func(*Options)) (*GetCalculatedAttributeForProfileOutput, error) {
	if params == nil {
		params = &GetCalculatedAttributeForProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCalculatedAttributeForProfile", params, optFns, c.addOperationGetCalculatedAttributeForProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCalculatedAttributeForProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCalculatedAttributeForProfileInput struct {

	// The unique name of the calculated attribute.
	//
	// This member is required.
	CalculatedAttributeName *string

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	noSmithyDocumentSerde
}

type GetCalculatedAttributeForProfileOutput struct {

	// The unique name of the calculated attribute.
	CalculatedAttributeName *string

	// The display name of the calculated attribute.
	DisplayName *string

	// Indicates whether the calculated attribute’s value is based on partial data. If
	// data is partial, it is set to true.
	IsDataPartial *string

	// The value of the calculated attribute.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCalculatedAttributeForProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCalculatedAttributeForProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCalculatedAttributeForProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCalculatedAttributeForProfile"); err != nil {
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
	if err = addOpGetCalculatedAttributeForProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCalculatedAttributeForProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCalculatedAttributeForProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCalculatedAttributeForProfile",
	}
}
