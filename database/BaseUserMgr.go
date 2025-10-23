package database

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"

	"strconv"
	"time"

	"github.com/dahaipublic/common/model"
	"github.com/dahaipublic/common/util"
	"github.com/dahaipublic/common/xstr"

	"gorm.io/gorm"

	xtime "github.com/dahaipublic/common/time"
)

// 通过设备id获取用户id
func GetUserIDByDeviceID(deviceID string) (userID uint64, errCode EErrCode) {
	if deviceID == "" {
		return
	}
	userID, errCode = BaseUserMgr.GetUserIDByDeviceID(deviceID)

	return
}

// 通过token获取用户id
func GetUserIDByToken(token string) (userID uint64, errCode EErrCode) {
	userID, errCode = BaseUserMgr.GetUserIDByToken(token)
	return
}

type CBaseUserMgr struct{}

// 通过id获取数据
func (this *CBaseUserMgr) GetUserByID(userID uint64) (data model.User, errCode EErrCode) {
	key := fmt.Sprintf(model.RedisKeyUserID, userID)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.User{}).Where("id = ?", userID).First(&data).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

// 通过设备id获取用户信息
func (this *CBaseUserMgr) GetUserByDeviceID(deviceID string) (user model.User, errCode EErrCode) {
	var id uint64
	//先获取用户id
	key := fmt.Sprintf(model.RedisKeyUserDeviceIDGetID, deviceID)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		id, _ = strconv.ParseUint(res.Val(), 10, 64)
	} else {
		err := ORMDB.Model(&model.User{}).Where("device_id = ? ", deviceID).Select("id").First(&id).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		Redis.GetRedis().Set(key, fmt.Sprintf("%d", id), time.Second*model.RedisExpiration)
	}
	//再获取用户信息
	user, errCode = this.GetUserByID(id)
	if errCode != No_Error {
		return
	}
	return
}

func GenerateUserCode() string {
	var num uint64
	err := binary.Read(rand.Reader, binary.BigEndian, &num)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%06d", num%1000000)
}

// GenerateSmsCode 生成 6 位短信验证码
func GenerateSmsCode() (string, EErrCode) {
	max := big.NewInt(1000000) // [0, 1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", Err_Unknown
	}
	return fmt.Sprintf("%06d", n.Int64()), No_Error
}

// 通过设备id获取用户信息
func (this *CBaseUserMgr) GetUserIDByDeviceID(deviceID string) (userID uint64, errCode EErrCode) {
	user, userErr := this.GetUserByDeviceID(deviceID)
	if userErr == No_Error {
		userID = user.ID
		return
	}
	tx := ORMDB.Begin()
	defer func() {
		if errCode != No_Error {
			tx.Rollback()
			return
		}
		tx.Commit()
		return
	}()
	//给没有昵称的游客分配一个昵称
	// var userTourist model.UserTourist
	// userTourist, err = this.GetNotUseTouristNickName(tx)
	// if err != nil {
	// 	return
	// }
	user.Nickname = util.GuestName + GenerateUserCode()
	user.DeviceID = deviceID
	user.RoleID = 1
	err := this.UserCreate(tx, &user)
	if err != No_Error {
		return
	}
	userID = user.ID
	return
}

// 通过token获取用户信息
func (this *CBaseUserMgr) GetUserIDByToken(token string) (userID uint64, errCode EErrCode) {
	res := Redis.RedisGet(fmt.Sprintf(model.RedisKeyUserToken, token))
	if res == "" {
		return
	}
	userID, _ = strconv.ParseUint(res, 10, 64)
	//处理在线奖励
	date := time.Now().Format("2006-01-02")
	redisRes := Redis.RedisGet(fmt.Sprintf(model.RedisKeyUserDailyOnlineRewards, date, userID))
	if redisRes == "" {
		tx := ORMDB.Begin()
		defer func() {
			if errCode != No_Error {
				tx.Rollback()
				return
			}
			tx.Commit()
			return
		}()
		var user model.User
		user, err := this.GetUserByID(userID)
		if err != No_Error {
			return
		}
		err = this.UserExperienceTask(tx, user, 0, model.UserExperienceTypeDailyOnline)
		if err != No_Error {
			return
		}
	}
	return
}

// 创建用户
func (this *CBaseUserMgr) UserCreate(tx *gorm.DB, data *model.User) (errCode EErrCode) {
	err := tx.Model(&model.User{}).Create(&data).Error
	if err != nil {

		errCode = Err_DB
		return
	}
	return
}

