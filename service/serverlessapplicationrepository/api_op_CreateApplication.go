// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an application, optionally including an AWS SAM file to create the first
// application version in the same call.
func (c *Client) CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) {
	stack := middleware.NewStack("CreateApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateApplicationMiddlewares(stack)
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
	addOpCreateApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplication(options.Region), middleware.Before)
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
			OperationName: "CreateApplication",
			Err:           err,
		}
	}
	out := result.(*CreateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInput struct {

	// The description of the application.Minimum length=1. Maximum length=256
	//
	// This member is required.
	Description *string

	// A valid identifier from https://spdx.org/licenses/ (https://spdx.org/licenses/).
	SpdxLicenseId *string

	// The local raw packaged AWS SAM template file of your application. The file has
	// the format file://<path>/<filename>.You can specify only one of templateBody and
	// templateUrl; otherwise an error results.
	TemplateBody *string

	// The name of the author publishing the app.Minimum length=1. Maximum
	// length=127.Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	//
	// This member is required.
	Author *string

	// A link to the S3 object containing the packaged AWS SAM template of your
	// application.You can specify only one of templateBody and templateUrl; otherwise
	// an error results.
	TemplateUrl *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// A link to the S3 object in Markdown language that contains a more detailed
	// description of the application and how it works.Maximum size 5 MBYou can specify
	// only one of readmeBody and readmeUrl; otherwise, an error results.
	ReadmeUrl *string

	// A link to the S3 object that contains the license of the app that matches the
	// spdxLicenseID value of your application.Maximum size 5 MBYou can specify only
	// one of licenseBody and licenseUrl; otherwise, an error results.
	LicenseUrl *string

	// Labels to improve discovery of apps in search results.Minimum length=1. Maximum
	// length=127. Maximum number of labels: 10Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []*string

	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string

	// The name of the application that you want to publish.Minimum length=1. Maximum
	// length=140Pattern: "[a-zA-Z0-9\\-]+";
	//
	// This member is required.
	Name *string

	// A local text file that contains the license of the app that matches the
	// spdxLicenseID value of your application. The file has the format
	// file://<path>/<filename>.Maximum size 5 MBYou can specify only one of
	// licenseBody and licenseUrl; otherwise, an error results.
	LicenseBody *string

	// A local text readme file in Markdown language that contains a more detailed
	// description of the application and how it works. The file has the format
	// file://<path>/<filename>.Maximum size 5 MBYou can specify only one of readmeBody
	// and readmeUrl; otherwise, an error results.
	ReadmeBody *string

	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string

	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.Maximum size 50 MB
	SourceCodeArchiveUrl *string
}

type CreateApplicationOutput struct {

	// The description of the application.Minimum length=1. Maximum length=256
	Description *string

	// The name of the author publishing the app.Minimum length=1. Maximum
	// length=127.Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	Author *string

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The date and time this resource was created.
	CreationTime *string

	// Version information about the application.
	Version *types.Version

	// A link to a license file of the app that matches the spdxLicenseID value of your
	// application.Maximum size 5 MB
	LicenseUrl *string

	// A link to the readme file in Markdown language that contains a more detailed
	// description of the application and how it works.Maximum size 5 MB
	ReadmeUrl *string

	// The URL to the public profile of a verified author. This URL is submitted by the
	// author.
	VerifiedAuthorUrl *string

	// Whether the author of this application has been verified. This means means that
	// AWS has made a good faith review, as a reasonable and prudent service provider,
	// of the information provided by the requester and has confirmed that the
	// requester's identity is as claimed.
	IsVerifiedAuthor *bool

	// The name of the application.Minimum length=1. Maximum length=140Pattern:
	// "[a-zA-Z0-9\\-]+";
	Name *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// Labels to improve discovery of apps in search results.Minimum length=1. Maximum
	// length=127. Maximum number of labels: 10Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []*string

	// A valid identifier from https://spdx.org/licenses/.
	SpdxLicenseId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "CreateApplication",
	}
}
