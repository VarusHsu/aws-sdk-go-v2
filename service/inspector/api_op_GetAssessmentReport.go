// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Produces an assessment report that includes detailed and comprehensive results
// of a specified assessment run.
func (c *Client) GetAssessmentReport(ctx context.Context, params *GetAssessmentReportInput, optFns ...func(*Options)) (*GetAssessmentReportOutput, error) {
	if params == nil {
		params = &GetAssessmentReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssessmentReport", params, optFns, c.addOperationGetAssessmentReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssessmentReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssessmentReportInput struct {

	// The ARN that specifies the assessment run for which you want to generate a
	// report.
	//
	// This member is required.
	AssessmentRunArn *string

	// Specifies the file format (html or pdf) of the assessment report that you want
	// to generate.
	//
	// This member is required.
	ReportFileFormat types.ReportFileFormat

	// Specifies the type of the assessment report that you want to generate. There
	// are two types of assessment reports: a finding report and a full report. For
	// more information, see [Assessment Reports].
	//
	// [Assessment Reports]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_reports.html
	//
	// This member is required.
	ReportType types.ReportType

	noSmithyDocumentSerde
}

type GetAssessmentReportOutput struct {

	// Specifies the status of the request to generate an assessment report.
	//
	// This member is required.
	Status types.ReportStatus

	// Specifies the URL where you can find the generated assessment report. This
	// parameter is only returned if the report is successfully generated.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssessmentReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAssessmentReport"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetAssessmentReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssessmentReport(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetAssessmentReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAssessmentReport",
	}
}
