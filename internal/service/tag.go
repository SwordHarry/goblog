package service

import (
	"github.com/jinzhu/gorm"
	"goblog/internal/dao"
	"goblog/internal/request"
	"goblog/pkg/app"
)

// service 层与接口校验层合并
func (svc *Service) CountTag(param *request.CountTagRequest) (int, error) {
	return svc.model.CountTag(param.Name, param.State)
}
func (svc *Service) GetTagList(param *request.TagListRequest, pager *app.Pager) ([]*dao.Tag, error) {
	return svc.model.ListTags(param.Name, param.State, app.GetPageOffset(pager.Page, pager.PageSize), pager.PageSize)
}

func (svc *Service) GetTagByName(param *request.CreateTagRequest) (*dao.Tag, error) {
	t, err := svc.model.GetTagByName(param.Name, param.State)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return t, nil
}

func (svc *Service) CreateTag(param *request.CreateTagRequest) error {
	return svc.model.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *request.UpdateTagRequest) error {
	return svc.model.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *request.DeleteTagRequest) error {
	return svc.model.DeleteTag(param.ID)
}
