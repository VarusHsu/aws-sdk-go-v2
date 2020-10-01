// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current status of an existing association or disassociation request.
// A ResourceNotFoundException is thrown when no recent association or
// disassociation request with the specified token is found, or when the server
// does not exist. A ValidationException is raised when parameters of the request
// are not valid.
func (c *Client) DescribeNodeAssociationStatus(ctx context.Context, params *DescribeNodeAssociationStatusInput, optFns ...func(*Options)) (*DescribeNodeAssociationStatusOutput, error) {
	stack := middleware.NewStack("DescribeNodeAssociationStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeNodeAssociationStatusMiddlewares(stack)
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
	addOpDescribeNodeAssociationStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(options.Region), middleware.Before)
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
			OperationName: "DescribeNodeAssociationStatus",
			Err:           err,
		}
	}
	out := result.(*DescribeNodeAssociationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeAssociationStatusInput struct {

	// The name of the server from which to disassociate the node.
	//
	// This member is required.
	ServerName *string

	// The token returned in either the AssociateNodeResponse or the
	// DisassociateNodeResponse.
	//
	// This member is required.
	NodeAssociationStatusToken *string
}

type DescribeNodeAssociationStatusOutput struct {

	// The status of the association or disassociation request. Possible values:
	//
	//     *
	// SUCCESS: The association or disassociation succeeded.
	//
	//     * FAILED: The
	// association or disassociation failed.
	//
	//     * IN_PROGRESS: The association or
	// disassociation is still in progress.
	NodeAssociationStatus types.NodeAssociationStatus

	// Attributes specific to the node association. In Puppet, the attibute
	// PUPPET_NODE_CERT contains the signed certificate (the result of the CSR).
	EngineAttributes []*types.EngineAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeNodeAssociationStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNodeAssociationStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNodeAssociationStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeNodeAssociationStatus",
	}
}
