package jsapi

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/consts"
	"hash"
	"io/ioutil"
	nethttp "net/http"
	neturl "net/url"
)

func (a *JsapiApiService) TradeBill(ctx context.Context, req TradeBillRequest) (
	resp *BillResponse,
	result *core.APIResult,
	err error,
) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/bill/tradebill"
	// Make sure All Required Params are properly set

	if req.BillDate == nil {
		return nil, nil, fmt.Errorf("field `BillDate` is required and must be specified in TradeBillRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("bill_date", core.ParameterToString(*req.BillDate, ""))
	if req.BillType != nil {
		localVarQueryParams.Add("bill_type", core.ParameterToString(*req.BillType, ""))
	}
	if req.TarType != nil {
		localVarQueryParams.Add("tar_type", core.ParameterToString(*req.TarType, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(
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

	// Extract PrepayResponse from Http Response
	resp = new(BillResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

func (a *JsapiApiService) DownloadURL(ctx context.Context, bill BillResponse) (
	resp []byte,
	result *core.APIResult,
	err error,
) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := *bill.DownloadURL
	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"text/plain"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.RequestWithoutValidate(
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

	resp, err = ioutil.ReadAll(result.Response.Body)
	if err != nil {
		return nil, result, err
	}

	err = validateHash(resp, *bill.HashType, *bill.HashValue)
	return
}

func validateHash(data []byte, hashType, hashValue string) error {
	var hasher hash.Hash
	if hashType == "SHA1" {
		hasher = sha1.New()
	}
	if hasher == nil {
		return nil
	}

	_, err := hasher.Write(data)
	if err != nil {
		return err
	}

	result := hasher.Sum(nil)
	if hashValue != hex.EncodeToString(result) {
		return fmt.Errorf("hashValue invalid, value: %s", hashValue)
	}
	return nil
}
