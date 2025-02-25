package entity

type ClientRepositoryInterface interface {
	Save(entity *Client) error
	FindByID(id int) (*Client, error)
	FindAll(page, limit int, sort string) ([]*Client, error)
	Update(entity *Client) error
	Delete(id int) error
}
