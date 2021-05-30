package sms

import (
	"fmt"
	"net/http"

	appconfig "github.com/jiaqi-yin/go-verification-code/src/app_config"
	"github.com/jiaqi-yin/go-verification-code/src/domain/sms"
	"github.com/jiaqi-yin/go-verification-code/src/utils"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	phoneCountryCode = "+61" // Australia mobile number
)

type SmsClientInterface interface {
	SendMessage(string, string) utils.RestErr
}

type smsClient struct {
	client *rest.RequestBuilder
}

func (c *smsClient) SendMessage(phone string, code string) utils.RestErr {
	var destination string
	if len(phone) == 10 {
		// Convert the 10 digit mobile number to the E.164 international format.
		destination = phoneCountryCode + phone[1:]
	} else {
		return utils.NewBadRequestError("phone_number should have 10 digit numbers")
	}

	messages := sms.Messages{
		Messages: []sms.Message{
			{
				Content:     fmt.Sprintf("Your verification code is %s", code),
				Destination: destination,
				Format:      "SMS",
			},
		},
	}

	response := c.client.Post("/v1/messages", messages)

	if response == nil || response.Response == nil {
		return utils.NewInternalServerError("invalid rest response when trying to send text messages")
	}

	if response.StatusCode > 299 {
		return utils.NewInternalServerError("invalide rest response status when trying to send text messages")
	}

	return nil
}

func NewSmsClient() SmsClientInterface {
	headers := make(http.Header)
	headers.Add("Authorization", "Basic "+appconfig.Config.SmsAuthorization)

	requestBuilder := &rest.RequestBuilder{
		Headers:     headers,
		BaseURL:     appconfig.Config.SmsBaseUrl,
		ContentType: rest.JSON,
	}

	return &smsClient{
		client: requestBuilder,
	}
}
