package casbinx

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type CasbinService struct {
	db *gorm.DB
}

var CasbinServiceApp = newCasbinService
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func newCasbinService(db *gorm.DB) *CasbinService {
	return &CasbinService{
		db: db,
	}
}

// Casbin
// @author: qgw
// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func (c *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(c.db)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// UpdateCasbin
// @author: qgw
// @function: UpdateCasbin
// @description: 更新casbin权限
// @param: authorityId string, casbinInfos []request.CasbinInfo
// @return: error
func (c *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	c.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := c.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasbinApi
// @author: qgw
// @function: UpdateCasbinApi
// @description: API更新随动
// @param: oldPath string, newPath string, oldMethod string, newMethod string
// @return: error
func (c *CasbinService) UpdateCasbinApi(oldPath string, newPath string,
	oldMethod string, newMethod string) error {
	err := c.db.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetCasbinByAuthorityId
// @author: qgw
// @function: GetPolicyPathByAuthorityId
// @description: 获取权限列表
// @param: authorityId string
// @return: pathMaps []request.CasbinInfo
func (c *CasbinService) GetCasbinByAuthorityId(AuthorityID uint) (pathMaps []CasbinInfo, err error) {
	e := c.Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list, err := e.GetFilteredPolicy(0, authorityId)
	if err != nil {
		return pathMaps, err
	}
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps, nil
}

// ClearCasbin
// @author: qgw
// @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func (c *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := c.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
