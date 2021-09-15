package model

const(
	PaginationFilterByCountry ="country"
)

type Pagination struct{
	FilterBy 	string
	Limit  		uint32
	Page   		uint32
}
