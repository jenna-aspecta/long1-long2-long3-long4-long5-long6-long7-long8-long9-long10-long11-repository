// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Returns the access control list (ACL) of an object. To use this operation, you
// must have s3:GetObjectAcl permissions or READ_ACP access to the object. For
// more information, see [Mapping of ACL permissions and access policy permissions]in the Amazon S3 User Guide
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// By default, GET returns ACL information about the current version of an object.
// To return ACL information about a different version, use the versionId
// subresource.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// requests to read ACLs are still supported and return the
// bucket-owner-full-control ACL with the owner being the account that created the
// bucket. For more information, see [Controlling object ownership and disabling ACLs]in the Amazon S3 User Guide.
//
// The following operations are related to GetObjectAcl :
//
// [GetObject]
//
// [GetObjectAttributes]
//
// [DeleteObject]
//
// [PutObject]
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Mapping of ACL permissions and access policy permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#acl-access-policy-permission-mapping
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Controlling object ownership and disabling ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
func (c *Client) GetObjectAcl(ctx context.Context, params *GetObjectAclInput, optFns ...func(*Options)) (*GetObjectAclOutput, error) {
	if params == nil {
		params = &GetObjectAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObjectAcl", params, optFns, c.addOperationGetObjectAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectAclInput struct {

	// The bucket name that contains the object for which to get the ACL information.
	//
	// Access points - When you use this action with an access point, you must provide
	// the alias of the access point in place of the bucket name or specify the access
	// point ARN. When using the access point ARN, you must direct requests to the
	// access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see [Using access points]in the Amazon S3 User Guide.
	//
	// [Using access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html
	//
	// This member is required.
	Bucket *string

	// The key of the object for which to get the ACL information.
	//
	// This member is required.
	Key *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. If either the
	// source or destination S3 bucket has Requester Pays enabled, the requester will
	// pay for corresponding charges to copy the object. For information about
	// downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays Buckets]in the Amazon S3 User
	// Guide.
	//
	// This functionality is not supported for directory buckets.
	//
	// [Downloading Objects in Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer types.RequestPayer

	// Version ID used to reference a specific version of the object.
	//
	// This functionality is not supported for directory buckets.
	VersionId *string

	noSmithyDocumentSerde
}

func (in *GetObjectAclInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.Key = in.Key

}

type GetObjectAclOutput struct {

	// A list of grants.
	Grants []types.Grant

	//  Container for the bucket owner's display name and ID.
	Owner *types.Owner

	// If present, indicates that the requester was successfully charged for the
	// request.
	//
	// This functionality is not supported for directory buckets.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetObjectAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetObjectAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectAcl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetObjectAcl"); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetObjectAclValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectAcl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetObjectAclUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *GetObjectAclInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetObjectAcl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetObjectAcl",
	}
}

// getGetObjectAclBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getGetObjectAclBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetObjectAclInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetObjectAclUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetObjectAclBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}