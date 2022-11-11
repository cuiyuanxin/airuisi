package app

import (
	"github.com/cuiyuanxin/airuisi/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

//type Pager struct {
//	Page      int `json:"pageNo"`
//	Limit     int `json:"pageSize"`
//	TotalRows int `json:"totalCount"`
//	TotalPage int `json:"totalPage"`
//}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//func (r *Response) ToResponse(data interface{}, err *errcode.Error) {
//	response := gin.H{
//		"code": err.Code(),
//		"msg":  err.Msg(),
//	}
//	if data != nil {
//		response["data"] = data
//	}
//	r.Ctx.JSON(http.StatusOK, response)
//}

//func (r *Response) ToResponseList(list interface{}, totalRows, totalPage int) {
//	r.Ctx.JSON(http.StatusOK, gin.H{
//		"list": list,
//		"pager": Pager{
//			Page:      GetPage(r.Ctx),
//			Limit:     GetLimit(r.Ctx),
//			TotalRows: totalRows,
//			TotalPage: totalPage,
//		},
//	})
//}

//func (r *Response) ToSuccessResponse(err *errcode.Error) {
//	response := gin.H{"code": err.Code(), "msg": err.Msg()}
//	details := err.Details()
//	if len(details) > 0 {
//		response["details"] = details
//	}
//
//	r.Ctx.JSON(http.StatusOK, response)
//}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
