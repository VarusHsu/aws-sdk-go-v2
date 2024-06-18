// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assign a registered instance to a layer.
//
//   - You can assign registered on-premises instances to any layer type.
//
//   - You can assign registered Amazon EC2 instances only to custom layers.
//
//   - You cannot use this action with instances that were created with OpsWorks
//     Stacks.
//
// Required Permissions: To use this action, an Identity and Access Management
// (IAM) user must have a Manage permissions level for the stack or an attached
// policy that explicitly grants permissions. For more information on user
// permissions, see [Managing User Permissions].
//
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) AssignInstance(ctx context.Context, params *AssignInstanceInput, optFns ...func(*Options)) (*AssignInstanceOutput, error) {
	if params == nil {
		params = &AssignInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssignInstance", params, optFns, c.addOperationAssignInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssignInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignInstanceInput struct {

	// The instance ID.
	//
	// This member is required.
	InstanceId *string

	// The layer ID, which must correspond to a custom layer. You cannot assign a
	// registered instance to a built-in layer.
	//
	// This member is required.
	LayerIds []string

	noSmithyDocumentSerde
}

type AssignInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssignInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssignInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssignInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssignInstance"); err != nil {
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
	if err = addOpAssignInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssignInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssignInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssignInstance",
	}
}
