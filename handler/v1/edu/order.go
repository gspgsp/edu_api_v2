package edu

import (
	"github.com/valyala/fasthttp"
	_"github.com/iGoogle-ink/gopay"
	//"github.com/iGoogle-ink/gopay/alipay"
	//"log"
	//"edu_api_v2/response"
)

/**
生成预订单
 */
func GenerateOrder(ctx *fasthttp.RequestCtx) {

	//app_id := "2016091200495858"
	//private_key := "MIIEpQIBAAKCAQEA3pmywRJv1zazvelXJM/LRArpqcWydr7EFdH3o5Ai/fAjjqYLoY+1oMgB+1GJqay9U6m0fz5/ImE8qac2qoIkKKLZGuhzcKUMXYLYAYJe/1cCr9YfDwIpJoAhCqYLk5yMUE0+jWH66lJWSZv7QDXD4ILOC+A3MVALkbl38a3T4NLxno5AmmWxwd+9+0EvPffU0Y76lIvDLnSFJm6w7ZTSLXI2p49ixEUs7ZkbpcyeaBJXMfa2BNBKMeMqbtOxX5D22+t+jeAGqe+Q/hj3GgM1DTXn8HRVn6EAtb9Gzcz2NxdV2LSH1KRMI36CILmF8fy4t1alMQNsbkFI78k8gZI9vwIDAQABAoIBAQDU/rJ5Zd/YSUJhxMWatq3jtKFyznV+g/jyfA991Wx2GXKciytz8yOy7c5foNydm2kphaftXWqfBfXay+JV8uMQcwpUYfob3gHf9FQl1Y56utIaWEMtxD6F1XJIUyuemdv09oSVGhzSW+iu1G7sOMXtrJvl5yxIpjP7w0EO4uprIwQyG+MGbC24Nd3lc0d55ea2WBSHy5GqgWwtKlCt11/MC2t0fZ01CbZYutrFCQ8RkYCzvDBqTvmILlKQbiemexlhNgIV8umjNInT5YkGAXkDXQhBVp53kIdh0MmuS9arxDOgyLwFPUA0/CjPkNdDYjSIHs9HEUiD8gQb3nHMikQBAoGBAO9v4ONxQBqB+xT+MyJ04tZW4uYl4mqKFclSCzF8ZcB3CBWmSIhJlauuQTOdOQr4GlMASW2EYYlyjAL7wSGVCPKz092joVy0SUOg4nZWi96Gj9I9zXLE77IRyxd+3S0CFfesdJ9laNppoIJKew89ZW9XpLk48qwqFOJBpIWVOSO/AoGBAO3/qh855VioTQuawE9hTqJjeRmWurb7Ex8qyAndoqRXNIR50fC2j8uzErRW099bu/8ge4PCT93UVii3TlfgbyOTwjIwjQ4KvQISUyc6Lcg7d9IwzRK/ZNpqIXn3S3BcRReok3p/OL52A1Iu3XOSx28miWHldK8+/aM3xgFipWYBAoGBAMIgdDYTaQTU48xWeQDzwcJEObyddKAqipVqNncNw7WzDFP5OI9/EaFbK7P8QfZ5x2YKMn5fuXDl9Uyp7yEaI2IbmKf3demipk/1jgaiDs8BTHQJ9ZuavjgiowXOruZB7aucIAXoA4Yp46AfYPyvK6GFbqFZ4a7ig4IZ9Jj1rpkjAoGAfxPhEVVHiWQr50NU5c2uP/TfJ3Oc9LDxD9Y3A1WNTzSk/QtStrWz5+8Q0Kv/6yY/JkWAzeSF4otLCeh+c1uJ3DQ8H7Jezp2bOCzyq2JcKarfGqMmDJmXTywV/dq2NCOAXNaNEH43nTLEKTCO2/QGeunXtybee+glY+4W0oyKhgECgYEA2/lixdQ55PXt7zj6XPdGMSmvtnlXVZVAMDwHdLdEEK+3jCXkivvc3/QyugjEJVR+nqWbJvwX4H/J0qRgnQU73XpB0M0mBbFXMAfkhzGdU4Y1hWbkI/rHN3cAG0YSZJ1bENCK9Ooc5IW8n1HPuftqRz8t/lRKVAmxplm3YAllgtA="
	//notify_url := "http://api.gsplovedss.xyz/api/v1/notify"
	//
	//client := alipay.NewClient(app_id, private_key, false)
	//client.SetCharset("UTF-8")
	//client.SetSignType("RSA2")
	//client.SetNotifyUrl(notify_url)
	//
	//body := make(gopay.BodyMap)
	//body.Set("subject", "手机网站测试支付")
	//body.Set("out_trade_no", "GZ201901301040355703")
	//body.Set("quit_url", "https://www.gopay.ink")
	//body.Set("total_amount", "100.00")
	//body.Set("product_code", "QUICK_WAP_WAY")
	//
	//url, err := client.TradeWapPay(body)
	//if err != nil {
	//	log.Printf("url err is:%s", err.Error())
	//	return
	//}
	//
	//response.New(ctx).SetData(url).JsonReturn()
}
