package base

const (
	CQSSC_API_PJ       = 1
	CQSSC_API_OFFICIAL = 2
	XJSSC_API_PJ       = 1
	XJSSC_API_OFFICIAL = 2
	TJSSC_API_PJ       = 1
	TJSSC_API_OFFICIAL = 2
	YNSSC_API_PJ       = 1
	YNSSC_API_OFFICIAL = 2
	BJPK_API_PJ        = 1
	BJPK_API_OFFICIAL  = 2
)
const (
	CQSSC_TYPE = 1 //重庆时时彩
	XJSSC_TYPE = 2 //新疆时时彩
	BJPK_TYPE  = 3 //北京pk
	TJSSC_TYPE = 4 //天津时时彩
	YNSSC_TYPE = 5 //云南时时彩
)

const (
	ZERO_BALL  = 1  //球0
	ONE_BALL   = 1  //球1
	TWO_BALL   = 2  //球2
	THRID_BALL = 3  //球3
	FOUR_BALL  = 4  //球4
	FIVE_BALL  = 5  //球5
	SIX_BALL   = 6  //球6
	SEVEN_BALL = 7  //球7
	EIGHT_BALL = 8  //球8
	NINE_BALL  = 9  //球9
	TEN_BALL   = 10 //球10
)
const (
	STATUS_YES = 1
	STATUS_NO  = 0
)
const (
	PJ_CQSSC = "http://www.j0024.com/lottery/getAllCqsscAutoList"
	PJ_XJSSC = "http://www.j0024.com/lottery/getAllXjsscAutoList"
	PJ_TJSSC = "http://www.j0024.com/lottery/getAllTjsscAutoList"
	PJ_YNSSC = "http://www.j0024.com/lottery/getAllYnsscAutoList"
	PJ_BJPK  = "http://www.j0024.com/lottery/getAllBjpkAutoList"
)
const (
	OFFICIAL_CQSSC = "http://f.apiplus.net/cqssc-20.json"
	OFFICIAL_XJSSC = "http://f.apiplus.net/xjssc-20.json"
	OFFICIAL_TJSSC = "http://f.apiplus.net/tjssc-20.json"
	OFFICIAL_YNSSC = "http://f.apiplus.net/ynssc-20.json"
	OFFICIAL_BJPK  = "http://f.apiplus.net/bjpk10-20.json"
)
const (
	T_CQSSC = "lottery_cqssc"
	T_XJSSC = "lottery_xjssc"
	T_TJSSC = "lottery_tjssc"
	T_YNSSC = "lottery_ynssc"
	T_BJPK  = "lottery_bjpk"
)
const (
	CQSSC_NAME = "重庆时时彩"
	XJSSC_NAME = "新疆时时彩"
	TJSSC_NAME = "天津时时彩"
	YNSSC_NAME = "云南时时彩"
	BJPK_NAME  = "北京pk"
)
const (
	SSC_SPLIT        = 5  //大小分界线
	SSC_TOTAL_SPLIT  = 23 //总和大小分界线
	BJKL8_SPLIT      = 6  //大小分界线
	BJPK_SPLIT       = 6
	BJPK_TOTAL_SPLIT = 12 //总和大小分界线
)

const (
	BIG_SMALL_TYPE = 1
	ODD_EVEN_TYP   = 2
	TOTAL_TYP      = 3
	STAR_TYP       = 4
	PRED_TYP       = 5
)
const (
	SEND_MSG_TO_FRIEND_URL = "http://localhost:8080/sendMsgToFr"
)
const (
	CQSSC_START_TIME = "10:00:00"
	CQSSC_END_TIME   = "2:00:00"
	XJSSC_START_TIME = "10:00:00"
	XJSSC_END_TIME   = "2:00:00"
	TJSSC_START_TIME = "9:00:00"
	TJSSC_END_TIME   = "23:00:00"
	YNSSC_START_TIME = "9:30:00"
	YNSSC_END_TIME   = "22:30:00"
	BJPK_START_TIME  = "9:00:00"
	BJPK_END_TIME    = "23:59:59"
)
