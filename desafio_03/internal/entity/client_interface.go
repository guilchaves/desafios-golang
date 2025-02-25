package entity

type ClientRepositoryInterface interface {
	Save(entity *Client) error
	FindByID(id int) (*Client, error)
	Update(entity *Client) error
}
