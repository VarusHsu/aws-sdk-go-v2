// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// An asynchronous API that updates a stream’s existing edge configuration. The
// Kinesis Video Stream will sync the stream’s edge configuration with the Edge
// Agent IoT Greengrass component that runs on an IoT Hub Device, setup at your
// premise. The time to sync can vary and depends on the connectivity of the Hub
// Device. The SyncStatus will be updated as the edge configuration is
// acknowledged, and synced with the Edge Agent. If this API is invoked for the
// first time, a new edge configuration will be created for the stream, and the
// sync status will be set to SYNCING . You will have to wait for the sync status
// to reach a terminal state such as: IN_SYNC , or SYNC_FAILED , before using this
// API again. If you invoke this API during the syncing process, a
// ResourceInUseException will be thrown. The connectivity of the stream’s edge
// configuration and the Edge Agent will be retried for 15 minutes. After 15
// minutes, the status will transition into the SYNC_FAILED state. To move an edge
// configuration from one device to another, use DeleteEdgeConfiguration to delete
// the current edge configuration. You can then invoke StartEdgeConfigurationUpdate
// with an updated Hub Device ARN.
func (c *Client) StartEdgeConfigurationUpdate(ctx context.Context, params *StartEdgeConfigurationUpdateInput, optFns ...func(*Options)) (*StartEdgeConfigurationUpdateOutput, error) {
	if params == nil {
		params = &StartEdgeConfigurationUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartEdgeConfigurationUpdate", params, optFns, c.addOperationStartEdgeConfigurationUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartEdgeConfigurationUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartEdgeConfigurationUpdateInput struct {

	// The edge configuration details required to invoke the update process.
	//
	// This member is required.
	EdgeConfig *types.EdgeConfig

	// The Amazon Resource Name (ARN) of the stream. Specify either the StreamName or
	// the StreamARN .
	StreamARN *string

	// The name of the stream whose edge configuration you want to update. Specify
	// either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type StartEdgeConfigurationUpdateOutput struct {

	// The timestamp at which a stream’s edge configuration was first created.
	CreationTime *time.Time

	// A description of the stream's edge configuration that will be used to sync with
	// the Edge Agent IoT Greengrass component. The Edge Agent component will run on an
	// IoT Hub Device setup at your premise.
	EdgeConfig *types.EdgeConfig

	// A description of the generated failure status.
	FailedStatusDetails *string

	// The timestamp at which a stream’s edge configuration was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string

	// The name of the stream from which the edge configuration was updated.
	StreamName *string

	// The current sync status of the stream's edge configuration. When you invoke
	// this API, the sync status will be set to the SYNCING state. Use the
	// DescribeEdgeConfiguration API to get the latest status of the edge configuration.
	SyncStatus types.SyncStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartEdgeConfigurationUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartEdgeConfigurationUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartEdgeConfigurationUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartEdgeConfigurationUpdate"); err != nil {
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
	if err = addOpStartEdgeConfigurationUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartEdgeConfigurationUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartEdgeConfigurationUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartEdgeConfigurationUpdate",
	}
}
