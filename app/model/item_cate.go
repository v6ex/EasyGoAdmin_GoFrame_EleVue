// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// ItemCate is the golang structure for table sys_item_cate.
type ItemCate internal.ItemCate

// Fill with you ideas below.

// 栏目查询条件
type ItemCateQueryReq struct {
	Name   string `p:"name"`    // 栏目名称
	ItemId int    `p:"item_id"` // 站点ID
}

// 添加站点
type ItemCateAddReq struct {
	Name    string `p:"name"        v:"required#栏目名称不能为空"`     // 栏目名称
	Pid     int    `p:"pid"`                                   // 父级ID
	ItemId  uint   `p:"item_id"     v:"required#请选择所属站点"`      // 站点ID
	Pinyin  string `p:"pinyin"      v:"required#栏目拼音(全拼)不能为空"` // 拼音(全)
	Code    string `p:"code"        v:"required#栏目拼音(简拼)不能为空"` // 拼音(简)
	IsCover int    `p:"is_cover"    v:"required#请选择是否有封面"`     // 是否有封面：1是 2否
	Cover   string `p:"cover"`                                 // 封面
	Status  int    `p:"status"      v:"required#请选择栏目状态"`      // 状态：1启用 2停用
	Note    string `p:"note"`                                  // 备注
	Sort    uint   `p:"sort"        v:"required#排序号不能为空"`      // 排序
}

// 修改
type ItemCateUpdateReq struct {
	Id      int    `p:"id" v:"required#主键ID不能为空"`
	Name    string `p:"name"        v:"required#栏目名称不能为空"`     // 栏目名称
	Pid     int    `p:"pid"`                                   // 父级ID
	ItemId  uint   `p:"item_id"     v:"required#请选择所属站点"`      // 站点ID
	Pinyin  string `p:"pinyin"      v:"required#栏目拼音(全拼)不能为空"` // 拼音(全)
	Code    string `p:"code"        v:"required#栏目拼音(简拼)不能为空"` // 拼音(简)
	IsCover int    `p:"is_cover"    v:"required#请选择是否有封面"`     // 是否有封面：1是 2否
	Cover   string `p:"cover"`                                 // 封面
	Status  int    `p:"status"      v:"required#请选择栏目状态"`      // 状态：1启用 2停用
	Note    string `p:"note"`                                  // 备注
	Sort    uint   `p:"sort"        v:"required#排序号不能为空"`      // 排序
}

// 删除栏目
type ItemCateDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

// 栏目信息
type ItemCateInfoVo struct {
	ItemCate
	ItemName string `json:"itemName"` // 栏目名称
}

// 栏目树结构
type CateTreeNode struct {
	ItemCate
	Children []*CateTreeNode `json:"children"` // 子栏目
}
