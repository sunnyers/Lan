package config

const SUCCESS = 200
/**
 * 非法的请求
 */
const ILLEGAL_REQUEST = 401
/**
 * 被禁止的请求
 */
const FORBIDDEN_REQUEST = 403
/**
 * 请求不存在
 */
const REQUEST_NOT_EXIST = 404
/**
 * 内部服务器错误
 */
const INTERNAL_SERVER_ERROR = 500
/**
 * 添加失败
 */
const DATA_INSERT_ERROR = 501
/**
 * 删除失败
 */
const DATA_DELETE_ERROR = 502
/**
 * 保存失败
 */
const DATA_UPDATE_ERROR = 503
/**
 * 数据查询失败
 */
const DATA_QUERY_ERROR = 504
/**
 * 系统错误
 */
const SYSTEM_ERROR = -1
/**
 * sign参数异常
 */
const COMMUNAL_PARAM_SIGN_INVALID=-10
/**
 * appkey参数异常
 */
const COMMUNAL_PARAM_APPKEY_INVALID=-11
/**
 * timestamp参数异常
 */
const COMMUNAL_PARAM_TIMESTAMP_INVALID=-12
/**
 * 参数不合法
 */
const PARAM_INVALID=-13
/**
 * session参数异常
 */
const COMMUNAL_PARAM_SESSIONID_INVALID=-14
/**
 * 参数不能为空
 */
const PARAM_EMPTY =-16
/**
 * 未登录
 */
const USER_NOT_LOGIN=-17
/**
 * 节目不存在
 */
const PLATFORM_INEXISTENCE = -1131
/**
 * 权限不足
 */
const PERMISSION_DENIED = -1132
/**
 * 重复开播
 */
const REPEAT_PLAY = -1133
/**
 * 重复停播
 */
const REPEAT_STOP = -1134
/**
 * 节目没有开播
 */
const PROGRAM_NOT_PLAY = -1135
/**
 * 进入直播间时候密码不正确
 */
const ROOM_PASSWORD_ERROR = -1136
/**
 * 房间人数已满
 */
const ROOM_FULL = -1137
/**
 * 你被踢出直播间，不能进入
 */
const ROOM_KICK = -1138
/**
 * 时长过大
 */
const TIME_OVERSIZE = -1139
/**
 * 内容超出指定长度
 */
const CONTENT_EXCEEDS = -1140
/**
 * 用户不在列表
 */
const ROOM_USER_INEXISTENCE = -1141
/**
 * 礼物类型不支持
 */
const GIFT_TYPE_ERROR = -1142
/**
 * 每条广播限制50个字！
 */
const BROADCAST_EXCEED = -1143
/**
 * 非法的操作类型
 */
const ILLEGAL_TYPE = -1144
/**
 * 昵称不能为空
 */
const NIKENAME_NOTEMPTY = -1145
/**
 * 不支持该开放平台
 */
const PLATFORM_ERROR = -1201
/**
 * 开放登录失败
 */
const OPEN_LOGIN_ERROR = -1202
/**
 * 非法的注册账号
 */
const USER_ILLEGALITY = -1203
/**
 * 两次密码不一致
 */
const PASSWORD_UNLIKE = -1204
/**
 * 验证码错误
 */
const VERIFICATION_ERROR = -1205
/**
 * 账号密码错误
 */
const USER_PASSWORD_ERROR = -1206
/**
 * 您输入的文本中包含非法字符
 */
const ILLEGAL_STRING = -1207
/**
 * 账户被禁用
 */
const USR_DISABLED = -1208
/**
 * 昵称已存在
 */
const NICKNAME_REPETITION = -1209
/**
 * 用户不存在
 */
const USER_INEXISTENCE = -1210
/**
 * 余额不足
 */
const NOT_SUFFICIENT_FUNDS = -1211
/**
 * 头像修改失败
 */
const USER_ICONS_ERROR = -1212
/**
 * 昵称长度不合规则
 */
const NICKNAME_LENGTH_ERROR = -1213
/**
 * 请不要上传超过1M的图片
 */
const IMAGE_SIZE_ERROR = -1214
/**
 * 内容长度不符合大小
 */
const CONTENT_SIZE_ERROR = -1215
/**
 * 账号已经存在
 */
const ACCOUNT_EXIST  = -1216
/**
 * 顺网登录失败
 */
const OPEN_LOGIN_BY_SHUNWANG_ERROR = -1217

/**
 * 手机号不正确
 */
const MOBILE_ERROR = -1218

/**
 * 发送短信失败
 */
const SEND_SMS_ERROR = -1219

/**
 * 举报失败
 */
const REPORT_ERROR = -1220

/**
 * 背包数量不足
 */
const NOT_BAG_NUMBER = -1221

