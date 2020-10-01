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

// Lists the attachments between your Direct Connect gateways and virtual
// interfaces. You must specify a Direct Connect gateway, a virtual interface, or
// both. If you specify a Direct Connect gateway, the response contains all virtual
// interfaces attached to the Direct Connect gateway. If you specify a virtual
// interface, the response contains all Direct Connect gateways attached to the
// virtual interface. If you specify both, the response contains the attachment
// between the Direct Connect gateway and the virtual interface.
func (c *Client) DescribeDirectConnectGatewayAttachments(ctx context.Context, params *DescribeDirectConnectGatewayAttachmentsInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewayAttachmentsOutput, error) {
	stack := middleware.NewStack("DescribeDirectConnectGatewayAttachments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDirectConnectGatewayAttachmentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAttachments(options.Region), middleware.Before)
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
			OperationName: "DescribeDirectConnectGatewayAttachments",
			Err:           err,
		}
	}
	out := result.(*DescribeDirectConnectGatewayAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewayAttachmentsInput struct {

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token provided in the previous call to retrieve the next page.
	NextToken *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string
}

type DescribeDirectConnectGatewayAttachmentsOutput struct {

	// The token to retrieve the next page.
	NextToken *string

	// The attachments.
	DirectConnectGatewayAttachments []*types.DirectConnectGatewayAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDirectConnectGatewayAttachmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGatewayAttachments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGatewayAttachments{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAttachments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGatewayAttachments",
	}
}
