// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleMenuDao is the data access object for table hg_admin_role_menu.
type AdminRoleMenuDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdminRoleMenuColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleMenuColumns defines and stores column names for table hg_admin_role_menu.
type AdminRoleMenuColumns struct {
	RoleId string // 角色ID
	MenuId string // 菜单ID
}

// adminRoleMenuColumns holds the columns for table hg_admin_role_menu.
var adminRoleMenuColumns = AdminRoleMenuColumns{
	RoleId: "role_id",
	MenuId: "menu_id",
}

// NewAdminRoleMenuDao creates and returns a new DAO object for table data access.
func NewAdminRoleMenuDao() *AdminRoleMenuDao {
	return &AdminRoleMenuDao{
		group:   "default",
		table:   "hg_admin_role_menu",
		columns: adminRoleMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleMenuDao) Columns() AdminRoleMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
