package source

import (
	"github.com/gookit/color"
	"github.com/madneal/gshark/global"
	"github.com/madneal/gshark/model"
	"time"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "success"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "setting", Name: "setting", Hidden: false, Component: "view/routerHolder.vue", Sort: 1, Meta: model.Meta{Title: "管理", Icon: "setting"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "rule", Name: "rule", Hidden: false, Component: "view/rule/rule.vue", Sort: 1, Meta: model.Meta{Title: "规则管理", Icon: "s-order"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "result", Name: "result", Hidden: false, Component: "view/routerHolder.vue", Sort: 0, Meta: model.Meta{Title: "搜索结果", Icon: "search"}},
	{GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "26", Path: "code", Name: "code", Hidden: false, Component: "view/searchResult/searchResult.vue", Sort: 1, Meta: model.Meta{Title: "代码", Icon: "document"}},
	{GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "token", Name: "token", Hidden: false, Component: "view/token/token.vue", Sort: 2, Meta: model.Meta{Title: "token管理", Icon: "collection"}},
	{GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "26", Path: "subdomain", Name: "subdomain", Hidden: false, Component: "view/subdomain/subdomain.vue", Sort: 2, Meta: model.Meta{Title: "子域名资产报告", Icon: "s-data"}},
	{GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "filter", Name: "filter", Hidden: false, Component: "view/filter/filter.vue", Sort: 3, Meta: model.Meta{Title: "过滤规则", Icon: "s-opportunity"}},
}

func (m *menu) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
