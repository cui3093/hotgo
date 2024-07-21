// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysDictDataDao is internal type for wrapping internal DAO implements.
type internalSysDictDataDao = *internal.SysDictDataDao

// sysDictDataDao is the data access object for table hg_sys_dict_data.
// You can define custom methods on it to extend its functionality as you wish.
type sysDictDataDao struct {
	internalSysDictDataDao
}

var (
	// SysDictData is globally public accessible object for table hg_sys_dict_data operations.
	SysDictData = sysDictDataDao{
		internal.NewSysDictDataDao(),
	}
)

// Fill with you ideas below.
