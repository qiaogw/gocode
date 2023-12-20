syntax = "v1"

import (
"replay/replay.api"
{{- range .Tables }}
	"api-desc/{{.TableUrl}}.api"
{{- end}}
	"admin-desc/api.api"
	"admin-desc/config.api"
	"admin-desc/dept.api"
	"admin-desc/dictdata.api"
	"admin-desc/dicttype.api"
	"admin-desc/loginlog.api"
	"admin-desc/menu.api"
	"admin-desc/migration.api"
	"admin-desc/operalog.api"
	"admin-desc/post.api"
	"admin-desc/role.api"
	"admin-desc/user.api"
	"admin-desc/login.api"
)

info (
	title: "{{.Service}}"//  add title
	desc: "{{.Service}}"//  add description
	author: "{{.Author}}"
	email: "{{.Email}}"
)

{{- range .Tables }}

{{- if .IsAuth }}
//需要登录
{{- else}}
//不需要登录
{{- end }}
@server(
	group : {{.TableUrl}}
	prefix : /{{.Db}}/{{.TableUrl}}
{{- if .IsAuth}}
	jwt: Auth
{{- end }}
)
service {{.Db}} {
	@doc(
		summary: "提取-{{.TableComment}}"
	)
	@handler Get{{.Table}}
	post  /get (Get{{.Table}}Request) returns(Get{{.Table}}Response)

	@doc(
		summary: "列表-{{.TableComment}}"
	)
	@handler List{{.Table}}
	post /list (List{{.Table}}Request) returns(List{{.Table}}Response )

	@doc(
		summary: "创建-{{.TableComment}}"
	)
	@handler Create{{.Table}}
	post  /create (Create{{.Table}}Request) returns(CommonResponse)

	@doc(
		summary: "更新-{{.TableComment}}"
	)
	@handler Update{{.Table}}
	post /update (Update{{.Table}}Request) returns(CommonResponse)

	@doc(
		summary: "删除-{{.TableComment}}"
	)
	@handler Delete{{.Table}}
	post  /delete (Delete{{.Table}}Request) returns(CommonResponse)

{{- if .IsImport}}
	@doc(
		summary: "导出-{{.TableComment}}"
	)
	@handler Export{{.Table}}
	post /export (List{{.Table}}Request) returns (CommonResponse)

	@doc(
	summary: "导出-{{.TableComment}}模板"
	)
	@handler ExportTemplate{{.Table}}
	post /exportTemplate (NullRequest) returns (CommonResponse)

	@doc(
		summary: "导入-{{.TableComment}}"
	)
	@handler Import{{.Table}}
	post /import (ImportRequest) returns (CommonResponse)
{{- end }}
}
{{- end }}


@server(
	group : adminauth
	prefix : /admin/auth
)

service {{.Database}} {
	@doc(
		summary: "验证码"
	)
	@handler Captcha //
	post /captcha (CaptchaRequest) returns (CommonResponse)

	@doc(
		summary: "用户登录"
	)
	@handler Login //
	post /login (LoginRequest) returns (CommonResponse)
}

@server(
	group : adminauth
	prefix : /admin/auth
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "用户验证"
	)
	@handler Verify //
	get /verify (CaptchaRequest) returns (VerifyResponse)
}

@server(
	group : adminapi
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-api"
	)
	@handler AdminGetApi //
	post /api/get (GetAdminApiRequest) returns (CommonResponse)

	@doc(
		summary: "提取-api-根据角色"
	)
	@handler AdminGetApiByRole //
	post /api/getApiByRole (GetAdminApiRequest) returns (CommonResponse)

	@doc(
		summary: "列表-api"
	)
	@handler AdminListApi //
	post /api/list (ListAdminApiRequest) returns (CommonResponse)

	@doc(
		summary: "创建-api"
	)
	@handler AdminCreateApi //
	post /api/create (CreateAdminApiRequest) returns (CommonResponse)

	@doc(
		summary: "更新-api"
	)
	@handler AdminUpdateApi //
	post /api/update (UpdateAdminApiRequest) returns (CommonResponse)

	@doc(
		summary: "删除-api"
	)
	@handler AdminDeleteApi //
	post /api/delete (DeleteAdminApiRequest) returns (CommonResponse)

}

