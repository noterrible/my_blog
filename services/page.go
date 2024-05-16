package services

import "my_blog/global"

type ListResp[T any] struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	List      []T `json:"list"`
}
type PageInfo struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
	SortKey  string `json:"sort_key"`
	Sort     string `json:"sort"`
	Debug    bool   `json:"debug"`
}
type Option struct {
	PageInfo
	Columns          []string `json:"columns"` //精确查询
	ColumnsValue     []string `json:"columns_value"`
	LikeColumns      []string `json:"like_columns"` //模糊查询
	LikeColumnsValue []string `json:"LikeColumnsValue"`
}

func GetList[T any](res []T, option Option) (ListResp[T], error) {
	res = make([]T, 0)
	//参数校验
	if option.Page <= 0 {
		option.Page = 1
	}
	if option.PageSize <= 0 {
		option.PageSize = 10
	}
	//拼接参数
	var whereString string
	for _, column := range option.Columns {
		whereString += "And " + column + " = ?"
	}
	var whereStringValue []interface{}
	for _, column := range option.ColumnsValue {
		whereStringValue = append(whereStringValue, column)
	}
	var likeWhereString string
	for _, column := range option.LikeColumns {
		likeWhereString += "And " + column + " like ?"
	}
	var likeWhereStringValue []interface{}
	for _, column := range option.LikeColumnsValue {
		likeWhereStringValue = append(likeWhereStringValue, "%"+column+"%")
	}
	var total int64
	query := global.DB.Model(res)
	if option.Debug == true {
		query = query.Debug()
	}
	filter := whereString + likeWhereString
	filterValue := append(whereStringValue, likeWhereStringValue)
	if len(filterValue) != 0 {
		query = query.Where("1=1 "+filter, filterValue)
	}
	offset := option.PageSize * (option.Page - 1)
	query = query.Count(&total)
	err := query.Limit(option.PageSize).Offset(offset).Order(option.SortKey + option.Sort).Find(&res).Error
	if err != nil {
		return ListResp[T]{}, err
	}
	//计算总页数
	totalPage := func() int {
		//计算总页数
		totalPage := len(res) / option.PageSize
		if len(res)%option.PageSize == 0 {
			return totalPage
		} else {
			return totalPage + 1
		}
	}()
	resp := ListResp[T]{
		Page:      option.Page,
		PageSize:  option.PageSize,
		Total:     int(total),
		TotalPage: totalPage,
		List:      res,
	}
	return resp, nil
}
