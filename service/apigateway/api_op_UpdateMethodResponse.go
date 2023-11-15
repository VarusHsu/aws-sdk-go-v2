// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing MethodResponse resource.
func (c *Client) UpdateMethodResponse(ctx context.Context, params *UpdateMethodResponseInput, optFns ...func(*Options)) (*UpdateMethodResponseOutput, error) {
	if params == nil {
		params = &UpdateMethodResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMethodResponse", params, optFns, c.addOperationUpdateMethodResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMethodResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update an existing MethodResponse resource.
type UpdateMethodResponseInput struct {

	// The HTTP verb of the Method resource.
	//
	// This member is required.
	HttpMethod *string

	// The Resource identifier for the MethodResponse resource.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The status code for the MethodResponse resource.
	//
	// This member is required.
	StatusCode *string

	// For more information about supported patch operations, see Patch Operations (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html)
	// .
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// Represents a method response of a given HTTP status code returned to the
// client. The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
type UpdateMethodResponseOutput struct {

	// Specifies the Model resources used for the response's content-type. Response
	// models are represented as a key/value map, with a content-type as the key and a
	// Model name as the value.
	ResponseModels map[string]string

	// A key-value map specifying required or optional response parameters that API
	// Gateway can send back to the caller. A key defines a method response header and
	// the value specifies whether the associated method response header is required or
	// not. The expression of the key must match the pattern
	// method.response.header.{name} , where name is a valid and unique header name.
	// API Gateway passes certain integration response data to the method response
	// headers specified here according to the mapping you prescribe in the API's
	// IntegrationResponse. The integration response data that can be mapped include an
	// integration response header expressed in integration.response.header.{name} , a
	// static value enclosed within a pair of single quotes (e.g., 'application/json'
	// ), or a JSON expression from the back-end response payload in the form of
	// integration.response.body.{JSON-expression} , where JSON-expression is a valid
	// JSON expression without the $ prefix.)
	ResponseParameters map[string]bool

	// The method response's status code.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMethodResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMethodResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMethodResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMethodResponse"); err != nil {
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
	if err = addOpUpdateMethodResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMethodResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMethodResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMethodResponse",
	}
}
