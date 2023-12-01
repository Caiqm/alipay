package alipay

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"hash"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrNullParams           = errors.New("alipay: bad params")
	ErrBadResponse          = errors.New("alipay: bad response")
	ErrSignNotFound         = errors.New("alipay: sign content not found")
	ErrAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

type Client struct {
	appId            string
	pubKey           string
	priKey           string
	host             string
	notifyVerifyHost string
	encryptIV        []byte
	encryptType      string
	encryptKey       []byte
	location         *time.Location
	Client           *http.Client
	onReceivedData   func(method string, data []byte)
	isProduction     bool
}

type OptionFunc func(c *Client)

// 初始化客户端
func New(appId, privateKey string, isProduction bool, opts ...OptionFunc) (nClient *Client, err error) {
	if appId == "" || privateKey == "" {
		return nil, ErrNullParams
	}
	nClient = &Client{}
	nClient.isProduction = isProduction
	nClient.appId = appId
	nClient.priKey = privateKey
	if nClient.isProduction {
		nClient.host = kProductionGateway
		nClient.notifyVerifyHost = kProductionMAPIGateway
	} else {
		nClient.host = kNewSandboxGateway
		nClient.notifyVerifyHost = kNewSandboxGateway
	}
	nClient.Client = http.DefaultClient
	nClient.location = time.Local
	for _, opt := range opts {
		if opt != nil {
			opt(nClient)
		}
	}
	return
}

// 加载支付宝公钥
func (c *Client) LoadAlipayCertPublicKey(s string) (err error) {
	if s == "" {
		err = ErrAliPublicKeyNotFound
		return
	}
	c.pubKey = s
	return
}

// SetEncryptKey 接口内容加密密钥 https://opendocs.alipay.com/common/02mse3
func (c *Client) SetEncryptKey(key string) error {
	if key == "" {
		return nil
	}
	var data, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	c.encryptIV = data[:16]
	c.encryptType = "AES"
	c.encryptKey = data
	return nil
}

// 请求参数
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	// 公共参数
	values.Add(kFieldAppId, c.appId)
	values.Add(kFieldMethod, param.APIName())
	values.Add(kFieldFormat, kFormat)
	values.Add(kFieldCharset, kCharset)
	values.Add(kFieldSignType, kSignTypeRSA2)
	values.Add(kFieldTimestamp, time.Now().In(c.location).Format(kTimeFormat))
	values.Add(kFieldVersion, kVersion)
	// 业务参数
	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	var content = string(jsonBytes)
	// 添加biz_content业务参数
	if content != kEmptyBizContent {
		values.Add(kFieldBizContent, content)
	}
	// 没有bizContent参数，公共方法添加参数
	var params = param.Params()
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v)
	}
	// 生成签名
	signContentBytes, _ := url.QueryUnescape(values.Encode())
	signature, err := c.sign([]byte(signContentBytes), kSignTypeRSA2)
	if err != nil {
		return nil, err
	}
	// 添加签名
	values.Add(kFieldSign, signature)
	return values, nil
}

// 生成签名
func (c *Client) sign(data []byte, SignType string) (signature string, err error) {
	if len(data) == 0 {
		err = ErrSignNotFound
		return
	}
	var h hash.Hash
	var hType crypto.Hash
	switch SignType {
	case "RSA":
		h = sha1.New()
		hType = crypto.SHA1
	case "RSA2":
		h = sha256.New()
		hType = crypto.SHA256
	}
	h.Write(data)
	d := h.Sum(nil)
	pk, err := c.parsePrivateKey()
	if err != nil {
		return
	}
	bs, err := rsa.SignPKCS1v15(rand.Reader, pk, hType, d)
	if err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(bs)
	return
}

// 验证签名内容
func (c *Client) parsePrivateKey() (pk *rsa.PrivateKey, err error) {
	encodedKey, _ := base64.StdEncoding.DecodeString(c.priKey)
	pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(encodedKey)
	if err != nil {
		pkcs8PrivateKey, err = x509.ParsePKCS1PrivateKey(encodedKey)
		if err != nil {
			err = fmt.Errorf("解析密钥失败，err: %v", err)
			return
		}
	}
	pk = pkcs8PrivateKey.(*rsa.PrivateKey)
	return
}

