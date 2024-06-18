// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an event data store. The required EventDataStore value is an ARN or the
// ID portion of the ARN. Other parameters are optional, but at least one optional
// parameter must be specified, or CloudTrail throws an error. RetentionPeriod is
// in days, and valid values are integers between 7 and 3653 if the BillingMode is
// set to EXTENDABLE_RETENTION_PRICING , or between 7 and 2557 if BillingMode is
// set to FIXED_RETENTION_PRICING . By default, TerminationProtection is enabled.
//
// For event data stores for CloudTrail events, AdvancedEventSelectors includes or
// excludes management or data events in your event data store. For more
// information about AdvancedEventSelectors , see [AdvancedEventSelectors].
//
// For event data stores for CloudTrail Insights events, Config configuration
// items, Audit Manager evidence, or non-Amazon Web Services events,
// AdvancedEventSelectors includes events of that type in your event data store.
//
// [AdvancedEventSelectors]: https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html
func (c *Client) UpdateEventDataStore(ctx context.Context, params *UpdateEventDataStoreInput, optFns ...func(*Options)) (*UpdateEventDataStoreOutput, error) {
	if params == nil {
		params = &UpdateEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventDataStore", params, optFns, c.addOperationUpdateEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventDataStoreInput struct {

	// The ARN (or the ID suffix of the ARN) of the event data store that you want to
	// update.
	//
	// This member is required.
	EventDataStore *string

	// The advanced event selectors used to select events for the event data store.
	// You can configure up to five advanced event selectors for each event data store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// You can't change the billing mode from EXTENDABLE_RETENTION_PRICING to
	// FIXED_RETENTION_PRICING . If BillingMode is set to EXTENDABLE_RETENTION_PRICING
	// and you want to use FIXED_RETENTION_PRICING instead, you'll need to stop
	// ingestion on the event data store and create a new event data store that uses
	// FIXED_RETENTION_PRICING .
	//
	// The billing mode for the event data store determines the cost for ingesting
	// events and the default and maximum retention period for the event data store.
	//
	// The following are the possible values:
	//
	//   - EXTENDABLE_RETENTION_PRICING - This billing mode is generally recommended if
	//   you want a flexible retention period of up to 3653 days (about 10 years). The
	//   default retention period for this billing mode is 366 days.
	//
	//   - FIXED_RETENTION_PRICING - This billing mode is recommended if you expect to
	//   ingest more than 25 TB of event data per month and need a retention period of up
	//   to 2557 days (about 7 years). The default retention period for this billing mode
	//   is 2557 days.
	//
	// For more information about CloudTrail pricing, see [CloudTrail Pricing] and [Managing CloudTrail Lake costs].
	//
	// [CloudTrail Pricing]: http://aws.amazon.com/cloudtrail/pricing/
	// [Managing CloudTrail Lake costs]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.html
	BillingMode types.BillingMode

	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.
	// The value can be an alias name prefixed by alias/ , a fully specified ARN to an
	// alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Disabling or deleting the KMS key, or removing CloudTrail permissions on the
	// key, prevents CloudTrail from logging events to the event data store, and
	// prevents users from querying the data in the event data store that was encrypted
	// with the key. After you associate an event data store with a KMS key, the KMS
	// key cannot be removed or changed. Before you disable or delete a KMS key that
	// you are using with an event data store, delete or back up your event data store.
	//
	// CloudTrail also supports KMS multi-Region keys. For more information about
	// multi-Region keys, see [Using multi-Region keys]in the Key Management Service Developer Guide.
	//
	// Examples:
	//
	//   - alias/MyAliasName
	//
	//   - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	//   - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//   - 12345678-1234-1234-1234-123456789012
	//
	// [Using multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
	KmsKeyId *string

	// Specifies whether an event data store collects events from all Regions, or only
	// from the Region in which it was created.
	MultiRegionEnabled *bool

	// The event data store name.
	Name *string

	// Specifies whether an event data store collects events logged for an
	// organization in Organizations.
	//
	// Only the management account for the organization can convert an organization
	// event data store to a non-organization event data store, or convert a
	// non-organization event data store to an organization event data store.
	OrganizationEnabled *bool

	// The retention period of the event data store, in days. If BillingMode is set to
	// EXTENDABLE_RETENTION_PRICING , you can set a retention period of up to 3653
	// days, the equivalent of 10 years. If BillingMode is set to
	// FIXED_RETENTION_PRICING , you can set a retention period of up to 2557 days, the
	// equivalent of seven years.
	//
	// CloudTrail Lake determines whether to retain an event by checking if the
	// eventTime of the event is within the specified retention period. For example, if
	// you set a retention period of 90 days, CloudTrail will remove events when the
	// eventTime is older than 90 days.
	//
	// If you decrease the retention period of an event data store, CloudTrail will
	// remove any events with an eventTime older than the new retention period. For
	// example, if the previous retention period was 365 days and you decrease it to
	// 100 days, CloudTrail will remove events with an eventTime older than 100 days.
	RetentionPeriod *int32

	// Indicates that termination protection is enabled and the event data store
	// cannot be automatically deleted.
	TerminationProtectionEnabled *bool

	noSmithyDocumentSerde
}

type UpdateEventDataStoreOutput struct {

	// The advanced event selectors that are applied to the event data store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The billing mode for the event data store.
	BillingMode types.BillingMode

	// The timestamp that shows when an event data store was first created.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	//  If Lake query federation is enabled, provides the ARN of the federation role
	// used to access the resources for the federated event data store.
	FederationRoleArn *string

	//  Indicates the [Lake query federation] status. The status is ENABLED if Lake query federation is
	// enabled, or DISABLED if Lake query federation is disabled. You cannot delete an
	// event data store if the FederationStatus is ENABLED .
	//
	// [Lake query federation]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html
	FederationStatus types.FederationStatus

	// Specifies the KMS key ID that encrypts the events delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	//
	//     arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Indicates whether the event data store includes events from all Regions, or
	// only from the Region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period, in days.
	RetentionPeriod *int32

	// The status of an event data store.
	Status types.EventDataStoreStatus

	// Indicates whether termination protection is enabled for the event data store.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when the event data store was last updated.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp .
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEventDataStore"); err != nil {
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
	if err = addOpUpdateEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEventDataStore",
	}
}
