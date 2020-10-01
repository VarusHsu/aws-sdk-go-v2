// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Disassociates a connection from a link aggregation group (LAG). The connection
// is interrupted and re-established as a standalone connection (the connection is
// not deleted; to delete the connection, use the DeleteConnection () request). If
// the LAG has associated virtual interfaces or hosted connections, they remain
// associated with the LAG. A disassociated connection owned by an AWS Direct
// Connect Partner is automatically converted to an interconnect. If disassociating
// the connection would cause the LAG to fall below its setting for minimum number
// of operational connections, the request fails, except when it's the last member
// of the LAG. If all connections are disassociated, the LAG continues to exist as
// an empty LAG with no physical connections.
func (c *Client) DisassociateConnectionFromLag(ctx context.Context, params *DisassociateConnectionFromLagInput, optFns ...func(*Options)) (*DisassociateConnectionFromLagOutput, error) {
	stack := middleware.NewStack("DisassociateConnectionFromLag", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateConnectionFromLagMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDisassociateConnectionFromLagValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateConnectionFromLag(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DisassociateConnectionFromLag",
			Err:           err,
		}
	}
	out := result.(*DisassociateConnectionFromLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateConnectionFromLagInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// The ID of the LAG.
	//
	// This member is required.
	LagId *string
}

// Information about an AWS Direct Connect connection.
type DisassociateConnectionFromLagOutput struct {

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string

	// The ID of the VLAN.
	Vlan *int32

	// The name of the connection.
	ConnectionName *string

	// The AWS Region where the connection is located.
	Region *string

	// The ID of the LAG.
	LagId *string

	// The ID of the connection.
	ConnectionId *string

	// The tags associated with the connection.
	Tags []*types.Tag

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string

	// The state of the connection. The following are the possible values:
	//
	//     *
	// ordering: The initial state of a hosted connection provisioned on an
	// interconnect. The connection stays in the ordering state until the owner of the
	// hosted connection confirms or declines the connection order.
	//
	//     * requested:
	// The initial state of a standard connection. The connection stays in the
	// requested state until the Letter of Authorization (LOA) is sent to the
	// customer.
	//
	//     * pending: The connection has been approved and is being
	// initialized.
	//
	//     * available: The network link is up and the connection is
	// ready for use.
	//
	//     * down: The network link is down.
	//
	//     * deleting: The
	// connection is being deleted.
	//
	//     * deleted: The connection has been deleted.
	//
	//
	// * rejected: A hosted connection in the ordering state enters the rejected state
	// if it is deleted by the customer.
	//
	//     * unknown: The state of the connection is
	// not available.
	ConnectionState types.ConnectionState

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string

	// The name of the AWS Direct Connect service provider associated with the
	// connection.
	PartnerName *string

	// The time of the most recent call to DescribeLoa () for this connection.
	LoaIssueTime *time.Time

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The location of the connection.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateConnectionFromLagMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateConnectionFromLag{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateConnectionFromLag{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateConnectionFromLag(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DisassociateConnectionFromLag",
	}
}
