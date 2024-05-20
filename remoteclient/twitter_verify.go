package remoteclient

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"memoo-backend/localconfig"
	"net"
	"net/http"
	"time"
)

func TwitterVerify(twitterUrl string) error {
	config := oauth1.NewConfig(localconfig.AppConfig.TwitterAttribute.TwitterConsumerKey, localconfig.AppConfig.TwitterAttribute.TwitterConsumerSecret)
	token := oauth1.NewToken(localconfig.AppConfig.TwitterAttribute.TwitterAccessToken, localconfig.AppConfig.TwitterAttribute.TwitterConsumerKey)
	//twitterUrlArr := strings.Split(twitterUrl, "/")
	//username := twitterUrlArr[len(twitterUrlArr)-1]
	// 创建 HTTP 客户端并设置超时
	httpClient := config.Client(oauth1.NoContext, token)

	// 设置连接和握手超时
	httpClient.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second, // 连接超时时间
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 20 * time.Second, // TLS 握手超时时间
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
	}

	// 创建 Twitter API 客户端
	twitterClient := twitter.NewClient(httpClient)

	// 验证账户真实性
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: twitter.Bool(true),
		SkipStatus:      twitter.Bool(false),
		IncludeEmail:    twitter.Bool(false),
	}
	user, _, err := twitterClient.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		fmt.Println("Error verifying Twitter account:", err)
		return err
	}

	fmt.Println("Verified Twitter account:", user.ScreenName)
	return nil
	//httpClient := config.Client(context.Background(), token)
	//twitterUrlArr := strings.Split(twitterUrl, "/")
	//username := twitterUrlArr[len(twitterUrlArr)-1]
	//client := twitter.NewClient(httpClient)
	//user, _, err := client.Users.Show(&twitter.UserShowParams{
	//	ScreenName: username,
	//})
	//if err != nil {
	//	return err
	//}
	//if user.Verified {
	//	return nil
	//} else {
	//	return errors.New("Twitter account is not verified")
	//}
}
