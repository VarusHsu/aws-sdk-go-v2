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

// Associates a new node with the server. For more information about how to
// disassociate a node, see DisassociateNode (). On a Chef server: This command is
// an alternative to knife bootstrap. Example (Chef): aws opsworks-cm
// associate-node --server-name MyServer --node-name MyManagedNode
// --engine-attributes "Name=CHEF_ORGANIZATION,Value=default"
// "Name=CHEF_NODE_PUBLIC_KEY,Value=public-key-pem" On a Puppet server, this
// command is an alternative to the puppet cert sign command that signs a Puppet
// node CSR. Example (Puppet): aws opsworks-cm associate-node --server-name
// MyServer --node-name MyManagedNode --engine-attributes
// "Name=PUPPET_NODE_CSR,Value=csr-pem" A node can can only be associated with
// servers that are in a HEALTHY state. Otherwise, an InvalidStateException is
// thrown. A ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid. The
// AssociateNode API call can be integrated into Auto Scaling configurations, AWS
// Cloudformation templates, or the user data of a server's instance.
func (c *Client) AssociateNode(ctx context.Context, params *AssociateNodeInput, optFns ...func(*Options)) (*AssociateNodeOutput, error) {
	stack := middleware.NewStack("AssociateNode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateNodeMiddlewares(stack)
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
	addOpAssociateNodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateNode(options.Region), middleware.Before)
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
			OperationName: "AssociateNode",
			Err:           err,
		}
	}
	out := result.(*AssociateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateNodeInput struct {

	// Engine attributes used for associating the node. Attributes accepted in a
	// AssociateNode request for Chef
	//
	//     * CHEF_ORGANIZATION: The Chef organization
	// with which the node is associated. By default only one organization named
	// default can exist.
	//
	//     * CHEF_NODE_PUBLIC_KEY: A PEM-formatted public key. This
	// key is required for the chef-client agent to access the Chef API.
	//
	// Attributes
	// accepted in a AssociateNode request for Puppet
	//
	//     * PUPPET_NODE_CSR: A
	// PEM-formatted certificate-signing request (CSR) that is created by the node.
	//
	// This member is required.
	EngineAttributes []*types.EngineAttribute

	// The name of the server with which to associate the node.
	//
	// This member is required.
	ServerName *string

	// The name of the node.
	//
	// This member is required.
	NodeName *string
}

type AssociateNodeOutput struct {

	// Contains a token which can be passed to the DescribeNodeAssociationStatus API
	// call to get the status of the association request.
	NodeAssociationStatusToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateNodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateNode{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateNode{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateNode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "AssociateNode",
	}
}
