// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a DataSource from a database hosted on an Amazon Redshift cluster. A
// DataSource references data that can be used to perform either CreateMLModel,
// CreateEvaluation, or CreateBatchPrediction operations.  <p>
// <code>CreateDataSourceFromRedshift</code> is an asynchronous operation. In
// response to <code>CreateDataSourceFromRedshift</code>, Amazon Machine Learning
// (Amazon ML) immediately returns and sets the <code>DataSource</code> status to
// <code>PENDING</code>. After the <code>DataSource</code> is created and ready for
// use, Amazon ML sets the <code>Status</code> parameter to <code>COMPLETED</code>.
// <code>DataSource</code> in <code>COMPLETED</code> or <code>PENDING</code> states
// can be used to perform only <code>CreateMLModel</code>,
// <code>CreateEvaluation</code>, or <code>CreateBatchPrediction</code> operations.
// </p> <p> If Amazon ML can't accept the input source, it sets the
// <code>Status</code> parameter to <code>FAILED</code> and includes an error
// message in the <code>Message</code> attribute of the <code>GetDataSource</code>
// operation response. </p> <p>The observations should be contained in the database
// hosted on an Amazon Redshift cluster and should be specified by a
// <code>SelectSqlQuery</code> query. Amazon ML executes an <code>Unload</code>
// command in Amazon Redshift to transfer the result set of the
// <code>SelectSqlQuery</code> query to <code>S3StagingLocation</code>.</p>
// <p>After the <code>DataSource</code> has been created, it's ready for use in
// evaluations and batch predictions. If you plan to use the
// <code>DataSource</code> to train an <code>MLModel</code>, the
// <code>DataSource</code> also requires a recipe. A recipe describes how each
// input variable will be used in training an <code>MLModel</code>. Will the
// variable be included or excluded from training? Will the variable be
// manipulated; for example, will it be combined with another variable or will it
// be split apart into word combinations? The recipe provides answers to these
// questions.</p> <p>You can't change an existing datasource, but you can copy and
// modify the settings from an existing Amazon Redshift datasource to create a new
// datasource. To do so, call <code>GetDataSource</code> for an existing datasource
// and copy the values to a <code>CreateDataSource</code> call. Change the settings
// that you want to change and make sure that all required fields have the
// appropriate values.</p>
func (c *Client) CreateDataSourceFromRedshift(ctx context.Context, params *CreateDataSourceFromRedshiftInput, optFns ...func(*Options)) (*CreateDataSourceFromRedshiftOutput, error) {
	stack := middleware.NewStack("CreateDataSourceFromRedshift", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDataSourceFromRedshiftMiddlewares(stack)
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
	addOpCreateDataSourceFromRedshiftValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSourceFromRedshift(options.Region), middleware.Before)
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
			OperationName: "CreateDataSourceFromRedshift",
			Err:           err,
		}
	}
	out := result.(*CreateDataSourceFromRedshiftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceFromRedshiftInput struct {

	// A user-supplied ID that uniquely identifies the DataSource.
	//
	// This member is required.
	DataSourceId *string

	// The compute statistics for a DataSource. The statistics are generated from the
	// observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if the
	// DataSource needs to be used for MLModel training.
	ComputeStatistics *bool

	// A user-supplied name or description of the DataSource.
	DataSourceName *string

	// The data specification of an Amazon Redshift DataSource:
	//
	//     *
	// DatabaseInformation -
	//
	//         * DatabaseName - The name of the Amazon Redshift
	// database.
	//
	//         * ClusterIdentifier - The unique ID for the Amazon Redshift
	// cluster.
	//
	//     * DatabaseCredentials - The AWS Identity and Access Management
	// (IAM) credentials that are used to connect to the Amazon Redshift database.
	//
	//
	// * SelectSqlQuery - The query that is used to retrieve the observation data for
	// the Datasource.
	//
	//     * S3StagingLocation - The Amazon Simple Storage Service
	// (Amazon S3) location for staging Amazon Redshift data. The data retrieved from
	// Amazon Redshift using the SelectSqlQuery query is stored in this location.
	//
	//
	// * DataSchemaUri - The Amazon S3 location of the DataSchema.
	//
	//     * DataSchema -
	// A JSON string representing the schema. This is not required if DataSchemaUri is
	// specified.
	//
	//     * DataRearrangement - A JSON string that represents the
	// splitting and rearrangement requirements for the DataSource. Sample -
	// "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// This member is required.
	DataSpec *types.RedshiftDataSpec

	// A fully specified role Amazon Resource Name (ARN). Amazon ML assumes the role on
	// behalf of the user to create the following:  <ul> <li> <p>A security group to
	// allow Amazon ML to execute the <code>SelectSqlQuery</code> query on an Amazon
	// Redshift cluster</p> </li> <li> <p>An Amazon S3 bucket policy to grant Amazon ML
	// read/write permissions on the <code>S3StagingLocation</code> </p> </li> </ul>
	//
	// This member is required.
	RoleARN *string
}

// Represents the output of a CreateDataSourceFromRedshift operation, and is an
// acknowledgement that Amazon ML received the request. The
// CreateDataSourceFromRedshift operation is asynchronous. You can poll for updates
// by using the GetBatchPrediction operation and checking the Status parameter.
type CreateDataSourceFromRedshiftOutput struct {

	// A user-supplied ID that uniquely identifies the datasource. This value should be
	// identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDataSourceFromRedshiftMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataSourceFromRedshift{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataSourceFromRedshift{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSourceFromRedshift(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "CreateDataSourceFromRedshift",
	}
}
