// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Clears any customization of a GatewayResponse () of a specified response type on
// the given RestApi () and resets it with the default settings.
func (c *Client) DeleteGatewayResponse(ctx context.Context, params *DeleteGatewayResponseInput, optFns ...func(*Options)) (*DeleteGatewayResponseOutput, error) {
	stack := middleware.NewStack("DeleteGatewayResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteGatewayResponseMiddlewares(stack)
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
	addOpDeleteGatewayResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGatewayResponse(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "DeleteGatewayResponse",
			Err:           err,
		}
	}
	out := result.(*DeleteGatewayResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Clears any customization of a GatewayResponse () of a specified response type on
// the given RestApi () and resets it with the default settings.
type DeleteGatewayResponseInput struct {

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Template *bool

	// [Required] The response type of the associated GatewayResponse (). Valid values
	// are
	//
	//     * ACCESS_DENIED
	//
	//     * API_CONFIGURATION_ERROR
	//
	//     *
	// AUTHORIZER_FAILURE
	//
	//     * AUTHORIZER_CONFIGURATION_ERROR
	//
	//     *
	// BAD_REQUEST_PARAMETERS
	//
	//     * BAD_REQUEST_BODY
	//
	//     * DEFAULT_4XX
	//
	//     *
	// DEFAULT_5XX
	//
	//     * EXPIRED_TOKEN
	//
	//     * INVALID_SIGNATURE
	//
	//     *
	// INTEGRATION_FAILURE
	//
	//     * INTEGRATION_TIMEOUT
	//
	//     * INVALID_API_KEY
	//
	//     *
	// MISSING_AUTHENTICATION_TOKEN
	//
	//     * QUOTA_EXCEEDED
	//
	//     * REQUEST_TOO_LARGE
	//
	//
	// * RESOURCE_NOT_FOUND
	//
	//     * THROTTLED
	//
	//     * UNAUTHORIZED
	//
	//     *
	// UNSUPPORTED_MEDIA_TYPE
	//
	// This member is required.
	ResponseType types.GatewayResponseType

	TemplateSkipList []*string

	Title *string

	Name *string
}

type DeleteGatewayResponseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteGatewayResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGatewayResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGatewayResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteGatewayResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteGatewayResponse",
	}
}
