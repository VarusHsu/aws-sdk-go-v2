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

// Creates a route calculator resource in your Amazon Web Services account.
//
// You can send requests to a route calculator resource to estimate travel time,
// distance, and get directions. A route calculator sources traffic and road
// network data from your chosen data provider.
//
// If your application is tracking or routing assets you use in your business,
// such as delivery vehicles or employees, you must not use Esri as your
// geolocation provider. See section 82 of the [Amazon Web Services service terms]for more details.
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms
func (c *Client) CreateRouteCalculator(ctx context.Context, params *CreateRouteCalculatorInput, optFns ...func(*Options)) (*CreateRouteCalculatorOutput, error) {
	if params == nil {
		params = &CreateRouteCalculatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRouteCalculator", params, optFns, c.addOperationCreateRouteCalculatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteCalculatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRouteCalculatorInput struct {

	// The name of the route calculator resource.
	//
	// Requirements:
	//
	//   - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.),
	//   and underscores (_).
	//
	//   - Must be a unique Route calculator resource name.
	//
	//   - No spaces allowed. For example, ExampleRouteCalculator .
	//
	// This member is required.
	CalculatorName *string

	// Specifies the data provider of traffic and road network data.
	//
	// This field is case-sensitive. Enter the valid values as shown. For example,
	// entering HERE returns an error.
	//
	// Valid values include:
	//
	//   - Esri – For additional information about [Esri]'s coverage in your region of
	//   interest, see [Esri details on street networks and traffic coverage].
	//
	// Route calculators that use Esri as a data source only calculate routes that are
	//   shorter than 400 km.
	//
	//   - Grab – Grab provides routing functionality for Southeast Asia. For
	//   additional information about [GrabMaps]' coverage, see [GrabMaps countries and areas covered].
	//
	//   - Here – For additional information about [HERE Technologies]' coverage in your region of
	//   interest, see [HERE car routing coverage]and [HERE truck routing coverage].
	//
	// For additional information , see [Data providers] on the Amazon Location Service Developer
	// Guide.
	//
	// [HERE car routing coverage]: https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html
	// [Esri]: https://docs.aws.amazon.com/location/latest/developerguide/esri.html
	// [HERE truck routing coverage]: https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html
	// [HERE Technologies]: https://docs.aws.amazon.com/location/latest/developerguide/HERE.html
	// [GrabMaps]: https://docs.aws.amazon.com/location/latest/developerguide/grab.html
	// [Esri details on street networks and traffic coverage]: https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm
	// [Data providers]: https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html
	// [GrabMaps countries and areas covered]: https://docs.aws.amazon.com/location/latest/developerguide/grab.html#grab-coverage-area
	//
	// This member is required.
	DataSource *string

	// The optional description for the route calculator resource.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage .
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// Applies one or more tags to the route calculator resource. A tag is a key-value
	// pair helps manage, identify, search, and filter your resources by labelling
	// them.
	//
	//   - For example: { "tag1" : "value1" , "tag2" : "value2" }
	//
	// Format: "key" : "value"
	//
	// Restrictions:
	//
	//   - Maximum 50 tags per resource
	//
	//   - Each resource tag must be unique with a maximum of one value.
	//
	//   - Maximum key length: 128 Unicode characters in UTF-8
	//
	//   - Maximum value length: 256 Unicode characters in UTF-8
	//
	//   - Can use alphanumeric characters (A–Z, a–z, 0–9), and the following
	//   characters: + - = . _ : / @.
	//
	//   - Cannot use "aws:" as a prefix for a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRouteCalculatorOutput struct {

	// The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN
	// when you specify a resource across all Amazon Web Services.
	//
	//   - Format example:
	//   arn:aws:geo:region:account-id:route-calculator/ExampleCalculator
	//
	// This member is required.
	CalculatorArn *string

	// The name of the route calculator resource.
	//
	//   - For example, ExampleRouteCalculator .
	//
	// This member is required.
	CalculatorName *string

	// The timestamp when the route calculator resource was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	//   - For example, 2020–07-2T12:15:20.000Z+01:00
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRouteCalculatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRouteCalculator"); err != nil {
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
	if err = addEndpointPrefix_opCreateRouteCalculatorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateRouteCalculatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRouteCalculator(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateRouteCalculatorMiddleware struct {
}

func (*endpointPrefix_opCreateRouteCalculatorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateRouteCalculatorMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.routes." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateRouteCalculatorMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateRouteCalculatorMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateRouteCalculator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRouteCalculator",
	}
}
