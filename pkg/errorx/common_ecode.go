package errorx

// All common ecode
var (
	OK = New(0, "success")

	CreateExistingUser    = New(112, "Create an already existing user")
	UserOrPasswordInvalid = New(113, "User/Password invalid")
	UserNotFound          = New(110, "该用户不存在!")
	UserDisabled          = New(111, "User is disabled")

	ServerErr = New(500, "服务器错误！")

	BadRequest          = New(400, "Invalid request")
	UNAUTHORIZED        = New(401, "Unauthorized")
	FORBIDDEN           = New(403, "Forbidden")
	VerificationCodeErr = New(405, "ErrVerificationCode")

	RoleSuperiorNotExist        = New(900, "所选上级不存在")
	RoleIsExist                 = New(901, "角色不存在")
	RoleSuperiorSameWithHimSelf = New(902, "所属上级和自己一致，数据异常")
	RoleHasBeenUsed             = New(903, "角色已被使用，无法删除")

	ValidaCodeErr = New(1000, "验证码已过期||验证码错误")

	MerchantRegisterErr     = New(1001, "注册失败！")
	MerchantLoginNotInitErr = New(1002, "尚未完成数据初始化，请重新注册！")
	MerchantIsRegisterErr   = New(1003, "用户已存在！")
	MerchantNotExistErr     = New(1004, "用户不存在！")
	GetMerchantInfoErr      = New(1005, "获取商家用户信息失败!")

	HotelUploadCoverErr     = New(1101, "至少上传一张图片!")
	UpdateHotelInfoErr      = New(1102, "保存酒店信息失败!")
	UpdateAptitudeInfoErr   = New(1103, "保存资质信息失败!")
	UpdateSettlementInfoErr = New(1104, "保存结算信息失败!")
	UpdateContractInfoErr   = New(1105, "保存合约信息失败!")
	ChildAccountIsExist     = New(1006, "该成员已添加至酒店!")

	GetHotelRoomListErr = New(1201, "获取房型列表失败!")
	CreateHotelRoomErr  = New(1202, "创建房型失败!")
	GetHotelCoverImgErr = New(1203, "获取酒店封面图失败!")
	SaveImgErr          = New(1204, "保存图片失败！")
	GetNotAuditImgErr   = New(1205, "获取未审核图片失败!")

	HotelRoomIsExistErr       = New(1300, "该房型已存在!")
	HotelRoomNotExistErr      = New(1301, "该房型不存在!")
	HotelNotExistErr          = New(1302, "该酒店不存在!")
	HotelRoomStateNotExistErr = New(1303, "该房态不存在！")

	GetOrderInfoErr = New(1400, "订单生成中!")
	CheckInHotelErr = New(1401, "确认入住失败,请联系管理员!")

	FSMCreateErr    = New(2000, "FSM create err")       //状态机创建失败
	FSMGetStatusErr = New(2100, "FSM status not found") //状态机状态

)
