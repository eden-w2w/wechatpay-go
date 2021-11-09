package jsapi

import (
	"context"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/option"
	"github.com/eden-w2w/wechatpay-go/utils"
	"log"
	"testing"
)

var (
	mchID                      string = "1615730640"                               // 商户号
	mchCertificateSerialNumber string = "6BA78EF54BC73369F56DFB1ABE1F4EC362C23F3F" // 商户证书序列号
	mchAPIv3Key                string = "7e577d434ffb7118b51058fd3d9056f7"         // 商户APIv3密钥
	mchPK                      string = "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCc74WajjWEWpwR\njAbmAxLStEUG637YDOQI98lrorP+ENEvIG6qm4hjcSu0gf3+W/2UIMgn5sX+WCNt\nOMpU76vAqNyozUnrXfngdoBrWfg+cQ0JMzHAByDnckvcMMWSyfL15xCt7Apk+VpK\nK2WUQZwHVIjZFZY8NK4cFfwPl9NqZWlkOB81QuCmCFtkr7AegiX15HSI0I2aGhCp\nAcRTYcdM7T7j+ZK071Oik7dXaBOR0KA1qXf0Z8eDQBtqzkUq6Zy2MqW0jDgW7/AS\n9Lrgmy5bm3r4KhAjf8Lj24QZrVNE1kzGunDub5a669FcryeThTF37qnLKnDeuRIL\nnVyZ5S7BAgMBAAECggEAK7Q8hcUyvDArpGtVhHq4pDsVug6dHXfBjYnL6xC9hXST\nfMdQamyz99WNcVB+NnbR0I2skAghfWp4OzOEjVDNiPK9uLiq7EQYkV0O+FZu3pzH\nH5fFNBAIJsxaufBPJUjeQ4Lcz0KbxnD2pw8c08PNuvDcx2/Ojeex0RqCPqludANo\nOeZpts759sFWfvsssBfxAOvWKAG8Y24v56sfObFKINyw1UnYE6hsqTIGlfOUSN2M\nXjuPAnrN33HkGv75P9geG4UVw2tzB+vU0MpIoFuiRyRrzJbcbGp2NulobC4JJ/Bq\nnhOHgxiKM7xk90pvncNepuz34FDb5vo+OQPAWNE1rQKBgQDPMDb/mOQmlebMfHk6\n7a412ePzPQIaSgPIRN/TpxMVLS5c+ftuhal8mCHIsuL7mf0kc/tV3pthUE/7sDdx\natWgcO/ilhcPY++XxsZdau5stDQ+v+RTD7gSxMrImbORpMY4sdOI9ug1AJUjWmqM\nbYNy3AuzdfME/zVa12ehKL0ynwKBgQDB6IENPzIKiZyKSbA6Qtfgx7tzbQ0S5DbV\neRaGISDC13R2Ijris4lxojQbG33PfHpwaX7uHa8bTqwEViEcu8CyLEbTcFss41M0\nMscKAZnPxJGf4P1ODXg7eH2L23NrKlt6UJjlUjXTpfar7/UlgRwmoFMGFK8F/GhG\n2Nc+ZpKCnwKBgGazea+9M+rEh4F5egZx5tceyNW7uh0Z70IIkgKyYM0wdjj0WWPx\nsY/nQWfg9I7PSCIXs9Be5hqY2uzh7rmzwW0kTmS9DXU9jnhnudB4vvL3aKZqLF3X\n/uCe1/4T4mUhjnhv/XF5IqWTQloh6YruhigbV1l/8BTcBLCg8ed68D31AoGAJ1Dl\nGbJ+ivRr8P4P5UA4VTMqJrUtuW7cT1xF3NxKsDbPPkSie+S7MXcZ6YUhdThf+vro\n6Y7LOlYxOo/cFt9aqsWfdq+JTQiqjiLdiIHale3dOPV9Zp4EispbakgMluX3tk0x\nMw4AK/GymhoFMp2C/7HyI4F+G1CaNAMzYUua2TECgYAfvS4JapxzIEKqL7I+mk92\nktmtAlgvH4B6UR3ex+r8G4Xr3gUKTHw9s5GGTpJSL5LdFVcpf/nzQD6fi1scvNAv\nTSTLIdxUuX9auOiZzBam6S3sE1d70OXbYBS2KIFwq/+ga6lu2xnCcwAHiGaRuUyM\nm1rPBK+s06SpeD/OEwRNYA==\n-----END PRIVATE KEY-----\n"

	client *core.Client
)

func init() {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey(mchPK)
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}
}

func TestTradeBill(t *testing.T) {
	svc := JsapiApiService{Client: client}
	resp, _, err := svc.TradeBill(
		context.Background(), TradeBillRequest{
			BillDate: core.String("2021-11-08"),
			BillType: nil,
			TarType:  nil,
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
	testDownloadURL(*resp, t)
}

func testDownloadURL(bill BillResponse, t *testing.T) {
	svc := JsapiApiService{Client: client}
	data, _, err := svc.DownloadURL(context.Background(), bill)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(data))
}
