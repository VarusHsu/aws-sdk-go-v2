// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the license configurations for your account.
func (c *Client) ListLicenseConfigurations(ctx context.Context, params *ListLicenseConfigurationsInput, optFns ...func(*Options)) (*ListLicenseConfigurationsOutput, error) {
	stack := middleware.NewStack("ListLicenseConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListLicenseConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLicenseConfigurations(options.Region), middleware.Before)
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
			OperationName: "ListLicenseConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListLicenseConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLicenseConfigurationsInput struct {

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Filters to scope the results. The following filters and logical operators are
	// supported:
	//
	//     * licenseCountingType - The dimension on which licenses are
	// counted (vCPU). Logical operators are EQUALS | NOT_EQUALS.
	//
	//     *
	// enforceLicenseCount - A Boolean value that indicates whether hard license
	// enforcement is used. Logical operators are EQUALS | NOT_EQUALS.
	//
	//     *
	// usagelimitExceeded - A Boolean value that indicates whether the available
	// licenses have been exceeded. Logical operators are EQUALS | NOT_EQUALS.
	Filters []*types.Filter

	// Amazon Resource Names (ARN) of the license configurations.
	LicenseConfigurationArns []*string

	// Token for the next set of results.
	NextToken *string
}

type ListLicenseConfigurationsOutput struct {

	// Token for the next set of results.
	NextToken *string

	// Information about the license configurations.
	LicenseConfigurations []*types.LicenseConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListLicenseConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListLicenseConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLicenseConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLicenseConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListLicenseConfigurations",
	}
}
