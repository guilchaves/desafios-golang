package entity

type ClientRepositoryInterface interface {
	Save(entity *Client) error
	FindByID(id uint) (*Client, error)
	Update(entity *Client) error
}
