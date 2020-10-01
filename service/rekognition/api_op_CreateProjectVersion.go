// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new version of a model and begins training. Models are managed as part
// of an Amazon Rekognition Custom Labels project. You can specify one training
// dataset and one testing dataset. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the version of the model. Training takes a while
// to complete. You can get the current status by calling DescribeProjectVersions
// (). Once training has successfully completed, call DescribeProjectVersions () to
// get the training results and evaluate the model. After evaluating the model, you
// start the model by calling StartProjectVersion (). This operation requires
// permissions to perform the rekognition:CreateProjectVersion action.
func (c *Client) CreateProjectVersion(ctx context.Context, params *CreateProjectVersionInput, optFns ...func(*Options)) (*CreateProjectVersionOutput, error) {
	stack := middleware.NewStack("CreateProjectVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProjectVersionMiddlewares(stack)
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
	addOpCreateProjectVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProjectVersion(options.Region), middleware.Before)
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
			OperationName: "CreateProjectVersion",
			Err:           err,
		}
	}
	out := result.(*CreateProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectVersionInput struct {

	// The Amazon S3 location to store the results of training.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// A name for the version of the model. This value must be unique.
	//
	// This member is required.
	VersionName *string

	// The ARN of the Amazon Rekognition Custom Labels project that manages the model
	// that you want to train.
	//
	// This member is required.
	ProjectArn *string

	// The dataset to use for training.
	//
	// This member is required.
	TrainingData *types.TrainingData

	// The dataset to use for testing.
	//
	// This member is required.
	TestingData *types.TestingData
}

type CreateProjectVersionOutput struct {

	// The ARN of the model version that was created. Use DescribeProjectVersion to get
	// the current status of the training operation.
	ProjectVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProjectVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProjectVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProjectVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProjectVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateProjectVersion",
	}
}
