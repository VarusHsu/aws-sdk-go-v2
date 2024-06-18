// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the MembershipId in an identity store.
//
// If you have administrator access to a member account, you can use this API from
// the member account. Read about [member accounts]in the Organizations User Guide.
//
// [member accounts]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html
func (c *Client) GetGroupMembershipId(ctx context.Context, params *GetGroupMembershipIdInput, optFns ...func(*Options)) (*GetGroupMembershipIdOutput, error) {
	if params == nil {
		params = &GetGroupMembershipIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupMembershipId", params, optFns, c.addOperationGetGroupMembershipIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupMembershipIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupMembershipIdInput struct {

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// An object that contains the identifier of a group member. Setting the UserID
	// field to the specific identifier for a user indicates that the user is a member
	// of the group.
	//
	// This member is required.
	MemberId types.MemberId

	noSmithyDocumentSerde
}

type GetGroupMembershipIdOutput struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The identifier for a GroupMembership in an identity store.
	//
	// This member is required.
	MembershipId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGroupMembershipIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetGroupMembershipId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetGroupMembershipId{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGroupMembershipId"); err != nil {
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
	if err = addOpGetGroupMembershipIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupMembershipId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGroupMembershipId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGroupMembershipId",
	}
}
