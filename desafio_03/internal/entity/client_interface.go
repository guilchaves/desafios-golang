package entity

type ClientRepositoryInterface interface {
	Save(client *Client) error
}
