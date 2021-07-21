package app

import "encoding/json"

const (
	WatchQuestion              = 31 //问题围观
	QuestionAnswer             = 32 //问题回答
	RewardQuestionAnswer       = 73 //悬赏问题回答
	RewardQuestionAdopt        = 74 //悬赏问题采纳
	RewardWatchQuestion        = 75 //悬赏问题围观
	RewardWatchAnswer          = 76 //悬赏回复围观
	ReplyComment               = 33 //回复评论(文章)
	ReplyCommentShortVideo     = 56 //回复评论(短视频)
	ReplyCommentLesson         = 72 //回复评论(直播课节)
	SpColumnAuditFailed        = 14 //专栏申请审核(未通过)
	ArticleReward              = 34 //文章赞赏
	ShortVideoReward           = 52 //短视频赞赏
	QuestionReward             = 51 //问答赞赏
	CourseReward               = 35 //课程赞赏
	ArticleComment             = 8  //文章评论
	ShortVideoComment          = 53 //短视频评论
	CourseAppraise             = 6  //课程评价
	LessonAppraise             = 71 //直播课节评价
	ArticlePay                 = 36 //文章付费
	ShortVideoPay              = 54 //短视频付费
	CoursePay                  = 37 //课程付费
	QuestionAsk                = 38 //向我提问
	WatchAnswer                = 39 //围观问答
	SpColumnAuditSuccess       = 40 //专栏申请审核（通过）
	ArticleAudit               = 17 //文章审核
	ShortVideoAudit            = 50 //短视频审核
	CourseAudit                = 16 //课程审核
	SpColumnInfoAudit          = 13 //专栏资料审核
	SpColumnCertAudit          = 15 //专栏认证审核
	OffArticle                 = 64 //下架文章
	OffSmallVideo              = 65 //下架视频
	OffCourse                  = 66 //下架课程
	OffQuestion                = 67 //下架问答
	DisabledLive               = 68 //禁用直播问
	LiveOpen                   = 41 //直播间开通
	ArticleNotice              = 42 //专栏通知(文章)
	ShortVideoNotice           = 55 //专栏通知(短视频)
	LiveNotice                 = 43 //专栏通知(直播)
	DistributionOpen           = 44 //开通推广
	DistributionProfit         = 45 //推广收益
	NewRedBag                  = 46 //拉新红包
	DistributionUpLevel        = 47 //推广升级
	VipRenew                   = 48 //会员续费
	VipTimeout                 = 49 //会员到期
	AnchorRemindBeforeDay      = 60 //开播提醒(主播提前一天开播提醒)
	AnchorRemindBeforeHalfHour = 61 //开播提醒(主播提前半小时开通提醒)
	UserRemindBeforeMinutes    = 62 //用户提前5分钟开播提醒
	UserRemindBySpColumn       = 63 //用户关注专栏开播提醒
	SystemNotice               = 0  //系统通知
	LiveCoursePay              = 81 //系列课付费
	LiveCourseReward           = 82 //系列课赞赏
	QuestionAnswerSettle       = 30 //问题回答付费
	SimStockOperateNotify      = 90 //模拟炒股操作股票通知
)

type JSSMessage struct {
	Id      string `json:"id"`
	Type    int    `json:"type"`
	State   int    `json:"state"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
	EvId    string `json:"evId"`
}

func InitAppMsg(id, content, userId, evId string, iType, state int) JSSMessage {
	return JSSMessage{
		Id: id,
		Type: iType,
		State: state,
		Content: content,
		UserId: userId,
		EvId: evId,
	}
}

func (msg JSSMessage)BuildToJson() string{
	bytes, err := json.Marshal(msg)
	if err != nil {
		return ""
	} else {
		return string(bytes)
	}
}
