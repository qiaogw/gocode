package adminmodel

import (
	"github.com/qiaogw/gocode/common/modelx"
	"time"
)

type ListUserReq struct {
	modelx.Pagination `search:"-"`
	Id                string    `json:"id" form:"id" search:"type:exact;column:id;table:admin_user"`
	PostId            string    `json:"postId" form:"postId" search:"type:exact;column:post_id;table:admin_user"`
	DeptId            string    `json:"deptId" form:"deptId" search:"type:exact;column:dept_id;table:admin_user"`
	Uuid              string    `json:"uuid" form:"uuid" search:"type:exact;column:uuid;table:admin_user"`
	Username          string    `json:"username" form:"username" search:"type:exact;column:username;table:admin_user"`
	Password          string    `json:"password" form:"password" search:"type:exact;column:password;table:admin_user"`
	NickName          string    `json:"nickName" form:"nickName" search:"type:exact;column:nick_name;table:admin_user"`
	RealName          string    `json:"realName" form:"realName" search:"type:exact;column:real_name;table:admin_user"`
	Mobile            string    `json:"mobile" form:"mobile" search:"type:exact;column:mobile;table:admin_user"`
	Avatar            string    `json:"avatar" form:"avatar" search:"type:exact;column:avatar;table:admin_user"`
	Gender            string    `json:"gender" form:"gender" search:"type:exact;column:gender;table:admin_user"`
	Email             string    `json:"email" form:"email" search:"type:exact;column:email;table:admin_user"`
	SideMode          string    `json:"sideMode" form:"sideMode" search:"type:exact;column:side_mode;table:admin_user"`
	BaseColor         string    `json:"baseColor" form:"baseColor" search:"type:exact;column:base_color;table:admin_user"`
	RoleId            string    `json:"roleId" form:"roleId" search:"type:exact;column:role_id;table:admin_user"`
	ActiveColor       string    `json:"activeColor" form:"activeColor" search:"type:exact;column:active_color;table:admin_user"`
	Sort              int64     `json:"sort" form:"sort" search:"type:exact;column:sort;table:admin_user"`
	Remark            string    `json:"remark" form:"remark" search:"type:exact;column:remark;table:admin_user"`
	Status            string    `json:"status" form:"status" search:"type:exact;column:status;table:admin_user"`
	CreateBy          string    `json:"createBy" form:"createBy" search:"type:exact;column:create_by;table:admin_user"`
	UpdateBy          string    `json:"updateBy" form:"updateBy" search:"type:exact;column:update_by;table:admin_user"`
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" search:"type:exact;column:created_at;table:admin_user"`
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" search:"type:exact;column:updated_at;table:admin_user"`
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" search:"type:exact;column:deleted_at;table:admin_user"`
	UserOrder
}

type UserOrder struct {
	IdOrder          string    `json:"id" form:"id" search:"type:order;column:id;table:admin_user"`
	PostIdOrder      string    `json:"postId" form:"postId" search:"type:order;column:post_id;table:admin_user"`
	DeptIdOrder      string    `json:"deptId" form:"deptId" search:"type:order;column:dept_id;table:admin_user"`
	UuidOrder        string    `json:"uuid" form:"uuid" search:"type:order;column:uuid;table:admin_user"`
	UsernameOrder    string    `json:"username" form:"username" search:"type:order;column:username;table:admin_user"`
	PasswordOrder    string    `json:"password" form:"password" search:"type:order;column:password;table:admin_user"`
	NickNameOrder    string    `json:"nickName" form:"nickName" search:"type:order;column:nick_name;table:admin_user"`
	RealNameOrder    string    `json:"realName" form:"realName" search:"type:order;column:real_name;table:admin_user"`
	MobileOrder      string    `json:"mobile" form:"mobile" search:"type:order;column:mobile;table:admin_user"`
	AvatarOrder      string    `json:"avatar" form:"avatar" search:"type:order;column:avatar;table:admin_user"`
	GenderOrder      string    `json:"gender" form:"gender" search:"type:order;column:gender;table:admin_user"`
	EmailOrder       string    `json:"email" form:"email" search:"type:order;column:email;table:admin_user"`
	SideModeOrder    string    `json:"sideMode" form:"sideMode" search:"type:order;column:side_mode;table:admin_user"`
	BaseColorOrder   string    `json:"baseColor" form:"baseColor" search:"type:order;column:base_color;table:admin_user"`
	RoleIdOrder      string    `json:"roleId" form:"roleId" search:"type:order;column:role_id;table:admin_user"`
	ActiveColorOrder string    `json:"activeColor" form:"activeColor" search:"type:order;column:active_color;table:admin_user"`
	SortOrder        int64     `json:"sort" form:"sort" search:"type:order;column:sort;table:admin_user"`
	RemarkOrder      string    `json:"remark" form:"remark" search:"type:order;column:remark;table:admin_user"`
	StatusOrder      string    `json:"status" form:"status" search:"type:order;column:status;table:admin_user"`
	CreateByOrder    string    `json:"createBy" form:"createBy" search:"type:order;column:create_by;table:admin_user"`
	UpdateByOrder    string    `json:"updateBy" form:"updateBy" search:"type:order;column:update_by;table:admin_user"`
	CreatedAtOrder   time.Time `json:"createdAt" form:"createdAt" search:"type:order;column:created_at;table:admin_user"`
	UpdatedAtOrder   time.Time `json:"updatedAt" form:"updatedAt" search:"type:order;column:updated_at;table:admin_user"`
	DeletedAtOrder   time.Time `json:"deletedAt" form:"deletedAt" search:"type:order;column:deleted_at;table:admin_user"`
}

func (m *ListUserReq) GetNeedSearch() interface{} {
	return *m
}
