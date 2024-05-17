package remoteclient

import (
	"context"
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"memoo-backend/localconfig"
)

func TwitterVerify(twitterUrl string) error {
	config := oauth1.NewConfig(localconfig.AppConfig.TwitterAttribute.TwitterConsumerKey, localconfig.AppConfig.TwitterAttribute.TwitterConsumerSecret)
	token := oauth1.NewToken(localconfig.AppConfig.TwitterAttribute.TwitterAccessToken, localconfig.AppConfig.TwitterAttribute.TwitterConsumerKey)
	httpClient := config.Client(context.Background(), token)
	//twitterUrlArr := strings.Split(twitterUrl, "/")
	//username := twitterUrlArr[len(twitterUrlArr)-1]
	client := twitter.NewClient(httpClient)
	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "@W8sFgv45Jt16576",
	})
	if err != nil {
		return err
	}
	if user.Verified {
		return nil
	} else {
		return errors.New("Twitter account is not verified")
	}
}
