// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a state machine's definition, its IAM role Amazon
// Resource Name (ARN), and configuration.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
//   - The following qualified state machine ARN refers to a Distributed Map state
//     with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// If you provide a qualified state machine ARN that refers to a Distributed Map
//
//	state, the request fails with ValidationException .
//
//	- The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// If you provide a qualified state machine ARN that refers to a version ARN or an
//
//	alias ARN, the request starts execution for that version or alias.
//
//	- The following unqualified state machine ARN refers to a state machine named
//	myStateMachine .
//
// arn::states:::stateMachine:
//
// This API action returns the details for a state machine version if the
// stateMachineArn you specify is a state machine version ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func (c *Client) DescribeStateMachine(ctx context.Context, params *DescribeStateMachineInput, optFns ...func(*Options)) (*DescribeStateMachineOutput, error) {
	if params == nil {
		params = &DescribeStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachine", params, optFns, c.addOperationDescribeStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine for which you want the
	// information.
	//
	// If you specify a state machine version ARN, this API returns details about that
	// version. The version ARN is a combination of state machine ARN and the version
	// number separated by a colon (:). For example, stateMachineARN:1 .
	//
	// This member is required.
	StateMachineArn *string

	noSmithyDocumentSerde
}

type DescribeStateMachineOutput struct {

	// The date the state machine is created.
	//
	// For a state machine version, creationDate is the date the version was created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon States Language definition of the state machine. See [Amazon States Language].
	//
	// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
	//
	// This member is required.
	Definition *string

	// The name of the state machine.
	//
	// A name must not contain:
	//
	//   - white space
	//
	//   - brackets < > { } [ ]
	//
	//   - wildcard characters ? *
	//
	//   - special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//   - control characters ( U+0000-001F , U+007F-009F )
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z,
	// a-z, - and _.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role used when creating this state
	// machine. (The IAM role maintains security by granting Step Functions access to
	// Amazon Web Services resources.)
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) that identifies the state machine.
	//
	// If you specified a state machine version ARN in your request, the API returns
	// the version ARN. The version ARN is a combination of state machine ARN and the
	// version number separated by a colon (:). For example, stateMachineARN:1 .
	//
	// This member is required.
	StateMachineArn *string

	// The type of the state machine ( STANDARD or EXPRESS ).
	//
	// This member is required.
	Type types.StateMachineType

	// The description of the state machine version.
	Description *string

	// A user-defined or an auto-generated string that identifies a Map state. This
	// parameter is present only if the stateMachineArn specified in input is a
	// qualified state machine ARN.
	Label *string

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// The revision identifier for the state machine.
	//
	// Use the revisionId parameter to compare between versions of a state machine
	// configuration used for executions without performing a diff of the properties,
	// such as definition and roleArn .
	RevisionId *string

	// The current status of the state machine.
	Status types.StateMachineStatus

	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStateMachine"); err != nil {
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
	if err = addOpDescribeStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStateMachine",
	}
}
