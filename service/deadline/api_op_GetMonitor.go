// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about the specified monitor.
func (c *Client) GetMonitor(ctx context.Context, params *GetMonitorInput, optFns ...func(*Options)) (*GetMonitorOutput, error) {
	if params == nil {
		params = &GetMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMonitor", params, optFns, c.addOperationGetMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMonitorInput struct {

	// The unique identifier for the monitor. This ID is returned by the CreateMonitor
	// operation.
	//
	// This member is required.
	MonitorId *string

	noSmithyDocumentSerde
}

type GetMonitorOutput struct {

	// The UNIX timestamp of the date and time that the monitor was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user name of the person that created the monitor.
	//
	// This member is required.
	CreatedBy *string

	// The name used to identify the monitor on the Deadline Cloud console.
	//
	// This member is required.
	DisplayName *string

	// The Amazon Resource Name (ARN) that the IAM Identity Center assigned to the
	// monitor when it was created.
	//
	// This member is required.
	IdentityCenterApplicationArn *string

	// The Amazon Resource Name (ARN) of the IAM Identity Center instance responsible
	// for authenticating monitor users.
	//
	// This member is required.
	IdentityCenterInstanceArn *string

	// The unique identifier for the monitor.
	//
	// This member is required.
	MonitorId *string

	// The Amazon Resource Name (ARN) of the IAM role for the monitor. Users of the
	// monitor use this role to access Deadline Cloud resources.
	//
	// This member is required.
	RoleArn *string

	// The subdomain used for the monitor URL. The full URL of the monitor is
	// subdomain.Region.deadlinecloud.amazonaws.com.
	//
	// This member is required.
	Subdomain *string

	// The complete URL of the monitor. The full URL of the monitor is
	// subdomain.Region.deadlinecloud.amazonaws.com.
	//
	// This member is required.
	Url *string

	// The UNIX timestamp of the last date and time that the monitor was updated.
	UpdatedAt *time.Time

	// The user name of the person that last updated the monitor.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMonitor"); err != nil {
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
	if err = addEndpointPrefix_opGetMonitorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMonitor(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMonitorMiddleware struct {
}

func (*endpointPrefix_opGetMonitorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMonitorMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetMonitorMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetMonitorMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMonitor",
	}
}
