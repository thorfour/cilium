// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRegionsRequest
type DescribeRegionsInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to display all Regions, including Regions that are disabled
	// for your account.
	AllRegions *bool `type:"boolean"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * endpoint - The endpoint of the Region (for example, ec2.us-east-1.amazonaws.com).
	//
	//    * region-name - The name of the Region (for example, us-east-1).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The names of the Regions. You can specify any Regions, whether they are enabled
	// and disabled for your account.
	RegionNames []string `locationName:"RegionName" locationNameList:"RegionName" type:"list"`
}

// String returns the string representation
func (s DescribeRegionsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRegionsResult
type DescribeRegionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Regions.
	Regions []Region `locationName:"regionInfo" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeRegionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRegions = "DescribeRegions"

// DescribeRegionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Regions that are enabled for your account, or all Regions.
//
// For a list of the Regions supported by Amazon EC2, see Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#ec2_region).
//
// For information about enabling and disabling Regions for your account, see
// Managing AWS Regions (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html)
// in the AWS General Reference.
//
//    // Example sending a request using DescribeRegionsRequest.
//    req := client.DescribeRegionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRegions
func (c *Client) DescribeRegionsRequest(input *DescribeRegionsInput) DescribeRegionsRequest {
	op := &aws.Operation{
		Name:       opDescribeRegions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRegionsInput{}
	}

	req := c.newRequest(op, input, &DescribeRegionsOutput{})
	return DescribeRegionsRequest{Request: req, Input: input, Copy: c.DescribeRegionsRequest}
}

// DescribeRegionsRequest is the request type for the
// DescribeRegions API operation.
type DescribeRegionsRequest struct {
	*aws.Request
	Input *DescribeRegionsInput
	Copy  func(*DescribeRegionsInput) DescribeRegionsRequest
}

// Send marshals and sends the DescribeRegions API request.
func (r DescribeRegionsRequest) Send(ctx context.Context) (*DescribeRegionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRegionsResponse{
		DescribeRegionsOutput: r.Request.Data.(*DescribeRegionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRegionsResponse is the response type for the
// DescribeRegions API operation.
type DescribeRegionsResponse struct {
	*DescribeRegionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRegions request.
func (r *DescribeRegionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
