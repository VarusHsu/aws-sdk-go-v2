// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies some configurations of the Network File System (NFS) transfer location
// that you're using with DataSync.
//
// For more information, see [Configuring transfers to or from an NFS file server].
//
// [Configuring transfers to or from an NFS file server]: https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html
func (c *Client) UpdateLocationNfs(ctx context.Context, params *UpdateLocationNfsInput, optFns ...func(*Options)) (*UpdateLocationNfsOutput, error) {
	if params == nil {
		params = &UpdateLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationNfs", params, optFns, c.addOperationUpdateLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationNfsInput struct {

	// Specifies the Amazon Resource Name (ARN) of the NFS transfer location that you
	// want to update.
	//
	// This member is required.
	LocationArn *string

	// Specifies how DataSync can access a location using the NFS protocol.
	MountOptions *types.NfsMountOptions

	// The DataSync agents that are connecting to a Network File System (NFS) location.
	OnPremConfig *types.OnPremConfig

	// Specifies the export path in your NFS file server that you want DataSync to
	// mount.
	//
	// This path (or a subdirectory of the path) is where DataSync transfers data to
	// or from. For information on configuring an export for DataSync, see [Accessing NFS file servers].
	//
	// [Accessing NFS file servers]: https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#accessing-nfs
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationNfsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLocationNfs"); err != nil {
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
	if err = addOpUpdateLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationNfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLocationNfs",
	}
}
