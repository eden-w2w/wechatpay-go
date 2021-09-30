// Copyright 2021 Tencent Inc. All rights reserved.
//
// 商家转账对外API
//
// - 场景及业务流程： 商户可通过该产品实现同时向多个用户微信零钱进行转账的操作，可用于发放奖金补贴、佣金货款结算、员工报销等场景。 [https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0](https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0) - 接入步骤：     - 商户在微信支付商户平台开通“批量转账到零钱”产品权限，并勾选“使用API方式发起转账”。     - 调用批量转账接口，对多个用户微信零钱发起转账。     - 调用查询批次接口，可获取到转账批次详情及当前状态。     - 调用查询明细接口，可获取到单条转账明细详情及当前状态。
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package transferbatch

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// CloseReasonType * `MERCHANT_REVOCATION` - 商户主动撤销, 商户主动撤销（页面方式） * `OVERDUE_CLOSE` - 系统超时关闭, 系统超时关闭，可能原因账户余额不足或其他错误
type CloseReasonType string

func (e CloseReasonType) Ptr() *CloseReasonType {
	return &e
}

// Enums of CloseReasonType
const (
	CLOSEREASONTYPE_MERCHANT_REVOCATION CloseReasonType = "MERCHANT_REVOCATION"
	CLOSEREASONTYPE_OVERDUE_CLOSE       CloseReasonType = "OVERDUE_CLOSE"
)

// FailReasonType * `ACCOUNT_FROZEN` - 账户冻结, 该用户账户被冻结 * `REAL_NAME_CHECK_FAIL` - 用户未实名, 收款人未实名认证，需要用户完成微信实名认证 * `NAME_NOT_CORRECT` - 用户姓名校验失败, 收款人姓名校验不通过，请核实信息 * `OPENID_INVALID` - Openid校验失败, Openid格式错误或者不属于商家公众账号 * `TRANSFER_QUOTA_EXCEED` - 超过用户单笔收款额度, 超过用户单笔收款额度，核实产品设置是否准确 * `DAY_RECEIVED_QUOTA_EXCEED` - 超过用户单日收款额度, 超过用户单日收款额度，核实产品设置是否准确 * `MONTH_RECEIVED_QUOTA_EXCEED` - 超过用户单月收款额度, 超过用户单月收款额度，核实产品设置是否准确 * `DAY_RECEIVED_COUNT_EXCEED` - 超过用户单日收款次数, 超过用户单日收款次数，核实产品设置是否准确 * `PRODUCT_AUTH_CHECK_FAIL` - 产品权限校验失败, 未开通该权限或权限被冻结，请核实产品权限状态 * `OVERDUE_CLOSE` - 转账关闭, 超过系统重试期，系统自动关闭 * `ID_CARD_NOT_CORRECT` - 用户身份证校验失败, 收款人身份证校验不通过，请核实信息 * `ACCOUNT_NOT_EXIST` - 用户账户不存在, 该用户账户不存在 * `TRANSFER_RISK` - 转账存在风险, 该笔转账可能存在风险，已被微信拦截
type FailReasonType string

func (e FailReasonType) Ptr() *FailReasonType {
	return &e
}

// Enums of FailReasonType
const (
	FAILREASONTYPE_ACCOUNT_FROZEN              FailReasonType = "ACCOUNT_FROZEN"
	FAILREASONTYPE_REAL_NAME_CHECK_FAIL        FailReasonType = "REAL_NAME_CHECK_FAIL"
	FAILREASONTYPE_NAME_NOT_CORRECT            FailReasonType = "NAME_NOT_CORRECT"
	FAILREASONTYPE_OPENID_INVALID              FailReasonType = "OPENID_INVALID"
	FAILREASONTYPE_TRANSFER_QUOTA_EXCEED       FailReasonType = "TRANSFER_QUOTA_EXCEED"
	FAILREASONTYPE_DAY_RECEIVED_QUOTA_EXCEED   FailReasonType = "DAY_RECEIVED_QUOTA_EXCEED"
	FAILREASONTYPE_MONTH_RECEIVED_QUOTA_EXCEED FailReasonType = "MONTH_RECEIVED_QUOTA_EXCEED"
	FAILREASONTYPE_DAY_RECEIVED_COUNT_EXCEED   FailReasonType = "DAY_RECEIVED_COUNT_EXCEED"
	FAILREASONTYPE_PRODUCT_AUTH_CHECK_FAIL     FailReasonType = "PRODUCT_AUTH_CHECK_FAIL"
	FAILREASONTYPE_OVERDUE_CLOSE               FailReasonType = "OVERDUE_CLOSE"
	FAILREASONTYPE_ID_CARD_NOT_CORRECT         FailReasonType = "ID_CARD_NOT_CORRECT"
	FAILREASONTYPE_ACCOUNT_NOT_EXIST           FailReasonType = "ACCOUNT_NOT_EXIST"
	FAILREASONTYPE_TRANSFER_RISK               FailReasonType = "TRANSFER_RISK"
)

