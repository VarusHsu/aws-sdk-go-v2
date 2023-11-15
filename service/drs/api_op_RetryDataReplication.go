// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// WARNING: RetryDataReplication is deprecated. Causes the data replication
// initiation sequence to begin immediately upon next Handshake for the specified
// Source Server ID, regardless of when the previous initiation started. This
// command will work only if the Source Server is stalled or is in a DISCONNECTED
// or STOPPED state.
//
// Deprecated: WARNING: RetryDataReplication is deprecated
func (c *Client) RetryDataReplication(ctx context.Context, params *RetryDataReplicationInput, optFns ...func(*Options)) (*RetryDataReplicationOutput, error) {
	if params == nil {
		params = &RetryDataReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetryDataReplication", params, optFns, c.addOperationRetryDataReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetryDataReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetryDataReplicationInput struct {

	// The ID of the Source Server whose data replication should be retried.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type RetryDataReplicationOutput struct {

	// The ARN of the Source Server.
	Arn *string

	// The Data Replication Info of the Source Server.
	DataReplicationInfo *types.DataReplicationInfo

	// The status of the last recovery launch of this Source Server.
	LastLaunchResult types.LastLaunchResult

	// The lifecycle information of this Source Server.
	LifeCycle *types.LifeCycle

	// The ID of the Recovery Instance associated with this Source Server.
	RecoveryInstanceId *string

	// Replication direction of the Source Server.
	ReplicationDirection types.ReplicationDirection

	// For EC2-originated Source Servers which have been failed over and then failed
	// back, this value will mean the ARN of the Source Server on the opposite
	// replication direction.
	ReversedDirectionSourceServerArn *string

	// Source cloud properties of the Source Server.
	SourceCloudProperties *types.SourceCloudProperties

	// ID of the Source Network which is protecting this Source Server's network.
	SourceNetworkID *string

	// The source properties of the Source Server.
	SourceProperties *types.SourceProperties

	// The ID of the Source Server.
	SourceServerID *string

	// The staging area of the source server.
	StagingArea *types.StagingArea

	// The tags associated with the Source Server.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetryDataReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRetryDataReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRetryDataReplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RetryDataReplication"); err != nil {
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
	if err = addOpRetryDataReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetryDataReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetryDataReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RetryDataReplication",
	}
}