@server(
	group : adminconfig
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-系统配置"
	)
	@handler AdminGetConfig //
	post /config/get (GetAdminConfigRequest) returns (CommonResponse)

	@doc(
		summary: "列表-系统配置"
	)
	@handler AdminListConfig //
	post /config/list (ListAdminConfigRequest) returns (CommonResponse)

	@doc(
		summary: "创建-系统配置"
	)
	@handler AdminCreateConfig //
	post /config/create (CreateAdminConfigRequest) returns (CommonResponse)

	@doc(
		summary: "更新-系统配置"
	)
	@handler AdminUpdateConfig //
	post /config/update (UpdateAdminConfigRequest) returns (CommonResponse)

	@doc(
		summary: "删除-系统配置"
	)
	@handler AdminDeleteConfig //
	post /config/delete (DeleteAdminConfigRequest) returns (CommonResponse)
}

@server(
	group : admindept
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-部门"
	)
	@handler AdminGetDept //
	post /dept/get (GetAdminDeptRequest) returns (CommonResponse)

	@doc(
		summary: "列表-部门"
	)
	@handler AdminListDept //
	post /dept/list (ListAdminDeptRequest) returns (CommonResponse)

	@doc(
		summary: "创建-部门"
	)
	@handler AdminCreateDept //
	post /dept/create (CreateAdminDeptRequest) returns (CommonResponse)

	@doc(
		summary: "更新-部门"
	)
	@handler AdminUpdateDept //
	post /dept/update (UpdateAdminDeptRequest) returns (CommonResponse)

	@doc(
		summary: "删除-部门"
	)
	@handler AdminDeleteDept //
	post /dept/delete (DeleteAdminDeptRequest) returns (CommonResponse)
}

@server(
	group : admindictdata
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-字典数据"
	)
	@handler AdminGetDictData //
	post /dictdata/get (GetAdminDictDataRequest) returns (CommonResponse)

	@doc(
		summary: "列表-字典数据"
	)
	@handler AdminListDictData //
	post /dictdata/list (ListAdminDictDataRequest) returns (CommonResponse)

	@doc(
		summary: "创建-字典数据"
	)
	@handler AdminCreateDictData //
	post /dictdata/create (CreateAdminDictDataRequest) returns (CommonResponse)

	@doc(
		summary: "更新-字典数据"
	)
	@handler AdminUpdateDictData //
	post /dictdata/update (UpdateAdminDictDataRequest) returns (CommonResponse)

	@doc(
		summary: "删除-字典数据"
	)
	@handler AdminDeleteDictData //
	post /dictdata/delete (DeleteAdminDictDataRequest) returns (CommonResponse)
}

@server(
	group : admindicttype
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-字典"
	)
	@handler AdminGetDictType //
	post /dicttype/get (GetAdminDictTypeRequest) returns (CommonResponse)

	@doc(
		summary: "列表-字典"
	)
	@handler AdminListDictType //
	post /dicttype/list (ListAdminDictTypeRequest) returns (CommonResponse)

	@doc(
		summary: "创建-字典"
	)
	@handler AdminCreateDictType //
	post /dicttype/create (CreateAdminDictTypeRequest) returns (CommonResponse)

	@doc(
		summary: "更新-字典"
	)
	@handler AdminUpdateDictType //
	post /dicttype/update (UpdateAdminDictTypeRequest) returns (CommonResponse)

	@doc(
		summary: "删除-字典"
	)
	@handler AdminDeleteDictType //
	post /dicttype/delete (DeleteAdminDictTypeRequest) returns (CommonResponse)
}