// GetTransferBatchByNoRequest
type GetTransferBatchByNoRequest struct {
	// 微信批次单号，微信商家转账系统返回的唯一标识
	BatchId *string `json:"batch_id"`
	// true-是；false-否，默认否。商户可选择是否查询指定状态的转账明细单，当转账批次单状态为“FINISHED”（已完成）时，才会返回满足条件的转账明细单
	NeedQueryDetail *bool `json:"need_query_detail"`
	// 该次请求资源的起始位置。返回的明细是按照设置的明细条数进行分页展示的，一次查询可能无法返回所有明细，我们使用该参数标识查询开始位置，默认值为0
	Offset *int64 `json:"offset,omitempty"`
	// 该次请求可返回的最大明细条数，最小20条，最大100条，不传则默认20条。不足20条按实际条数返回
	Limit *int64 `json:"limit,omitempty"`
	// 查询指定状态的转账明细单   ALL:全部。需要同时查询转账成功和转账失败的明细单   SUCCESS:转账成功。只查询转账成功的明细单   FAIL:转账失败。只查询转账失败的明细单
	DetailStatus *string `json:"detail_status,omitempty"`
}

func (o GetTransferBatchByNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.BatchId == nil {
		return nil, fmt.Errorf("field `BatchId` is required and must be specified in GetTransferBatchByNoRequest")
	}
	toSerialize["batch_id"] = o.BatchId

	if o.NeedQueryDetail == nil {
		return nil, fmt.Errorf("field `NeedQueryDetail` is required and must be specified in GetTransferBatchByNoRequest")
	}
	toSerialize["need_query_detail"] = o.NeedQueryDetail

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.DetailStatus != nil {
		toSerialize["detail_status"] = o.DetailStatus
	}
	return json.Marshal(toSerialize)
}

func (o GetTransferBatchByNoRequest) String() string {
	var ret string
	if o.BatchId == nil {
		ret += "BatchId:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchId:%v, ", *o.BatchId)
	}

	if o.NeedQueryDetail == nil {
		ret += "NeedQueryDetail:<nil>, "
	} else {
		ret += fmt.Sprintf("NeedQueryDetail:%v, ", *o.NeedQueryDetail)
	}

	if o.Offset == nil {
		ret += "Offset:<nil>, "
	} else {
		ret += fmt.Sprintf("Offset:%v, ", *o.Offset)
	}

	if o.Limit == nil {
		ret += "Limit:<nil>, "
	} else {
		ret += fmt.Sprintf("Limit:%v, ", *o.Limit)
	}

	if o.DetailStatus == nil {
		ret += "DetailStatus:<nil>"
	} else {
		ret += fmt.Sprintf("DetailStatus:%v", *o.DetailStatus)
	}

	return fmt.Sprintf("GetTransferBatchByNoRequest{%s}", ret)
}

func (o GetTransferBatchByNoRequest) Clone() *GetTransferBatchByNoRequest {
	ret := GetTransferBatchByNoRequest{}

	if o.BatchId != nil {
		ret.BatchId = new(string)
		*ret.BatchId = *o.BatchId
	}

	if o.NeedQueryDetail != nil {
		ret.NeedQueryDetail = new(bool)
		*ret.NeedQueryDetail = *o.NeedQueryDetail
	}

	if o.Offset != nil {
		ret.Offset = new(int64)
		*ret.Offset = *o.Offset
	}

	if o.Limit != nil {
		ret.Limit = new(int64)
		*ret.Limit = *o.Limit
	}

	if o.DetailStatus != nil {
		ret.DetailStatus = new(string)
		*ret.DetailStatus = *o.DetailStatus
	}

	return &ret
}

// GetTransferBatchByOutNoRequest
type GetTransferBatchByOutNoRequest struct {
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
	// true-是；false-否，默认否。商户可选择是否查询指定状态的转账明细单，当转账批次单状态为“FINISHED”（已完成）时，才会返回满足条件的转账明细单
	NeedQueryDetail *bool `json:"need_query_detail"`
	// 该次请求资源（转账明细单）的起始位置，从0开始，默认值为0
	Offset *int64 `json:"offset,omitempty"`
	// 该次请求可返回的最大资源（转账明细单）条数，最小20条，最大100条，不传则默认20条。不足20条按实际条数返回
	Limit *int64 `json:"limit,omitempty"`
	// 查询指定状态的转账明细单   ALL:全部。需要同时查询转账成功和转账失败的明细单   SUCCESS:转账成功。只查询转账成功的明细单   FAIL:转账失败。只查询转账失败的明细单
	DetailStatus *string `json:"detail_status,omitempty"`
}

func (o GetTransferBatchByOutNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in GetTransferBatchByOutNoRequest")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo

	if o.NeedQueryDetail == nil {
		return nil, fmt.Errorf("field `NeedQueryDetail` is required and must be specified in GetTransferBatchByOutNoRequest")
	}
	toSerialize["need_query_detail"] = o.NeedQueryDetail

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.DetailStatus != nil {
		toSerialize["detail_status"] = o.DetailStatus
	}
	return json.Marshal(toSerialize)
}

func (o GetTransferBatchByOutNoRequest) String() string {
	var ret string
	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v, ", *o.OutBatchNo)
	}

	if o.NeedQueryDetail == nil {
		ret += "NeedQueryDetail:<nil>, "
	} else {
		ret += fmt.Sprintf("NeedQueryDetail:%v, ", *o.NeedQueryDetail)
	}

	if o.Offset == nil {
		ret += "Offset:<nil>, "
	} else {
		ret += fmt.Sprintf("Offset:%v, ", *o.Offset)
	}

	if o.Limit == nil {
		ret += "Limit:<nil>, "
	} else {
		ret += fmt.Sprintf("Limit:%v, ", *o.Limit)
	}

	if o.DetailStatus == nil {
		ret += "DetailStatus:<nil>"
	} else {
		ret += fmt.Sprintf("DetailStatus:%v", *o.DetailStatus)
	}

	return fmt.Sprintf("GetTransferBatchByOutNoRequest{%s}", ret)
}

