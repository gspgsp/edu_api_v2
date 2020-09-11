package edu

import (
	"github.com/valyala/fasthttp"
	_ "github.com/smartwalle/alipay"
	"edu_api_v2/response"
	"github.com/smartwalle/alipay"
	"log"
	"time"
	"strconv"
	"edu_api_v2/config"
)

/**
生成预订单
 */
func GenerateOrder(ctx *fasthttp.RequestCtx) {

	client, err := alipay.New(config.Config.Alipay.AppId, config.Config.Alipay.PrivateKey, config.Config.Alipay.IsProduction)
	client.LoadAliPayPublicKey(config.Config.Alipay.AliPubliceKey)

	if err != nil {
		log.Printf("init client err is:%s", err.Error())
		return
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = config.Config.Alipay.NotifyUrl
	p.ReturnURL = config.Config.Alipay.ReturnUrl
	p.Subject = "go手机网站测试支付"
	p.OutTradeNo = strconv.FormatInt(time.Now().Unix(), 10)
	p.TotalAmount = "110"
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		log.Printf("trade err is:%s", err.Error())
		return
	}

	var payUrl = url.String()

	response.New(ctx).SetData(payUrl).JsonReturn()
}
