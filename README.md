# alipay
支付宝 AliPay SDK for Go, 集成简单，功能完善中，持续更新！

## 目录结构
alipay/  
├── internal/  
│   ├── config/             # 配置管理  
│   └── util/               # 工具函数  
├── pkg/  
│   ├── alipay/             # 支付宝支付功能实现  
│   ├── client/             # HTTP 客户端封装  
│   └── model/              # 模型定义  
├── cmd/  
│   └── example/            # 示例代码  
├── docs/                   # SDK 文档  
├── tests/                  # 单元测试和集成测试  
├── LICENSE                 # 许可证文件  
├── README.md               # SDK 使用文档  
└── go.mod                  # Go modules 文件  

* **internal** 目录包含了一些供 SDK 内部使用的功能，例如配置管理和工具函数。
* **pkg** 目录包含了公共的支付宝支付功能实现、HTTP 客户端封装以及模型定义等。
* **cmd** 目录包含了示例代码，用于演示 SDK 的使用方法。
* **docs** 目录用于存放 SDK 的文档，可以包含使用指南、API 文档等。
* **tests** 目录包含了单元测试和集成测试代码。
* **LICENSE** 文件包含了 SDK 的许可证信息。
* **README.md** 文件用于介绍 SDK，并提供使用指南和其他相关信息。
* **go.mod** 文件是 Go modules 的配置文件，用于管理 SDK 的依赖项。

## 下载

#### 启用 Go module【推荐】
在您的项目中的go.mod文件内添加这行代码  
```go
require require github.com/leiphp/alipay v1.0.3
```

并且在项目中使用 "github.com/leiphp/alipay" 引用 AliPay Go SDK。

例如
```go
import (
    "github.com/leiphp/alipay/pkg/alipay"
    "github.com/leiphp/alipay/pkg/xid"
)
```

## 帮助
在集成的过程中有遇到问题，欢迎加 QQ  1124378213 讨论  
支持公钥证书加签方式和普通密钥加签方式，该项目以**采用证书加密**为例  

