package model

type CronGetFootballTableLiveRsp struct {
	Code    int64               `json:"code"`
	Results []FootballTableLive `json:"results"`
}

type FootballTableLive struct {
	SeasonId uint64           `json:"season_id" mapstructure:"season_id" msgpack:"season_id"`
	Tables   []FootballTables `json:"tables" mapstructure:"tables" msgpack:"tables"`
}
type FootballTables struct {
	ID         uint64               `json:"id" mapstructure:"id" msgpack:"id"`                         //积分榜表id
	Conference string               `json:"conference" mapstructure:"conference" msgpack:"conference"` //分区信息（极少部分赛事才有，比如美职联）
	Group      uint64               `json:"group" mapstructure:"group" msgpack:"group"`                //不为0表示分组赛的第几组，1-A、2-B以此类推
	StageID    uint64               `json:"stage_id" mapstructure:"stage_id" msgpack:"stage_id"`       //所属阶段id
	Rows       []FootballTablesRows `json:"rows"`
}
type FootballTablesRows struct {
	TeamID           uint64 `json:"team_id" mapstructure:"team_id" msgpack:"team_id"`                                  //球队id
	PromotionID      uint64 `json:"promotion_id" mapstructure:"promotion_id" msgpack:"promotion_id"`                   //升降级id
	Points           uint64 `json:"points" mapstructure:"points" msgpack:"points"`                                     //积分
	Position         uint64 `json:"position" mapstructure:"position" msgpack:"position"`                               //排名
	DeductPoints     uint64 `json:"deduct_points" mapstructure:"deduct_points" msgpack:"deduct_points"`                //扣除积分
	Note             string `json:"note" mapstructure:"note" msgpack:"note"`                                           //说明
	Total            uint64 `json:"total" mapstructure:"total" msgpack:"total"`                                        //比赛场次
	Won              uint64 `json:"won" mapstructure:"won" msgpack:"won"`                                              //胜的场次
	Draw             uint64 `json:"draw" mapstructure:"draw" msgpack:"draw"`                                           //平的场次
	Loss             uint64 `json:"loss" mapstructure:"loss" msgpack:"loss"`                                           //负的场次
	Goals            uint64 `json:"goals" mapstructure:"goals" msgpack:"goals"`                                        //进球
	GoalsAgainst     uint64 `json:"goals_against" mapstructure:"goals_against" msgpack:"goals_against"`                //失球
	GoalDiff         uint64 `json:"goal_diff" mapstructure:"goal_diff" msgpack:"goal_diff"`                            //净胜球
	HomePoints       uint64 `json:"home_points" mapstructure:"home_points" msgpack:"home_points"`                      //主场积分
	HomePosition     uint64 `json:"home_position" mapstructure:"home_position" msgpack:"home_position"`                //主场排名
	HomeTotal        uint64 `json:"home_total" mapstructure:"home_total" msgpack:"home_total"`                         //主场比赛场次
	HomeWon          uint64 `json:"home_won" mapstructure:"home_won" msgpack:"home_won"`                               //主场胜的场次
	HomeDraw         uint64 `json:"home_draw" mapstructure:"home_draw" msgpack:"home_draw"`                            //主场平的场次
	HomeLoss         uint64 `json:"home_loss" mapstructure:"home_loss" msgpack:"home_loss"`                            //主场负的场次
	HomeGoals        uint64 `json:"home_goals" mapstructure:"home_goals" msgpack:"home_goals"`                         //主场进球
	HomeGoalsAgainst uint64 `json:"home_goals_against" mapstructure:"home_goals_against" msgpack:"home_goals_against"` //主场失球
	HomeGoalDiff     uint64 `json:"home_goal_diff" mapstructure:"home_goal_diff" msgpack:"home_goal_diff"`             //主场净胜球
	AwayPoints       uint64 `json:"away_points" mapstructure:"away_points" msgpack:"away_points"`                      //客场积分
	AwayPosition     uint64 `json:"away_position" mapstructure:"away_position" msgpack:"away_position"`                //客场排名
	AwayTotal        uint64 `json:"away_total" mapstructure:"away_total" msgpack:"away_total"`                         //客场比赛场次
	AwayWon          uint64 `json:"away_won" mapstructure:"away_won" msgpack:"away_won"`                               //客场胜的场次
	AwayDraw         uint64 `json:"away_draw" mapstructure:"away_draw" msgpack:"away_draw"`                            //客场平的场次
	AwayLoss         uint64 `json:"away_loss" mapstructure:"away_loss" msgpack:"away_loss"`                            //客场负的场次
	AwayGoals        uint64 `json:"away_goals" mapstructure:"away_goals" msgpack:"away_goals"`                         //客场进球
	AwayGoalsAgainst uint64 `json:"away_goals_against" mapstructure:"away_goals_against" msgpack:"away_goals_against"` //客场失球
	AwayGoalDiff     uint64 `json:"away_goal_diff" mapstructure:"away_goal_diff" msgpack:"away_goal_diff"`             //客场净胜球

}
type CronGetBasketballCompetitionTableDetailRsp struct {
	Code    int64                       `json:"code"`
	Results BasketballCompetitionResult `json:"results"`
}
type BasketballCompetitionResult struct {
	//SeasonId uint64                        `json:"season_id" mapstructure:"season_id" msgpack:"season_id"`
	Tables []BasketballCompetitionTables `json:"tables" mapstructure:"tables" msgpack:"tables"`
}
type BasketballCompetitionTables struct {
	ID      int64                  `json:"id"`
	Scope   int64                  `json:"scope"` //统计范围，1-赛季、2-预选赛、3-小组赛、4-季前赛、5-常规赛、6-淘汰赛(季后赛)、0-无
	Name    string                 `json:"name"`
	StageID int64                  `json:"stage_id"`
	Rows    []BasketballTablesRows `json:"rows" mapstructure:"rows" msgpack:"rows"`
}

