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

// Gets face detection results for a Amazon Rekognition Video analysis started by
// StartFaceDetection (). Face detection with Amazon Rekognition Video is an
// asynchronous operation. You start face detection by calling StartFaceDetection
// () which returns a job identifier (JobId). When the face detection operation
// finishes, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to
// StartFaceDetection. To get the results of the face detection operation, first
// check that the status value published to the Amazon SNS topic is SUCCEEDED. If
// so, call GetFaceDetection () and pass the job identifier (JobId) from the
// initial call to StartFaceDetection. GetFaceDetection returns an array of
// detected faces (Faces) sorted by the time the faces were detected. Use
// MaxResults parameter to limit the number of labels returned. If there are more
// results than specified in MaxResults, the value of NextToken in the operation
// response contains a pagination token for getting the next set of results. To get
// the next page of results, call GetFaceDetection and populate the NextToken
// request parameter with the token value returned from the previous call to
// GetFaceDetection.
func (c *Client) GetFaceDetection(ctx context.Context, params *GetFaceDetectionInput, optFns ...func(*Options)) (*GetFaceDetectionOutput, error) {
	stack := middleware.NewStack("GetFaceDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFaceDetectionMiddlewares(stack)
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
	addOpGetFaceDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFaceDetection(options.Region), middleware.Before)
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
			OperationName: "GetFaceDetection",
			Err:           err,
		}
	}
	out := result.(*GetFaceDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFaceDetectionInput struct {

	// If the previous response was incomplete (because there are more faces to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of faces.
	NextToken *string

	// Unique identifier for the face detection job. The JobId is returned from
	// StartFaceDetection.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32
}

type GetFaceDetectionOutput struct {

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// An array of faces detected in the video. Each element contains a detected face's
	// details and the time, in milliseconds from the start of the video, the face was
	// detected.
	Faces []*types.FaceDetection

	// The current status of the face detection job.
	JobStatus types.VideoJobStatus

	// If the response is truncated, Amazon Rekognition returns this token that you can
	// use in the subsequent request to retrieve the next set of faces.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFaceDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFaceDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFaceDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFaceDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetFaceDetection",
	}
}
