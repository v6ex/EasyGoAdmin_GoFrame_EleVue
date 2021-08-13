// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

//跨域处理中间件
func CORS(r *ghttp.Request) {
	fmt.Println("跨域处理中间件")
	r.Response.CORSDefault()
	// 前置中间件
	r.Middleware.Next()
}
