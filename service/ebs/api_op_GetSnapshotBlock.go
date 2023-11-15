// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Returns the data in a block in an Amazon Elastic Block Store snapshot. You
// should always retry requests that receive server ( 5xx ) error responses, and
// ThrottlingException and RequestThrottledException client error responses. For
// more information see Error retries (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/error-retries.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) GetSnapshotBlock(ctx context.Context, params *GetSnapshotBlockInput, optFns ...func(*Options)) (*GetSnapshotBlockOutput, error) {
	if params == nil {
		params = &GetSnapshotBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshotBlock", params, optFns, c.addOperationGetSnapshotBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnapshotBlockInput struct {

	// The block index of the block in which to read the data. A block index is a
	// logical index in units of 512 KiB blocks. To identify the block index, divide
	// the logical offset of the data in the logical volume by the block size (logical
	// offset of data/ 524288 ). The logical offset of the data must be 512 KiB
	// aligned.
	//
	// This member is required.
	BlockIndex *int32

	// The block token of the block from which to get data. You can obtain the
	// BlockToken by running the ListChangedBlocks or ListSnapshotBlocks operations.
	//
	// This member is required.
	BlockToken *string

	// The ID of the snapshot containing the block from which to get data. If the
	// specified snapshot is encrypted, you must have permission to use the KMS key
	// that was used to encrypt the snapshot. For more information, see Using
	// encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// This member is required.
	SnapshotId *string

	noSmithyDocumentSerde
}

type GetSnapshotBlockOutput struct {

	// The data content of the block.
	BlockData io.ReadCloser

	// The checksum generated for the block, which is Base64 encoded.
	Checksum *string

	// The algorithm used to generate the checksum for the block, such as SHA256.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The size of the data in the block.
	DataLength *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSnapshotBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSnapshotBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSnapshotBlock{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSnapshotBlock"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSnapshotBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshotBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSnapshotBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSnapshotBlock",
	}
}
