package gormx

import (
	"fmt"
	"gorm.io/gorm"
)

type DataPermission struct {
	DataScope  string
	UserId     interface{}
	DeptId     interface{}
	RoleId     interface{}
	DataFilter string //自定义条件create_by in 用户id范围
}

// PermissionData 数据权限
func PermissionData(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if !config.ApplicationConfig.EnableDP {
		// 	return db
		// }
		switch p.DataScope {
		case "2": //自定数据权限
			where := fmt.Sprintf("(%s)", p.DataFilter)
			return db.Where(tableName + ".create_by in " + where)
		case "3": //本部门数据权限
			return db.Where(tableName+".create_by in "+
				"(SELECT user_id from admin_user where dept_id = ? )", p.DeptId)
		case "4": //本部门及以下数据权限
			return db.Where(tableName+".create_by in "+
				"(SELECT user_id from admin_user "+
				"where admin_user.dept_id in "+
				"(with RECURSIVE temp_child as "+
				"(select * from admin_dept where id = ?  union all "+
				"select c.* from admin_dept as c,temp_child t where c.parent_id=t.id) "+
				"select id from temp_child)", p.DeptId)
		case "5": //仅本人数据权限
			return db.Where(tableName+".create_by = ?", p.UserId)
		default:
			return db
		}
	}
}