/**
 * 用户无权开播（关播）节目
 */
const NOT_LIVE_USER = -1222

/**
 * 其他端开播",不能修改流地址
 */
const NOT_UPDATE_STREAM = -1223

/**
 * 节目封面修改失败
 */
const PROGRAM_COVER_ERROR = -1224

/**
 * 绑定手机号失败
 */
const BIND_MOBILE_ERROR = -1225

/**
 * 主播申请未提交
 */
const APPLY_NOT_SUBMIT = -1226

/**
 * 图片大小不符合规定
 */
const IMAGE_SIZE_NOT_MATCH = -1227

/**
 * 未开放主播修改节目封面权限
 */
const NOT_OPEN_UPDATE_COVER_AUTH = -1228

/**
 * 氧币金额不足打掉当前跑道
 */
const RUNWAY_NOT_MATCH = -1229

/**
 * 不满足送礼条件
 */
const NOT_MATCH_SEND_GIFT = -1230

/**
 * 盛天登录失败
 */
const OPEN_LOGIN_BY_SHENGTIAN_ERROR = -1231

/**
 * 签到失败
 */
const SIGN_DO_ERROR = -1232

/**
 * 补签失败
 */
const MEND_DO_ERROR = -1233

/**
 * 该手机号已被设为登录号
 */
const PHONE_IS_BIND_LOGIN = -1234

/**
 * 登录号设置失败
 */
const LOGIN_ACCOUNT_SETTING_ERROR = -1235

/**
 * 用户不是主播
 */
const USER_NOT_ANCHOR = -1236

/**
 * 红包个数错误
 */
const RED_PACKED_QUANTITY_ERROR = -1237

/**
 * 红包金额错误
 */
const RED_PACKED_AMOUNT_ERROR = -1238

/**
 * 未绑定手机号
 */
const NOT_BIND_PHONE = -1239


var codeMap = map[int]string {
	200 : "success",
	401 : "非法请求",
	403 : "被禁止的请求",
	404 : "请求不存在",
	500 : "内部服务器错误",
	501 : "添加失败",
	502 : "删除失败",
	503 : "修改失败",
	504 : "查询失败",

	-1  : "System error",
	-10 : "sign参数异常",
	-11 : "appkey参数异常",
	-12 : "timestamp参数异常",
	-13 : "参数不合法",
	-14 : "session参数异常",
	-16 : "参数不能为空",
	-17 : "用户未登录",

	-1131 : "节目不存在",
	-1132 : "权限不足",
	-1133 : "重复开播",
	-1134 : "重复停播",
	-1135 : "节目没有开播",
	-1136 : "进入直播间时候密码不正确",
	-1137 : "房间人数已满",
	-1138 : "你被踢出直播间，不能进入",
	-1139 : "时长过大",
	-1140 : "内容超出指定长度",
	-1141 : "用户不在列表",
	-1142 : "礼物类型不支持",
	-1143 : "每条广播限制50个字！",
	-1144 : "非法的操作类型",
	-1145 : "昵称不能为空",

	-1201 : "不支持该开放平台",
	-1202 : "开放登录失败",
	-1203 : "非法的注册账号",
	-1204 : "两次密码不一致",
	-1205 : "验证码错误",
	-1206 : "账号密码错误",
	-1207 : "您输入的文本中包含非法字符",
	-1208 : "账户被禁用",
	-1209 : "昵称已存在",
	-1210 : "用户不存在",
	-1211 : "余额不足",
	-1212 : "头像修改失败",
	-1213 : "昵称长度不合规则",
	-1214 : "请不要上传超过1M的图片",
	-1215 : "内容大小不符合",
	-1216 : "账号已经存在",
	-1217 : "顺网登录失败",
	-1218 : "手机号不正确",
	-1219 : "发送短信失败",
	-1220 : "举报失败",
	-1221 : "背包数量不足",
	-1222 : "用户无权开播(关闭)节目",
	-1223 : "其他端开播,不能修改流地址",
	-1224 : "节目封面修改失败",
	-1225 : "绑定手机号失败",
	-1226 : "主播申请未提交",
	-1227 : "图片大小不合法规定",
	-1228 : "未开放修改主播封面权限",
	-1229 : "氧币金额不足打掉当前跑道",
	-1230 : "不满足送礼条件,不是守护",

	-1231 : "米果登录失败",
	-1232 : "签到失败",
	-1233 : "补签失败",
	-1234 : "该手机已被设为登录号",
	-1235 : "登录号设置失败",
	-1236 : "用户不是主播",
	-1237 : "红包个数错误",
	-1238 : "红包金额错误",
	-1239 : "未绑定手机号",
}

func GetDesc(code int) string{
	if v,ok := codeMap[code]; ok{
		return v
	}
	return ""
}