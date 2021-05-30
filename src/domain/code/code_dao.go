package code

import (
	"strconv"
	"time"

	"github.com/jiaqi-yin/go-verification-code/src/clients/redis"
	"github.com/jiaqi-yin/go-verification-code/src/utils"
)

const (
	codeExpiration = 2 * time.Minute
	countKeyPrefix = "count:"
	codeKeyPrefix  = "code:"
	retryLimit     = 3
)

func getCountKey(phoneNo string) string {
	return countKeyPrefix + phoneNo
}

func getCodeKey(phoneNo string) string {
	return codeKeyPrefix + phoneNo
}

func (codeGenerator *CodeGenerator) Generate() (string, utils.RestErr) {
	countKey := getCountKey(codeGenerator.Phone)
	codeKey := getCodeKey(codeGenerator.Phone)

	count, _ := redis.Client.Get(countKey)
	if count == "" {
		expiration, err := utils.TimeDiffInSecBetweenNowAndTomorrow()
		if err != nil {
			return "", utils.NewInternalServerError(err.Error())
		}
		redis.Client.Set(countKey, "1", expiration)
		verificationCode := utils.EncodeToString(6)
		redis.Client.Set(codeKey, verificationCode, codeExpiration)
		return verificationCode, nil
	} else {
		retries, err := strconv.Atoi(count)
		if err != nil {
			return "", utils.NewInternalServerError(err.Error())
		}
		if retries < retryLimit {
			verificationCode := utils.EncodeToString(6)
			redis.Client.Set(codeKey, verificationCode, codeExpiration)
			redis.Client.Incr(countKey)
			return verificationCode, nil
		} else {
			return "", utils.NewBadRequestError("You are exceeding the limit today. Please try it tomorrow.")
		}
	}
}

func (codeVerifier *CodeVerifier) Verify() utils.RestErr {
	inputCode := codeVerifier.Code
	codeKey := getCodeKey(codeVerifier.Phone)
	redisCode, _ := redis.Client.Get(codeKey)
	if redisCode == "" {
		return utils.NewBadRequestError("verification code does not exist in redis")
	} else {
		if inputCode == redisCode {
			return nil
		} else {
			return utils.NewBadRequestError("verification code does not match")
		}
	}
}
