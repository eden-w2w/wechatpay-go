package coupons

import (
	"context"
	"fmt"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/consts"
	"github.com/eden-w2w/wechatpay-go/services"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

type CouponApiService services.Service

func (s *CouponApiService) GetStocks(ctx context.Context, req GetStocksRequest) (
	resp *StocksPagination,
	result *core.APIResult,
	err error,
) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/stocks"
	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("stock_creator_mchid", core.ParameterToString(req.StockCreatorMchID, ""))
	localVarQueryParams.Add("offset", core.ParameterToString(req.Offset, ""))
	localVarQueryParams.Add("limit", core.ParameterToString(req.Limit, ""))
	if req.CreateStartTime != nil {
		localVarQueryParams.Add("create_start_time", core.ParameterToString(*req.CreateStartTime, ""))
	}
	if req.CreateEndTime != nil {
		localVarQueryParams.Add("create_end_time", core.ParameterToString(*req.CreateEndTime, ""))
	}
	if req.Status != nil {
		localVarQueryParams.Add("status", core.ParameterToString(*req.Status, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(
		ctx,
		localVarHTTPMethod,
		localVarPath,
		localVarHeaderParams,
		localVarQueryParams,
		localVarPostBody,
		localVarHTTPContentType,
	)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(StocksPagination)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (s *CouponApiService) GiveCoupon(ctx context.Context, req GiveCouponRequest) (
	resp *GiveCouponResponse,
	result *core.APIResult,
	err error,
) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	if req.OpenID == nil {
		return nil, nil, fmt.Errorf("field `OpenID` is required and must be specified in GiveCouponRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/favor/users/{openid}/coupons"
	// Build Path with Path Params
	localVarPath = strings.Replace(
		localVarPath,
		"{openid}",
		neturl.PathEscape(core.ParameterToString(*req.OpenID, "")),
		-1,
	)
	req.OpenID = nil

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = s.Client.Request(
		ctx,
		localVarHTTPMethod,
		localVarPath,
		localVarHeaderParams,
		localVarQueryParams,
		localVarPostBody,
		localVarHTTPContentType,
	)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(GiveCouponResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