// 请求主方法
func (c *Client) doRequest(method string, param Param, result interface{}) (err error) {
	// 创建一个请求
	req, _ := http.NewRequest(method, c.host, nil)
	// 判断参数是否为空
	if param != nil {
		var values url.Values
		values, err = c.URLValues(param)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(strings.NewReader(values.Encode()))
	}
	// 添加header头
	req.Header.Add("Content-Type", kContentType)
	// 发起请求数据
	rsp, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	var apiName = param.APIName()
	var bizFieldName = strings.Replace(apiName, ".", "_", -1) + kResponseSuffix
	return c.decode(bodyBytes, bizFieldName, param.NeedVerify(), result)
}

// 解密返回数据
func (c *Client) decode(data []byte, bizFieldName string, needVerifySign bool, result interface{}) (err error) {
	var raw = make(map[string]json.RawMessage)
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var signBytes = raw[kFieldSign]
	var certBytes = raw[kFieldAlyPayCertSN]
	var bizBytes = raw[bizFieldName]
	var errBytes = raw[kErrorResponse]
	if len(certBytes) > 1 {
		certBytes = certBytes[1 : len(certBytes)-1]
	}
	if len(signBytes) > 1 {
		signBytes = signBytes[1 : len(signBytes)-1]
	}
	// 判断是否有错误
	if len(bizBytes) == 0 {
		if len(errBytes) > 0 {
			var rErr *Error
			if err = json.Unmarshal(errBytes, &rErr); err != nil {
				return err
			}
			return rErr
		}
		return ErrBadResponse
	}
	// 数据解密
	var plaintext []byte
	if plaintext, err = c.decrypt(bizBytes); err != nil {
		return err
	}
	// 验证签名
	if needVerifySign {
		if c.onReceivedData != nil {
			c.onReceivedData(bizFieldName, plaintext)
		}
		if len(signBytes) == 0 {
			// 没有签名数据，返回的内容一般为错误信息
			var rErr *Error
			if err = json.Unmarshal(plaintext, &rErr); err != nil {
				return err
			}
			return rErr
		}
		fmt.Println(string(certBytes))
		// 验证签名
		if err = c.Verify(bizBytes, signBytes); err != nil {
			return err
		}
	}
	// 返回数据
	if err = json.Unmarshal(plaintext, result); err != nil {
		return err
	}
	return nil
}

// 解码数据
func (c *Client) decrypt(data []byte) ([]byte, error) {
	var plaintext = data
	if len(data) > 1 && data[0] == '"' {
		var ciphertext, err = base64decode(data[1 : len(data)-1])
		if err != nil {
			return nil, err
		}
		block, err := aes.NewCipher(c.encryptKey)
		if err != nil {
			return nil, err
		}
		var mode = cipher.NewCBCDecrypter(block, c.encryptIV)
		mode.CryptBlocks(ciphertext, ciphertext)
		// 去除填充
		padding := int(ciphertext[len(ciphertext)-1])
		plaintext = ciphertext[:len(ciphertext)-padding]
		if err != nil {
			return nil, err
		}
	}
	return plaintext, nil
}

// 验证密钥
func (c *Client) Verify(signContent, sign []byte) (err error) {
	// 步骤1，加载RSA的公钥
	aliPublicKey := c.formatAlipayPublicKey(c.pubKey)
	block, _ := pem.Decode([]byte(aliPublicKey))
	// keyByts, _ := base64.StdEncoding.DecodeString(publicKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	// 步骤2，计算代签名字串的SHA1哈希
	hashed := sha256.Sum256(signContent)
	hashs := crypto.SHA256
	digest := hashed[:]
	// 步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	data, _ := base64decode(sign)
	// 步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
	err = rsa.VerifyPKCS1v15(rsaPub, hashs, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return
	}
	return
}

// 格式化支付宝普通支付宝公钥
func (c *Client) formatAlipayPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// 请求接口
func (c *Client) Request(payload *Payload, result interface{}) (err error) {
	return c.doRequest(http.MethodPost, payload, result)
}

func (c *Client) OnReceivedData(fn func(method string, data []byte)) {
	c.onReceivedData = fn
}

func base64decode(data []byte) ([]byte, error) {
	var dBuf = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dBuf, data)
	return dBuf[:n], err
}
