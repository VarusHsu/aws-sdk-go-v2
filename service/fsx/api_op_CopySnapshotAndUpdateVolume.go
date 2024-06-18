// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing volume by using a snapshot from another Amazon FSx for
// OpenZFS file system. For more information, see [on-demand data replication]in the Amazon FSx for OpenZFS
// User Guide.
//
// [on-demand data replication]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/on-demand-replication.html
func (c *Client) CopySnapshotAndUpdateVolume(ctx context.Context, params *CopySnapshotAndUpdateVolumeInput, optFns ...func(*Options)) (*CopySnapshotAndUpdateVolumeOutput, error) {
	if params == nil {
		params = &CopySnapshotAndUpdateVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopySnapshotAndUpdateVolume", params, optFns, c.addOperationCopySnapshotAndUpdateVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopySnapshotAndUpdateVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotAndUpdateVolumeInput struct {

	// The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify
	// Amazon Web Services resources. We require an ARN when you need to specify a
	// resource unambiguously across all of Amazon Web Services. For more information,
	// see [Amazon Resource Names (ARNs)]in the Amazon Web Services General Reference.
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	SourceSnapshotARN *string

	// Specifies the ID of the volume that you are copying the snapshot to.
	//
	// This member is required.
	VolumeId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// Specifies the strategy to use when copying data from a snapshot to the volume.
	//
	//   - FULL_COPY - Copies all data from the snapshot to the volume.
	//
	//   - INCREMENTAL_COPY - Copies only the snapshot data that's changed since the
	//   previous replication.
	//
	// CLONE isn't a valid copy strategy option for the CopySnapshotAndUpdateVolume
	// operation.
	CopyStrategy types.OpenZFSCopyStrategy

	// Confirms that you want to delete data on the destination volume that wasn’t
	// there during the previous snapshot replication.
	//
	// Your replication will fail if you don’t include an option for a specific type
	// of data and that data is on your destination. For example, if you don’t include
	// DELETE_INTERMEDIATE_SNAPSHOTS and there are intermediate snapshots on the
	// destination, you can’t copy the snapshot.
	//
	//   - DELETE_INTERMEDIATE_SNAPSHOTS - Deletes snapshots on the destination volume
	//   that aren’t on the source volume.
	//
	//   - DELETE_CLONED_VOLUMES - Deletes snapshot clones on the destination volume
	//   that aren't on the source volume.
	//
	//   - DELETE_INTERMEDIATE_DATA - Overwrites snapshots on the destination volume
	//   that don’t match the source snapshot that you’re copying.
	Options []types.UpdateOpenZFSVolumeOption

	noSmithyDocumentSerde
}

type CopySnapshotAndUpdateVolumeOutput struct {

	// A list of administrative actions for the file system that are in process or
	// waiting to be processed. Administrative actions describe changes to the Amazon
	// FSx system.
	AdministrativeActions []types.AdministrativeAction

	// The lifecycle state of the destination volume.
	Lifecycle types.VolumeLifecycle

	// The ID of the volume that you copied the snapshot to.
	VolumeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopySnapshotAndUpdateVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopySnapshotAndUpdateVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopySnapshotAndUpdateVolume{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopySnapshotAndUpdateVolume"); err != nil {
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
	if err = addIdempotencyToken_opCopySnapshotAndUpdateVolumeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopySnapshotAndUpdateVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshotAndUpdateVolume(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCopySnapshotAndUpdateVolume struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCopySnapshotAndUpdateVolume) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCopySnapshotAndUpdateVolume) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CopySnapshotAndUpdateVolumeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CopySnapshotAndUpdateVolumeInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCopySnapshotAndUpdateVolumeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCopySnapshotAndUpdateVolume{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCopySnapshotAndUpdateVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopySnapshotAndUpdateVolume",
	}
}