func (o GetTransferBatchByOutNoRequest) Clone() *GetTransferBatchByOutNoRequest {
	ret := GetTransferBatchByOutNoRequest{}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	if o.NeedQueryDetail != nil {
		ret.NeedQueryDetail = new(bool)
		*ret.NeedQueryDetail = *o.NeedQueryDetail
	}

	if o.Offset != nil {
		ret.Offset = new(int64)
		*ret.Offset = *o.Offset
	}

	if o.Limit != nil {
		ret.Limit = new(int64)
		*ret.Limit = *o.Limit
	}

	if o.DetailStatus != nil {
		ret.DetailStatus = new(string)
		*ret.DetailStatus = *o.DetailStatus
	}

	return &ret
}

// GetTransferDetailByNoRequest
type GetTransferDetailByNoRequest struct {
	// 微信批次单号，微信商家转账系统返回的唯一标识
	BatchId *string `json:"batch_id"`
	// 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识
	DetailId *string `json:"detail_id"`
}

func (o GetTransferDetailByNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.BatchId == nil {
		return nil, fmt.Errorf("field `BatchId` is required and must be specified in GetTransferDetailByNoRequest")
	}
	toSerialize["batch_id"] = o.BatchId

	if o.DetailId == nil {
		return nil, fmt.Errorf("field `DetailId` is required and must be specified in GetTransferDetailByNoRequest")
	}
	toSerialize["detail_id"] = o.DetailId
	return json.Marshal(toSerialize)
}

func (o GetTransferDetailByNoRequest) String() string {
	var ret string
	if o.BatchId == nil {
		ret += "BatchId:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchId:%v, ", *o.BatchId)
	}

	if o.DetailId == nil {
		ret += "DetailId:<nil>"
	} else {
		ret += fmt.Sprintf("DetailId:%v", *o.DetailId)
	}

	return fmt.Sprintf("GetTransferDetailByNoRequest{%s}", ret)
}

func (o GetTransferDetailByNoRequest) Clone() *GetTransferDetailByNoRequest {
	ret := GetTransferDetailByNoRequest{}

	if o.BatchId != nil {
		ret.BatchId = new(string)
		*ret.BatchId = *o.BatchId
	}

	if o.DetailId != nil {
		ret.DetailId = new(string)
		*ret.DetailId = *o.DetailId
	}

	return &ret
}

// GetTransferDetailByOutNoRequest
type GetTransferDetailByOutNoRequest struct {
	// 商户系统内部区分转账批次单下不同转账明细单的唯一标识
	OutDetailNo *string `json:"out_detail_no"`
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
}

func (o GetTransferDetailByOutNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutDetailNo == nil {
		return nil, fmt.Errorf("field `OutDetailNo` is required and must be specified in GetTransferDetailByOutNoRequest")
	}
	toSerialize["out_detail_no"] = o.OutDetailNo

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in GetTransferDetailByOutNoRequest")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo
	return json.Marshal(toSerialize)
}

func (o GetTransferDetailByOutNoRequest) String() string {
	var ret string
	if o.OutDetailNo == nil {
		ret += "OutDetailNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutDetailNo:%v, ", *o.OutDetailNo)
	}

	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>"
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v", *o.OutBatchNo)
	}

	return fmt.Sprintf("GetTransferDetailByOutNoRequest{%s}", ret)
}

func (o GetTransferDetailByOutNoRequest) Clone() *GetTransferDetailByOutNoRequest {
	ret := GetTransferDetailByOutNoRequest{}

	if o.OutDetailNo != nil {
		ret.OutDetailNo = new(string)
		*ret.OutDetailNo = *o.OutDetailNo
	}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	return &ret
}

// InitiateBatchTransferRequest
type InitiateBatchTransferRequest struct {
	// 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	Appid *string `json:"appid"`
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
	// 该笔批量转账的名称
	BatchName *string `json:"batch_name"`
	// 转账说明，UTF8编码，最多允许32个字符
	BatchRemark *string `json:"batch_remark"`
	// 转账金额单位为“分”。转账总金额必须与批次内所有明细转账金额之和保持一致，否则无法发起转账操作
	TotalAmount *int64 `json:"total_amount"`
	// 一个转账批次单最多发起三千笔转账。转账总笔数必须与批次内所有明细之和保持一致，否则无法发起转账操作
	TotalNum *int64 `json:"total_num"`
	// 发起批量转账的明细列表，最多三千笔
	TransferDetailList []TransferDetailInput `json:"transfer_detail_list,omitempty"`
}

