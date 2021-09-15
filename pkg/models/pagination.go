package model

const (
	PaginationFilterByCountry = "country"
)

/*Pagination model for pagination
FilterBy could be by  countries or state (Nok , ok)
*/
type Pagination struct {
	FilterBy string
	Limit    uint32
	Page     uint32
}
