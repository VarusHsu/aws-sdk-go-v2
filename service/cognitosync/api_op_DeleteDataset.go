// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specific dataset. The dataset will be deleted permanently, and the
// action can't be undone. Datasets that this dataset was merged with will no
// longer report the merge. Any subsequent operation on this dataset will result in
// a ResourceNotFoundException. This API can be called with temporary user
// credentials provided by Cognito Identity or with developer credentials.
func (c *Client) DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) {
	stack := middleware.NewStack("DeleteDataset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDatasetMiddlewares(stack)
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
	addOpDeleteDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataset(options.Region), middleware.Before)
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
			OperationName: "DeleteDataset",
			Err:           err,
		}
	}
	out := result.(*DeleteDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete the specific dataset.
type DeleteDatasetInput struct {

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string
}

// Response to a successful DeleteDataset request.
type DeleteDatasetOutput struct {

	// A collection of data for an identity pool. An identity pool can have multiple
	// datasets. A dataset is per identity and can be general or associated with a
	// particular entity in an application (like a saved game). Datasets are
	// automatically created if they don't exist. Data is synced by dataset, and a
	// dataset can hold up to 1MB of key-value pairs.
	Dataset *types.Dataset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDatasetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataset{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "DeleteDataset",
	}
}
