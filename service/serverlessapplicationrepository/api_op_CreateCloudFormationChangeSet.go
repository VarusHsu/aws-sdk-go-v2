// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AWS CloudFormation change set for the given application.
func (c *Client) CreateCloudFormationChangeSet(ctx context.Context, params *CreateCloudFormationChangeSetInput, optFns ...func(*Options)) (*CreateCloudFormationChangeSetOutput, error) {
	if params == nil {
		params = &CreateCloudFormationChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCloudFormationChangeSet", params, optFns, c.addOperationCreateCloudFormationChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCloudFormationChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCloudFormationChangeSetInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	//
	// This member is required.
	StackName *string

	// A list of values that you must specify before you can deploy certain
	// applications. Some applications might include resources that can affect
	// permissions in your AWS account, for example, by creating new AWS Identity and
	// Access Management (IAM) users. For those applications, you must explicitly
	// acknowledge their capabilities by specifying this parameter.
	//
	// The only valid values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM,
	// CAPABILITY_RESOURCE_POLICY, and CAPABILITY_AUTO_EXPAND.
	//
	// The following resources require you to specify CAPABILITY_IAM or
	// CAPABILITY_NAMED_IAM: [AWS::IAM::Group], [AWS::IAM::InstanceProfile], [AWS::IAM::Policy], and [AWS::IAM::Role]. If the application contains IAM resources,
	// you can specify either CAPABILITY_IAM or CAPABILITY_NAMED_IAM. If the
	// application contains IAM resources with custom names, you must specify
	// CAPABILITY_NAMED_IAM.
	//
	// The following resources require you to specify CAPABILITY_RESOURCE_POLICY: [AWS::Lambda::Permission], [AWS::IAM:Policy], [AWS::ApplicationAutoScaling::ScalingPolicy]
	// , [AWS::S3::BucketPolicy], [AWS::SQS::QueuePolicy], and [AWS::SNS:TopicPolicy].
	//
	// Applications that contain one or more nested applications require you to
	// specify CAPABILITY_AUTO_EXPAND.
	//
	// If your application template contains any of the above resources, we recommend
	// that you review all permissions associated with the application before
	// deploying. If you don't specify this parameter for an application that requires
	// capabilities, the call will fail.
	//
	// [AWS::SQS::QueuePolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html
	// [AWS::ApplicationAutoScaling::ScalingPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
	// [AWS::S3::BucketPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html
	// [AWS::IAM::InstanceProfile]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
	// [AWS::IAM::Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
	// [AWS::IAM::Group]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
	// [AWS::IAM::Role]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
	// [AWS::IAM:Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
	// [AWS::SNS:TopicPolicy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
	// [AWS::Lambda::Permission]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
	Capabilities []string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	ChangeSetName *string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	ClientToken *string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	Description *string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	NotificationArns []string

	// A list of parameter values for the parameters of the application.
	ParameterOverrides []types.ParameterValue

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	ResourceTypes []string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	RollbackConfiguration *types.RollbackConfiguration

	// The semantic version of the application:
	//
	// [https://semver.org/]
	//
	// [https://semver.org/]: https://semver.org/
	SemanticVersion *string

	// This property corresponds to the parameter of the same name for the AWS
	// CloudFormation [CreateChangeSet]API.
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet
	Tags []types.Tag

	// The UUID returned by CreateCloudFormationTemplate.
	//
	// Pattern:
	// [0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}
	TemplateId *string

	noSmithyDocumentSerde
}

type CreateCloudFormationChangeSetOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the change set.
	//
	// Length constraints: Minimum length of 1.
	//
	// Pattern: ARN:[-a-zA-Z0-9:/]*
	ChangeSetId *string

	// The semantic version of the application:
	//
	// [https://semver.org/]
	//
	// [https://semver.org/]: https://semver.org/
	SemanticVersion *string

	// The unique ID of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCloudFormationChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCloudFormationChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCloudFormationChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCloudFormationChangeSet"); err != nil {
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
	if err = addOpCreateCloudFormationChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFormationChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCloudFormationChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCloudFormationChangeSet",
	}
}
