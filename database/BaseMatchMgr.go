package database

import (
	"common"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	//	. "common/conf"
	"common/model"
	"common/xstr"
)

type CBaseMatchMgr struct {
}

func (this *CBaseMatchMgr) GetMatchByID(id uint64, sportsType int8) (data model.MatchAdmin, errCode common.EErrCode) {
	if sportsType == model.MatchTypeFootball {
		footballMatch, _ := this.GetFootballMatchByID(id)
		data = footballMatch.MatchAdmin
	} else if sportsType == model.MatchTypeBasketball {
		basketballMatch, _ := this.GetBasketballMatchByID(id)
		data = basketballMatch.MatchAdmin
	}
	return
}

// 通过id获取数据
func (this *CBaseMatchMgr) GetBasketballMatchByID(id uint64) (data model.BasketballMatchAdmin, errCode common.EErrCode) {

	err := ORMDB.Model(&model.BasketballMatchAdmin{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		errCode = common.Err_DB
		return
	}
	// jsonVal, _ := json.Marshal(data)
	// Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	// }
	return
}

// 通过id获取数据
func (this *CBaseMatchMgr) GetFootballMatchByID(id uint64) (data model.FootballMatchAdmin, errCode common.EErrCode) {
	// key := fmt.Sprintf(model.RedisKeyMatchID, id)
	// res := Redis.GetRedis().Get(key)
	// if res.Err() == nil {
	// 	err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
	// 	JsonErrTest(err, &errCode)
	// 	return
	// } else {
	err := ORMDB.Model(&model.FootballMatchAdmin{}).Where("id = ?", id).First(&data)
	if err != nil {
		errCode = common.Err_DB
		return
	}
	// jsonVal, _ := json.Marshal(data)
	// Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	// }
	return
}

func (this *CBaseMatchMgr) GetMatchInfoByMatchID(matchID uint64) (data model.MatchInfo, errCode common.EErrCode) {
	key := fmt.Sprintf(model.RedisKeyMatchInfoMatchID, matchID)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
		common.JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.MatchInfo{}).Where("match_id = ?", matchID).Order("id desc").First(&data).Error
		if err != nil {
			errCode = common.Err_DB
			return
		}

		jsonVal, _ := json.Marshal(data)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

func (this *CBaseMatchMgr) MatchStatusNotStartOrCompletion(sportsType int8) (matchStatusIDNotStart, matchStatusIDCompletion int8) {
	if sportsType == model.MatchTypeFootball {
		matchStatusIDNotStart = model.MatchFootballStatusIDNotStart
		matchStatusIDCompletion = model.MatchFootballStatusIDCompletion
	} else if sportsType == model.MatchTypeBasketball {
		matchStatusIDNotStart = model.MatchBasketballStatusIDNotStart
		matchStatusIDCompletion = model.MatchBasketballStatusIDCompletion
	}
	return
}

// 获取比赛直播数据
func (this *CBaseMatchMgr) GetMatchLive(matchID uint64) (rsp interface{}, errCode common.EErrCode) {
	key := fmt.Sprintf(model.RedisKeyMatchLive, matchID)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &rsp)
		common.JsonErrTest(err, &errCode)
		return
	}
	matchInfo, _ := this.GetMatchInfoByMatchID(matchID)
	if matchInfo.Live != "" {
		err := json.Unmarshal([]byte(matchInfo.Live), &rsp)
		if err != nil {
			errCode = common.Err_JsonDecode
			return
		}

		Redis.GetRedis().Set(key, matchInfo.Live, time.Second*model.RedisExpirationOneDay)
	}
	return
}

