package dao

import (
	"ayzy_platform/common/database"

	"log"
)

// Create 创建实体
func Create(value interface{}) error {
	return database.DB.Create(value).Error
}

// Save 保存实体
func Save(value interface{}) error {
	return database.DB.Save(value).Error
}

// // Updates 更新实体
// func (b *BaseRepository) Updates(model interface{}, value interface{}) error {
// 	return database.DB.Model(model).Updates(value).Error
// }

// DeleteByWhere 根据条件删除实体
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := database.DB.Where(where).Delete(model)
	err = db.Error
	if err != nil {
		log.Println("删除实体出错", err)
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID 根据id删除实体
func DeleteByID(model interface{}, id int) error {
	return database.DB.Where("id=?", id).Delete(model).Error
}

// DeleteByIDS 根据多个id删除多个实体
func DeleteByIDS(model interface{}, ids []int) (count int64, err error) {
	db := database.DB.Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		log.Println("删除多个实体出错", err)
		return
	}
	count = db.RowsAffected
	return
}

// First 根据条件获取一个实体
func First(where interface{}, out interface{}, selects ...string) error {
	db := database.DB.Where(where)
	if len(selects) > 0 {
		for _, sel := range selects {
			db = db.Select(sel)
		}
	}
	return db.First(out).Error
}

// FirstByID 根据条件获取一个实体
func FirstByID(out interface{}, id int64) error {
	return database.DB.First(out, id).Error
}

// Find 根据条件返回数据
func Find(where interface{}, out interface{}, sel string, orders ...string) error {
	db := database.DB.Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// GetPages 分页返回数据
func GetPages(model interface{}, out interface{}, pageIndex, pageSize int, totalCount *int64, where interface{}, orders ...string) error {
	db := database.DB.Model(model).Where(model)
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		log.Println("查询总数出错", err)
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList 查询 model 中的一个列作为切片
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return database.DB.Model(model).Where(where).Pluck(fieldName, out).Error
}

//GetTransaction 获取事务
/*
func GetTransaction() *gorm.DB {
	return database.DB.Begin()
}
*/
