package util

import (
	"common/model"
	"sort"
	"strconv"
	"strings"
)

// 解析主/客队球员数据
func ParsePlayerData(raw []interface{}) []model.BasketballLineupTeamPlayerDetail {
	var players []model.BasketballLineupTeamPlayerDetail
	for _, item := range raw {
		playerArr := item.([]interface{})
		stats := strings.Split(playerArr[6].(string), "^")

		fg := strings.Split(stats[1], "-")
		threep := strings.Split(stats[2], "-")
		ft := strings.Split(stats[3], "-")
		id := uint64(toInt(playerArr[0]))

		players = append(players, model.BasketballLineupTeamPlayerDetail{
			ID:           id,
			Name:         playerArr[3].(string),
			Logo:         playerArr[4].(string),
			Num:          playerArr[5].(string),
			PlayTime:     stats[0],
			FGMade:       toInt(fg[0]),
			FGAttempt:    toInt(fg[1]),
			ThreePMade:   toInt(threep[0]),
			ThreePAttemp: toInt(threep[1]),
			FTMade:       toInt(ft[0]),
			FTAttempt:    toInt(ft[1]),
			ORB:          toInt(stats[4]),
			DRB:          toInt(stats[5]),
			TRB:          toInt(stats[6]),
			AST:          toInt(stats[7]),
			STL:          toInt(stats[8]),
			BLK:          toInt(stats[9]),
			TO:           toInt(stats[10]),
			PF:           toInt(stats[11]),
			PlusMinus:    toInt(stats[12]),
			PTS:          toInt(stats[13]),
			Played:       toInt(stats[14]),
			OnCourt:      toInt(stats[15]),
			Substitute:   toInt(stats[16]),
		})
	}
	sort.Slice(players, func(i, j int) bool {
		// 先按 Substitute 排序，0 在前
		if players[i].Substitute != players[j].Substitute {
			return players[i].Substitute < players[j].Substitute
		}
		// 如果 Substitute 相同，可以按其他字段再排序，比如位置或得分
		return players[i].PTS > players[j].PTS
	})
	return players
}

// 转 int
func toInt(v interface{}) int {
	switch val := v.(type) {
	case string:
		i, _ := strconv.Atoi(val)
		return i
	case float64:
		return int(val)
	default:
		return 0
	}
}

// 解析整队数据
func ParseTeamStats(s string) model.BasketballLineupTeamStats {
	parts := strings.Split(s, "^")
	fg := strings.Split(safeGet(parts, 1), "-")
	three := strings.Split(safeGet(parts, 2), "-")
	ft := strings.Split(safeGet(parts, 3), "-")

	return model.BasketballLineupTeamStats{
		FGMade:       toInt(safeGet(fg, 0)),
		FGAttempt:    toInt(safeGet(fg, 1)),
		ThreeMade:    toInt(safeGet(three, 0)),
		ThreeAttempt: toInt(safeGet(three, 1)),
		FTMade:       toInt(safeGet(ft, 0)),
		FTAttempt:    toInt(safeGet(ft, 1)),
		OffRebound:   toInt(safeGet(parts, 4)),
		DefRebound:   toInt(safeGet(parts, 5)),
		TotalRebound: toInt(safeGet(parts, 6)),
		Assists:      toInt(safeGet(parts, 7)),
		Steals:       toInt(safeGet(parts, 8)),
		Blocks:       toInt(safeGet(parts, 9)),
		Turnovers:    toInt(safeGet(parts, 10)),
		Fouls:        toInt(safeGet(parts, 11)),
		Points:       toInt(safeGet(parts, 13)), // 不够长时自动返回 0
	}
}
func safeGet(parts []string, idx int) string {
	if idx >= 0 && idx < len(parts) {
		return parts[idx]
	}
	return "0"
}

// 获取最高得分、助攻和篮板的球员
func GetHighestStats(players []model.BasketballLineupTeamPlayerDetail, statType string) model.BasketballLineupHighestStatsPlayerDetail {
	var playerDetail model.BasketballLineupTeamPlayerDetail
	var highestPlayer model.BasketballLineupHighestStatsPlayerDetail
	for _, player := range players {
		switch statType {
		case "PTS":
			if player.PTS > playerDetail.PTS {
				highestPlayer = model.BasketballLineupHighestStatsPlayerDetail{
					Logo:  player.Logo,
					Num:   player.Num,
					Name:  player.Name,
					Score: player.PTS,
				}
			}
		case "AST":
			if player.AST > playerDetail.AST {
				highestPlayer = model.BasketballLineupHighestStatsPlayerDetail{
					Logo:  player.Logo,
					Num:   player.Num,
					Name:  player.Name,
					Score: player.AST,
				}
			}
		case "TRB":
			if player.TRB > playerDetail.TRB {
				highestPlayer = model.BasketballLineupHighestStatsPlayerDetail{
					Logo:  player.Logo,
					Num:   player.Num,
					Name:  player.Name,
					Score: player.TRB,
				}
			}
		}
	}
	return highestPlayer
}
func GetNameByScope(scope int64) string {
	switch scope {
	case 1:
		return "SZN"
	case 2:
		return "ELE"
	case 3:
		return "GRP"
	case 4:
		return "PRE"
	case 5:
		return "RS"
	case 6:
		return "PO/KO"
	case 0:
		return ""
	default:
		return "Unknown"
	}
}