// // 获取比赛指数数据
// func (this *CBaseMatchMgr) GetMatchOdds(matchID uint64) (rsp interface{}, errCode EErrCode) {
// 	matchInfo, _ := this.GetMatchInfoByMatchID(matchID)
// 	if matchInfo.Odds != "" {
// 		err := json.Unmarshal([]byte(matchInfo.Odds), &rsp)
// 		JsonErrTest(err, &errCode)
// 	}
// 	key := fmt.Sprintf(model.RedisKeyOddsMatchID, matchID)
// 	res := Redis.RedisGet(key)
// 	if res != "" {
// 		matchOddsData := make(map[string]map[string][]interface{})
// 		json.Unmarshal([]byte(res), &matchOddsData)
// 		for companyID, _ := range matchOddsData {
// 			for oddsType, _ := range matchOddsData[companyID] {
// 				minData, _ := Redis.GetRedis().ZRangeByScoreWithScores(fmt.Sprintf(model.RedisKeyOddsMatchIDCompanyIDType, matchID, companyID, oddsType), &redis.ZRangeBy{
// 					Min:    "-inf",
// 					Max:    "+inf",
// 					Offset: 0,
// 					Count:  1,
// 				}).Result()
// 				maxData, _ := Redis.GetRedis().ZRevRangeByScoreWithScores(fmt.Sprintf(model.RedisKeyOddsMatchIDCompanyIDType, matchID, companyID, oddsType), &redis.ZRangeBy{
// 					Min:    "-inf",
// 					Max:    "+inf",
// 					Offset: 0,
// 					Count:  1,
// 				}).Result()
// 				if len(minData) > 0 {
// 					if minData[0].Member != "" {
// 						var minOddsData interface{}
// 						json.Unmarshal([]byte(minData[0].Member.(string)), &minOddsData)
// 						matchOddsData[companyID][oddsType] = append(matchOddsData[companyID][oddsType], minOddsData)
// 					}
// 				}
// 				if len(maxData) > 0 {
// 					if maxData[0].Member != "" {
// 						var maxOddsData interface{}
// 						json.Unmarshal([]byte(maxData[0].Member.(string)), &maxOddsData)
// 						matchOddsData[companyID][oddsType] = append(matchOddsData[companyID][oddsType], maxOddsData)
// 					}
// 				}
// 			}
// 		}
// 		rsp = matchOddsData
// 	}

// 	return
// }

// 创建
func (this *CBaseMatchMgr) MatchAnnouncementCreate(tx *gorm.DB, matchAnnouncement *model.MatchAnnouncement) (errCode common.EErrCode) {
	err := tx.Model(&model.MatchAnnouncement{}).Create(&matchAnnouncement).Error
	if err != nil {
		errCode = common.Err_DB
		return
	}

	jsonVal, _ := json.Marshal(matchAnnouncement)
	Redis.GetRedis().Set(fmt.Sprintf(model.RedisKeyMatchAnnouncementMatchID, matchAnnouncement.MatchID), string(jsonVal), time.Second*model.RedisExpirationOneDay)
	return
}

func (this *CBaseMatchMgr) GetMatchAnnouncementByMatchID(matchID uint64) (rsp model.MatchAnnouncement, errCode common.EErrCode) {
	key := fmt.Sprintf(model.RedisKeyMatchAnnouncementMatchID, matchID)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err := json.Unmarshal(xstr.String2Bytes(res.Val()), &rsp)
		common.JsonErrTest(err, &errCode)
		return
	} else {
		err := ORMDB.Model(&model.MatchAnnouncement{}).Where("match_id = ?", matchID).Order("id desc").First(&rsp).Error
		if err != nil {
			errCode = common.Err_DB
			return
		}
		jsonVal, _ := json.Marshal(rsp)
		Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	}
	return
}

// copy from  api/param
type BasketballMatchLive struct {
	ID             uint64         `json:"id"`
	Score          []interface{}  `json:"score"`
	Timer          [4]int64       `json:"timer"`
	OverTimeScores [][]int64      `json:"over_time_scores"`
	Stats          [][3]int64     `json:"stats"`
	Tlive          [][]string     `json:"tlive"`
	Players        [4]interface{} `json:"players"`
}

