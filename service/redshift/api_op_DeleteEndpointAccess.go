// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a Redshift-managed VPC endpoint.
func (c *Client) DeleteEndpointAccess(ctx context.Context, params *DeleteEndpointAccessInput, optFns ...func(*Options)) (*DeleteEndpointAccessOutput, error) {
	if params == nil {
		params = &DeleteEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEndpointAccess", params, optFns, c.addOperationDeleteEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEndpointAccessInput struct {

	// The Redshift-managed VPC endpoint to delete.
	//
	// This member is required.
	EndpointName *string

	noSmithyDocumentSerde
}

// Describes a Redshift-managed VPC endpoint.
type DeleteEndpointAccessOutput struct {

	// The DNS address of the endpoint.
	Address *string

	// The cluster identifier of the cluster associated with the endpoint.
	ClusterIdentifier *string

	// The time (UTC) that the endpoint was created.
	EndpointCreateTime *time.Time

	// The name of the endpoint.
	EndpointName *string

	// The status of the endpoint.
	EndpointStatus *string

	// The port number on which the cluster accepts incoming connections.
	Port *int32

	// The Amazon Web Services account ID of the owner of the cluster.
	ResourceOwner *string

	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName *string

	// The connection endpoint for connecting to an Amazon Redshift cluster through
	// the proxy.
	VpcEndpoint *types.VpcEndpoint

	// The security groups associated with the endpoint.
	VpcSecurityGroups []types.VpcSecurityGroupMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteEndpointAccess"); err != nil {
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
	if err = addOpDeleteEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteEndpointAccess",
	}
}
