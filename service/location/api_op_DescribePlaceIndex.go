// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the place index resource details.
func (c *Client) DescribePlaceIndex(ctx context.Context, params *DescribePlaceIndexInput, optFns ...func(*Options)) (*DescribePlaceIndexOutput, error) {
	if params == nil {
		params = &DescribePlaceIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePlaceIndex", params, optFns, c.addOperationDescribePlaceIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePlaceIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePlaceIndexInput struct {

	// The name of the place index resource.
	//
	// This member is required.
	IndexName *string

	noSmithyDocumentSerde
}

type DescribePlaceIndexOutput struct {

	// The timestamp for when the place index resource was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// The data provider of geospatial data. Values can be one of the following:
	//
	//   - Esri
	//
	//   - Grab
	//
	//   - Here
	//
	// For more information about data providers, see [Amazon Location Service data providers].
	//
	// [Amazon Location Service data providers]: https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html
	//
	// This member is required.
	DataSource *string

	// The specified data storage option for requesting Places.
	//
	// This member is required.
	DataSourceConfiguration *types.DataSourceConfiguration

	// The optional description for the place index resource.
	//
	// This member is required.
	Description *string

	// The Amazon Resource Name (ARN) for the place index resource. Used to specify a
	// resource across Amazon Web Services.
	//
	//   - Format example: arn:aws:geo:region:account-id:place-index/ExamplePlaceIndex
	//
	// This member is required.
	IndexArn *string

	// The name of the place index resource being described.
	//
	// This member is required.
	IndexName *string

	// The timestamp for when the place index resource was last updated in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	UpdateTime *time.Time

	// No longer used. Always returns RequestBasedUsage .
	//
	// Deprecated: Deprecated. Always returns RequestBasedUsage.
	PricingPlan types.PricingPlan

	// Tags associated with place index resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePlaceIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePlaceIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePlaceIndex{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePlaceIndex"); err != nil {
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
	if err = addEndpointPrefix_opDescribePlaceIndexMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribePlaceIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlaceIndex(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribePlaceIndexMiddleware struct {
}

func (*endpointPrefix_opDescribePlaceIndexMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribePlaceIndexMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.places." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDescribePlaceIndexMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribePlaceIndexMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribePlaceIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePlaceIndex",
	}
}
