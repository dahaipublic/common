package database

import (
	"common"
	"common/model"
)

type CBaseTeamMgr struct{}

// 通过id获取数据
func (this *CBaseTeamMgr) GetTeamByID(id uint64, sportsType int8) (data model.Team, err error) {
	if sportsType == model.MatchTypeFootball {
		footballTeam, _ := this.GetFootballTeamByID(id)
		data = footballTeam.Team
	} else if sportsType == model.MatchTypeBasketball {
		basketballTeam, _ := this.GetBasketballTeamByID(id)
		data = basketballTeam.Team
	}

	if data.ShortNameZh == "" {
		data.ShortNameZh = data.NameZh
	}
	return
}

// 通过id获取数据
func (this *CBaseTeamMgr) GetTeamNameAndLogByID(id uint64, sportsType int8, lanID common.ELanDef) (name, shortName, logo string) {
	data, _ := this.GetTeamByID(id, sportsType)
	logo = data.Logo
	switch lanID {
	//case ELan_zh_CN:
	case common.ELan_en_US:

		name = data.NameEn
		if data.ShortNameEn != "" {
			shortName = data.ShortNameEn
		} else {
			shortName = data.NameEn
		}
	case common.ELan_tr_TR:
		if data.NameTr == "" {
			name = data.NameEn
		} else {
			name = data.NameTr
			data.ShortNameEn = data.NameTr
		}

		if data.ShortNameEn != "" {
			shortName = data.ShortNameEn
		} else {
			shortName = data.NameEn
		}

	case common.ELan_ar_EG:
		if data.NameAr == "" {
			name = data.NameEn

		} else {
			name = data.NameAr
			data.ShortNameEn = data.NameAr
		}

		if data.ShortNameEn != "" {
			shortName = data.ShortNameEn
		} else {
			shortName = data.NameEn
		}
	default:
		name = data.NameEn
		shortName = data.ShortNameEn

	}

	// 土耳其没有简写先用英文
	return name, shortName, logo
}

// 通过id获取数据
func (this *CBaseTeamMgr) GetBasketballTeamByID(id uint64) (data model.BasketballTeam, errCode common.EErrCode) {
	//key := fmt.Sprintf(model.RedisKeyBasketballTeamID, id)
	//res := Redis.GetRedis().Get(key)
	//if res.Err() == nil {
	//	err := json.Unmarshal(xstr.String2Bytes(res.Val()), &data)
	//	JsonErrTest(err, &errCode)
	//	return
	//}

	ORMDB.Model(&model.BasketballTeam{}).Where("id = ?", id).First(&data)
	//jsonVal, _ := json.Marshal(&data)
	//Redis.GetRedis().Set(key, string(jsonVal), time.Second*model.RedisExpirationOneDay)
	return
}

func (this *CBaseTeamMgr) GetFootballTeamByID(id uint64) (data model.FootballTeam, errCode common.EErrCode) {

	ORMDB.Model(&model.FootballTeam{}).Where("id = ?", id).First(&data)
	return
}
