// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Inspects the clinical text for protected health information (PHI) entities and
// returns the entity category, location, and confidence score for each entity.
// Amazon Comprehend Medical only detects entities in English language texts.
func (c *Client) DetectPHI(ctx context.Context, params *DetectPHIInput, optFns ...func(*Options)) (*DetectPHIOutput, error) {
	stack := middleware.NewStack("DetectPHI", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectPHIMiddlewares(stack)
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
	addOpDetectPHIValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectPHI(options.Region), middleware.Before)
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
			OperationName: "DetectPHI",
			Err:           err,
		}
	}
	out := result.(*DetectPHIOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectPHIInput struct {

	// A UTF-8 text string containing the clinical content being examined for PHI
	// entities. Each string must contain fewer than 20,000 bytes of characters.
	//
	// This member is required.
	Text *string
}

type DetectPHIOutput struct {

	// The collection of PHI entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity text,
	// the entity category, where the entity text begins and ends, and the level of
	// confidence that Amazon Comprehend Medical has in its detection.
	//
	// This member is required.
	Entities []*types.Entity

	// If the result of the previous request to DetectPHI was truncated, include the
	// PaginationToken to fetch the next page of PHI entities.
	PaginationToken *string

	// The version of the model used to analyze the documents. The version number looks
	// like X.X.X. You can use this information to track the model used for a
	// particular batch of documents.
	//
	// This member is required.
	ModelVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectPHIMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectPHI{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectPHI{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectPHI(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DetectPHI",
	}
}