@server(
	group : adminloginlog
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {

	@doc(
		summary: "提取-登录日志"
	)
	@handler AdminGetLoginLog //
	post /loginlog/get (GetLoginLogRequest) returns (CommonResponse)

	@doc(
		summary: "列表-登录日志"
	)
	@handler AdminListLoginLog //
	post /loginlog/list (ListLoginLogRequest) returns (CommonResponse)

	@doc(
		summary: "创建-登录日志"
	)
	@handler AdminCreateLoginLog //
	post /loginlog/create (CreateLoginLogRequest) returns (CommonResponse)

	@doc(
		summary: "更新-登录日志"
	)
	@handler AdminUpdateLoginLog //
	post /loginlog/update (UpdateLoginLogRequest) returns (CommonResponse)

	@doc(
		summary: "删除-登录日志"
	)
	@handler AdminDeleteLoginLog //
	post /loginlog/delete (DeleteLoginLogRequest) returns (CommonResponse)

}

@server(
	group : adminmenu
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-用户菜单"
	)
	@handler AdminGetMenuSelf //
	post /menu/getMenu (GetAdminMenuSelfRequest) returns (CommonResponse)

	@doc(
		summary: "提取-用户菜单-根据角色"
	)
	@handler AdminGetMenuByRole //
	post /menu/getMenuByRole (GetAdminMenuRequest) returns (CommonResponse)

	@doc(
		summary: "提取-菜单"
	)
	@handler AdminGetMenu //
	post /menu/get (GetAdminMenuRequest) returns (CommonResponse)

	@doc(
		summary: "列表-菜单"
	)
	@handler AdminListMenu //
	post /menu/list (ListAdminMenuRequest) returns (CommonResponse)

	@doc(
		summary: "创建-菜单"
	)
	@handler AdminCreateMenu //
	post /menu/create (CreateAdminMenuRequest) returns (CommonResponse)

	@doc(
		summary: "更新-菜单"
	)
	@handler AdminUpdateMenu //
	post /menu/update (UpdateAdminMenuRequest) returns (CommonResponse)

	@doc(
		summary: "删除-菜单"
	)
	@handler AdminDeleteMenu //
	post /menu/delete (DeleteAdminMenuRequest) returns (CommonResponse)
}

@server(
	group : adminmigration
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-版本"
	)
	@handler AdminGetMigration //
	post /migration/get (GetAdminMigrationRequest) returns (CommonResponse)

	@doc(
		summary: "列表-版本"
	)
	@handler AdminListMigration //
	post /migration/list (ListAdminMigrationRequest) returns (CommonResponse)

	@doc(
		summary: "创建-版本"
	)
	@handler AdminCreateMigration //
	post /migration/create (CreateAdminMigrationRequest) returns (CommonResponse)

	@doc(
		summary: "更新-版本"
	)
	@handler AdminUpdateMigration //
	post /migration/update (UpdateAdminMigrationRequest) returns (CommonResponse)

	@doc(
		summary: "删除-版本"
	)
	@handler AdminDeleteMigration //
	post /migration/delete (DeleteAdminMigrationRequest) returns (CommonResponse)
}

@server(
	group : adminoperalog
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-操作日志"
	)
	@handler AdminGetOperaLog //
	post /operalog/get (GetAdminOperaLogRequest) returns (CommonResponse)

	@doc(
		summary: "列表-操作日志"
	)
	@handler AdminListOperaLog //
	post /operalog/list (ListAdminOperaLogRequest) returns (CommonResponse)

	@doc(
		summary: "创建-操作日志"
	)
	@handler AdminCreateOperaLog //
	post /operalog/create (CreateAdminOperaLogRequest) returns (CommonResponse)

	@doc(
		summary: "更新-操作日志"
	)
	@handler AdminUpdateOperaLog //
	post /operalog/update (UpdateAdminOperaLogRequest) returns (CommonResponse)

	@doc(
		summary: "删除-操作日志"
	)
	@handler AdminDeleteOperaLog //
	post /operalog/delete (DeleteAdminOperaLogRequest) returns (CommonResponse)
}