## 加签方式说明
主要分为以下两种加签方式：[详情](https://opendocs.alipay.com/common/02kf5p?pathHash=81acb065)

| 接口加签方式 | 密钥 | 证书 |
| --------- | --------- | --------- |
| 特性 | 安全性、真实性 | 安全性、真实性、抗抵赖 |
| 相应文件 | 应用私钥、应用公钥、支付宝公钥 | 应用私钥、应用公钥、应用公钥证书、支付宝公钥证书、支付宝根证书 |
| 适用场景 | 除资金支出类场景外，都可以使用密钥加签方式 | 红包（原 “现金红包”）、转账到支付宝账户 必须使用证书加签方式，其他场景可以选择使用密钥或证书 |

注意：需先完成配置后才可对接口进行加签。

## 设置证书加签方式
在接口调用前，您需要在支付宝开放平台配置接口加签方式。[详情](https://opendocs.alipay.com/common/056zub?pathHash=91c49771) 
1. 设置加签方式，选择 证书 > 下一步。  
2. 生成 CSR 文件。  
    a. 下载并安装 [支付宝开放平台密钥工具](https://opendocs.alipay.com/common/02kipk)。  
    b. 打开密钥工具，进入 [生成密钥](https://opendocs.alipay.com/common/02kipl) 功能。  
    c. 加签方式选择 证书，加签算法选择 RSA2。（开放平台支持 SM2 算法，但暂未公开配置，如有需要请联系支付宝技术支持）  
    d. 填写 组织/公司，必须与开放平台主账号名称完全相同。  
    e. 点击 生成CSR文件，可以点击 打开文件位置 查看生成的应用私钥、应用公钥和 CSR文件。  
    f. 返回开放平台控制台中，点击 下一步。
3. 上传生成的 CSR 文件，选择 证书到期后的处理方式（默认为自动签发），设置 安全联系人 信息，签署 开放平台服务协议。点击 确认上传。此时需要输入短信验证码或支付密码，完成安全验证。  
4. 证书配置完成。

## 下载证书
配置完成接口加签方式就可以在最后一步下载证书，点击全部下载到本地
* **应用公钥证书**
```go
appCertPublicKey_2017020905582002.crt
```
* **支付宝公钥证书**
 ```go
 alipayCertPublicKey_RSA2.crt
 ```
* **支付宝根证书**
 ```go
 alipayRootCert.crt
 ```

## 项目初始化

**下面用到的 privateKey 需要特别注意一下，如果是通过“支付宝开发平台开发助手”创建的CSR文件，在 CSR 文件所在的目录下会生成相应的私钥文件，我们需要使用该私钥进行签名。**

```go
var client, err = alipay.New(appID, privateKey, isProduction)
```

#### 关于应用私钥 (privateKey)

应用私钥是我们通过工具生成的私钥，调用支付宝接口的时候，我们需要使用该私钥对参数进行签名。

#### 关于 alipay.New() 函数中的最后一个参数 isProduction

支付宝提供了用于开发时测试的 sandbox 环境，对接的时候需要注意相关的 app id 和密钥是 sandbox 环境还是 production 环境的。如果是 sandbox 环境，本参数应该传 false，否则为 true。

### 公钥证书加签方式

如果采用公钥证书方式进行验证签名，需要调用以下几个方法加载证书信息，所有证书都是从支付宝创建的应用处下载

```go
client.LoadAppCertPublicKeyFromFile("/路径/appCertPublicKey_2017020905582002.crt") // 加载应用公钥证书
client.LoadAliPayRootCertFromFile("/路径/alipayRootCert.crt")                // 加载支付宝根证书
client.LoadAlipayCertPublicKeyFromFile("/路径/alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
```

### 普通密钥加签方式

需要注意此处用到的公钥是**支付宝公钥**，不是我们用工具生成的应用公钥。

[如何查看支付宝公钥？](https://opendocs.alipay.com/common/057aqe)

```go
client.LoadAliPayPublicKey("aliPublicKey")
```

特别注意：**证书加签方式**和**普通密钥加签方式**不能同时存在，只能选择其中一种。

### 接口内容加密

详细内容访问 [接口内容加密方式](https://opendocs.alipay.com/common/02mse3) 进行了解。

如果需要开启该功能，只需调用一下 SetEncryptKey() 方法。

```go
client.SetEncryptKey("key")
```

如果不需要开启该功能，则不用调用该方法。

### 签名验证

内部已实现对支付宝返回的数据进行签名验证，详细信息请参考[自行实现验签](https://doc.open.alipay.com/docs/doc.htm?docType=1&articleId=106120)。

需要自行对签名验证的场景有 **同步回调(return_url)** 和 **异步通知(notify_url)** 的 HTTP 处理函数。

#### 同步回调(return_url)

发起支付(网页支付)的时候，如果有提供 ReturnURL 参数，那么支付成功之后，支付宝会将浏览器重定向到该 URL，并附带上相关的参数。

```go
var p = alipay.TradeWapPay{}
p.ReturnURL = "http://xxx/return"
```

对支付宝提供的参数进行签名验证：

```go
http.HandleFunc("/return", func (writer http.ResponseWriter, request *http.Request) {
request.ParseForm()
if err := client.VerifySign(request.Form); err != nil {
// 如果 err 不为空，则表示验签失败
fmt.Println(err)
return
}
// 业务处理
}
```

#### 异步通知(notify_url)

有支付或者其它动作发生后，支付宝服务器会调用我们提供的 NotifyURL，并向其传递相关的信息。参考[手机网站支付结果异步通知](https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.XM5C4a&treeId=203&articleId=105286&docType=1)。

```go
var p = alipay.TradeWapPay{}
p.NotifyURL = "http://xxx/return"
```

解析通知并验证签名：

```go
http.HandleFunc("/notify", func (writer http.ResponseWriter, request *http.Request) {
request.ParseForm()

// DecodeNotification 内部已调用 VerifySign 方法验证签名
var noti, err = client.DecodeNotification(request.Form)
if err != nil {
// 错误处理
fmt.Println(err)
return
}
// 业务处理
// 如果通知消息没有问题，我们需要确认收到通知消息，不然支付宝后续会继续推送相同的消息
alipay.ACKNotification(writer)
})
```

#### 支持 RSA2 签名及验证

采用 RSA2 签名，不再提供 RSA 的支持。

#### 特别注意

提供给支付宝的 NotifyURL 和 ReturnURL 最好不要附带任何参数，支付宝在生成签名信息的时候不会包含 URL 中的参数，而 VerifySign() 方法在验证签名的时候会将收到的所有参数一起验证。

如果确实需要附带参数，可以在调用 VerifySign() 方法前，将附带的参数从 request.Form 中删除。

## 已实现接口

```
中线(-)后面的名称是该接口在 Client 结构体中对应的方法名。
```

* **手机网站支付接口**

  alipay.trade.wap.pay - **TradeWapPay()**

* **电脑网站支付**

  alipay.trade.page.pay - **TradePagePay()**

* **统一收单线下交易查询**

  alipay.trade.query - **TradeQuery()**

* **统一收单交易支付接口**

  alipay.trade.pay - **TradePay()**

* **统一收单交易创建接口**

  alipay.trade.create - **TradeCreate()**

* **统一收单线下交易预创建**

  alipay.trade.precreate - **TradePreCreate()**

* **统一收单交易撤销接口**

  alipay.trade.cancel - **TradeCancel()**

* **统一收单交易关闭接口**

  alipay.trade.close - **TradeClose()**

* **统一收单交易退款接口**

  alipay.trade.refund - **TradeRefund()**

* **App 支付接口**

  alipay.trade.app.pay - **TradeAppPay()**

* **统一收单交易退款查询**

  alipay.trade.fastpay.refund.query - **TradeFastpayRefundQuery()**

* **支付宝订单信息同步接口**

  alipay.trade.orderinfo.sync - **TradeOrderInfoSync()**

* **单笔转账到支付宝账户接口**

  alipay.fund.trans.toaccount.transfer - **FundTransToAccountTransfer()**

* **查询转账订单接口**

  alipay.fund.trans.order.query - **FundTransOrderQuery()**

* **资金授权发码接口**

  alipay.fund.auth.order.voucher.create - **FundAuthOrderVoucherCreate()**

* **资金授权操作查询接口**

  alipay.fund.auth.operation.detail.query - **FundAuthOperationDetailQuery()**

* **资金授权撤销接口**

  alipay.fund.auth.operation.cancel - **FundAuthOperationCancel()**

* **资金授权解冻接口**

  alipay.fund.auth.order.unfreeze - **FundAuthOrderUnfreeze()**

* **资金授权冻结接口**

  alipay.fund.auth.order.freeze - **FundAuthOrderFreeze()**

* **线上资金授权冻结接口**

  alipay.fund.auth.order.app.freeze - **FundAuthOrderAppFreeze()**

* **查询对账单下载地址**

  alipay.data.dataservice.bill.downloadurl.query - **BillDownloadURLQuery()**

* **支付宝商家账户当前余额查询**

  alipay.data.bill.balance.query - **BillBalanceQuery()**

* **身份认证初始化服务**

  alipay.user.certify.open.initialize - **UserCertifyOpenInitialize()**

* **身份认证开始认证**

  alipay.user.certify.open.certify - **UserCertifyOpenCertify()**

* **身份认证记录查询**

  alipay.user.certify.open.query - **UserCertifyOpenQuery()**

* **用户信息授权(网站支付宝登录快速接入)**

  生成授权链接 - **PublicAppAuthorize()**

* **换取授权访问令牌**

  alipay.system.oauth.token - **SystemOauthToken()**

* **支付宝会员授权信息查询**

  alipay.user.info.share - **UserInfoShare()**

* **App支付宝登录**

  com.alipay.account.auth - **AccountAuth()**

* **支付宝个人协议页面签约**

  alipay.user.agreement.page.sign - **AgreementPageSign()**

* **支付宝个人代扣协议查询**

  alipay.user.agreement.query - **AgreementQuery()**

* **支付宝个人代扣协议解约**

  alipay.user.agreement.unsign - **AgreementUnsign()**

* **支单笔转账接口**

  alipay.fund.trans.uni.transfer - **FundTransUniTransfer()**

* **转账业务单据查询接口**

  alipay.fund.trans.common.query - **FundTransCommonQuery()**

* **支付宝资金账户资产查询接口**

  alipay.fund.account.query - **FundAccountQuery()**

* **小程序获取会员手机号数据解析**

  my.getPhoneNumber - **DecodePhoneNumber()**

## 集成流程

从[支付宝开放平台](https://open.alipay.com/)申请创建相关的应用，使用自己的支付宝账号登录即可。

### 沙箱环境

支付宝开放平台为每一个应用提供了沙箱环境，供开发人员开发测试使用。

沙箱环境是独立的，每一个应用都会有一个商家账号和买家账号。

#### 沙箱环境网关地址

沙箱环境目前有两个网关地址：

* 老地址: [https://openapi.alipaydev.com/gateway.do](https://openapi.alipaydev.com/gateway.do)
* 新地址: [https://openapi-sandbox.dl.alipaydev.com/gateway.do](https://openapi-sandbox.dl.alipaydev.com/gateway.do)

大家在对接的时候一定要确认清楚是新地址还是老地址。

本 SDK 目前默认使用的是 **【新地址】**，如果需要使用老地址，只需要在初始化的时候通过 alipay.WithPastSandboxGateway() 指定即可。

```go
alipay.New(appId, privateKey, isProduction, alipay.WithPastSandboxGateway())
```

### 应用信息配置

参考[官网文档](https://docs.open.alipay.com/200/105894) 进行应用的配置。

本 SDK 中的签名方法默认为 **RSA2**，采用支付宝提供的 [RSA签名&验签工具](https://docs.open.alipay.com/291/105971) 生成秘钥时，秘钥长度推荐 **2048**。所以在支付宝管理后台请注意配置 **RSA2(SHA256)密钥**。

生成秘钥对之后，将公钥提供给支付宝（通过支付宝后台上传）对我们请求的数据进行签名验证，我们的代码中将使用私钥对请求数据签名。

请参考 [如何生成 RSA 密钥](https://docs.open.alipay.com/291/105971)。

### 创建 Wap 支付

```go
var privateKey = "xxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
var client, err = alipay.New(appId, privateKey, false)
if err != nil {
fmt.Println(err)
return
}

// 加载应用公钥证书
if err = client.LoadAppCertPublicKeyFromFile("appCertPublicKey_2017011104995404.crt"); err != nil {
// 错误处理
}

// 加载支付宝根证书
if err = client.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
// 错误处理
}

// 加载支付宝公钥证书
if err = client.LoadAlipayCertPublicKeyFromFile("alipayCertPublicKey_RSA2.crt"); err != nil {
// 错误处理
}

// 加载内容密钥，可选
if err = client.SetEncryptKey("FtVd5SgrsUzYQRAPBmejHQ=="); err != nil {
// 错误处理
}

var p = alipay.TradeWapPay{}
p.NotifyURL = "http://xxx"
p.ReturnURL = "http://xxx"
p.Subject = "标题"
p.OutTradeNo = "传递一个唯一单号"
p.TotalAmount = "10.00"
p.ProductCode = "QUICK_WAP_WAY"

var url, err = client.TradeWapPay(p)
if err != nil {
fmt.Println(err)
}

// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
var payURL = url.String()
fmt.Println(payURL)
```

## 自定义请求

对于本库还未实现接口，可使用 alipay.Payload 结构体作为参数调用 alipay.Client 结构体的 Request() 方法。

```go
var p = alipay.NewPayload("这里是接口名称，如：alipay.trade.query")
// 添加公共请求参数，如：app_auth_token
p.AddParam("key", "value")
// 添加请求参数(业务相关)
p.AddBizField("key", "value")

// 应用场景一：需要向支付宝服务器发起网络请求，如交易查询接口（alipay.trade.query）
var result map[string]interface{}
// result 也可以为结构体，可参照 alipay.TradeQueryRsp
var err = client.Request(p, &result)
if err != nil {
...
}

// 应用场景二：只需要生成 URL，如网页支付（alipay.trade.page.pay）
var url, err = clietn.BuildURL(p)
...

// 应用场景三：只需要对参数进行签名，如 App 支付（alipay.trade.app.pay）
var s, err = client.EncodeParam(p)
...
```

[更多信息](https://github.com/leiphp/alipay/)

## 致谢
项目起初参考了很多开源项目的解决方案，开源不易，感谢分享
- 感谢**smartwalle**分享：[alipay](https://github.com/smartwalle/alipay)

## 示例

[网页支付](https://www.100txy.com/alidonate)

## 感谢你的支持

如果对你有帮助，请我喝杯咖啡吧！

![支付宝](https://www.100txy.com/static/images/icon/alipayimg.jpg)
![微信](https://www.100txy.com/static/images/icon/weipayimg.jpg)

## License

This project is licensed under the MIT License.
