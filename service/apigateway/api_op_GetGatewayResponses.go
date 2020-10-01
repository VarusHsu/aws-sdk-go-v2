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

// Gets the GatewayResponses () collection on the given RestApi (). If an API
// developer has not added any definitions for gateway responses, the result will
// be the API Gateway-generated default GatewayResponses () collection for the
// supported response types.
func (c *Client) GetGatewayResponses(ctx context.Context, params *GetGatewayResponsesInput, optFns ...func(*Options)) (*GetGatewayResponsesOutput, error) {
	stack := middleware.NewStack("GetGatewayResponses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetGatewayResponsesMiddlewares(stack)
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
	addOpGetGatewayResponsesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGatewayResponses(options.Region), middleware.Before)
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
			OperationName: "GetGatewayResponses",
			Err:           err,
		}
	}
	out := result.(*GetGatewayResponsesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the GatewayResponses () collection on the given RestApi (). If an API
// developer has not added any definitions for gateway responses, the result will
// be the API Gateway-generated default GatewayResponses () collection for the
// supported response types.
type GetGatewayResponsesInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500. The GatewayResponses () collection does not support
	// pagination and the limit does not apply here.
	Limit *int32

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	Name *string

	TemplateSkipList []*string

	Title *string

	// The current pagination position in the paged result set. The GatewayResponse ()
	// collection does not support pagination and the position does not apply here.
	Position *string

	Template *bool
}

// The collection of the GatewayResponse () instances of a RestApi () as a
// responseType-to-GatewayResponse () object map of key-value pairs. As such,
// pagination is not supported for querying this collection. For more information
// about valid gateway response types, see Gateway Response Types Supported by API
// Gateway
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/supported-gateway-response-types.html)
// Example:
// Get the collection of gateway responses of an API
//
// Request
//
// This example request
// shows how to retrieve the GatewayResponses () collection from an API. GET
// /restapis/o81lxisefl/gatewayresponses HTTP/1.1 Host:
// beta-apigateway.us-east-1.amazonaws.com Content-Type: application/json
// X-Amz-Date: 20170503T220604Z Authorization: AWS4-HMAC-SHA256
// Credential={access-key-id}/20170503/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date,
// Signature=59b42fe54a76a5de8adf2c67baa6d39206f8e9ad49a1d77ccc6a5da3103a398a
// Cache-Control: no-cache Postman-Token:
// 5637af27-dc29-fc5c-9dfe-0645d52cb515
// Response
//
// The successful operation returns
// the 200 OK status code and a payload similar to the following: { "_links": {
// "curies": { "href":
// "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-gatewayresponse-{rel}.html",
// "name": "gatewayresponse", "templated": true }, "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses" }, "first": { "href":
// "/restapis/o81lxisefl/gatewayresponses" }, "gatewayresponse:by-type": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "item": [ { "href": "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE"
// }, { "href": "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND" }, {
// "href": "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/THROTTLED" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR" }, {
// "href": "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" }, {
// "href": "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE" }, { "href":
// "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE" } ] }, "_embedded": {
// "item": [ { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_FAILURE" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "INTEGRATION_FAILURE", "statusCode": "504" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/RESOURCE_NOT_FOUND" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "RESOURCE_NOT_FOUND", "statusCode": "404" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/REQUEST_TOO_LARGE" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "REQUEST_TOO_LARGE", "statusCode": "413" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/THROTTLED" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/THROTTLED" } }, "defaultResponse": true,
// "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "THROTTLED",
// "statusCode": "429" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNSUPPORTED_MEDIA_TYPE" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "UNSUPPORTED_MEDIA_TYPE", "statusCode": "415" }, { "_links": {
// "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_CONFIGURATION_ERROR" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "AUTHORIZER_CONFIGURATION_ERROR", "statusCode": "500" }, {
// "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX" }, "gatewayresponse:put": {
// "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated":
// true }, "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/DEFAULT_5XX" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "DEFAULT_5XX" },
// { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX" }, "gatewayresponse:put": {
// "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated":
// true }, "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/DEFAULT_4XX" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "DEFAULT_4XX" },
// { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_PARAMETERS" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "BAD_REQUEST_PARAMETERS", "statusCode": "400" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/BAD_REQUEST_BODY" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType":
// "BAD_REQUEST_BODY", "statusCode": "400" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN" }, "gatewayresponse:put":
// { "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated":
// true }, "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/EXPIRED_TOKEN" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "EXPIRED_TOKEN",
// "statusCode": "403" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED" }, "gatewayresponse:put":
// { "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated":
// true }, "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/ACCESS_DENIED" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "ACCESS_DENIED",
// "statusCode": "403" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INVALID_API_KEY" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType":
// "INVALID_API_KEY", "statusCode": "403" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED" }, "gatewayresponse:put": {
// "href": "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated":
// true }, "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/UNAUTHORIZED" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType": "UNAUTHORIZED",
// "statusCode": "401" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/API_CONFIGURATION_ERROR" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "API_CONFIGURATION_ERROR", "statusCode": "500" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/QUOTA_EXCEEDED" } }, "defaultResponse":
// true, "responseParameters": {}, "responseTemplates": { "application/json":
// "{\"message\":$context.error.messageString}" }, "responseType":
// "QUOTA_EXCEEDED", "statusCode": "429" }, { "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INTEGRATION_TIMEOUT" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "INTEGRATION_TIMEOUT", "statusCode": "504" }, { "_links": {
// "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/MISSING_AUTHENTICATION_TOKEN" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "MISSING_AUTHENTICATION_TOKEN", "statusCode": "403" }, {
// "_links": { "self": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/INVALID_SIGNATURE" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "INVALID_SIGNATURE", "statusCode": "403" }, { "_links": {
// "self": { "href": "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE" },
// "gatewayresponse:put": { "href":
// "/restapis/o81lxisefl/gatewayresponses/{response_type}", "templated": true },
// "gatewayresponse:update": { "href":
// "/restapis/o81lxisefl/gatewayresponses/AUTHORIZER_FAILURE" } },
// "defaultResponse": true, "responseParameters": {}, "responseTemplates": {
// "application/json": "{\"message\":$context.error.messageString}" },
// "responseType": "AUTHORIZER_FAILURE", "statusCode": "500" } ] } }
//     </div>
// <div class="seeAlso"> <a
// href="https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html">Customize
// Gateway Responses</a> </div>
type GetGatewayResponsesOutput struct {

	// Returns the entire collection, because of no pagination support.
	Items []*types.GatewayResponse

	// The current pagination position in the paged result set. The GatewayResponse ()
	// collection does not support pagination and the position does not apply here.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetGatewayResponsesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetGatewayResponses{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGatewayResponses{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetGatewayResponses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetGatewayResponses",
	}
}
