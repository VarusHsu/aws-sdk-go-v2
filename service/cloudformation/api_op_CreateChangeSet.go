// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a list of changes that will be applied to a stack so that you can
// review the changes before executing them. You can create a change set for a
// stack that doesn't exist or an existing stack. If you create a change set for a
// stack that doesn't exist, the change set shows all of the resources that
// CloudFormation will create. If you create a change set for an existing stack,
// CloudFormation compares the stack's information with the information that you
// submit in the change set and lists the differences. Use change sets to
// understand which resources CloudFormation will create or change, and how it will
// change resources in an existing stack, before you create or update a stack.
//
// To create a change set for a stack that doesn't exist, for the ChangeSetType
// parameter, specify CREATE . To create a change set for an existing stack,
// specify UPDATE for the ChangeSetType parameter. To create a change set for an
// import operation, specify IMPORT for the ChangeSetType parameter. After the
// CreateChangeSet call successfully completes, CloudFormation starts creating the
// change set. To check the status of the change set or to review it, use the DescribeChangeSet
// action.
//
// When you are satisfied with the changes the change set will make, execute the
// change set by using the ExecuteChangeSetaction. CloudFormation doesn't make changes until you
// execute the change set.
//
// To create a change set for the entire stack hierarchy, set IncludeNestedStacks
// to True .
func (c *Client) CreateChangeSet(ctx context.Context, params *CreateChangeSetInput, optFns ...func(*Options)) (*CreateChangeSetOutput, error) {
	if params == nil {
		params = &CreateChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChangeSet", params, optFns, c.addOperationCreateChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateChangeSet action.
type CreateChangeSetInput struct {

	// The name of the change set. The name must be unique among all change sets that
	// are associated with the specified stack.
	//
	// A change set name can contain only alphanumeric, case sensitive characters, and
	// hyphens. It must start with an alphabetical character and can't exceed 128
	// characters.
	//
	// This member is required.
	ChangeSetName *string

	// The name or the unique ID of the stack for which you are creating a change set.
	// CloudFormation generates the change set by comparing this stack's information
	// with the information that you submit, such as a modified template or different
	// parameter input values.
	//
	// This member is required.
	StackName *string

	// In some cases, you must explicitly acknowledge that your stack template
	// contains certain capabilities in order for CloudFormation to create the stack.
	//
	//   - CAPABILITY_IAM and CAPABILITY_NAMED_IAM
	//
	// Some stack templates might include resources that can affect permissions in
	//   your Amazon Web Services account; for example, by creating new Identity and
	//   Access Management (IAM) users. For those stacks, you must explicitly acknowledge
	//   this by specifying one of these capabilities.
	//
	// The following IAM resources require you to specify either the CAPABILITY_IAM or
	//   CAPABILITY_NAMED_IAM capability.
	//
	//   - If you have IAM resources, you can specify either capability.
	//
	//   - If you have IAM resources with custom names, you must specify
	//   CAPABILITY_NAMED_IAM .
	//
	//   - If you don't specify either of these capabilities, CloudFormation returns
	//   an InsufficientCapabilities error.
	//
	// If your stack template contains these resources, we suggest that you review all
	//   permissions associated with them and edit their permissions if necessary.
	//
	// [AWS::IAM::AccessKey]
	//
	// [AWS::IAM::Group]
	//
	// [AWS::IAM::InstanceProfile]
	//
	// [AWS::IAM::Policy]
	//
	// [AWS::IAM::Role]
	//
	// [AWS::IAM::User]
	//
	// [AWS::IAM::UserToGroupAddition]
	//
	// For more information, see [Acknowledging IAM resources in CloudFormation templates].
	//
	//   - CAPABILITY_AUTO_EXPAND
	//
	// Some template contain macros. Macros perform custom processing on templates;
	//   this can include simple actions like find-and-replace operations, all the way to
	//   extensive transformations of entire templates. Because of this, users typically
	//   create a change set from the processed template, so that they can review the
	//   changes resulting from the macros before actually creating the stack. If your
	//   stack template contains one or more macros, and you choose to create a stack
	//   directly from the processed template, without first reviewing the resulting
	//   changes in a change set, you must acknowledge this capability. This includes the
	//   [AWS::Include]and [AWS::Serverless]transforms, which are macros hosted by CloudFormation.
	//
	// This capacity doesn't apply to creating change sets, and specifying it when
	//   creating change sets has no effect.
	//
	// If you want to create a stack from a stack template that contains macros and
	//   nested stacks, you must create or update the stack directly from the template
	//   using the CreateStackor UpdateStackaction, and specifying this capability.
	//
	// For more information about macros, see [Using CloudFormation macros to perform custom processing on templates].
	//
	// Only one of the Capabilities and ResourceType parameters can be specified.
	//
	// [AWS::IAM::AccessKey]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
	// [AWS::Include]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html
	// [AWS::IAM::User]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
	// [AWS::IAM::InstanceProfile]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
	// [Acknowledging IAM resources in CloudFormation templates]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities
	// [AWS::IAM::Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
	// [AWS::IAM::Group]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
	// [AWS::IAM::UserToGroupAddition]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
	// [AWS::IAM::Role]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
	// [Using CloudFormation macros to perform custom processing on templates]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html
	// [AWS::Serverless]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html
	Capabilities []types.Capability

	// The type of change set operation. To create a change set for a new stack,
	// specify CREATE . To create a change set for an existing stack, specify UPDATE .
	// To create a change set for an import operation, specify IMPORT .
	//
	// If you create a change set for a new stack, CloudFormation creates a stack with
	// a unique stack ID, but no template or resources. The stack will be in the [REVIEW_IN_PROGRESS]state
	// until you execute the change set.
	//
	// By default, CloudFormation specifies UPDATE . You can't use the UPDATE type to
	// create a change set for a new stack or the CREATE type to create a change set
	// for an existing stack.
	//
	// [REVIEW_IN_PROGRESS]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-describing-stacks.html#d0e11995
	ChangeSetType types.ChangeSetType

	// A unique identifier for this CreateChangeSet request. Specify this token if you
	// plan to retry requests so that CloudFormation knows that you're not attempting
	// to create another change set with the same name. You might retry CreateChangeSet
	// requests to ensure that CloudFormation successfully received them.
	ClientToken *string

	// A description to help you identify this change set.
	Description *string

	// Indicates if the change set imports resources that already exist.
	//
	// This parameter can only import resources that have custom names in templates.
	// For more information, see [name type]in the CloudFormation User Guide. To import resources
	// that do not accept custom names, such as EC2 instances, use the resource import
	// feature instead. For more information, see [Bringing existing resources into CloudFormation management]in the CloudFormation User Guide.
	//
	// [Bringing existing resources into CloudFormation management]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import.html
	// [name type]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html
	ImportExistingResources *bool

	// Creates a change set for the all nested stacks specified in the template. The
	// default behavior of this action is set to False . To include nested sets in a
	// change set, specify True .
	IncludeNestedStacks *bool

	// The Amazon Resource Names (ARNs) of Amazon Simple Notification Service (Amazon
	// SNS) topics that CloudFormation associates with the stack. To remove all
	// associated notification topics, specify an empty list.
	NotificationARNs []string

	// Determines what action will be taken if stack creation fails. If this parameter
	// is specified, the DisableRollback parameter to the [ExecuteChangeSet] API operation must not be
	// specified. This must be one of these values:
	//
	//   - DELETE - Deletes the change set if the stack creation fails. This is only
	//   valid when the ChangeSetType parameter is set to CREATE . If the deletion of
	//   the stack fails, the status of the stack is DELETE_FAILED .
	//
	//   - DO_NOTHING - if the stack creation fails, do nothing. This is equivalent to
	//   specifying true for the DisableRollback parameter to the [ExecuteChangeSet]API operation.
	//
	//   - ROLLBACK - if the stack creation fails, roll back the stack. This is
	//   equivalent to specifying false for the DisableRollback parameter to the [ExecuteChangeSet]API
	//   operation.
	//
	// For nested stacks, when the OnStackFailure parameter is set to DELETE for the
	// change set for the parent stack, any failure in a child stack will cause the
	// parent stack creation to fail and all stacks to be deleted.
	//
	// [ExecuteChangeSet]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteChangeSet.html
	OnStackFailure types.OnStackFailure

	// A list of Parameter structures that specify input parameters for the change
	// set. For more information, see the Parameterdata type.
	Parameters []types.Parameter

	// The template resource types that you have permissions to work with if you
	// execute this change set, such as AWS::EC2::Instance , AWS::EC2::* , or
	// Custom::MyCustomInstance .
	//
	// If the list of resource types doesn't include a resource type that you're
	// updating, the stack update fails. By default, CloudFormation grants permissions
	// to all resource types. Identity and Access Management (IAM) uses this parameter
	// for condition keys in IAM policies for CloudFormation. For more information, see
	// [Controlling access with Identity and Access Management]in the CloudFormation User Guide.
	//
	// Only one of the Capabilities and ResourceType parameters can be specified.
	//
	// [Controlling access with Identity and Access Management]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html
	ResourceTypes []string

	// The resources to import into your stack.
	ResourcesToImport []types.ResourceToImport

	// The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role
	// that CloudFormation assumes when executing the change set. CloudFormation uses
	// the role's credentials to make calls on your behalf. CloudFormation uses this
	// role for all future operations on the stack. Provided that users have permission
	// to operate on the stack, CloudFormation uses this role even if the users don't
	// have permission to pass it. Ensure that the role grants least permission.
	//
	// If you don't specify a value, CloudFormation uses the role that was previously
	// associated with the stack. If no role is available, CloudFormation uses a
	// temporary session that is generated from your user credentials.
	RoleARN *string

	// The rollback triggers for CloudFormation to monitor during stack creation and
	// updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Key-value pairs to associate with this stack. CloudFormation also propagates
	// these tags to resources in the stack. You can specify a maximum of 50 tags.
	Tags []types.Tag

	// A structure that contains the body of the revised template, with a minimum
	// length of 1 byte and a maximum length of 51,200 bytes. CloudFormation generates
	// the change set by comparing this template with the template of the stack that
	// you specified.
	//
	// Conditional: You must specify only TemplateBody or TemplateURL .
	TemplateBody *string

	// The location of the file that contains the revised template. The URL must point
	// to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket or
	// a Systems Manager document. CloudFormation generates the change set by comparing
	// this template with the stack that you specified. The location for an Amazon S3
	// bucket must start with https:// .
	//
	// Conditional: You must specify only TemplateBody or TemplateURL .
	TemplateURL *string

	// Whether to reuse the template that's associated with the stack to create the
	// change set.
	UsePreviousTemplate *bool

	noSmithyDocumentSerde
}

// The output for the CreateChangeSet action.
type CreateChangeSetOutput struct {

	// The Amazon Resource Name (ARN) of the change set.
	Id *string

	// The unique ID of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChangeSet"); err != nil {
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
	if err = addOpCreateChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChangeSet",
	}
}
