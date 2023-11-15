// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information the Amazon CloudWatch Internet Monitor has created and stored
// about a health event for a specified monitor. This information includes the
// impacted locations, and all the information related to the event, by location.
// The information returned includes the impact on performance, availability, and
// round-trip time, information about the network providers (ASNs), the event type,
// and so on. Information rolled up at the global traffic level is also returned,
// including the impact type and total traffic impact.
func (c *Client) GetHealthEvent(ctx context.Context, params *GetHealthEventInput, optFns ...func(*Options)) (*GetHealthEventOutput, error) {
	if params == nil {
		params = &GetHealthEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHealthEvent", params, optFns, c.addOperationGetHealthEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHealthEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHealthEventInput struct {

	// The internally generated identifier of a health event. Because EventID contains
	// the forward slash (“/”) character, you must URL-encode the EventID field in the
	// request URL.
	//
	// This member is required.
	EventId *string

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	noSmithyDocumentSerde
}

type GetHealthEventOutput struct {

	// The Amazon Resource Name (ARN) of the event.
	//
	// This member is required.
	EventArn *string

	// The internally generated identifier of a health event.
	//
	// This member is required.
	EventId *string

	// The type of impairment of a specific health event.
	//
	// This member is required.
	ImpactType types.HealthEventImpactType

	// The locations affected by a health event.
	//
	// This member is required.
	ImpactedLocations []types.ImpactedLocation

	// The time when a health event was last updated or recalculated.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The time when a health event started.
	//
	// This member is required.
	StartedAt *time.Time

	// The status of a health event.
	//
	// This member is required.
	Status types.HealthEventStatus

	// The time when a health event was created.
	CreatedAt *time.Time

	// The time when a health event was resolved. If the health event is still active,
	// the end time is not set.
	EndedAt *time.Time

	// The threshold percentage for a health score that determines, along with other
	// configuration information, when Internet Monitor creates a health event when
	// there's an internet issue that affects your application end users.
	HealthScoreThreshold float64

	// The impact on total traffic that a health event has, in increased latency or
	// reduced availability. This is the percentage of how much latency has increased
	// or availability has decreased during the event, compared to what is typical for
	// traffic from this client location to the Amazon Web Services location using this
	// client network.
	PercentOfTotalTrafficImpacted *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHealthEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetHealthEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetHealthEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetHealthEvent"); err != nil {
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
	if err = addOpGetHealthEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthEvent(options.Region), middleware.Before); err != nil {
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
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetHealthEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetHealthEvent",
	}
}
