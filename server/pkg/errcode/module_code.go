package errcode

var (
	ErrorUsersCheckMobileIsNo      = NewError(20010000, "手机号不存在")
	ErrorUsersCheckMobileIsDeleted = NewError(20010001, "手机号已禁用")
	ErrorUsersSendCodeFail         = NewError(20010002, "手机验证码发送失败")
	ErrorUsersCodeFault            = NewError(20010003, "验证码错误")
	ErrorUsersCodeOverdue          = NewError(20010004, "验证码已过期")
	ErrorUsersCodeStatus           = NewError(20010005, "验证码还在有效期内，请勿多次申请")
)