func (o InitiateBatchTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["appid"] = o.Appid

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo

	if o.BatchName == nil {
		return nil, fmt.Errorf("field `BatchName` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["batch_name"] = o.BatchName

	if o.BatchRemark == nil {
		return nil, fmt.Errorf("field `BatchRemark` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["batch_remark"] = o.BatchRemark

	if o.TotalAmount == nil {
		return nil, fmt.Errorf("field `TotalAmount` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["total_amount"] = o.TotalAmount

	if o.TotalNum == nil {
		return nil, fmt.Errorf("field `TotalNum` is required and must be specified in InitiateBatchTransferRequest")
	}
	toSerialize["total_num"] = o.TotalNum

	if o.TransferDetailList != nil {
		toSerialize["transfer_detail_list"] = o.TransferDetailList
	}
	return json.Marshal(toSerialize)
}

func (o InitiateBatchTransferRequest) String() string {
	var ret string
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v, ", *o.OutBatchNo)
	}

	if o.BatchName == nil {
		ret += "BatchName:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchName:%v, ", *o.BatchName)
	}

	if o.BatchRemark == nil {
		ret += "BatchRemark:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchRemark:%v, ", *o.BatchRemark)
	}

	if o.TotalAmount == nil {
		ret += "TotalAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TotalAmount:%v, ", *o.TotalAmount)
	}

	if o.TotalNum == nil {
		ret += "TotalNum:<nil>, "
	} else {
		ret += fmt.Sprintf("TotalNum:%v, ", *o.TotalNum)
	}

	ret += fmt.Sprintf("TransferDetailList:%v", o.TransferDetailList)

	return fmt.Sprintf("InitiateBatchTransferRequest{%s}", ret)
}

func (o InitiateBatchTransferRequest) Clone() *InitiateBatchTransferRequest {
	ret := InitiateBatchTransferRequest{}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	if o.BatchName != nil {
		ret.BatchName = new(string)
		*ret.BatchName = *o.BatchName
	}

	if o.BatchRemark != nil {
		ret.BatchRemark = new(string)
		*ret.BatchRemark = *o.BatchRemark
	}

	if o.TotalAmount != nil {
		ret.TotalAmount = new(int64)
		*ret.TotalAmount = *o.TotalAmount
	}

	if o.TotalNum != nil {
		ret.TotalNum = new(int64)
		*ret.TotalNum = *o.TotalNum
	}

	if o.TransferDetailList != nil {
		ret.TransferDetailList = make([]TransferDetailInput, len(o.TransferDetailList))
		for i, item := range o.TransferDetailList {
			ret.TransferDetailList[i] = *item.Clone()
		}
	}

	return &ret
}

// InitiateBatchTransferResponse
type InitiateBatchTransferResponse struct {
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
	// 微信批次单号，微信商家转账系统返回的唯一标识
	BatchId *string `json:"batch_id"`
	// 批次受理成功时返回，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	CreateTime *time.Time `json:"create_time"`
}

func (o InitiateBatchTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in InitiateBatchTransferResponse")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo

	if o.BatchId == nil {
		return nil, fmt.Errorf("field `BatchId` is required and must be specified in InitiateBatchTransferResponse")
	}
	toSerialize["batch_id"] = o.BatchId

	if o.CreateTime == nil {
		return nil, fmt.Errorf("field `CreateTime` is required and must be specified in InitiateBatchTransferResponse")
	}
	toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)
	return json.Marshal(toSerialize)
}

func (o InitiateBatchTransferResponse) String() string {
	var ret string
	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v, ", *o.OutBatchNo)
	}

	if o.BatchId == nil {
		ret += "BatchId:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchId:%v, ", *o.BatchId)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>"
	} else {
		ret += fmt.Sprintf("CreateTime:%v", *o.CreateTime)
	}

	return fmt.Sprintf("InitiateBatchTransferResponse{%s}", ret)
}

