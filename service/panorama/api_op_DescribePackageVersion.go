// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a package version.
func (c *Client) DescribePackageVersion(ctx context.Context, params *DescribePackageVersionInput, optFns ...func(*Options)) (*DescribePackageVersionOutput, error) {
	if params == nil {
		params = &DescribePackageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackageVersion", params, optFns, c.addOperationDescribePackageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackageVersionInput struct {

	// The version's ID.
	//
	// This member is required.
	PackageId *string

	// The version's version.
	//
	// This member is required.
	PackageVersion *string

	// The version's owner account.
	OwnerAccount *string

	// The version's patch version.
	PatchVersion *string

	noSmithyDocumentSerde
}

type DescribePackageVersionOutput struct {

	// Whether the version is the latest available.
	//
	// This member is required.
	IsLatestPatch bool

	// The version's ID.
	//
	// This member is required.
	PackageId *string

	// The version's name.
	//
	// This member is required.
	PackageName *string

	// The version's version.
	//
	// This member is required.
	PackageVersion *string

	// The version's patch version.
	//
	// This member is required.
	PatchVersion *string

	// The version's status.
	//
	// This member is required.
	Status types.PackageVersionStatus

	// The account ID of the version's owner.
	OwnerAccount *string

	// The ARN of the package.
	PackageArn *string

	// The version's registered time.
	RegisteredTime *time.Time

	// The version's status description.
	StatusDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePackageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePackageVersion"); err != nil {
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
	if err = addOpDescribePackageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackageVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePackageVersion",
	}
}
