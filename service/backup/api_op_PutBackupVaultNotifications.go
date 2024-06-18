// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Turns on notifications on a backup vault for the specified topic and events.
func (c *Client) PutBackupVaultNotifications(ctx context.Context, params *PutBackupVaultNotificationsInput, optFns ...func(*Options)) (*PutBackupVaultNotificationsOutput, error) {
	if params == nil {
		params = &PutBackupVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBackupVaultNotifications", params, optFns, c.addOperationPutBackupVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBackupVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBackupVaultNotificationsInput struct {

	// An array of events that indicate the status of jobs to back up resources to the
	// backup vault.
	//
	// For common use cases and code samples, see [Using Amazon SNS to track Backup events].
	//
	// The following events are supported:
	//
	//   - BACKUP_JOB_STARTED | BACKUP_JOB_COMPLETED
	//
	//   - COPY_JOB_STARTED | COPY_JOB_SUCCESSFUL | COPY_JOB_FAILED
	//
	//   - RESTORE_JOB_STARTED | RESTORE_JOB_COMPLETED | RECOVERY_POINT_MODIFIED
	//
	//   - S3_BACKUP_OBJECT_FAILED | S3_RESTORE_OBJECT_FAILED
	//
	// The list below shows items that are deprecated events (for reference) and are
	// no longer in use. They are no longer supported and will not return statuses or
	// notifications. Refer to the list above for current supported events.
	//
	// [Using Amazon SNS to track Backup events]: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	// This member is required.
	BackupVaultEvents []types.BackupVaultEvent

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s
	// events; for example, arn:aws:sns:us-west-2:111122223333:MyVaultTopic .
	//
	// This member is required.
	SNSTopicArn *string

	noSmithyDocumentSerde
}

type PutBackupVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBackupVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutBackupVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutBackupVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBackupVaultNotifications"); err != nil {
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
	if err = addOpPutBackupVaultNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBackupVaultNotifications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutBackupVaultNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBackupVaultNotifications",
	}
}
