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

// Creates a tracker resource in your Amazon Web Services account, which lets you
// retrieve current and historical location of devices.
func (c *Client) CreateTracker(ctx context.Context, params *CreateTrackerInput, optFns ...func(*Options)) (*CreateTrackerOutput, error) {
	if params == nil {
		params = &CreateTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTracker", params, optFns, c.addOperationCreateTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrackerInput struct {

	// The name for the tracker resource.
	//
	// Requirements:
	//
	//   - Contain only alphanumeric characters (A-Z, a-z, 0-9) , hyphens (-), periods
	//   (.), and underscores (_).
	//
	//   - Must be a unique tracker resource name.
	//
	//   - No spaces allowed. For example, ExampleTracker .
	//
	// This member is required.
	TrackerName *string

	// An optional description for the tracker resource.
	Description *string

	// Whether to enable position UPDATE events from this tracker to be sent to
	// EventBridge.
	//
	// You do not need enable this feature to get ENTER and EXIT events for geofences
	// with this tracker. Those events are always sent to EventBridge.
	EventBridgeEnabled *bool

	// Enables GeospatialQueries for a tracker that uses a [Amazon Web Services KMS customer managed key].
	//
	// This parameter is only used if you are using a KMS customer managed key.
	//
	// If you wish to encrypt your data using your own KMS customer managed key, then
	// the Bounding Polygon Queries feature will be disabled by default. This is
	// because by using this feature, a representation of your device positions will
	// not be encrypted using the your KMS managed key. The exact device position,
	// however; is still encrypted using your managed key.
	//
	// You can choose to opt-in to the Bounding Polygon Quseries feature. This is done
	// by setting the KmsKeyEnableGeospatialQueries parameter to true when creating or
	// updating a Tracker.
	//
	// [Amazon Web Services KMS customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html
	KmsKeyEnableGeospatialQueries *bool

	// A key identifier for an [Amazon Web Services KMS customer managed key]. Enter a key ID, key ARN, alias name, or alias ARN.
	//
	// [Amazon Web Services KMS customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html
	KmsKeyId *string

	// Specifies the position filtering for the tracker resource.
	//
	// Valid values:
	//
	//   - TimeBased - Location updates are evaluated against linked geofence
	//   collections, but not every location update is stored. If your update frequency
	//   is more often than 30 seconds, only one update per 30 seconds is stored for each
	//   unique device ID.
	//
	//   - DistanceBased - If the device has moved less than 30 m (98.4 ft), location
	//   updates are ignored. Location updates within this area are neither evaluated
	//   against linked geofence collections, nor stored. This helps control costs by
	//   reducing the number of geofence evaluations and historical device positions to
	//   paginate through. Distance-based filtering can also reduce the effects of GPS
	//   noise when displaying device trajectories on a map.
	//
	//   - AccuracyBased - If the device has moved less than the measured accuracy,
	//   location updates are ignored. For example, if two consecutive updates from a
	//   device have a horizontal accuracy of 5 m and 10 m, the second update is ignored
	//   if the device has moved less than 15 m. Ignored location updates are neither
	//   evaluated against linked geofence collections, nor stored. This can reduce the
	//   effects of GPS noise when displaying device trajectories on a map, and can help
	//   control your costs by reducing the number of geofence evaluations.
	//
	// This field is optional. If not specified, the default value is TimeBased .
	PositionFiltering types.PositionFiltering

	// No longer used. If included, the only allowed value is RequestBasedUsage .
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// This parameter is no longer used.
	//
	// Deprecated: Deprecated. No longer allowed.
	PricingPlanDataSource *string

	// Applies one or more tags to the tracker resource. A tag is a key-value pair
	// helps manage, identify, search, and filter your resources by labelling them.
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

type CreateTrackerOutput struct {

	// The timestamp for when the tracker resource was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// The Amazon Resource Name (ARN) for the tracker resource. Used when you need to
	// specify a resource across all Amazon Web Services.
	//
	//   - Format example: arn:aws:geo:region:account-id:tracker/ExampleTracker
	//
	// This member is required.
	TrackerArn *string

	// The name of the tracker resource.
	//
	// This member is required.
	TrackerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTracker"); err != nil {
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
	if err = addEndpointPrefix_opCreateTrackerMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTracker(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateTrackerMiddleware struct {
}

func (*endpointPrefix_opCreateTrackerMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateTrackerMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.tracking." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateTrackerMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateTrackerMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateTracker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTracker",
	}
}