func (o InitiateBatchTransferResponse) Clone() *InitiateBatchTransferResponse {
	ret := InitiateBatchTransferResponse{}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	if o.BatchId != nil {
		ret.BatchId = new(string)
		*ret.BatchId = *o.BatchId
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	return &ret
}

// TransferBatchEntity
type TransferBatchEntity struct {
	// 转账批次单基本信息
	TransferBatch *TransferBatchGet `json:"transfer_batch"`
	// 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回。包括微信明细单号、明细状态信息
	TransferDetailList []TransferDetailCompact `json:"transfer_detail_list,omitempty"`
}

func (o TransferBatchEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TransferBatch == nil {
		return nil, fmt.Errorf("field `TransferBatch` is required and must be specified in TransferBatchEntity")
	}
	toSerialize["transfer_batch"] = o.TransferBatch

	if o.TransferDetailList != nil {
		toSerialize["transfer_detail_list"] = o.TransferDetailList
	}
	return json.Marshal(toSerialize)
}

func (o TransferBatchEntity) String() string {
	var ret string
	ret += fmt.Sprintf("TransferBatch:%v, ", o.TransferBatch)

	ret += fmt.Sprintf("TransferDetailList:%v", o.TransferDetailList)

	return fmt.Sprintf("TransferBatchEntity{%s}", ret)
}

func (o TransferBatchEntity) Clone() *TransferBatchEntity {
	ret := TransferBatchEntity{}

	if o.TransferBatch != nil {
		ret.TransferBatch = o.TransferBatch.Clone()
	}

	if o.TransferDetailList != nil {
		ret.TransferDetailList = make([]TransferDetailCompact, len(o.TransferDetailList))
		for i, item := range o.TransferDetailList {
			ret.TransferDetailList[i] = *item.Clone()
		}
	}

	return &ret
}

// TransferBatchGet
type TransferBatchGet struct {
	// 微信支付分配的商户号
	Mchid *string `json:"mchid"`
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
	// 微信批次单号，微信商家转账系统返回的唯一标识
	BatchId *string `json:"batch_id"`
	// 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	Appid *string `json:"appid"`
	// ACCEPTED:已受理。批次已受理成功，若发起批量转账的30分钟后，转账批次单仍处于该状态，可能原因是商户账户余额不足等。商户可查询账户资金流水，若该笔转账批次单的扣款已经发生，则表示批次已经进入转账中，请再次查单确认   PROCESSING:转账中。已开始处理批次内的转账明细单   FINISHED:已完成。批次内的所有转账明细单都已处理完成   CLOSED:已关闭。可查询具体的批次关闭原因确认
	BatchStatus *string `json:"batch_status"`
	// API:API方式发起   WEB:页面方式发起
	BatchType *string `json:"batch_type"`
	// 该笔批量转账的名称
	BatchName *string `json:"batch_name"`
	// 转账说明，UTF8编码，最多允许32个字符
	BatchRemark *string `json:"batch_remark"`
	// 如果批次单状态为“CLOSED”（已关闭），则有关闭原因 * `MERCHANT_REVOCATION` - 商户主动撤销 * `OVERDUE_CLOSE` - 系统超时关闭
	CloseReason *CloseReasonType `json:"close_reason,omitempty"`
	// 转账金额单位为“分”
	TotalAmount *int64 `json:"total_amount"`
	// 一个转账批次单最多发起三千笔转账
	TotalNum *int64 `json:"total_num"`
	// 批次受理成功时返回，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 批次最近一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 转账成功的金额，单位为“分”。当批次状态为“PROCESSING”（转账中）时，转账成功金额随时可能变化
	SuccessAmount *int64 `json:"success_amount,omitempty"`
	// 转账成功的笔数。当批次状态为“PROCESSING”（转账中）时，转账成功笔数随时可能变化
	SuccessNum *int64 `json:"success_num,omitempty"`
	// 转账失败的金额，单位为“分”
	FailAmount *int64 `json:"fail_amount,omitempty"`
	// 转账失败的笔数
	FailNum *int64 `json:"fail_num,omitempty"`
}

func (o TransferBatchGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in TransferBatchGet")
	}
	toSerialize["mchid"] = o.Mchid

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in TransferBatchGet")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo

	if o.BatchId == nil {
		return nil, fmt.Errorf("field `BatchId` is required and must be specified in TransferBatchGet")
	}
	toSerialize["batch_id"] = o.BatchId

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in TransferBatchGet")
	}
	toSerialize["appid"] = o.Appid

	if o.BatchStatus == nil {
		return nil, fmt.Errorf("field `BatchStatus` is required and must be specified in TransferBatchGet")
	}
	toSerialize["batch_status"] = o.BatchStatus

	if o.BatchType == nil {
		return nil, fmt.Errorf("field `BatchType` is required and must be specified in TransferBatchGet")
	}
	toSerialize["batch_type"] = o.BatchType

	if o.BatchName == nil {
		return nil, fmt.Errorf("field `BatchName` is required and must be specified in TransferBatchGet")
	}
	toSerialize["batch_name"] = o.BatchName

	if o.BatchRemark == nil {
		return nil, fmt.Errorf("field `BatchRemark` is required and must be specified in TransferBatchGet")
	}
	toSerialize["batch_remark"] = o.BatchRemark

	if o.CloseReason != nil {
		toSerialize["close_reason"] = o.CloseReason
	}

	if o.TotalAmount == nil {
		return nil, fmt.Errorf("field `TotalAmount` is required and must be specified in TransferBatchGet")
	}
	toSerialize["total_amount"] = o.TotalAmount

	if o.TotalNum == nil {
		return nil, fmt.Errorf("field `TotalNum` is required and must be specified in TransferBatchGet")
	}
	toSerialize["total_num"] = o.TotalNum

	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)
	}

	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime.Format(time.RFC3339)
	}

	if o.SuccessAmount != nil {
		toSerialize["success_amount"] = o.SuccessAmount
	}

	if o.SuccessNum != nil {
		toSerialize["success_num"] = o.SuccessNum
	}

	if o.FailAmount != nil {
		toSerialize["fail_amount"] = o.FailAmount
	}

	if o.FailNum != nil {
		toSerialize["fail_num"] = o.FailNum
	}
	return json.Marshal(toSerialize)
}

func (o TransferBatchGet) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v, ", *o.OutBatchNo)
	}

	if o.BatchId == nil {
		ret += "BatchId:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchId:%v, ", *o.BatchId)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.BatchStatus == nil {
		ret += "BatchStatus:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchStatus:%v, ", *o.BatchStatus)
	}

	if o.BatchType == nil {
		ret += "BatchType:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchType:%v, ", *o.BatchType)
	}

	if o.BatchName == nil {
		ret += "BatchName:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchName:%v, ", *o.BatchName)
	}

	if o.BatchRemark == nil {
		ret += "BatchRemark:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchRemark:%v, ", *o.BatchRemark)
	}

	if o.CloseReason == nil {
		ret += "CloseReason:<nil>, "
	} else {
		ret += fmt.Sprintf("CloseReason:%v, ", *o.CloseReason)
	}

	if o.TotalAmount == nil {
		ret += "TotalAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TotalAmount:%v, ", *o.TotalAmount)
	}

	if o.TotalNum == nil {
		ret += "TotalNum:<nil>, "
	} else {
		ret += fmt.Sprintf("TotalNum:%v, ", *o.TotalNum)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	if o.UpdateTime == nil {
		ret += "UpdateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("UpdateTime:%v, ", *o.UpdateTime)
	}

	if o.SuccessAmount == nil {
		ret += "SuccessAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("SuccessAmount:%v, ", *o.SuccessAmount)
	}

	if o.SuccessNum == nil {
		ret += "SuccessNum:<nil>, "
	} else {
		ret += fmt.Sprintf("SuccessNum:%v, ", *o.SuccessNum)
	}

	if o.FailAmount == nil {
		ret += "FailAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("FailAmount:%v, ", *o.FailAmount)
	}

	if o.FailNum == nil {
		ret += "FailNum:<nil>"
	} else {
		ret += fmt.Sprintf("FailNum:%v", *o.FailNum)
	}

	return fmt.Sprintf("TransferBatchGet{%s}", ret)
}

