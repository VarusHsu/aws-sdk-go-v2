// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Allows a caller in a secure environment to assume a role with permission to
// access Amazon S3. In order to vend such credentials, Lake Formation assumes the
// role associated with a registered location, for example an Amazon S3 bucket,
// with a scope down policy which restricts the access to a single prefix.
func (c *Client) GetTemporaryGlueTableCredentials(ctx context.Context, params *GetTemporaryGlueTableCredentialsInput, optFns ...func(*Options)) (*GetTemporaryGlueTableCredentialsOutput, error) {
	if params == nil {
		params = &GetTemporaryGlueTableCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTemporaryGlueTableCredentials", params, optFns, c.addOperationGetTemporaryGlueTableCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTemporaryGlueTableCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTemporaryGlueTableCredentialsInput struct {

	// The ARN identifying a table in the Data Catalog for the temporary credentials
	// request.
	//
	// This member is required.
	TableArn *string

	// A structure representing context to access a resource (column names, query ID,
	// etc).
	AuditContext *types.AuditContext

	// The time period, between 900 and 21,600 seconds, for the timeout of the
	// temporary credentials.
	DurationSeconds *int32

	// Filters the request based on the user having been granted a list of specified
	// permissions on the requested resource(s).
	Permissions []types.Permission

	// A list of supported permission types for the table. Valid values are
	// COLUMN_PERMISSION and CELL_FILTER_PERMISSION .
	SupportedPermissionTypes []types.PermissionType

	noSmithyDocumentSerde
}

type GetTemporaryGlueTableCredentialsOutput struct {

	// The access key ID for the temporary credentials.
	AccessKeyId *string

	// The date and time when the temporary credentials expire.
	Expiration *time.Time

	// The secret key for the temporary credentials.
	SecretAccessKey *string

	// The session token for the temporary credentials.
	SessionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTemporaryGlueTableCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTemporaryGlueTableCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTemporaryGlueTableCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTemporaryGlueTableCredentials"); err != nil {
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
	if err = addOpGetTemporaryGlueTableCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemporaryGlueTableCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTemporaryGlueTableCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTemporaryGlueTableCredentials",
	}
}
