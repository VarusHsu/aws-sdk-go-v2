// Code generated by smithy-go-codegen DO NOT EDIT.

package rbin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rbin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Locks a retention rule. A locked retention rule can't be modified or deleted.
func (c *Client) LockRule(ctx context.Context, params *LockRuleInput, optFns ...func(*Options)) (*LockRuleOutput, error) {
	if params == nil {
		params = &LockRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LockRule", params, optFns, c.addOperationLockRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LockRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type LockRuleInput struct {

	// The unique ID of the retention rule.
	//
	// This member is required.
	Identifier *string

	// Information about the retention rule lock configuration.
	//
	// This member is required.
	LockConfiguration *types.LockConfiguration

	noSmithyDocumentSerde
}

type LockRuleOutput struct {

	// The retention rule description.
	Description *string

	// The unique ID of the retention rule.
	Identifier *string

	// Information about the retention rule lock configuration.
	LockConfiguration *types.LockConfiguration

	// The lock state for the retention rule.
	//   - locked - The retention rule is locked and can't be modified or deleted.
	//   - pending_unlock - The retention rule has been unlocked but it is still within
	//   the unlock delay period. The retention rule can be modified or deleted only
	//   after the unlock delay period has expired.
	//   - unlocked - The retention rule is unlocked and it can be modified or deleted
	//   by any user with the required permissions.
	//   - null - The retention rule has never been locked. Once a retention rule has
	//   been locked, it can transition between the locked and unlocked states only; it
	//   can never transition back to null .
	LockState types.LockState

	// Information about the resource tags used to identify resources that are
	// retained by the retention rule.
	ResourceTags []types.ResourceTag

	// The resource type retained by the retention rule.
	ResourceType types.ResourceType

	// Information about the retention period for which the retention rule is to
	// retain resources.
	RetentionPeriod *types.RetentionPeriod

	// The state of the retention rule. Only retention rules that are in the available
	// state retain resources.
	Status types.RuleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationLockRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpLockRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpLockRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "LockRule"); err != nil {
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
	if err = addOpLockRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLockRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opLockRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "LockRule",
	}
}
