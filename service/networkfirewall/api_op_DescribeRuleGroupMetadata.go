// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// High-level information about a rule group, returned by operations like create
// and describe. You can use the information provided in the metadata to retrieve
// and manage a rule group. You can retrieve all objects for a rule group by
// calling DescribeRuleGroup.
func (c *Client) DescribeRuleGroupMetadata(ctx context.Context, params *DescribeRuleGroupMetadataInput, optFns ...func(*Options)) (*DescribeRuleGroupMetadataOutput, error) {
	if params == nil {
		params = &DescribeRuleGroupMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRuleGroupMetadata", params, optFns, c.addOperationDescribeRuleGroupMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRuleGroupMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRuleGroupMetadataInput struct {

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	RuleGroupArn *string

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	RuleGroupName *string

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules.
	//
	// This setting is required for requests that do not include the RuleGroupARN .
	Type types.RuleGroupType

	noSmithyDocumentSerde
}

type DescribeRuleGroupMetadataOutput struct {

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	//
	// This member is required.
	RuleGroupArn *string

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	//
	// This member is required.
	RuleGroupName *string

	// The maximum operating resources that this rule group can use. Rule group
	// capacity is fixed at creation. When you update a rule group, you are limited to
	// this capacity. When you reference a rule group from a firewall policy, Network
	// Firewall reserves this capacity for the rule group.
	//
	// You can retrieve the capacity that would be required for a rule group before
	// you create the rule group by calling CreateRuleGroupwith DryRun set to TRUE .
	Capacity *int32

	// Returns the metadata objects for the specified rule group.
	Description *string

	// The last time that the rule group was changed.
	LastModifiedTime *time.Time

	// Additional options governing how Network Firewall handles the rule group. You
	// can only use these for stateful rule groups.
	StatefulRuleOptions *types.StatefulRuleOptions

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules.
	//
	// This setting is required for requests that do not include the RuleGroupARN .
	Type types.RuleGroupType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRuleGroupMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRuleGroupMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRuleGroupMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRuleGroupMetadata"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRuleGroupMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRuleGroupMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRuleGroupMetadata",
	}
}