// 修改用户信息
func (this *CBaseUserMgr) UserUpdateDB(tx *gorm.DB, id uint64, updateMap map[string]interface{}) (errCode EErrCode) {
	err := tx.Model(&model.User{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		errCode = Err_DB
		return
	}
	//删除缓存
	errorCode := this.DeleteUserCache(id)
	if errorCode != No_Error {
		return
	}
	return
}

// 删除用户信息缓存
func (this *CBaseUserMgr) DeleteUserCache(userID uint64) (errCode EErrCode) {
	res := Redis.GetRedis().Del(fmt.Sprintf(model.RedisKeyUserID, userID))
	if res.Err() != nil {
		errCode = Err_DB
		return
	}
	return
}

func (this *CBaseUserMgr) GetUserRoleByID(id uint64) (data model.UserRole, errCode EErrCode) {
	key := fmt.Sprintf(model.RedisKeyUserRoleID, id)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.UserRole{}).Where("id = ?", id).First(&data).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

func (this *CBaseUserMgr) GetUserRoleByLevel(level int64) (data model.UserRole, errCode EErrCode) {
	key := fmt.Sprintf(model.RedisKeyUserRoleLevel, level)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.UserRole{}).Where("level = ?", level).First(&data).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

func (this *CBaseUserMgr) CreateUserRoleLog(tx *gorm.DB, data *model.UserRoleLog) (errCode EErrCode) {
	err := tx.Model(&model.UserRoleLog{}).Create(&data).Error
	if err != nil {
		errCode = Err_DB
		return
	}
	return
}

func (this *CBaseUserMgr) GetUserLevelByLevel(level int64) (data model.UserLevel, errCode EErrCode) {
	key := fmt.Sprintf(model.RedisKeyUserLevelLevel, level)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.UserLevel{}).Where("level = ?", level).First(&data).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

func (this *CBaseUserMgr) CreateUserLevelLog(tx *gorm.DB, data *model.UserLevelLog) (errCode EErrCode) {
	err := tx.Model(&model.UserLevelLog{}).Create(&data).Error
	if err != nil {
		errCode = Err_DB
		return
	}
	return
}

func (this *CBaseUserMgr) UserLevelJudgeUpgrade(tx *gorm.DB, user model.User, lastExperience int64, userExperienceLogID uint64) (errCode EErrCode) {
	nextUserLevel, _ := this.GetUserLevelByLevel(user.Level + 1)
	if nextUserLevel.ID <= 0 || lastExperience < nextUserLevel.Experience {
		return
	}
	var userLevelLog model.UserLevelLog
	userLevelLog.UserExperienceLogID = userExperienceLogID
	userLevelLog.UserID = user.ID
	userLevelLog.BeforeLevel = user.Level
	userLevelLog.LastLevel = user.Level + 1
	errCode = this.CreateUserLevelLog(tx, &userLevelLog)
	if errCode != No_Error {
		return
	}
	//修改等级
	updateMap := make(map[string]interface{})
	updateMap["level"] = userLevelLog.LastLevel
	errCode = this.UserUpdateDB(tx, user.ID, updateMap)
	if errCode != No_Error {
		return
	}
	//判断角色是否变动
	errCode = this.UserRoleJudgeUpgrade(tx, user, userLevelLog.LastLevel, userLevelLog.ID)
	if errCode != No_Error {
		return
	}
	return
}

func (this *CBaseUserMgr) UserRoleJudgeUpgrade(tx *gorm.DB, user model.User, lastLevel int64, userLevelLogID uint64) (errCode EErrCode) {
	userRole, _ := this.GetUserRoleByLevel(lastLevel)
	if userRole.ID <= 0 {
		return
	}
	var userRoleLog model.UserRoleLog
	userRoleLog.UserLevelLogID = userLevelLogID
	userRoleLog.UserID = user.ID
	userRoleLog.BeforeRoleID = user.SecondaryRoleID
	userRoleLog.LastRoleID = userRole.ID
	errCode = this.CreateUserRoleLog(tx, &userRoleLog)
	if errCode != No_Error {
		return
	}
	//修改角色
	updateMap := make(map[string]interface{})
	updateMap["secondary_role_id"] = userRole.ID
	errCode = this.UserUpdateDB(tx, user.ID, updateMap)
	if errCode != No_Error {
		return
	}
	return
}

func (this *CBaseUserMgr) UserExperienceTask(tx *gorm.DB, user model.User, fromUserID uint64, experienceType int8) (errCode EErrCode) {
	var userExperienceLog model.UserExperienceLog
	userExperienceLog.UserID = user.ID
	userExperienceLog.FromUserID = fromUserID
	userExperienceLog.Type = experienceType
	//普通用户才能做任务
	if user.RoleID != model.UserRoleIDUser {
		return
	}
	userExperienceLog.BeforeExperienceNum = user.Experience
	if experienceType == model.UserExperienceTypeInviteNewUser {
		//获取邀请的奖励
		experienceTaskInviteNewUserRewards := this.GetConfigByKey("experience_task_invite_new_user_rewards")
		if experienceTaskInviteNewUserRewards == "" || experienceTaskInviteNewUserRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskInviteNewUserRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeSignIn {
		signInUserExperienceLog, _ := this.GetUserExperienceLogByType(user.ID, experienceType)
		timestamp := xtime.TodayZeroTimestamp()
		key := ""
		//如果没有上一条签到记录,获取上一条签到记录不是昨天的,或者已经七天了那按第一天签到
		if signInUserExperienceLog.ID <= 0 || (signInUserExperienceLog.CreatedAt < timestamp && signInUserExperienceLog.CreatedAt >= timestamp-86400) || signInUserExperienceLog.ContinuousDays == 7 {
			key = "experience_task_sign_in_one_rewards"
			userExperienceLog.ContinuousDays = 1
		} else {
			if signInUserExperienceLog.ContinuousDays == 1 {
				key = "experience_task_sign_in_two_rewards"
			} else if signInUserExperienceLog.ContinuousDays == 2 {
				key = "experience_task_sign_in_three_rewards"
			} else if signInUserExperienceLog.ContinuousDays == 3 {
				key = "experience_task_sign_in_four_rewards"
			} else if signInUserExperienceLog.ContinuousDays == 4 {
				key = "experience_task_sign_in_five_rewards"
			} else if signInUserExperienceLog.ContinuousDays == 5 {
				key = "experience_task_sign_in_six_rewards"
			} else if signInUserExperienceLog.ContinuousDays == 6 {
				key = "experience_task_sign_in_seven_rewards"
			}
			userExperienceLog.ContinuousDays = signInUserExperienceLog.ContinuousDays + 1
		}
		experienceTaskSignInOneRewards := this.GetConfigByKey(key)
		if experienceTaskSignInOneRewards == "" || experienceTaskSignInOneRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskSignInOneRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeUpdateHeaderImg {
		updateHeaderImgUserExperienceLog, _ := this.GetUserExperienceLogByType(user.ID, experienceType)
		if updateHeaderImgUserExperienceLog.ID > 0 {
			return
		}
		experienceTaskUpdateHeaderImgRewards := this.GetConfigByKey("experience_task_update_header_img_rewards")
		if experienceTaskUpdateHeaderImgRewards == "" || experienceTaskUpdateHeaderImgRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskUpdateHeaderImgRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeUpdateSignature {
		updateSignatureUserExperienceLog, _ := this.GetUserExperienceLogByType(user.ID, experienceType)
		if updateSignatureUserExperienceLog.ID > 0 {
			return
		}
		experienceTaskUpdateSignatureRewards := this.GetConfigByKey("experience_task_update_signature_rewards")
		if experienceTaskUpdateSignatureRewards == "" || experienceTaskUpdateSignatureRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskUpdateSignatureRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeUpdateNickname {
		updateNicknameUserExperienceLog, _ := this.GetUserExperienceLogByType(user.ID, experienceType)
		if updateNicknameUserExperienceLog.ID > 0 {
			return
		}
		experienceTaskUpdateNicknameRewards := this.GetConfigByKey("experience_task_update_nickname_rewards")
		if experienceTaskUpdateNicknameRewards == "" || experienceTaskUpdateNicknameRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskUpdateNicknameRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeRegister {
		experienceTaskRegisterRewards := this.GetConfigByKey("experience_task_register_rewards")
		if experienceTaskRegisterRewards == "" || experienceTaskRegisterRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskRegisterRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeDailyInteractive {
		date := time.Now().Format("2006-01-02")
		sendMessageNum := Redis.RedisGet(fmt.Sprintf(model.RedisKeyUserSendMessageNum, date, user.ID))
		if sendMessageNum != "20" {
			return
		}
		experienceTaskDailyInteractiveRewards := this.GetConfigByKey("experience_task_daily_interactive_rewards")
		if experienceTaskDailyInteractiveRewards == "" || experienceTaskDailyInteractiveRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskDailyInteractiveRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeWatchLive {
		date := time.Now().Format("2006-01-02")
		watchLiveRewards := Redis.RedisGet(fmt.Sprintf(model.RedisKeyMatchWatchLiveRewards, date, user.ID))
		if watchLiveRewards != "" {
			return
		}
		defer func() {
			if errCode == No_Error {
				Redis.RedisSet(fmt.Sprintf(model.RedisKeyMatchWatchLiveRewards, date, user.ID), fmt.Sprintf("%d", time.Now().Unix()), model.RedisExpirationOneDay*time.Second)
			}
		}()
		experienceTaskWatchLive10MinuteRewards := this.GetConfigByKey("experience_task_watch_live_10_minute_rewards")
		if experienceTaskWatchLive10MinuteRewards == "" || experienceTaskWatchLive10MinuteRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskWatchLive10MinuteRewards, 10, 64)
	} else if experienceType == model.UserExperienceTypeDailyOnline {
		date := time.Now().Format("2006-01-02")
		dailyOnlineRewards := Redis.RedisGet(fmt.Sprintf(model.RedisKeyUserDailyOnlineRewards, date, user.ID))
		if dailyOnlineRewards != "" {
			return
		}
		defer func() {
			if errCode == No_Error {
				Redis.RedisSet(fmt.Sprintf(model.RedisKeyUserDailyOnlineRewards, date, user.ID), fmt.Sprintf("%d", time.Now().Unix()), model.RedisExpirationOneDay*time.Second)
			}
		}()
		experienceTaskDailyOnlineRewards := this.GetConfigByKey("experience_task_daily_online_rewards")
		if experienceTaskDailyOnlineRewards == "" || experienceTaskDailyOnlineRewards == "0" {
			return
		}
		userExperienceLog.Num, _ = strconv.ParseInt(experienceTaskDailyOnlineRewards, 10, 64)
	}
	//创建记录
	userExperienceLog.LastExperienceNum = userExperienceLog.BeforeExperienceNum + userExperienceLog.Num
	errCode = this.CreateUserExperienceLog(tx, &userExperienceLog)
	if errCode != No_Error {
		return
	}
	//修改经验值
	updateMap := make(map[string]interface{})
	updateMap["experience"] = userExperienceLog.LastExperienceNum
	errCode = this.UserUpdateDB(tx, user.ID, updateMap)
	if errCode != No_Error {
		return
	}
	//判断升级
	errCode = this.UserLevelJudgeUpgrade(tx, user, userExperienceLog.LastExperienceNum, userExperienceLog.ID)
	if errCode != No_Error {
		return
	}
	return
}

// \dao\user_experience_log.go
// 判断奖励是否领取
func (this *CBaseUserMgr) JudgeAwardWhetherReceive(userID uint64, experienceType int8) (b bool) {
	userExperienceLog, _ := this.GetUserExperienceLogByType(userID, experienceType)
	if userExperienceLog.ID <= 0 {
		return
	}
	//今日零点时间戳
	if experienceType == model.UserExperienceTypeRegister || experienceType == model.UserExperienceTypeUpdateHeaderImg || experienceType == model.UserExperienceTypeUpdateNickname || experienceType == model.UserExperienceTypeUpdateSignature {
		if userExperienceLog.ID > 0 {
			b = true
			return
		}
	} else {
		timestamp := xtime.TodayZeroTimestamp()
		if userExperienceLog.CreatedAt >= timestamp {
			b = true
			return
		}
	}
	return
}

func (this *CBaseUserMgr) GetUserExperienceLogByType(userID uint64, experienceType int8) (data model.UserExperienceLog, errCode EErrCode) {
	key := fmt.Sprintf(model.RedisKeyUserExperienceLogType, userID, experienceType)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.UserExperienceLog{}).Where("user_id = ? and type = ?", userID, experienceType).Order("id desc").First(&data).Error
		if err != nil {
			errCode = Err_DB
			return
		}
		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

func (this *CBaseUserMgr) CreateUserExperienceLog(tx *gorm.DB, data *model.UserExperienceLog) (errCode EErrCode) {
	err := tx.Model(&model.UserExperienceLog{}).Create(&data).Error
	if err != nil {
		errCode = Err_DB
		return
	}
	//清除缓存
	errCode = this.DeleteUserExperienceLogCache(data.UserID, data.Type)
	if errCode != No_Error {
		return
	}
	return
}

func (this *CBaseUserMgr) DeleteUserExperienceLogCache(userID uint64, experienceType int8) (errCode EErrCode) {
	res := Redis.GetRedis().Del(fmt.Sprintf(model.RedisKeyUserExperienceLogType, userID, experienceType))
	if res.Err() != nil {
		errCode = Err_DB
		return
	}
	return
}

// 系统参数
func (this *CBaseUserMgr) GetConfigByKey(key string) (value string) {
	redisKey := fmt.Sprintf(model.RedisKeyConfigKey, key)
	res := Redis.GetRedis().Get(redisKey)
	if res.Err() == nil {
		value = res.Val()
		return
	}
	ORMDB.Model(&model.Config{}).Where("`key` = ?", key).Select("value").Scan(&value)
	Redis.GetRedis().Set(redisKey, value, time.Second*model.RedisExpiration)
	return
}

var BaseUserMgr = &CBaseUserMgr{}
