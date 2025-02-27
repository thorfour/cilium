// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeHostsRequest
type DescribeHostsInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	//
	//    * auto-placement - Whether auto-placement is enabled or disabled (on |
	//    off).
	//
	//    * availability-zone - The Availability Zone of the host.
	//
	//    * client-token - The idempotency token that you provided when you allocated
	//    the host.
	//
	//    * host-reservation-id - The ID of the reservation assigned to this host.
	//
	//    * instance-type - The instance type size that the Dedicated Host is configured
	//    to support.
	//
	//    * state - The allocation state of the Dedicated Host (available | under-assessment
	//    | permanent-failure | released | released-permanent-failure).
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filter []Filter `locationName:"filter" locationNameList:"Filter" type:"list"`

	// The IDs of the Dedicated Hosts. The IDs are used for targeted instance launches.
	HostIds []string `locationName:"hostId" locationNameList:"item" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given
	// a larger value than 500, you receive an error.
	//
	// You cannot specify this parameter and the host IDs parameter in the same
	// request.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeHostsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeHostsResult
type DescribeHostsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Dedicated Hosts.
	Hosts []Host `locationName:"hostSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeHostsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHosts = "DescribeHosts"

// DescribeHostsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified Dedicated Hosts or all your Dedicated Hosts.
//
// The results describe only the Dedicated Hosts in the Region you're currently
// using. All listed instances consume capacity on your Dedicated Host. Dedicated
// Hosts that have recently been released are listed with the state released.
//
//    // Example sending a request using DescribeHostsRequest.
//    req := client.DescribeHostsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeHosts
func (c *Client) DescribeHostsRequest(input *DescribeHostsInput) DescribeHostsRequest {
	op := &aws.Operation{
		Name:       opDescribeHosts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeHostsInput{}
	}

	req := c.newRequest(op, input, &DescribeHostsOutput{})
	return DescribeHostsRequest{Request: req, Input: input, Copy: c.DescribeHostsRequest}
}

// DescribeHostsRequest is the request type for the
// DescribeHosts API operation.
type DescribeHostsRequest struct {
	*aws.Request
	Input *DescribeHostsInput
	Copy  func(*DescribeHostsInput) DescribeHostsRequest
}

// Send marshals and sends the DescribeHosts API request.
func (r DescribeHostsRequest) Send(ctx context.Context) (*DescribeHostsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHostsResponse{
		DescribeHostsOutput: r.Request.Data.(*DescribeHostsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeHostsRequestPaginator returns a paginator for DescribeHosts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeHostsRequest(input)
//   p := ec2.NewDescribeHostsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeHostsPaginator(req DescribeHostsRequest) DescribeHostsPaginator {
	return DescribeHostsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeHostsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeHostsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeHostsPaginator struct {
	aws.Pager
}

func (p *DescribeHostsPaginator) CurrentPage() *DescribeHostsOutput {
	return p.Pager.CurrentPage().(*DescribeHostsOutput)
}

// DescribeHostsResponse is the response type for the
// DescribeHosts API operation.
type DescribeHostsResponse struct {
	*DescribeHostsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHosts request.
func (r *DescribeHostsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
