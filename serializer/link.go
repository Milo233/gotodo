package serializer

import "gotodo/model"

// BuildUserResponse 序列化用户响应
func BuildLinkResponse(link model.Link) LinkResponse {
	return LinkResponse {
		Data: BuildLink(link),
	}
}

// User 用户序列化器
type Link struct {
	Address	string   `json:"address"`
	Name  	string `json:"name"`
}

// UserResponse 单个用户序列化
type LinkResponse struct {
	Response
	Data Link `json:"data"`
}

// BuildUser 序列化用户
func BuildLink(link model.Link) Link {
	return Link{
		Address:        link.Address,
		Name:  link.Name,
	}
}
