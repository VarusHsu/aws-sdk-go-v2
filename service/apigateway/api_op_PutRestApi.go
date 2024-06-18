// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// A feature of the API Gateway control service for updating an existing API with
// an input of external API definitions. The update can take the form of merging
// the supplied definition into the existing API or overwriting the existing API.
func (c *Client) PutRestApi(ctx context.Context, params *PutRestApiInput, optFns ...func(*Options)) (*PutRestApiOutput, error) {
	if params == nil {
		params = &PutRestApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRestApi", params, optFns, c.addOperationPutRestApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A PUT request to update an existing API, with external API definitions
// specified as the request body.
type PutRestApiInput struct {

	// The PUT request body containing external API definitions. Currently, only
	// OpenAPI definition JSON/YAML files are supported. The maximum size of the API
	// definition file is 6MB.
	//
	// This member is required.
	Body []byte

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A query parameter to indicate whether to rollback the API update ( true ) or not
	// ( false ) when a warning is encountered. The default value is false .
	FailOnWarnings bool

	// The mode query parameter to specify the update mode. Valid values are "merge"
	// and "overwrite". By default, the update mode is "merge".
	Mode types.PutMode

	// Custom header parameters as part of the request. For example, to exclude
	// DocumentationParts from an imported API, set ignore=documentation as a
	// parameters value, as in the AWS CLI command of aws apigateway import-rest-api
	// --parameters ignore=documentation --body
	// 'file:///path/to/imported-api-body.json' .
	Parameters map[string]string

	noSmithyDocumentSerde
}

// Represents a REST API.
type PutRestApiOutput struct {

	// The source of the API key for metering requests according to a usage plan.
	// Valid values are: > HEADER to read the API key from the X-API-Key header of a
	// request. AUTHORIZER to read the API key from the UsageIdentifierKey from a
	// custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The API's description.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint bool

	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string

	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a
	// null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// The API's name.
	Name *string

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string

	// The API's root resource ID.
	RootResourceId *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRestApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRestApi"); err != nil {
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
	if err = addOpPutRestApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRestApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRestApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRestApi",
	}
}
