package entity

type ClientRepositoryInterface interface {
	Create(entity *Client) error
	FindAll(page, limit int, sort string) ([]Client, error)
	FindByID(id int) (*Client, error)
	Update(entity *Client) error
	Delete(id uint) error
}
