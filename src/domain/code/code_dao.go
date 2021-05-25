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

func (codeGenerator *CodeGenerator) Generate() utils.RestErr {
	phoneNo := codeGenerator.Phone
	countKey := countKeyPrefix + phoneNo
	codeKey := codeKeyPrefix + phoneNo

	count, _ := redis.Client.Get(countKey)
	if count == "" {
		expiration, err := utils.TimeDiffInSecBetweenNowAndTomorrow()
		if err != nil {
			return utils.NewInternalServerError(err.Error())
		}
		redis.Client.Set(countKey, "1", expiration)
		verificationCode := utils.EncodeToString(6)
		redis.Client.Set(codeKey, verificationCode, codeExpiration)
		return nil
	} else {
		retries, err := strconv.Atoi(count)
		if err != nil {
			return utils.NewInternalServerError(err.Error())
		}
		if retries < retryLimit {
			verificationCode := utils.EncodeToString(6)
			redis.Client.Set(codeKey, verificationCode, codeExpiration)
			redis.Client.Incr(countKey)
			return nil
		} else {
			return utils.NewBadRequestError("You are exceeding the limit. Please try it tomorrow.")
		}
	}
}
