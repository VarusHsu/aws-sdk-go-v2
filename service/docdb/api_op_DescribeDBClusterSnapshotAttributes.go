// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of cluster snapshot attribute names and values for a manual DB
// cluster snapshot.
//
// When you share snapshots with other Amazon Web Services accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon Web Services accounts that are authorized to copy or restore
// the manual cluster snapshot. If all is included in the list of values for the
// restore attribute, then the manual cluster snapshot is public and can be copied
// or restored by all Amazon Web Services accounts.
func (c *Client) DescribeDBClusterSnapshotAttributes(ctx context.Context, params *DescribeDBClusterSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotAttributesOutput, error) {
	if params == nil {
		params = &DescribeDBClusterSnapshotAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterSnapshotAttributes", params, optFns, c.addOperationDescribeDBClusterSnapshotAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterSnapshotAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterSnapshotAttributes.
type DescribeDBClusterSnapshotAttributesInput struct {

	// The identifier for the cluster snapshot to describe the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	noSmithyDocumentSerde
}

type DescribeDBClusterSnapshotAttributesOutput struct {

	// Detailed information about the attributes that are associated with a cluster
	// snapshot.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterSnapshotAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterSnapshotAttributes"); err != nil {
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
	if err = addOpDescribeDBClusterSnapshotAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterSnapshotAttributes",
	}
}
