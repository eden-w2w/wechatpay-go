// Copyright 2021 Tencent Inc. All rights reserved.
//
// 境内普通商户退款API
//
// 境内普通商户退款功能涉及的API文档
//
// API version: 1.1.1

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package refunddomestic_test

import (
	"context"
	"log"

	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/option"
	"github.com/eden-w2w/wechatpay-go/services/refunddomestic"
	"github.com/eden-w2w/wechatpay-go/utils"
)

func ExampleRefundsApiService_Create() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, result, err := svc.Create(ctx,
		refunddomestic.CreateRequest{
			SubMchid:      core.String("1900000109"),
			TransactionId: core.String("1217752501201407033233368018"),
			OutTradeNo:    core.String("1217752501201407033233368018"),
			OutRefundNo:   core.String("1217752501201407033233368018"),
			Reason:        core.String("商品已售完"),
			NotifyUrl:     core.String("https://weixin.qq.com"),
			FundsAccount:  refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				From: []refunddomestic.FundsFromItem{refunddomestic.FundsFromItem{
					Account: refunddomestic.ACCOUNT_AVAILABLE.Ptr(),
					Amount:  core.Int64(444),
				}},
				Refund: core.Int64(888),
				Total:  core.Int64(888),
			},
			GoodsDetail: []refunddomestic.GoodsDetail{refunddomestic.GoodsDetail{
				GoodsName:        core.String("iPhone6s 16G"),
				MerchantGoodsId:  core.String("1217752501201407033233368018"),
				RefundAmount:     core.Int64(528800),
				RefundQuantity:   core.Int64(1),
				UnitPrice:        core.Int64(528800),
				WechatpayGoodsId: core.String("1001"),
			}},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call Create err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleRefundsApiService_QueryByOutRefundNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, result, err := svc.QueryByOutRefundNo(ctx,
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: core.String("1217752501201407033233368018"),
			SubMchid:    core.String("1900000109"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryByOutRefundNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
