// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the direct dependencies for a package version. The dependencies are
// returned as PackageDependency
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html)
// objects. CodeArtifact extracts the dependencies for a package version from the
// metadata file for the package format (for example, the package.json file for npm
// packages and the pom.xml file for Maven). Any package version dependencies that
// are not listed in the configuration file are not returned.
func (c *Client) ListPackageVersionDependencies(ctx context.Context, params *ListPackageVersionDependenciesInput, optFns ...func(*Options)) (*ListPackageVersionDependenciesOutput, error) {
	stack := middleware.NewStack("ListPackageVersionDependencies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPackageVersionDependenciesMiddlewares(stack)
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
	addOpListPackageVersionDependenciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersionDependencies(options.Region), middleware.Before)
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
			OperationName: "ListPackageVersionDependencies",
			Err:           err,
		}
	}
	out := result.(*ListPackageVersionDependenciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionDependenciesInput struct {

	// The domain that contains the repository that contains the requested package
	// version dependencies.
	//
	// This member is required.
	Domain *string

	// The name of the package versions' package.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2).
	//
	// This member is required.
	PackageVersion *string

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The format of the package with the requested dependencies. The valid package
	// types are:
	//
	//     * npm: A Node Package Manager (npm) package.
	//
	//     * pypi: A
	// Python Package Index (PyPI) package.
	//
	//     * maven: A Maven package that contains
	// compiled code in a distributable format, such as a JAR file.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the repository that contains the requested package version.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
}

type ListPackageVersionDependenciesOutput struct {

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string

	// A format that specifies the type of the package that contains the returned
	// dependencies. The valid values are:
	//
	//     * npm
	//
	//     * pypi
	//
	//     * maven
	Format types.PackageFormat

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The returned list of PackageDependency
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html)
	// objects.
	Dependencies []*types.PackageDependency

	// The name of the package that contains the returned package versions
	// dependencies.
	Package *string

	// The version of the package that is specified in the request.
	Version *string

	// The current revision associated with the package version.
	VersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPackageVersionDependenciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersionDependencies{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersionDependencies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPackageVersionDependencies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackageVersionDependencies",
	}
}
