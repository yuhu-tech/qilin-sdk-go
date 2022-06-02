package http

import (
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/yuhu-tech/qilin-sdk-go/internal/encrypto"
)

const (
	AuthFailureTempale = "认证失败 %s"
	ErrorParams        = "参数格式有误 %s"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderXYuhuDate     = "x-yuhu-date"
	Tag                 = "YUHU1-HMAC-SHA256"
	EndFlag             = "yuhu1_request"

	YuhuDateLayout       = "20060102T150405Z"
	CredentialTimeLayout = "20060102"
)

type PayloadMaker interface {
	Payload() string
}

type Authenticator struct {
	// s signer.Interface
	ak, sk string

	// authBuilder strings.Builder
}

func NewAuthenticator(ak, sk string) (*Authenticator, error) {
	if ak == "" || sk == "" {
		return nil, errors.New("ak, sk need no empty")
	}
	return &Authenticator{
		ak: ak,
		sk: sk,
		// pool
		// authBuilder: strings.Builder{},
	}, nil
}

func (a Authenticator) GenerateAuthHeader(region, payload, operation string) (map[string]string, error) {
	now := time.Now()
	yuhuDate, credentialTime := now.Format(YuhuDateLayout), now.Format(CredentialTimeLayout)

	toBeSignData, err := encrypto.HmacSha256Replace(Tag, yuhuDate, payload)
	if err != nil {
		return nil, err
	}
	// 4.2 获取privateKey
	tagHeader := Tag[:strings.Index(Tag, "-")] // YUHU1
	requestTimeDay := yuhuDate[:strings.Index(yuhuDate, "T")]
	privateKey, err := encrypto.HmacSha256Replace(tagHeader+a.sk, requestTimeDay, region, operation, EndFlag)
	if err != nil {
		return nil, err
	}

	// 4.3 计算签名
	signature := hex.EncodeToString(encrypto.Hmac(privateKey, toBeSignData))

	// 生成 header
	headers := make(map[string]string, 2)
	authBuilder := strings.Builder{}
	authBuilder.WriteString(Tag)
	authBuilder.WriteString(" Credential=")
	authBuilder.WriteString(a.ak)
	authBuilder.WriteString("/")
	authBuilder.WriteString(credentialTime)
	authBuilder.WriteString("/" + region) //"/cn-shanghai-1/"
	authBuilder.WriteString("/" + operation)
	authBuilder.WriteString("/yuhu1_request,")
	authBuilder.WriteString("Signature=")
	authBuilder.WriteString(signature)

	headers[HeaderAuthorization] = authBuilder.String()
	headers[HeaderXYuhuDate] = yuhuDate

	return headers, nil
}

// public static Header[] initHeader(String ak, String sk,
//                                       String data, String methed,
//                                       Date date) {
//         Date d = new Date();
//         SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMdd'T'HHmmss'Z'");
//         SimpleDateFormat sdf2 = new SimpleDateFormat("yyyyMMdd");
//         if(date!=null){
//             d =date;
//         }
//         byte[] rest = BSHA("YUHU1-HMAC-SHA256".getBytes(StandardCharsets.UTF_8),
//                 sdf.format(d).getBytes(StandardCharsets.UTF_8));
//         byte[] payload = data.getBytes(StandardCharsets.UTF_8);
//         rest = BSHA(rest,payload);
//         byte[] Date = sdf2.format(d).getBytes(StandardCharsets.UTF_8);
//         byte[] Area = "cn-shanghai-1".getBytes(StandardCharsets.UTF_8);
//         byte[] Service = methed.getBytes(StandardCharsets.UTF_8);
//         byte[] EndFlag = "yuhu1_request".getBytes(StandardCharsets.UTF_8);
//         byte[] SK = ("YUHU1"+sk).getBytes(StandardCharsets.UTF_8);
//         byte[] private_key =  BSHA(BSHA(BSHA(BSHA(SK,Date),Area),Service),EndFlag);
//         byte[] result = BSHA(private_key,rest);
//         String  Authorization ="YUHU1-HMAC-SHA256 Credential="+ak+"/"+sdf2.format(d)+"/cn" +
//                 "-shanghai-1/"+methed+"/yuhu1_request," +
//                 "Signature="+toHexStr(result)+"";
//         Header[] headers = new Header[2];
//         headers[0] = new BasicHeader("Authorization",Authorization);
//         headers[1] = new BasicHeader("x-yuhu-date",sdf.format(d));
//         return headers;
//     }