//	type CronGetBasketballTableLiveRsp struct {
//		Code    int64                 `json:"code"`
//		Results []BasketballTableLive `json:"results"`
//	}
//
//	type BasketballTableLive struct {
//		SeasonId uint64             `json:"season_id" mapstructure:"season_id" msgpack:"season_id"`
//		Tables   []BasketballTables `json:"tables" mapstructure:"tables" msgpack:"tables"`
//	}
//
//	type BasketballTables struct {
//		ID         uint64                 `json:"id" mapstructure:"id" msgpack:"id"`                         //积分榜表id
//		Conference string                 `json:"conference" mapstructure:"conference" msgpack:"conference"` //分区信息（极少部分赛事才有，比如美职联）
//		Group      uint64                 `json:"group" mapstructure:"group" msgpack:"group"`                //不为0表示分组赛的第几组，1-A、2-B以此类推
//		StageID    uint64                 `json:"stage_id" mapstructure:"stage_id" msgpack:"stage_id"`       //所属阶段id
//		Rows       []BasketballTablesRows `json:"rows"`
//	}
type BasketballTablesRows struct {
	TeamID           uint64  `json:"team_id" mapstructure:"team_id" msgpack:"team_id"`                                  //球队id
	PromotionID      uint64  `json:"promotion_id" mapstructure:"promotion_id" msgpack:"promotion_id"`                   //升降级id
	Position         uint64  `json:"position" mapstructure:"position" msgpack:"position"`                               //排名
	Group            uint64  `json:"group" mapstructure:"group" msgpack:"group"`                                        //分组，1.季后赛 2.附加赛资格
	Note             string  `json:"note" mapstructure:"note" msgpack:"note"`                                           //说明
	Won              uint64  `json:"won" mapstructure:"won" msgpack:"won"`                                              //胜场=主场胜+客场胜
	Lost             uint64  `json:"lost" mapstructure:"lost" msgpack:"lost"`                                           //负场=主场负+客场负
	WonRate          float64 `json:"won_rate" mapstructure:"won_rate" msgpack:"won_rate"`                               //胜率
	GameBack         string  `json:"game_back" mapstructure:"game_back" msgpack:"game_back"`                            //胜场差，可能不存在
	PointsAvg        float64 `json:"points_avg" mapstructure:"points_avg" msgpack:"points_avg"`                         //场均得分，可能不存在
	PointsAgainstAvg float64 `json:"points_against_avg" mapstructure:"points_against_avg" msgpack:"points_against_avg"` //场均失分，可能不存在
	DiffAvg          float64 `json:"diff_avg" mapstructure:"diff_avg" msgpack:"diff_avg"`                               //场均净胜，可能不存在
	Streaks          uint64  `json:"streaks" mapstructure:"streaks" msgpack:"streaks"`                                  //近期战绩（正连胜，负连败），可能不存在
	Home             string  `json:"home" mapstructure:"home" msgpack:"home"`                                           //主场：主场胜-主场负，可能不存在
	Away             string  `json:"away" mapstructure:"away" msgpack:"away"`                                           //客场：客场胜-客场负，可能不存在
	Division         string  `json:"division" mapstructure:"division" msgpack:"division"`                               //赛区胜-负，该球队在该赛区的胜负数据，可能不存在
	Conference       string  `json:"conference" mapstructure:"conference" msgpack:"conference"`                         //东（西）部胜-负，该球队在东或西部的胜负数据，可能不存在
	Last10           string  `json:"last_10" mapstructure:"last_10" msgpack:"last_10"`                                  //近10场胜-负（在该赛季只打了5场比赛：4-1），可能不存在
	Points           uint64  `json:"points" mapstructure:"points" msgpack:"points"`                                     //积分，杯赛存在
	PointsFor        uint64  `json:"points_for" mapstructure:"points_for" msgpack:"points_for"`                         //总得分，杯赛存在
	PointsAgt        uint64  `json:"points_agt" mapstructure:"points_agt" msgpack:"points_agt"`                         //总失分，杯赛存在

}
