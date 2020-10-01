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

// Updates the specified application.
func (c *Client) UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) {
	stack := middleware.NewStack("UpdateApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateApplicationMiddlewares(stack)
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
	addOpUpdateApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplication(options.Region), middleware.Before)
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
			OperationName: "UpdateApplication",
			Err:           err,
		}
	}
	out := result.(*UpdateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The name of the author publishing the app.Minimum length=1. Maximum
	// length=127.Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	Author *string

	// The description of the application.Minimum length=1. Maximum length=256
	Description *string

	// A link to the readme file in Markdown language that contains a more detailed
	// description of the application and how it works.Maximum size 5 MB
	ReadmeUrl *string

	// A text readme file in Markdown language that contains a more detailed
	// description of the application and how it works.Maximum size 5 MB
	ReadmeBody *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// Labels to improve discovery of apps in search results.Minimum length=1. Maximum
	// length=127. Maximum number of labels: 10Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []*string
}

type UpdateApplicationOutput struct {

	// The date and time this resource was created.
	CreationTime *string

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The name of the author publishing the app.Minimum length=1. Maximum
	// length=127.Pattern "^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$";
	Author *string

	// The URL to the public profile of a verified author. This URL is submitted by the
	// author.
	VerifiedAuthorUrl *string

	// Version information about the application.
	Version *types.Version

	// The name of the application.Minimum length=1. Maximum length=140Pattern:
	// "[a-zA-Z0-9\\-]+";
	Name *string

	// A link to a license file of the app that matches the spdxLicenseID value of your
	// application.Maximum size 5 MB
	LicenseUrl *string

	// Whether the author of this application has been verified. This means means that
	// AWS has made a good faith review, as a reasonable and prudent service provider,
	// of the information provided by the requester and has confirmed that the
	// requester's identity is as claimed.
	IsVerifiedAuthor *bool

	// A link to the readme file in Markdown language that contains a more detailed
	// description of the application and how it works.Maximum size 5 MB
	ReadmeUrl *string

	// The description of the application.Minimum length=1. Maximum length=256
	Description *string

	// A valid identifier from https://spdx.org/licenses/.
	SpdxLicenseId *string

	// A URL with more information about the application, for example the location of
	// your GitHub repository for the application.
	HomePageUrl *string

	// Labels to improve discovery of apps in search results.Minimum length=1. Maximum
	// length=127. Maximum number of labels: 10Pattern: "^[a-zA-Z0-9+\\-_:\\/@]+$";
	Labels []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "UpdateApplication",
	}
}
