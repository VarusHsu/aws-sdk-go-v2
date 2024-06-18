// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds new columns to the table or updates one of the table's settings, for
// example capacity mode, auto scaling, encryption, point-in-time recovery, or ttl
// settings. Note that you can only update one specific table setting per update
// operation.
func (c *Client) UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) {
	if params == nil {
		params = &UpdateTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTable", params, optFns, c.addOperationUpdateTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTableInput struct {

	// The name of the keyspace the specified table is stored in.
	//
	// This member is required.
	KeyspaceName *string

	// The name of the table.
	//
	// This member is required.
	TableName *string

	// For each column to be added to the specified table:
	//
	//   - name - The name of the column.
	//
	//   - type - An Amazon Keyspaces data type. For more information, see [Data types]in the
	//   Amazon Keyspaces Developer Guide.
	//
	// [Data types]: https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types
	AddColumns []types.ColumnDefinition

	// The optional auto scaling settings to update for a table in provisioned
	// capacity mode. Specifies if the service can manage throughput capacity of a
	// provisioned table automatically on your behalf. Amazon Keyspaces auto scaling
	// helps you provision throughput capacity for variable workloads efficiently by
	// increasing and decreasing your table's read and write capacity automatically in
	// response to application traffic.
	//
	// If auto scaling is already enabled for the table, you can use UpdateTable to
	// update the minimum and maximum values or the auto scaling policy settings
	// independently.
	//
	// For more information, see [Managing throughput capacity automatically with Amazon Keyspaces auto scaling] in the Amazon Keyspaces Developer Guide.
	//
	// [Managing throughput capacity automatically with Amazon Keyspaces auto scaling]: https://docs.aws.amazon.com/keyspaces/latest/devguide/autoscaling.html
	AutoScalingSpecification *types.AutoScalingSpecification

	// Modifies the read/write throughput capacity mode for the table. The options are:
	//
	//   - throughputMode:PAY_PER_REQUEST and
	//
	//   - throughputMode:PROVISIONED - Provisioned capacity mode requires
	//   readCapacityUnits and writeCapacityUnits as input.
	//
	// The default is throughput_mode:PAY_PER_REQUEST .
	//
	// For more information, see [Read/write capacity modes] in the Amazon Keyspaces Developer Guide.
	//
	// [Read/write capacity modes]: https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html
	CapacitySpecification *types.CapacitySpecification

	// Enables client-side timestamps for the table. By default, the setting is
	// disabled. You can enable client-side timestamps with the following option:
	//
	//   - status: "enabled"
	//
	// Once client-side timestamps are enabled for a table, this setting cannot be
	// disabled.
	ClientSideTimestamps *types.ClientSideTimestamps

	// The default Time to Live setting in seconds for the table.
	//
	// For more information, see [Setting the default TTL value for a table] in the Amazon Keyspaces Developer Guide.
	//
	// [Setting the default TTL value for a table]: https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl
	DefaultTimeToLive *int32

	// Modifies the encryption settings of the table. You can choose one of the
	// following KMS key (KMS key):
	//
	//   - type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.
	//
	//   - type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and is
	//   created, owned, and managed by you. This option requires the
	//   kms_key_identifier of the KMS key in Amazon Resource Name (ARN) format as
	//   input.
	//
	// The default is AWS_OWNED_KMS_KEY .
	//
	// For more information, see [Encryption at rest] in the Amazon Keyspaces Developer Guide.
	//
	// [Encryption at rest]: https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html
	EncryptionSpecification *types.EncryptionSpecification

	// Modifies the pointInTimeRecovery settings of the table. The options are:
	//
	//   - status=ENABLED
	//
	//   - status=DISABLED
	//
	// If it's not specified, the default is status=DISABLED .
	//
	// For more information, see [Point-in-time recovery] in the Amazon Keyspaces Developer Guide.
	//
	// [Point-in-time recovery]: https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html
	PointInTimeRecovery *types.PointInTimeRecovery

	// The Region specific settings of a multi-Regional table.
	ReplicaSpecifications []types.ReplicaSpecification

	// Modifies Time to Live custom settings for the table. The options are:
	//
	//   - status:enabled
	//
	//   - status:disabled
	//
	// The default is status:disabled . After ttl is enabled, you can't disable it for
	// the table.
	//
	// For more information, see [Expiring data by using Amazon Keyspaces Time to Live (TTL)] in the Amazon Keyspaces Developer Guide.
	//
	// [Expiring data by using Amazon Keyspaces Time to Live (TTL)]: https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html
	Ttl *types.TimeToLive

	noSmithyDocumentSerde
}

type UpdateTableOutput struct {

	// The Amazon Resource Name (ARN) of the modified table.
	//
	// This member is required.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTable"); err != nil {
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
	if err = addOpUpdateTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTable",
	}
}
