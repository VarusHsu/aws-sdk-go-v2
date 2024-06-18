// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an IP access control group.
//
// An IP access control group provides you with the ability to control the IP
// addresses from which users are allowed to access their WorkSpaces. To specify
// the CIDR address ranges, add rules to your IP access control group and then
// associate the group with your directory. You can add rules when you create the
// group or at any time using AuthorizeIpRules.
//
// There is a default IP access control group associated with your directory. If
// you don't associate an IP access control group with your directory, the default
// group is used. The default group includes a default rule that allows users to
// access their WorkSpaces from anywhere. You cannot modify the default IP access
// control group for your directory.
func (c *Client) CreateIpGroup(ctx context.Context, params *CreateIpGroupInput, optFns ...func(*Options)) (*CreateIpGroupOutput, error) {
	if params == nil {
		params = &CreateIpGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIpGroup", params, optFns, c.addOperationCreateIpGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIpGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIpGroupInput struct {

	// The name of the group.
	//
	// This member is required.
	GroupName *string

	// The description of the group.
	GroupDesc *string

	// The tags. Each WorkSpaces resource can have a maximum of 50 tags.
	Tags []types.Tag

	// The rules to add to the group.
	UserRules []types.IpRuleItem

	noSmithyDocumentSerde
}

type CreateIpGroupOutput struct {

	// The identifier of the group.
	GroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIpGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIpGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIpGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIpGroup"); err != nil {
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
	if err = addOpCreateIpGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIpGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIpGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIpGroup",
	}
}
