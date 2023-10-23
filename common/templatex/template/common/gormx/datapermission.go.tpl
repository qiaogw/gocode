package gormx

import (
	"fmt"
	"gorm.io/gorm"
)

type DataPermission struct {
	DataScope string
	UserId    int
	DeptId    int
	RoleId    int
}

//Permission 数据权限
func Permission(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if !config.ApplicationConfig.EnableDP {
		// 	return db
		// }
		switch p.DataScope {
		case "2":
			return db.Where(tableName+".create_by in "+
				"(select sys_user.user_id from sys_role_dept "+
				"left join sys_user on sys_user.dept_id=sys_role_dept.dept_id "+
				"where sys_role_dept.role_id = ?)", p.RoleId)
		case "3":
			return db.Where(tableName+".create_by in "+
				"(SELECT user_id from sys_user where dept_id = ? )", p.DeptId)
		case "4":
			return db.Where(tableName+".create_by in "+
				"(SELECT user_id from sys_user "+
				"where sys_user.dept_id in "+
				"(select dept_id from sys_dept where dept_path like ? ))",
				"%/"+fmt.Sprint(p.DeptId)+"/%")
		case "5":
			return db.Where(tableName+".create_by = ?", p.UserId)
		default:
			return db
		}
	}
}