func (o TransferBatchGet) Clone() *TransferBatchGet {
	ret := TransferBatchGet{}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	if o.BatchId != nil {
		ret.BatchId = new(string)
		*ret.BatchId = *o.BatchId
	}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.BatchStatus != nil {
		ret.BatchStatus = new(string)
		*ret.BatchStatus = *o.BatchStatus
	}

	if o.BatchType != nil {
		ret.BatchType = new(string)
		*ret.BatchType = *o.BatchType
	}

	if o.BatchName != nil {
		ret.BatchName = new(string)
		*ret.BatchName = *o.BatchName
	}

	if o.BatchRemark != nil {
		ret.BatchRemark = new(string)
		*ret.BatchRemark = *o.BatchRemark
	}

	if o.CloseReason != nil {
		ret.CloseReason = new(CloseReasonType)
		*ret.CloseReason = *o.CloseReason
	}

	if o.TotalAmount != nil {
		ret.TotalAmount = new(int64)
		*ret.TotalAmount = *o.TotalAmount
	}

	if o.TotalNum != nil {
		ret.TotalNum = new(int64)
		*ret.TotalNum = *o.TotalNum
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.UpdateTime != nil {
		ret.UpdateTime = new(time.Time)
		*ret.UpdateTime = *o.UpdateTime
	}

	if o.SuccessAmount != nil {
		ret.SuccessAmount = new(int64)
		*ret.SuccessAmount = *o.SuccessAmount
	}

	if o.SuccessNum != nil {
		ret.SuccessNum = new(int64)
		*ret.SuccessNum = *o.SuccessNum
	}

	if o.FailAmount != nil {
		ret.FailAmount = new(int64)
		*ret.FailAmount = *o.FailAmount
	}

	if o.FailNum != nil {
		ret.FailNum = new(int64)
		*ret.FailNum = *o.FailNum
	}

	return &ret
}

// TransferDetailCompact
type TransferDetailCompact struct {
	// 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识
	DetailId *string `json:"detail_id"`
	// 商户系统内部区分转账批次单下不同转账明细单的唯一标识
	OutDetailNo *string `json:"out_detail_no"`
	// PROCESSING:转账中。正在处理中，转账结果尚未明确   SUCCESS:转账成功   FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）
	DetailStatus *string `json:"detail_status"`
}

func (o TransferDetailCompact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.DetailId == nil {
		return nil, fmt.Errorf("field `DetailId` is required and must be specified in TransferDetailCompact")
	}
	toSerialize["detail_id"] = o.DetailId

	if o.OutDetailNo == nil {
		return nil, fmt.Errorf("field `OutDetailNo` is required and must be specified in TransferDetailCompact")
	}
	toSerialize["out_detail_no"] = o.OutDetailNo

	if o.DetailStatus == nil {
		return nil, fmt.Errorf("field `DetailStatus` is required and must be specified in TransferDetailCompact")
	}
	toSerialize["detail_status"] = o.DetailStatus
	return json.Marshal(toSerialize)
}

func (o TransferDetailCompact) String() string {
	var ret string
	if o.DetailId == nil {
		ret += "DetailId:<nil>, "
	} else {
		ret += fmt.Sprintf("DetailId:%v, ", *o.DetailId)
	}

	if o.OutDetailNo == nil {
		ret += "OutDetailNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutDetailNo:%v, ", *o.OutDetailNo)
	}

	if o.DetailStatus == nil {
		ret += "DetailStatus:<nil>"
	} else {
		ret += fmt.Sprintf("DetailStatus:%v", *o.DetailStatus)
	}

	return fmt.Sprintf("TransferDetailCompact{%s}", ret)
}

func (o TransferDetailCompact) Clone() *TransferDetailCompact {
	ret := TransferDetailCompact{}

	if o.DetailId != nil {
		ret.DetailId = new(string)
		*ret.DetailId = *o.DetailId
	}

	if o.OutDetailNo != nil {
		ret.OutDetailNo = new(string)
		*ret.OutDetailNo = *o.OutDetailNo
	}

	if o.DetailStatus != nil {
		ret.DetailStatus = new(string)
		*ret.DetailStatus = *o.DetailStatus
	}

	return &ret
}