@server(
	group : adminpost
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-职务"
	)
	@handler AdminGetPost //
	post /post/get (GetAdminPostRequest) returns (CommonResponse)

	@doc(
		summary: "列表-职务"
	)
	@handler AdminListPost //
	post /post/list (ListAdminPostRequest) returns (CommonResponse)

	@doc(
		summary: "创建-职务"
	)
	@handler AdminCreatePost //
	post /post/create (CreateAdminPostRequest) returns (CommonResponse)

	@doc(
		summary: "更新-职务"
	)
	@handler AdminUpdatePost //
	post /post/update (UpdateAdminPostRequest) returns (CommonResponse)

	@doc(
		summary: "删除-职务"
	)
	@handler AdminDeletePost //
	post /post/delete (DeleteAdminPostRequest) returns (CommonResponse)
}

@server(
	group : adminrole
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-角色"
	)
	@handler AdminGetRole //
	post /role/get (GetAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "列表-角色"
	)
	@handler AdminListRole //
	post /role/list (ListAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "创建-角色"
	)
	@handler AdminCreateRole //
	post /role/create (CreateAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "更新-角色"
	)
	@handler AdminUpdateRole //
	post /role/update (UpdateAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "删除-角色"
	)
	@handler AdminDeleteRole //
	post /role/delete (DeleteAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "授权-角色-菜单"
	)
	@handler AdminSetRoleMenu //
	post /role/setRoleMenu (SetAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "授权-角色-Api"
	)
	@handler AdminSetRoleApi //
	post /role/setRoleApi (SetAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "添加-角色-用户"
	)
	@handler AdminSetRoleUsers //
	post /role/setRoleUsers (UpdateRoleUsersRequest) returns (CommonResponse)
}

@server(
	group : adminuser
	prefix : /admin
	jwt: Auth
	middleware: RefreshToken // 路由中间件声明
)
service {{.Database}} {
	@doc(
		summary: "提取-用户自身信息"
	)
	@handler AdminGetMe //
	post /user/getMe (GetmeAdminRequest) returns (CommonResponse)

	@doc(
		summary: "提取-用户"
	)
	@handler AdminGetUser //
	post /user/get (GetAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "列表-用户"
	)
	@handler AdminListUser //
	post /user/list (ListAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "创建-用户"
	)
	@handler AdminCreateUser //
	post /user/create (CreateAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "更新-用户"
	)
	@handler AdminUpdateUser //
	post /user/update (UpdateAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "删除-用户"
	)
	@handler AdminDeleteUser //
	post /user/delete (DeleteAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "修改自己密码"
	)
	@handler AdminSetPassword //
	post /user/setPassword (SetPasswordRequest) returns (CommonResponse)

	@doc(
		summary: "管理员重置密码"
	)
	@handler AdminResetPassword //
	post /user/resetPassword (ResetPasswordRequest) returns (CommonResponse)

	@doc(
		summary: "设置自己角色"
	)
	@handler AdminSetMeRole
	post /user/setMeRole (SetMeAdminRoleRequest) returns (CommonResponse)

	@doc(
		summary: "列表 用户 无此角色"
	)
	@handler AdminListNoUser
	post /user/listNoUser (ListNoAdminUserRequest) returns (CommonResponse)

	@doc(
		summary: "导出-用户"
	)
	@handler AdminExportUser //
	post /user/export (ListAdminUserRequest) returns (ExportResponse)

	@doc(
		summary: "导出-用户模板"
	)
	@handler AdminExportTemplateUser //
	post /user/exportTemplate (NullRequest) returns (ExportResponse)

	@doc(
		summary: "导入-用户"
	)
	@handler AdminImportUser //
	post /user/import (ImportRequest) returns (CommonResponse)

	@doc(
		summary: "列表-用户树"
	)
	@handler AdminTreeUser //
	post /user/tree (ListAdminUserRequest) returns (CommonResponse)
}