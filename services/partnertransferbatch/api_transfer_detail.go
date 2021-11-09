// Copyright 2021 Tencent Inc. All rights reserved.
//
// 服务商批量转账API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package partnertransferbatch

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/consts"
	"github.com/eden-w2w/wechatpay-go/services"
)

type TransferDetailApiService services.Service

// GetTransferDetailByNo 微信支付明细单号查询明细单
//
// ## 查询接口说明
// 微信支付明细单号查单接口。转账处理后延迟一段时间（异步进行转账），服务商可以通过该接口查询单笔转账明细单。返回消息中包含微信支付明细单号、明细状态、转账金额、失败原因、收款用户姓名、用户OpenID等信息。
//
// 接口限频：
// 单个服务商（查询转账明细单）50QPS，如果超过频率限制，会报错FREQUENCY_LIMITED，请降低频率请求。
//
// 注意事项：
// - API只支持查询最近30天内的转账明细单，30天之前的转账明细单请登录商户平台查询。
// - 转账明细单中涉及金额的字段单位为“分”。
// - 如果查询单号对应的数据不存在，那么数据不存在的原因可能是：（1）转账还在处理中；（2）转账批次单受理失败或还未开始处理导致转账明细单没有落地。在上述情况下，服务商首先需要检查该微信支付明细单号是否确实是自己发起，以及是否是该批次下的，如果服务商确认是自己发起且是该批次下的，则请服务商不要直接当做转账失败处理，请服务商隔几分钟再尝试查询（请勿转账和查询并发处理）。如果服务商误把还在转账处理中的明细单直接当转账失败处理，服务商应当自行承担因此产生的所有损失和责任。
// - 如果遇到回包返回新的错误码，请务必不要换单重试，请联系客服确认转账情况。如果有新的错误码，会更新到此API文档中。
// - 错误码描述字段message只供人工定位问题时做参考，系统实现时请不要依赖这个字段来做自动化处理。
func (a *TransferDetailApiService) GetTransferDetailByNo(ctx context.Context, req GetTransferDetailByNoRequest) (resp *TransferDetailEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.BatchId == nil {
		return nil, nil, fmt.Errorf("field `BatchId` is required and must be specified in GetTransferDetailByNoRequest")
	}
	if req.DetailId == nil {
		return nil, nil, fmt.Errorf("field `DetailId` is required and must be specified in GetTransferDetailByNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/partner-transfer/batches/batch-id/{batch_id}/details/detail-id/{detail_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"batch_id"+"}", neturl.PathEscape(core.ParameterToString(*req.BatchId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"detail_id"+"}", neturl.PathEscape(core.ParameterToString(*req.DetailId, "")), -1)

	// Make sure All Required Params are properly set

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract TransferDetailEntity from Http Response
	resp = new(TransferDetailEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	// 对应答中隐私字段进行解密
	err = a.Client.DecryptResponse(ctx, resp)
	if err != nil {
		return resp, result, err
	}

	return resp, result, nil
}

// GetTransferDetailByOutNo 商家明细单号查询明细单
//
// ## 查询接口说明
// 商户明细单号查单接口。转账处理后延迟一段时间（异步进行转账），服务商可以通过该接口查询单笔转账明细单。返回消息中包含微信明细单号、明细状态、转账金额、失败原因、收款用户姓名、用户OpenID等信息。
//
// 接口限频：
// 单个服务商（查询转账明细单）50QPS，如果超过频率限制，会报错FREQUENCY_LIMITED，请降低频率请求。
//
// 注意事项：
// - API只支持查询最近30天内的转账明细单，30天之前的转账明细单请登录商户平台查询。
// - 转账明细单中涉及金额的字段单位为“分”。
// - 如果查询单号对应的数据不存在，那么数据不存在的原因可能是：（1）转账还在处理中；（2）转账批次单受理失败或还未开始处理导致转账明细单没有落地。在上述情况下，服务商首先需要检查该商家明细单号是否确实是自己发起，以及是否是该批次下的，如果服务商确认是自己发起且是该批次下的，则请服务商不要直接当做转账失败处理，请服务商隔几分钟再尝试查询（请勿转账和查询并发处理）。如果服务商误把还在转账处理中的明细单直接当转账失败处理，服务商应当自行承担因此产生的所有损失和责任。
// - 如果遇到回包返回新的错误码，请务必不要换单重试，请联系客服确认转账情况。如果有新的错误码，会更新到此API文档中。
// - 错误码描述字段message只供人工定位问题时做参考，系统实现时请不要依赖这个字段来做自动化处理。
func (a *TransferDetailApiService) GetTransferDetailByOutNo(ctx context.Context, req GetTransferDetailByOutNoRequest) (resp *TransferDetailEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutBatchNo == nil {
		return nil, nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in GetTransferDetailByOutNoRequest")
	}
	if req.OutDetailNo == nil {
		return nil, nil, fmt.Errorf("field `OutDetailNo` is required and must be specified in GetTransferDetailByOutNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/partner-transfer/batches/out-batch-no/{out_batch_no}/details/out-detail-no/{out_detail_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_batch_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutBatchNo, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"out_detail_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutDetailNo, "")), -1)

	// Make sure All Required Params are properly set

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract TransferDetailEntity from Http Response
	resp = new(TransferDetailEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	// 对应答中隐私字段进行解密
	err = a.Client.DecryptResponse(ctx, resp)
	if err != nil {
		return resp, result, err
	}

	return resp, result, nil
}
