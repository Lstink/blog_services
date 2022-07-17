package dao

import (
	"github.com/lstink/blog/internal/model"
	"github.com/lstink/blog/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	// 实例化一个标签，并指明id
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	// 声明一个map
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	// 如果姓名不为空，则赋值到map上
	if name != "" {
		values["name"] = name
	}
	// 执行更新，传递参数，并将结果返回
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
