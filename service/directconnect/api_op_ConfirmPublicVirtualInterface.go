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
)

// Accepts ownership of a public virtual interface created by another AWS account.
// After the virtual interface owner makes this call, the specified virtual
// interface is created and made available to handle traffic.
func (c *Client) ConfirmPublicVirtualInterface(ctx context.Context, params *ConfirmPublicVirtualInterfaceInput, optFns ...func(*Options)) (*ConfirmPublicVirtualInterfaceOutput, error) {
	stack := middleware.NewStack("ConfirmPublicVirtualInterface", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpConfirmPublicVirtualInterfaceMiddlewares(stack)
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
	addOpConfirmPublicVirtualInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmPublicVirtualInterface(options.Region), middleware.Before)
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
			OperationName: "ConfirmPublicVirtualInterface",
			Err:           err,
		}
	}
	out := result.(*ConfirmPublicVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConfirmPublicVirtualInterfaceInput struct {

	// The ID of the virtual interface.
	//
	// This member is required.
	VirtualInterfaceId *string
}

type ConfirmPublicVirtualInterfaceOutput struct {

	// The state of the virtual interface. The following are the possible values:
	//
	//
	// * confirming: The creation of the virtual interface is pending confirmation from
	// the virtual interface owner. If the owner of the virtual interface is different
	// from the owner of the connection on which it is provisioned, then the virtual
	// interface will remain in this state until it is confirmed by the virtual
	// interface owner.
	//
	//     * verifying: This state only applies to public virtual
	// interfaces. Each public virtual interface needs validation before the virtual
	// interface can be created.
	//
	//     * pending: A virtual interface is in this state
	// from the time that it is created until the virtual interface is ready to forward
	// traffic.
	//
	//     * available: A virtual interface that is able to forward
	// traffic.
	//
	//     * down: A virtual interface that is BGP down.
	//
	//     * deleting: A
	// virtual interface is in this state immediately after calling
	// DeleteVirtualInterface () until it can no longer forward traffic.
	//
	//     *
	// deleted: A virtual interface that cannot forward traffic.
	//
	//     * rejected: The
	// virtual interface owner has declined creation of the virtual interface. If a
	// virtual interface in the Confirming state is deleted by the virtual interface
	// owner, the virtual interface enters the Rejected state.
	//
	//     * unknown: The
	// state of the virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpConfirmPublicVirtualInterfaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmPublicVirtualInterface{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmPublicVirtualInterface{}, middleware.After)
}

func newServiceMetadataMiddleware_opConfirmPublicVirtualInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "ConfirmPublicVirtualInterface",
	}
}
