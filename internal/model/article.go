package model

import (
	"github.com/jinzhu/gorm"
	"goblog/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

// 专门用于 swagger 显示返回值
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a *Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(a).Updates(values).Where("id = ? and is_del = ?", a.ID, 0).Error
}

func (a *Article) Get(db *gorm.DB) (*Article, error) {
	var article *Article
	db = db.Where("id = ? and state = ? and is_del = ?", a.ID, a.State, 0)
	// first 方法用于查询单条记录
	err := db.First(article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

// list articles with pagination
func (a *Article) ListArticles(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var result []*Article
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// TODO: have a bug
	err := db.Where("state = ?", a.State).Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (a *Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", a.ID, 0).Delete(a).Error
}

func (a *Article) CountArticles(db *gorm.DB, state uint8) (int, error) {
	var count int
	err := db.Table(a.TableName()).Where("state = ?", state).Count(&count).Error
	return count, err
}

// 关联查询
type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	Title         string
	Desc          string
	CoverImageUrl string
	Content       string
}

// list articles by tag id
func (a *Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{
		"ar.id as article_id",
		"ar.title as article_title",
		"ar.desc as article_desc",
		"ar.cover_image_url",
		"ar.content",
	}
	fields = append(fields, []string{
		"t.id as tag_id",
		"t.name as tag_name",
	}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	articleTag := new(ArticleTag)
	tag := new(Tag)
	article := new(Article)
	// select：指定要检索的字段，若不指定，则为 select *
	// joins：指定关联查询的语句
	// rows：执行sql语句并获取查询结果
	rows, err := db.Select(fields).Table(articleTag.TableName()+" as at").
		Joins("left join `"+tag.TableName()+"` as t on at.tag_id = t.id").
		Joins("left join `"+article.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.`tag_id` = ? and ar.state = ? and ar.is_del = ?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(
			&r.ArticleID,
			&r.Title,
			&r.Desc,
			&r.CoverImageUrl,
			&r.Content,
			&r.TagID,
			&r.TagName,
		); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (a *Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	articleTag := new(ArticleTag)
	tag := new(Tag)
	err := db.Table(articleTag.TableName()+" as at").
		Joins("left join `"+tag.TableName()+"` as t on at.tag_id= t.id").
		Joins("left join `"+a.TableName()+"` as ar on at.article_id = ar.id").
		Where("at.`tag_id` = ? and ar.state = ? and ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