// TransferDetailEntity
type TransferDetailEntity struct {
	// 微信支付分配的商户号
	Mchid *string `json:"mchid"`
	// 商户系统内部的商家批次单号，在商户系统内部唯一
	OutBatchNo *string `json:"out_batch_no"`
	// 微信批次单号，微信商家转账系统返回的唯一标识
	BatchId *string `json:"batch_id"`
	// 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	Appid *string `json:"appid"`
	// 商户系统内部区分转账批次单下不同转账明细单的唯一标识
	OutDetailNo *string `json:"out_detail_no"`
	// 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识
	DetailId *string `json:"detail_id"`
	// PROCESSING:转账中。正在处理中，转账结果尚未明确   SUCCESS:转账成功   FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）
	DetailStatus *string `json:"detail_status"`
	// 转账金额单位为“分”
	TransferAmount *int64 `json:"transfer_amount"`
	// 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	TransferRemark *string `json:"transfer_remark"`
	// 如果转账失败则有失败原因 * `ACCOUNT_FROZEN` - 账户冻结 * `REAL_NAME_CHECK_FAIL` - 用户未实名 * `NAME_NOT_CORRECT` - 用户姓名校验失败 * `OPENID_INVALID` - Openid校验失败 * `TRANSFER_QUOTA_EXCEED` - 超过用户单笔收款额度 * `DAY_RECEIVED_QUOTA_EXCEED` - 超过用户单日收款额度 * `MONTH_RECEIVED_QUOTA_EXCEED` - 超过用户单月收款额度 * `DAY_RECEIVED_COUNT_EXCEED` - 超过用户单日收款次数 * `PRODUCT_AUTH_CHECK_FAIL` - 产品权限校验失败 * `OVERDUE_CLOSE` - 转账关闭 * `ID_CARD_NOT_CORRECT` - 用户身份证校验失败 * `ACCOUNT_NOT_EXIST` - 用户账户不存在 * `TRANSFER_RISK` - 转账存在风险
	FailReason *FailReasonType `json:"fail_reason,omitempty"`
	// 商户appid下，某用户的openid
	Openid *string `json:"openid"`
	// 收款方姓名。采用标准RSA算法，公钥由微信侧提供
	UserName **os.File `json:"user_name" encryption:"EM_APIV3"`
	// 转账发起的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	InitiateTime *time.Time `json:"initiate_time"`
	// 明细最后一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime *time.Time `json:"update_time"`
}

