// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores the specified table to the specified point in time within the
// earliest_restorable_timestamp and the current time. For more information about
// restore points, see Time window for PITR continuous backups (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_window)
// in the Amazon Keyspaces Developer Guide. Any number of users can execute up to 4
// concurrent restores (any type of restore) in a given account. When you restore
// using point in time recovery, Amazon Keyspaces restores your source table's
// schema and data to the state based on the selected timestamp
// (day:hour:minute:second) to a new table. The Time to Live (TTL) settings are
// also restored to the state based on the selected timestamp. In addition to the
// table's schema, data, and TTL settings, RestoreTable restores the capacity
// mode, encryption, and point-in-time recovery settings from the source table.
// Unlike the table's schema data and TTL settings, which are restored based on the
// selected timestamp, these settings are always restored based on the table's
// settings as of the current time or when the table was deleted. You can also
// overwrite these settings during restore:
//   - Read/write capacity mode
//   - Provisioned throughput capacity settings
//   - Point-in-time (PITR) settings
//   - Tags
//
// For more information, see PITR restore settings (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_settings)
// in the Amazon Keyspaces Developer Guide. Note that the following settings are
// not restored, and you must configure them manually for the new table:
//   - Automatic scaling policies (for tables that use provisioned capacity mode)
//   - Identity and Access Management (IAM) policies
//   - Amazon CloudWatch metrics and alarms
func (c *Client) RestoreTable(ctx context.Context, params *RestoreTableInput, optFns ...func(*Options)) (*RestoreTableOutput, error) {
	if params == nil {
		params = &RestoreTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreTable", params, optFns, c.addOperationRestoreTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableInput struct {

	// The keyspace name of the source table.
	//
	// This member is required.
	SourceKeyspaceName *string

	// The name of the source table.
	//
	// This member is required.
	SourceTableName *string

	// The name of the target keyspace.
	//
	// This member is required.
	TargetKeyspaceName *string

	// The name of the target table.
	//
	// This member is required.
	TargetTableName *string

	// Specifies the read/write throughput capacity mode for the target table. The
	// options are:
	//   - throughputMode:PAY_PER_REQUEST
	//   - throughputMode:PROVISIONED - Provisioned capacity mode requires
	//   readCapacityUnits and writeCapacityUnits as input.
	// The default is throughput_mode:PAY_PER_REQUEST . For more information, see
	// Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
	// in the Amazon Keyspaces Developer Guide.
	CapacitySpecificationOverride *types.CapacitySpecification

	// Specifies the encryption settings for the target table. You can choose one of
	// the following KMS key (KMS key):
	//   - type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.
	//   - type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and is
	//   created, owned, and managed by you. This option requires the
	//   kms_key_identifier of the KMS key in Amazon Resource Name (ARN) format as
	//   input.
	// The default is type:AWS_OWNED_KMS_KEY . For more information, see Encryption at
	// rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
	// in the Amazon Keyspaces Developer Guide.
	EncryptionSpecificationOverride *types.EncryptionSpecification

	// Specifies the pointInTimeRecovery settings for the target table. The options
	// are:
	//   - status=ENABLED
	//   - status=DISABLED
	// If it's not specified, the default is status=DISABLED . For more information,
	// see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)
	// in the Amazon Keyspaces Developer Guide.
	PointInTimeRecoveryOverride *types.PointInTimeRecovery

	// The restore timestamp in ISO 8601 format.
	RestoreTimestamp *time.Time

	// A list of key-value pair tags to be attached to the restored table. For more
	// information, see Adding tags and labels to Amazon Keyspaces resources (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
	// in the Amazon Keyspaces Developer Guide.
	TagsOverride []types.Tag

	noSmithyDocumentSerde
}

type RestoreTableOutput struct {

	// The Amazon Resource Name (ARN) of the restored table.
	//
	// This member is required.
	RestoredTableARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRestoreTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRestoreTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RestoreTable"); err != nil {
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
	if err = addOpRestoreTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RestoreTable",
	}
}
