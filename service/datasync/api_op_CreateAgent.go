// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates an DataSync agent that you've deployed in your storage environment.
// The activation process associates the agent with your Amazon Web Services
// account. If you haven't deployed an agent yet, see the following topics to learn
// more:
//   - Agent requirements (https://docs.aws.amazon.com/datasync/latest/userguide/agent-requirements.html)
//   - Create an agent (https://docs.aws.amazon.com/datasync/latest/userguide/configure-agent.html)
//
// If you're transferring between Amazon Web Services storage services, you don't
// need a DataSync agent.
func (c *Client) CreateAgent(ctx context.Context, params *CreateAgentInput, optFns ...func(*Options)) (*CreateAgentOutput, error) {
	if params == nil {
		params = &CreateAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAgent", params, optFns, c.addOperationCreateAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateAgentRequest
type CreateAgentInput struct {

	// Specifies your DataSync agent's activation key. If you don't have an activation
	// key, see Activate your agent (https://docs.aws.amazon.com/datasync/latest/userguide/activate-agent.html)
	// .
	//
	// This member is required.
	ActivationKey *string

	// Specifies a name for your agent. You can see this name in the DataSync console.
	AgentName *string

	// Specifies the Amazon Resource Name (ARN) of the security group that protects
	// your task's network interfaces (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// when using a virtual private cloud (VPC) endpoint (https://docs.aws.amazon.com/datasync/latest/userguide/choose-service-endpoint.html#choose-service-endpoint-vpc)
	// . You can only specify one ARN.
	SecurityGroupArns []string

	// Specifies the ARN of the subnet where you want to run your DataSync task when
	// using a VPC endpoint. This is the subnet where DataSync creates and manages the
	// network interfaces (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// for your transfer. You can only specify one ARN.
	SubnetArns []string

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least one tag for your agent.
	Tags []types.TagListEntry

	// Specifies the ID of the VPC endpoint that you want your agent to connect to.
	// For example, a VPC endpoint ID looks like vpce-01234d5aff67890e1 . The VPC
	// endpoint you use must include the DataSync service name (for example,
	// com.amazonaws.us-east-2.datasync ).
	VpcEndpointId *string

	noSmithyDocumentSerde
}

// CreateAgentResponse
type CreateAgentOutput struct {

	// The ARN of the agent that you just activated. Use the ListAgents (https://docs.aws.amazon.com/datasync/latest/userguide/API_ListAgents.html)
	// operation to return a list of agents in your Amazon Web Services account and
	// Amazon Web Services Region.
	AgentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAgent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAgent"); err != nil {
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
	if err = addOpCreateAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAgent",
	}
}
