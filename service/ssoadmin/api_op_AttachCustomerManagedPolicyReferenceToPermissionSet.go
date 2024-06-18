// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches the specified customer managed policy to the specified PermissionSet.
func (c *Client) AttachCustomerManagedPolicyReferenceToPermissionSet(ctx context.Context, params *AttachCustomerManagedPolicyReferenceToPermissionSetInput, optFns ...func(*Options)) (*AttachCustomerManagedPolicyReferenceToPermissionSetOutput, error) {
	if params == nil {
		params = &AttachCustomerManagedPolicyReferenceToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachCustomerManagedPolicyReferenceToPermissionSet", params, optFns, c.addOperationAttachCustomerManagedPolicyReferenceToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachCustomerManagedPolicyReferenceToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachCustomerManagedPolicyReferenceToPermissionSetInput struct {

	// Specifies the name and path of a customer managed policy. You must have an IAM
	// policy that matches the name and path in each Amazon Web Services account where
	// you want to deploy your permission set.
	//
	// This member is required.
	CustomerManagedPolicyReference *types.CustomerManagedPolicyReference

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet .
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type AttachCustomerManagedPolicyReferenceToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachCustomerManagedPolicyReferenceToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachCustomerManagedPolicyReferenceToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachCustomerManagedPolicyReferenceToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AttachCustomerManagedPolicyReferenceToPermissionSet"); err != nil {
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
	if err = addOpAttachCustomerManagedPolicyReferenceToPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachCustomerManagedPolicyReferenceToPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachCustomerManagedPolicyReferenceToPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AttachCustomerManagedPolicyReferenceToPermissionSet",
	}
}
