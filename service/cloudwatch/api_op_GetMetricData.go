// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// You can use the GetMetricData API to retrieve as many as 500 different metrics
// in a single request, with a total of as many as 100,800 data points. You can
// also optionally perform math expressions on the values of the returned
// statistics, to create new time series that represent new insights into your
// data. For example, using Lambda metrics, you could divide the Errors metric by
// the Invocations metric to get an error rate time series. For more information
// about metric math expressions, see Metric Math Syntax and Functions
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax)
// in the Amazon CloudWatch User Guide.  <p>Calls to the <code>GetMetricData</code>
// API have a different pricing structure than calls to
// <code>GetMetricStatistics</code>. For more information about pricing, see <a
// href="https://aws.amazon.com/cloudwatch/pricing/">Amazon CloudWatch
// Pricing</a>.</p> <p>Amazon CloudWatch retains metric data as follows:</p> <ul>
// <li> <p>Data points with a period of less than 60 seconds are available for 3
// hours. These data points are high-resolution metrics and are available only for
// custom metrics that have been defined with a <code>StorageResolution</code> of
// 1.</p> </li> <li> <p>Data points with a period of 60 seconds (1-minute) are
// available for 15 days.</p> </li> <li> <p>Data points with a period of 300
// seconds (5-minute) are available for 63 days.</p> </li> <li> <p>Data points with
// a period of 3600 seconds (1 hour) are available for 455 days (15 months).</p>
// </li> </ul> <p>Data points that are initially published with a shorter period
// are aggregated together for long-term storage. For example, if you collect data
// using a period of 1 minute, the data remains available for 15 days with 1-minute
// resolution. After 15 days, this data is still available, but is aggregated and
// retrievable only with a resolution of 5 minutes. After 63 days, the data is
// further aggregated and is available with a resolution of 1 hour.</p> <p>If you
// omit <code>Unit</code> in your request, all data that was collected with any
// unit is returned, along with the corresponding units that were specified when
// the data was reported to CloudWatch. If you specify a unit, the operation
// returns only data that was collected with that unit specified. If you specify a
// unit that does not match the data collected, the results of the operation are
// null. CloudWatch does not perform unit conversions.</p>
func (c *Client) GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
	stack := middleware.NewStack("GetMetricData", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetMetricDataMiddlewares(stack)
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
	addOpGetMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricData(options.Region), middleware.Before)
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
			OperationName: "GetMetricData",
			Err:           err,
		}
	}
	out := result.(*GetMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricDataInput struct {

	// The metric queries to be returned. A single GetMetricData call can include as
	// many as 500 MetricDataQuery structures. Each of these structures can specify
	// either a metric to retrieve, or a math expression to perform on retrieved data.
	//
	// This member is required.
	MetricDataQueries []*types.MetricDataQuery

	// The time stamp indicating the earliest data to be returned. The value specified
	// is inclusive; results include data points with the specified time stamp.
	// CloudWatch rounds the specified time stamp as follows:
	//
	//     * Start time less
	// than 15 days ago - Round down to the nearest whole minute. For example, 12:32:34
	// is rounded down to 12:32:00.
	//
	//     * Start time between 15 and 63 days ago -
	// Round down to the nearest 5-minute clock interval. For example, 12:32:34 is
	// rounded down to 12:30:00.
	//
	//     * Start time greater than 63 days ago - Round
	// down to the nearest 1-hour clock interval. For example, 12:32:34 is rounded down
	// to 12:00:00.
	//
	// If you set Period to 5, 10, or 30, the start time of your request
	// is rounded down to the nearest time that corresponds to even 5-, 10-, or
	// 30-second divisions of a minute. For example, if you make a query at (HH:mm:ss)
	// 01:05:23 for the previous 10-second period, the start time of your request is
	// rounded down and you receive data from 01:05:10 to 01:05:20. If you make a query
	// at 15:07:17 for the previous 5 minutes of data, using a period of 5 seconds, you
	// receive data timestamped between 15:02:15 and 15:07:15. For better performance,
	// specify StartTime and EndTime values that align with the value of the metric's
	// Period and sync up with the beginning and end of an hour. For example, if the
	// Period of a metric is 5 minutes, specifying 12:05 or 12:30 as StartTime can get
	// a faster response from CloudWatch than setting 12:07 or 12:29 as the StartTime.
	//
	// This member is required.
	StartTime *time.Time

	// Include this value, if it was returned by the previous call, to get the next set
	// of data points.
	NextToken *string

	// The maximum number of data points the request should return before paginating.
	// If you omit this, the default of 100,800 is used.
	MaxDatapoints *int32

	// The order in which data points should be returned. TimestampDescending returns
	// the newest data first and paginates when the MaxDatapoints limit is reached.
	// TimestampAscending returns the oldest data first and paginates when the
	// MaxDatapoints limit is reached.
	ScanBy types.ScanBy

	// The time stamp indicating the latest data to be returned. The value specified is
	// exclusive; results include data points up to the specified time stamp. For
	// better performance, specify StartTime and EndTime values that align with the
	// value of the metric's Period and sync up with the beginning and end of an hour.
	// For example, if the Period of a metric is 5 minutes, specifying 12:05 or 12:30
	// as EndTime can get a faster response from CloudWatch than setting 12:07 or 12:29
	// as the EndTime.
	//
	// This member is required.
	EndTime *time.Time
}

type GetMetricDataOutput struct {

	// Contains a message about this GetMetricData operation, if the operation results
	// in such a message. An example of a message that might be returned is Maximum
	// number of allowed metrics exceeded. If there is a message, as much of the
	// operation as possible is still executed. A message appears here only if it is
	// related to the global GetMetricData operation. Any message about a specific
	// metric returned by the operation appears in the MetricDataResult object returned
	// for that metric.
	Messages []*types.MessageData

	// The metrics that are returned, including the metric name, namespace, and
	// dimensions.
	MetricDataResults []*types.MetricDataResult

	// A token that marks the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetMetricDataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricData{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricData{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetMetricData",
	}
}
