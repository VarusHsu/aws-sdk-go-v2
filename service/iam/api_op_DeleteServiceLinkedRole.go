// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits a service-linked role deletion request and returns a DeletionTaskId ,
// which you can use to check the status of the deletion. Before you call this
// operation, confirm that the role has no active sessions and that any resources
// used by the role in the linked service are deleted. If you call this operation
// more than once for the same service-linked role and an earlier deletion task is
// not complete, then the DeletionTaskId of the earlier request is returned. If
// you submit a deletion request for a service-linked role whose linked service is
// still accessing a resource, then the deletion task fails. If it fails, the
// GetServiceLinkedRoleDeletionStatus operation returns the reason for the failure,
// usually including the resources that must be deleted. To delete the
// service-linked role, you must first remove those resources from the linked
// service and then submit the deletion request again. Resources are specific to
// the service that is linked to the role. For more information about removing
// resources from a service, see the Amazon Web Services documentation (http://docs.aws.amazon.com/)
// for your service. For more information about service-linked roles, see Roles
// terms and concepts: Amazon Web Services service-linked role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role)
// in the IAM User Guide.
func (c *Client) DeleteServiceLinkedRole(ctx context.Context, params *DeleteServiceLinkedRoleInput, optFns ...func(*Options)) (*DeleteServiceLinkedRoleOutput, error) {
	if params == nil {
		params = &DeleteServiceLinkedRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServiceLinkedRole", params, optFns, c.addOperationDeleteServiceLinkedRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceLinkedRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceLinkedRoleInput struct {

	// The name of the service-linked role to be deleted.
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type DeleteServiceLinkedRoleOutput struct {

	// The deletion task identifier that you can use to check the status of the
	// deletion. This identifier is returned in the format task/aws-service-role/// .
	//
	// This member is required.
	DeletionTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServiceLinkedRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteServiceLinkedRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteServiceLinkedRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteServiceLinkedRole"); err != nil {
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
	if err = addOpDeleteServiceLinkedRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServiceLinkedRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteServiceLinkedRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteServiceLinkedRole",
	}
}
