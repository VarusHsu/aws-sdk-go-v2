// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a dashboard from a template. To first create a template, see the
// CreateTemplate () API operation. A dashboard is an entity in QuickSight that
// identifies QuickSight reports, created from analyses. You can share QuickSight
// dashboards. With the right permissions, you can create scheduled email reports
// from them. The CreateDashboard, DescribeDashboard, and ListDashboardsByUser API
// operations act on the dashboard entity. If you have the correct permissions, you
// can create a dashboard from a template that exists in a different AWS account.
func (c *Client) CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) {
	stack := middleware.NewStack("CreateDashboard", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDashboardMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateDashboardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDashboard(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateDashboard",
			Err:           err,
		}
	}
	out := result.(*CreateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDashboardInput struct {

	// A structure that contains the permissions of the dashboard. You can use this
	// structure for granting permissions with principal and action information.
	Permissions []*types.ResourcePermission

	// The ID for the dashboard, also added to the IAM policy.
	//
	// This member is required.
	DashboardId *string

	// The display name of the dashboard.
	//
	// This member is required.
	Name *string

	// A description for the first version of the dashboard being created.
	VersionDescription *string

	// The ID of the AWS account where you want to create the dashboard.
	//
	// This member is required.
	AwsAccountId *string

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the dashboard.
	Tags []*types.Tag

	// The Amazon Resource Name (ARN) of the theme that is being used for this
	// dashboard. If you add a value for this field, it overrides the value that is
	// used in the source entity. The theme ARN must exist in the same AWS account
	// where you create the dashboard.
	ThemeArn *string

	// The parameters for the creation of the dashboard, which you want to use to
	// override the default settings. A dashboard can have any type of parameters, and
	// some parameters might accept multiple values.
	Parameters *types.Parameters

	// The entity that you are using as a source when you create the dashboard. In
	// SourceEntity, you specify the type of object you're using as source. You can
	// only create a dashboard from a template, so you use a SourceTemplate entity. If
	// you need to create a dashboard from an analysis, first convert the analysis to a
	// template by using the CreateTemplate () API operation. For SourceTemplate,
	// specify the Amazon Resource Name (ARN) of the source template. The
	// SourceTemplateARN can contain any AWS Account and any QuickSight-supported AWS
	// Region. Use the DataSetReferences entity within SourceTemplate to list the
	// replacement datasets for the placeholders listed in the original. The schema in
	// each dataset must match its placeholder.
	//
	// This member is required.
	SourceEntity *types.DashboardSourceEntity

	// Options for publishing the dashboard when you create it:
	//
	//     *
	// AvailabilityStatus for AdHocFilteringOption - This status can be either ENABLED
	// or DISABLED. When this is set to DISABLED, QuickSight disables the left filter
	// pane on the published dashboard, which can be used for ad hoc (one-time)
	// filtering. This option is ENABLED by default.
	//
	//     * AvailabilityStatus for
	// ExportToCSVOption - This status can be either ENABLED or DISABLED. The visual
	// option to export data to .csv format isn't enabled when this is set to DISABLED.
	// This option is ENABLED by default.
	//
	//     * VisibilityState for
	// SheetControlsOption - This visibility state can be either COLLAPSED or EXPANDED.
	// This option is COLLAPSED by default.
	DashboardPublishOptions *types.DashboardPublishOptions
}

type CreateDashboardOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// The ID for the dashboard.
	DashboardId *string

	// The status of the dashboard creation request.
	CreationStatus types.ResourceStatus

	// The ARN of the dashboard, including the version number of the first version that
	// is created.
	VersionArn *string

	// The ARN of the dashboard.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDashboardMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDashboard{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDashboard{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDashboard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateDashboard",
	}
}
