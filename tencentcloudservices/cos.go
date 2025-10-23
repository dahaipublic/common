package tencentcloudservices

// import (
// 	"github.com/dahaipublic/common/conf"
// 	"context"
// 	"fmt"
// 	"mime/multipart"
// 	"net/http"
// 	"net/url"

// 	"github.com/tencentyun/cos-go-sdk-v5"
// )

// // 腾讯云上传图片
// func CosUploadImager(pathName string, fd multipart.File) (imagerUrl string, err error) {
// 	u, _ := url.Parse(fmt.Sprintf(conf.Conf.TencentCloud.Cos.UrlFormat, conf.Conf.TencentCloud.Cos.BucketName, conf.Conf.TencentCloud.Cos.Region))
// 	b := &cos.BaseURL{BucketURL: u}
// 	c := cos.NewClient(b, &http.Client{
// 		Transport: &cos.AuthorizationTransport{
// 			SecretID:  conf.Conf.TencentCloud.SecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://www.tencentcloud.com/document/product/598/32675
// 			SecretKey: conf.Conf.TencentCloud.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://www.tencentcloud.com/document/product/598/32675
// 		},
// 	})
// 	_, err = c.Object.Put(context.Background(), pathName, fd, nil)
// 	if err != nil {
// 		return
// 	}
// 	imagerUrl = fmt.Sprintf(conf.Conf.TencentCloud.Cos.UrlFormat, conf.Conf.TencentCloud.Cos.BucketName, conf.Conf.TencentCloud.Cos.Region) + "/" + pathName
// 	return
// }
