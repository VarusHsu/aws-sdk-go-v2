// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a usage plan key of a given key identifier.
func (c *Client) GetUsagePlanKey(ctx context.Context, params *GetUsagePlanKeyInput, optFns ...func(*Options)) (*GetUsagePlanKeyOutput, error) {
	if params == nil {
		params = &GetUsagePlanKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlanKey", params, optFns, c.addOperationGetUsagePlanKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get a usage plan key of a given key identifier.
type GetUsagePlanKeyInput struct {

	// The key Id of the to-be-retrieved UsagePlanKey resource representing a plan
	// customer.
	//
	// This member is required.
	KeyId *string

	// The Id of the UsagePlan resource representing the usage plan containing the
	// to-be-retrieved UsagePlanKey resource representing a plan customer.
	//
	// This member is required.
	UsagePlanId *string

	noSmithyDocumentSerde
}

// Represents a usage plan key to identify a plan customer.
type GetUsagePlanKeyOutput struct {

	// The Id of a usage plan key.
	Id *string

	// The name of a usage plan key.
	Name *string

	// The type of a usage plan key. Currently, the valid key type is API_KEY .
	Type *string

	// The value of a usage plan key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsagePlanKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlanKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlanKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUsagePlanKey"); err != nil {
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
	if err = addOpGetUsagePlanKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlanKey(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetUsagePlanKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUsagePlanKey",
	}
}