func (o TransferDetailEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Mchid == nil {
		return nil, fmt.Errorf("field `Mchid` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["mchid"] = o.Mchid

	if o.OutBatchNo == nil {
		return nil, fmt.Errorf("field `OutBatchNo` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["out_batch_no"] = o.OutBatchNo

	if o.BatchId == nil {
		return nil, fmt.Errorf("field `BatchId` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["batch_id"] = o.BatchId

	if o.Appid == nil {
		return nil, fmt.Errorf("field `Appid` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["appid"] = o.Appid

	if o.OutDetailNo == nil {
		return nil, fmt.Errorf("field `OutDetailNo` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["out_detail_no"] = o.OutDetailNo

	if o.DetailId == nil {
		return nil, fmt.Errorf("field `DetailId` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["detail_id"] = o.DetailId

	if o.DetailStatus == nil {
		return nil, fmt.Errorf("field `DetailStatus` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["detail_status"] = o.DetailStatus

	if o.TransferAmount == nil {
		return nil, fmt.Errorf("field `TransferAmount` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["transfer_amount"] = o.TransferAmount

	if o.TransferRemark == nil {
		return nil, fmt.Errorf("field `TransferRemark` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["transfer_remark"] = o.TransferRemark

	if o.FailReason != nil {
		toSerialize["fail_reason"] = o.FailReason
	}

	if o.Openid == nil {
		return nil, fmt.Errorf("field `Openid` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["openid"] = o.Openid

	if o.UserName == nil {
		return nil, fmt.Errorf("field `UserName` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["user_name"] = o.UserName

	if o.InitiateTime == nil {
		return nil, fmt.Errorf("field `InitiateTime` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["initiate_time"] = o.InitiateTime.Format(time.RFC3339)

	if o.UpdateTime == nil {
		return nil, fmt.Errorf("field `UpdateTime` is required and must be specified in TransferDetailEntity")
	}
	toSerialize["update_time"] = o.UpdateTime.Format(time.RFC3339)
	return json.Marshal(toSerialize)
}

func (o TransferDetailEntity) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.OutBatchNo == nil {
		ret += "OutBatchNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutBatchNo:%v, ", *o.OutBatchNo)
	}

	if o.BatchId == nil {
		ret += "BatchId:<nil>, "
	} else {
		ret += fmt.Sprintf("BatchId:%v, ", *o.BatchId)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}

	if o.OutDetailNo == nil {
		ret += "OutDetailNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutDetailNo:%v, ", *o.OutDetailNo)
	}

	if o.DetailId == nil {
		ret += "DetailId:<nil>, "
	} else {
		ret += fmt.Sprintf("DetailId:%v, ", *o.DetailId)
	}

	if o.DetailStatus == nil {
		ret += "DetailStatus:<nil>, "
	} else {
		ret += fmt.Sprintf("DetailStatus:%v, ", *o.DetailStatus)
	}

	if o.TransferAmount == nil {
		ret += "TransferAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferAmount:%v, ", *o.TransferAmount)
	}

	if o.TransferRemark == nil {
		ret += "TransferRemark:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferRemark:%v, ", *o.TransferRemark)
	}

	if o.FailReason == nil {
		ret += "FailReason:<nil>, "
	} else {
		ret += fmt.Sprintf("FailReason:%v, ", *o.FailReason)
	}

	if o.Openid == nil {
		ret += "Openid:<nil>, "
	} else {
		ret += fmt.Sprintf("Openid:%v, ", *o.Openid)
	}

	if o.UserName == nil {
		ret += "UserName:<nil>, "
	} else {
		ret += fmt.Sprintf("UserName:%v, ", *o.UserName)
	}

	if o.InitiateTime == nil {
		ret += "InitiateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("InitiateTime:%v, ", *o.InitiateTime)
	}

	if o.UpdateTime == nil {
		ret += "UpdateTime:<nil>"
	} else {
		ret += fmt.Sprintf("UpdateTime:%v", *o.UpdateTime)
	}

	return fmt.Sprintf("TransferDetailEntity{%s}", ret)
}

func (o TransferDetailEntity) Clone() *TransferDetailEntity {
	ret := TransferDetailEntity{}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.OutBatchNo != nil {
		ret.OutBatchNo = new(string)
		*ret.OutBatchNo = *o.OutBatchNo
	}

	if o.BatchId != nil {
		ret.BatchId = new(string)
		*ret.BatchId = *o.BatchId
	}

	if o.Appid != nil {
		ret.Appid = new(string)
		*ret.Appid = *o.Appid
	}

	if o.OutDetailNo != nil {
		ret.OutDetailNo = new(string)
		*ret.OutDetailNo = *o.OutDetailNo
	}

	if o.DetailId != nil {
		ret.DetailId = new(string)
		*ret.DetailId = *o.DetailId
	}

	if o.DetailStatus != nil {
		ret.DetailStatus = new(string)
		*ret.DetailStatus = *o.DetailStatus
	}

	if o.TransferAmount != nil {
		ret.TransferAmount = new(int64)
		*ret.TransferAmount = *o.TransferAmount
	}

	if o.TransferRemark != nil {
		ret.TransferRemark = new(string)
		*ret.TransferRemark = *o.TransferRemark
	}

	if o.FailReason != nil {
		ret.FailReason = new(FailReasonType)
		*ret.FailReason = *o.FailReason
	}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	if o.UserName != nil {
		ret.UserName = new(*os.File)
		*ret.UserName = *o.UserName
	}

	if o.InitiateTime != nil {
		ret.InitiateTime = new(time.Time)
		*ret.InitiateTime = *o.InitiateTime
	}

	if o.UpdateTime != nil {
		ret.UpdateTime = new(time.Time)
		*ret.UpdateTime = *o.UpdateTime
	}

	return &ret
}

// TransferDetailInput
type TransferDetailInput struct {
	// 商户系统内部区分转账批次单下不同转账明细单的唯一标识
	OutDetailNo *string `json:"out_detail_no"`
	// 转账金额单位为“分”
	TransferAmount *int64 `json:"transfer_amount"`
	// 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	TransferRemark *string `json:"transfer_remark"`
	// 商户appid下，某用户的openid
	Openid *string `json:"openid"`
	// 收款方姓名。采用标准RSA算法，公钥由微信侧提供
	UserName **os.File `json:"user_name" encryption:"EM_APIV3"`
	// 收款方身份证号，可不用填（采用标准RSA算法，公钥由微信侧提供）
	UserIdCard **os.File `json:"user_id_card,omitempty" encryption:"EM_APIV3"`
}

func (o TransferDetailInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutDetailNo == nil {
		return nil, fmt.Errorf("field `OutDetailNo` is required and must be specified in TransferDetailInput")
	}
	toSerialize["out_detail_no"] = o.OutDetailNo

	if o.TransferAmount == nil {
		return nil, fmt.Errorf("field `TransferAmount` is required and must be specified in TransferDetailInput")
	}
	toSerialize["transfer_amount"] = o.TransferAmount

	if o.TransferRemark == nil {
		return nil, fmt.Errorf("field `TransferRemark` is required and must be specified in TransferDetailInput")
	}
	toSerialize["transfer_remark"] = o.TransferRemark

	if o.Openid == nil {
		return nil, fmt.Errorf("field `Openid` is required and must be specified in TransferDetailInput")
	}
	toSerialize["openid"] = o.Openid

	if o.UserName == nil {
		return nil, fmt.Errorf("field `UserName` is required and must be specified in TransferDetailInput")
	}
	toSerialize["user_name"] = o.UserName

	if o.UserIdCard != nil {
		toSerialize["user_id_card"] = o.UserIdCard
	}
	return json.Marshal(toSerialize)
}

func (o TransferDetailInput) String() string {
	var ret string
	if o.OutDetailNo == nil {
		ret += "OutDetailNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutDetailNo:%v, ", *o.OutDetailNo)
	}

	if o.TransferAmount == nil {
		ret += "TransferAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferAmount:%v, ", *o.TransferAmount)
	}

	if o.TransferRemark == nil {
		ret += "TransferRemark:<nil>, "
	} else {
		ret += fmt.Sprintf("TransferRemark:%v, ", *o.TransferRemark)
	}

	if o.Openid == nil {
		ret += "Openid:<nil>, "
	} else {
		ret += fmt.Sprintf("Openid:%v, ", *o.Openid)
	}

	if o.UserName == nil {
		ret += "UserName:<nil>, "
	} else {
		ret += fmt.Sprintf("UserName:%v, ", *o.UserName)
	}

	if o.UserIdCard == nil {
		ret += "UserIdCard:<nil>"
	} else {
		ret += fmt.Sprintf("UserIdCard:%v", *o.UserIdCard)
	}

	return fmt.Sprintf("TransferDetailInput{%s}", ret)
}

func (o TransferDetailInput) Clone() *TransferDetailInput {
	ret := TransferDetailInput{}

	if o.OutDetailNo != nil {
		ret.OutDetailNo = new(string)
		*ret.OutDetailNo = *o.OutDetailNo
	}

	if o.TransferAmount != nil {
		ret.TransferAmount = new(int64)
		*ret.TransferAmount = *o.TransferAmount
	}

	if o.TransferRemark != nil {
		ret.TransferRemark = new(string)
		*ret.TransferRemark = *o.TransferRemark
	}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	if o.UserName != nil {
		ret.UserName = new(*os.File)
		*ret.UserName = *o.UserName
	}

	if o.UserIdCard != nil {
		ret.UserIdCard = new(*os.File)
		*ret.UserIdCard = *o.UserIdCard
	}

	return &ret
}
