// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you attach a resource-based policy to a resource, RAM automatically
// creates a resource share of featureSet = CREATED_FROM_POLICY with a managed
// permission that has the same IAM permissions as the original resource-based
// policy. However, this type of managed permission is visible to only the resource
// share owner, and the associated resource share can't be modified by using RAM.
// This operation promotes the resource share to a STANDARD resource share that is
// fully manageable in RAM. When you promote a resource share, you can then manage
// the resource share in RAM and it becomes visible to all of the principals you
// shared it with. Before you perform this operation, you should first run
// PromotePermissionCreatedFromPolicy to ensure that you have an appropriate
// customer managed permission that can be associated with this resource share
// after its is promoted. If this operation can't find a managed permission that
// exactly matches the existing CREATED_FROM_POLICY permission, then this
// operation fails.
func (c *Client) PromoteResourceShareCreatedFromPolicy(ctx context.Context, params *PromoteResourceShareCreatedFromPolicyInput, optFns ...func(*Options)) (*PromoteResourceShareCreatedFromPolicyOutput, error) {
	if params == nil {
		params = &PromoteResourceShareCreatedFromPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PromoteResourceShareCreatedFromPolicy", params, optFns, c.addOperationPromoteResourceShareCreatedFromPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PromoteResourceShareCreatedFromPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PromoteResourceShareCreatedFromPolicyInput struct {

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share to promote.
	//
	// This member is required.
	ResourceShareArn *string

	noSmithyDocumentSerde
}

type PromoteResourceShareCreatedFromPolicyOutput struct {

	// A return value of true indicates that the request succeeded. A value of false
	// indicates that the request failed.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPromoteResourceShareCreatedFromPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPromoteResourceShareCreatedFromPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "PromoteResourceShareCreatedFromPolicy",
	}
}
