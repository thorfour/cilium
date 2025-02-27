// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetEbsDefaultKmsKeyIdRequest
type GetEbsDefaultKmsKeyIdInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s GetEbsDefaultKmsKeyIdInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetEbsDefaultKmsKeyIdResult
type GetEbsDefaultKmsKeyIdOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the default CMK for encryption by default.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`
}

// String returns the string representation
func (s GetEbsDefaultKmsKeyIdOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetEbsDefaultKmsKeyId = "GetEbsDefaultKmsKeyId"

// GetEbsDefaultKmsKeyIdRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the default customer master key (CMK) for EBS encryption by default
// for your account in this Region. You can change the default CMK for encryption
// by default using ModifyEbsDefaultKmsKeyId or ResetEbsDefaultKmsKeyId.
//
// For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using GetEbsDefaultKmsKeyIdRequest.
//    req := client.GetEbsDefaultKmsKeyIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetEbsDefaultKmsKeyId
func (c *Client) GetEbsDefaultKmsKeyIdRequest(input *GetEbsDefaultKmsKeyIdInput) GetEbsDefaultKmsKeyIdRequest {
	op := &aws.Operation{
		Name:       opGetEbsDefaultKmsKeyId,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetEbsDefaultKmsKeyIdInput{}
	}

	req := c.newRequest(op, input, &GetEbsDefaultKmsKeyIdOutput{})
	return GetEbsDefaultKmsKeyIdRequest{Request: req, Input: input, Copy: c.GetEbsDefaultKmsKeyIdRequest}
}

// GetEbsDefaultKmsKeyIdRequest is the request type for the
// GetEbsDefaultKmsKeyId API operation.
type GetEbsDefaultKmsKeyIdRequest struct {
	*aws.Request
	Input *GetEbsDefaultKmsKeyIdInput
	Copy  func(*GetEbsDefaultKmsKeyIdInput) GetEbsDefaultKmsKeyIdRequest
}

// Send marshals and sends the GetEbsDefaultKmsKeyId API request.
func (r GetEbsDefaultKmsKeyIdRequest) Send(ctx context.Context) (*GetEbsDefaultKmsKeyIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEbsDefaultKmsKeyIdResponse{
		GetEbsDefaultKmsKeyIdOutput: r.Request.Data.(*GetEbsDefaultKmsKeyIdOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEbsDefaultKmsKeyIdResponse is the response type for the
// GetEbsDefaultKmsKeyId API operation.
type GetEbsDefaultKmsKeyIdResponse struct {
	*GetEbsDefaultKmsKeyIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEbsDefaultKmsKeyId request.
func (r *GetEbsDefaultKmsKeyIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