// 获取比赛阵容和统计数据
func (this *CBaseMatchMgr) GetMatchLineup(matchID uint64, sportsType int8) (rsp interface{}, errCode common.EErrCode) {
	if sportsType == model.MatchTypeFootball {
		key := fmt.Sprintf(model.RedisKeyFootballMatchLineup, matchID)
		res := Redis.GetRedis().Get(key)
		if res.Err() == nil {
			err := json.Unmarshal(xstr.String2Bytes(res.Val()), &rsp)
			common.JsonErrTest(err, &errCode)
			return
		}
		matchInfo, _ := this.GetMatchInfoByMatchID(matchID)
		if matchInfo.Lineup != "" {
			json.Unmarshal([]byte(matchInfo.Lineup), &rsp)
			Redis.GetRedis().Set(key, matchInfo.Lineup, time.Second*model.RedisExpirationOneDay)
		}
	} else if sportsType == model.MatchTypeBasketball {
		key := fmt.Sprintf(model.RedisKeyBasketballMatchLineup, matchID)
		res := Redis.GetRedis().Get(key)
		if res.Err() == nil {
			err := json.Unmarshal(xstr.String2Bytes(res.Val()), &rsp)
			common.JsonErrTest(err, &errCode)
			return
		}
		matchInfo, _ := this.GetMatchInfoByMatchID(matchID)
		if matchInfo.Live != "" {
			var basketballMatchLive BasketballMatchLive
			json.Unmarshal([]byte(matchInfo.Live), &basketballMatchLive)
			rsp = basketballMatchLive.Players
			dataStr, _ := json.Marshal(rsp)
			Redis.GetRedis().Set(key, string(dataStr), time.Second*model.RedisExpirationOneDay)
		}
	}
	return
}

// 修改
func (this *CBaseMatchMgr) MatchInfoUpdateDB(tx *gorm.DB, matchID uint64, updateMap map[string]interface{}) (errCode common.EErrCode) {
	err := tx.Model(&model.MatchInfo{}).Where("match_id = ?", matchID).Updates(updateMap).Error
	if err != nil {
		errCode = common.Err_DB
		return
	}

	errCode = this.MatchInfoDeleteCache(matchID)
	return
}

// 删除缓存
func (this *CBaseMatchMgr) MatchInfoDeleteCache(matchID uint64) (errCode common.EErrCode) {
	res := Redis.GetRedis().Del(fmt.Sprintf(model.RedisKeyMatchInfoMatchID, matchID))
	common.RedisErrTest(res.Err(), &errCode)
	// if res.Err() != nil {
	// 	err = res.Err()
	// 	return
	// }
	return
}

var BaseMatchMgr = &CBaseMatchMgr{}

// 发送信息
// 过滤
func GetSensitiveWordInspect() (rsp []model.SensitiveWord, err error) {
	key := fmt.Sprintf(model.RedisKeySensitiveWord)
	res := Redis.GetRedis().Get(key)
	if res.Err() == nil {
		err = json.Unmarshal(xstr.String2Bytes(res.Val()), &rsp)
		return
	} else {
		err = ORMDB.Model(&model.SensitiveWord{}).Find(&rsp).Error
		if err != nil {
			return
		}
		jsonVal, _ := json.Marshal(rsp)
		Redis.GetRedis().Set(key, string(jsonVal), 0)
	}

	return
}

// 敏感词检查
func SensitiveWordInspect(content string) (b bool, err error) {
	list, _ := GetSensitiveWordInspect()
	if len(list) <= 0 {
		return
	}
	for _, v := range list {
		if v.SensitiveWord != "" {
			if strings.Contains(content, v.SensitiveWord) {
				b = true
				return
			}
		}
		if v.PinyinStatus == 1 && v.SensitiveWordPinyin != "" {
			if strings.Contains(content, v.SensitiveWordPinyin) {
				b = true
				return
			}
		}
	}
	return
}
