// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"easygoadmin/app/model/internal"
)

// Role is the golang structure for table sys_role.
type Role internal.Role

// Fill with you ideas below.

type RolePageReq struct {
	Name  string `p:"name"`  // 角色名称
	Page  int    `p:"page"`  // 页码
	Limit int    `p:"limit"` // 每页数
}

type RoleAddReq struct {
	Name   string `p:"name" v:"required#角色名称不能为空"`
	Code   string `p:"code" v:"required#角色编码不能为空"`
	Status int    `p:"status" v:"required#角色状态不能为空"`
	Sort   int    `p:"sort" v:"required#角色排序不能为空"`
	Note   string `p:"note"` // 备注
}

type RoleUpdateReq struct {
	Id     int    `p:"id" v:"required#主键ID不能为空"`
	Name   string `p:"name" v:"required#角色名称不能为空"`
	Code   string `p:"code" v:"required#角色编码不能为空"`
	Status int    `p:"status" v:"required#角色状态不能为空"`
	Sort   int    `p:"sort" v:"required#角色排序不能为空"`
	Note   string `p:"note"` // 备注
}

type RoleDeleteReq struct {
	Ids string `p:"ids" v:"required#请选择需要删除的数据记录"`
}

type RoleStatusReq struct {
	Id     int `p:"id" v:"required#主键ID不能为空"`
	Status int `p:"status"    v:"required#角色状态不能为空"`
}
