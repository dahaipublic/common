package model

const (
	CronGetDataNum  = 1000 //定时获取纳米数据数量
	CronGetDataNum2 = 500

	//比赛类型
	MatchTypePlayback   int8 = -2 // 回放
	MatchTypeResult     int8 = -1 // 赛果
	MatchTypeHot        int8 = 0  //热门
	MatchTypeFootball   int8 = 1  //足球
	MatchTypeBasketball int8 = 2  //篮球

	//比赛状态
	// 0 比赛异常，说明：暂未判断具体原因的异常比赛，可能但不限于：腰斩、取消等等，建议隐藏处理
	MatchFootballStatusIDNotStart int8 = 1 //足球未开赛状态
	// 2 上半场
	MatchFootballStatusIDMidfield   int8 = 3 //足球中场状态
	MatchFootballStatusIDSecondHalf int8 = 4 //足球下半场状态
	// 5	加时赛
	// 6	加时赛(弃用)
	// 7	点球决战
	MatchFootballStatusIDCompletion int8 = 8 //足球完场状态
	// 9	推迟
	// 10	中断
	// 11	腰斩
	MatchFootballStatusIDCancel int8 = 12 // 12	取消
	// 13	待定

	// 0 比赛异常，说明：暂未判断具体原因的异常比赛，可能但不限于：腰斩、取消等等，建议隐藏处理
	MatchBasketballStatusIDNotStart int8 = 1 //篮球未开赛状态
	// 2	第一节
	// 3	第一节完
	// 4	第二节
	MatchBasketballStatusIDSection2Finish int8 = 5 //篮球第二节完状态
	MatchBasketballStatusIDSection3       int8 = 6 //篮球第三节状态
	// 7	第三节完
	// 8	第四节
	// 9	加时
	MatchBasketballStatusIDCompletion int8 = 10 //篮球完场状态
	// 11	中断
	MatchBasketballStatusIDCancel int8 = 12 // 12	取消
	// 13	延期
	// 14	腰斩
	// 15	待定

	//redis
	RedisExpiration          = 7200       //2个小时
	RedisExpirationFourHours = 14400      //4个小时
	RedisExpirationOneDay    = 86400      //1天
	RedisExpirationOneWeek   = 86400 * 7  //1周
	RedisExpirationOneMonth  = 86400 * 30 //1周

	//activity
	RedisKeyActivityList = "activity:list"
	RedisKeyActivityID   = "activity:id:%d"

	//activity_category
	RedisKeyActivityCategoryID = "activityCategory:id:%d"

	//advertisementInfo
	RedisKeyAdvertisementMatchInfo = "advertisementMatch:info:%d:%d:%d"
	RedisKeyAdvertisementID        = "advertisement:id:%d"

	//app_advertisementInfo
	RedisKeyAppAdvertisementListType = "appAdvertisement:list:%d"
	RedisKeyAppAdvertisementList     = "appAdvertisement:list"
	RedisKeyAppAdvertisementID       = "appAdvertisement:id:%d"

	//config
	RedisKeyConfigKey = "config:key:%s"

	//basketball_competition
	RedisKeyBasketballCompetitionID = "basketballCompetition:id:%d"

	//basketball_team
	RedisKeyBasketballTeamID = "basketballTeam:id:%d"

	//football_competition
	RedisKeyFootballCompetitionID = "footballCompetition:id:%d"

	//football_team
	RedisKeyFootballTeamID = "footballTeam:id:%d"

	//coach
	RedisKeyFootballCoachID   = "footballCoach:id:%d"
	RedisKeyBasketballCoachID = "basketballCoach:id:%d"

	//hot_competition
	RedisKeyHotCompetition     = "hotCompetition:%d:%d"
	RedisKeyHotCompetitionList = "hotCompetition:list:%d"

	//recommend_competition
	RedisKeyRecommendCompetition     = "hotCompetition:%d:%d"
	RedisKeyRecommendCompetitionList = "hotCompetition:list:%d"

	RedisKeyBallEGCompetition     = "ballEGCompetition:%d:%d"
	RedisKeyBallEGCompetitionList = "ballEGCompetition:list:%d"
	//match
	RedisKeyMatchID             = "match:id:%d"
	RedisKeyMatchVideo          = "match:video:%d:%d"
	RedisKeyHotMatchList        = "hotMatch:list"        //热门列表
	RedisKeyResultMatchList     = "resultMatch:list"     //赛果列表
	RedisKeyPlaybackMatchList   = "playbackMatch:list"   //回放列表
	RedisKeyFootballMatchList   = "footballMatch:list"   //足球比赛列表
	RedisKeyBasketballMatchList = "basketballMatch:list" //篮球比赛列表

	//RedisKeyFootballMatchListDateCompetitionID   = "footballMatch:list:%d:%d"   //足球赛事比赛列表
	//RedisKeyBasketballMatchListDateCompetitionID = "basketballMatch:list:%d:%d" //篮球赛事比赛列表

	RedisKeyHotMatchNewList = "hotMatchNew:list" //热门比赛列表
	//RedisKeyHotMatchNBAList                            = "hotMatchNew:nba"                   //热门比赛列表
	RedisKeyRecommendMatchList                         = "recommendMatch:list"               //推荐比赛列表
	RedisKeyBallEGMatchList                            = "ballEGMatch:list"                  //推荐比赛列表
	RedisKeyFootballMatchAnchorLiveList                = "footballMatch:anchorLiveList"      //足球比赛主播直播列表
	RedisKeyBasketballMatchAnchorLiveList              = "basketballMatch:anchorLiveList"    //篮球比赛主播直播列表
	RedisKeyFootballMatchAnchorLiveListCompetitionID   = "footballMatch:anchorLiveList:%d"   //足球赛事比赛主播直播列表
	RedisKeyBasketballMatchAnchorLiveListCompetitionID = "basketballMatch:anchorLiveList:%d" //篮球赛事比赛主播直播列表
	RedisKeyMatchSearch                                = "match:search:%s"                   //比赛搜索
	RedisKeyMatchWatchLiveRewards                      = "match:watchLiveRewards:%s:%d"      //每日在线奖励
	RedisKeyMatchSnapshot                              = "match:snapshot:%d"                 //比赛截图

	//match_info
	RedisKeyMatchInfoMatchID       = "matchInfo:%d"
	RedisKeyMatchInfoType          = "matchInfo:type:%d:%d"     //比赛详情类型(matchID:userID)
	RedisKeyMatchMessage           = "match:message:%d:%v"      //比赛聊天数据
	RedisKeyMatchMessageIndex      = "match:message:index:%d"   //比赛聊天数据
	RedisKeyMatchLive              = "match:live:%d"            //比赛直播数据
	RedisKeyMatchFootballTableLive = "match:football:tables:%d" //足球积分榜数据

	RedisKeyMatchAnalysisLan      = "match:analysis:%d:%d"       //比赛分析数据
	RedisKeyMatchLineupLan        = "match:lineup:%d:%d"         //比赛阵容数据
	RedisKeyMatchLineupGetDataLan = "match:lineup:getData:%d:%d" //比赛获取阵容数据缓存

	RedisKeyMatchAnalysis       = "match:analysis:%d"        //比赛分析数据
	RedisKeyFootballMatchLineup = "match:football:lineup:%d" //足球比赛阵容数据

	RedisKeyBasketballMatchLineup      = "match:basketball:lineup:%d" //比赛阵容数据
	RedisKeyMatchBasketballTableDetail = "match:basketball:tables:%d" //足球积分榜数据
	RedisKeyMatchLineupGetData         = "match:lineup:getData:%d"    //比赛获取阵容数据缓存

	//match_live
	RedisKeyMatchLiveMatchID          = "matchLive:matchID:%d"
	RedisKeyMatchLiveMatchIDAndUserID = "matchLive:matchID:userID:%d_%d"

	//odds
	RedisKeyOddsMatchIDType          = "match:odds:%d:%s"
	RedisKeyOddsMatchID              = "match:odds:%d" //比赛指数matchID
	RedisKeyOddsMatchIDCompanyIDType = "match:odds:%d:%s:%s"

	//sensitive_word
	RedisKeySensitiveWord = "sensitiveWord"

	//登录验证
	RedisKeyFailedUser = "adminUser:login:fail:user:"
	RedisKeyLockUser   = "adminUser:login:lock:user:"
	RedisKeyFailedIP   = "adminUser:login:fail:ip:"
	RedisKeyLockIP     = "adminUser:login:lock:ip:"

	//user
	RedisKeyUserToken              = "user:token:%s"
	RedisKeyUserID                 = "user:%d"
	RedisKeyUserDeviceIDGetID      = "user:device:getID:%s" //通过设备号获取用户id
	RedisKeyUserPhoneGetID         = "user:phone:getID:%s"  //通过手机号获取用户id
	RedisKeyUserSendMessageNum     = "user:sendMessageNum:%s:%d"
	RedisKeyUserDailyOnlineRewards = "user:dailyOnlineRewards:%s:%d" //每日在线奖励

	//socket
	RedisKeySocketDeviceIDClientID = "socket:deviceIDClientID:%s" //设备id下有多个客户端id
	RedisKeySocketClientRoom       = "socket:clientRoom:%s"       //客户端房间信息
	RedisKeySocketMatchRoomClient  = "socket:matchRoom:client:%s" //比赛房间客户端信息
	RedisKeySocketServerList       = "socket:serverList"

	//后台图形验证码
	RedisKeyAdminUserCaptchaImage = "adminUser:captchaImage:%s"
	RedisKeyAdminUserToken        = "adminUser:token:%s"

	//hot_competition
	RedisKeyOpeninstallChannelCode = "openinstall:channelCoded:%s"

	//anchor
	RedisKeyAnchorUserID   = "anchor:userID:%d"
	RedisKeyAnchorHotList  = "anchor:hotList"     //热门主播列表
	RedisKeyAnchorLiveList = "anchor:liveList:%d" //主播直播列表

	//user_attention
	RedisKeyUserAttention     = "userAttention:followUserID:fansUserID:%d:%d" //通过关注id和粉丝id获取信息
	RedisKeyUserAttentionList = "userAttention:list:%d"                       //用户关注列表

	//emoticons
	RedisKeyEmoticonsType = "emoticons:type:%d" //用户关注列表
	RedisKeyEmoticonsCode = "emoticons:code:%s"

	//Sms
	RedisKeySms    = "smsCode:%s:%s"     //短信验证码key
	RedisKeySmsNum = "smsCode:num:%s:%s" //短信验证码次数

	//match_announcement
	RedisKeyMatchAnnouncementMatchID = "matchAnnouncement:matchID:%d"

	//hot_match
	RedisKeyHotMatchMatchID = "hotMatch:matchID:%d"

	//AppVersion
	RedisKeyAppVersionDeviceType = "appVersion:deviceType:%s"

	//watch_history
	RedisKeyWatchHistoryUserIDOrMatchID = "watchHistory:%d:%d"
	RedisKeyWatchHistoryListUserID      = "watchHistory:list:%d"

	//user_level
	RedisKeyUserLevelLevel = "userLevel:level:%d"

	//user_role
	RedisKeyUserRoleID    = "userRole:id:%d"
	RedisKeyUserRoleLevel = "userRole:level:%d"

	//user_experience_log
	RedisKeyUserExperienceLogType = "userExperienceLog:type:%d:%d"

	//statistics
	RedisKeyStatisticsActiveUser       = "statistics:activeUser:%s:%s"    //今日活跃人数
	RedisKeyStatisticsOnlineUser       = "statistics:onlineUser:%s:%d:%s" //在线人数
	RedisKeyStatisticsClientIDDeviceID = "statistics:clientID:%s"         //客户端id判断的设备id
	RedisKeyStatisticsDeviceID         = "statistics:deviceID:%s:%d:%s"   //客户端id判断的设备id

	RedisKeyStatisticsMatchOnlineUser        = "statistics:matchOnlineUser:%s:%d:%d"  //比赛房间在线人数
	RedisKeyStatisticsMatchClientIDDeviceID  = "statistics:matchClientID:%s"          //比赛房间客户端id判断的设备id
	RedisKeyStatisticsMatchDeviceID          = "statistics:matchDeviceID:%s:%d:%d:%s" //比赛房间客户端id判断的设备id
	RedisKeyStatisticsMatchHighestOnlineUser = "statistics:matchHighestOnlineUser:%d" //比赛房间最高在线人数

	//app_function_control
	RedisKeyAppFunctionControl = "appFunctionControl:list" //今日活跃人数

	//分布式锁
	RedisKeySetNX = "setNx:%s"

	//区分纳米的源还是主播
	Nami_Live   = "live"
	Anchor_Live = "anchor"

	RedisKeyMatchVideoPushUrl = "nami:match:video:pushUrl"
	RedisKeyMatchVideoPlayUrl = "nami:match:video:playUrl"

	RedisKeyMobile1Min = "sms:mobile:%s:1min"
)
