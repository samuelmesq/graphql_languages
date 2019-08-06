package sensedata

import (
	"context"
	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

type viewer struct {
	ID graphql.ID
	name string
	avatar Image
	notifications []Notification
}

type Image struct {
	href  string
	title string
}

type Notification struct {
	title string
	kind  string
}

type Contract struct {
	amount     float64
	originalId string
	ownerId    graphql.ID
}

type PaginatedContracts struct {
	totalCount int32
	nodes []*Contract
}

type ContractFilter struct {
	Filter *struct{
		Outer *string
	}
}


func (i *Image) Href() string {
	return i.href
}

func (i *Image) Title() string {
	return i.title
}

func (n *Notification) Title() string {
	return n.title
}

func (n *Notification) Kind() string {
	return n.kind

}

func (r *Resolver) Viewer() *viewer {
	return &viewer{
		ID: "2",
		name: "LALA",
	}
}

var contracts = []*Contract{
	{
		amount: 155.4,
		originalId: "144",
		ownerId: "1",
	},
	{
		amount: 2000.0,
		originalId: "200",
		ownerId: "2",
	},
}


func (r *Resolver) Contracts(ctx context.Context, args ContractFilter) *PaginatedContracts {
	if (*args.Filter.Outer == "ALL") {
		return &PaginatedContracts{
			totalCount: int32(len(contracts)),
			nodes: contracts,
		}
	}

	filtered := make([]*Contract, 0)


	for _, c := range contracts {
		if (c.ownerId == "1") {
			filtered = append(filtered, c)
		}
	}

	return &PaginatedContracts{
		totalCount: int32(len(filtered)),
		nodes: filtered,
	}
}

func (v *viewer) Name() string {
	return v.name
}

func (v *viewer) Avatar() *Image {
	return &Image{title: "Something", href: "https://via.placeholder.com/150"}
}

var notifications = []*Notification{
	{
		title: "Something New",
		kind: "INFO",
	},
	{
		title: "Something Wrong",
		kind: "ERROR",
	},
}

func (v *viewer) Notifications() []*Notification {
	if (v.ID == "2") {
		return make([]*Notification, 0)
	}

	return notifications
}

func (p *PaginatedContracts) TotalCount() int32 {
	return p.totalCount
}

func (p *PaginatedContracts) Nodes() []*Contract {
	return p.nodes
}

func (c *Contract) Amount() float64 {
	return c.amount
}

func (c *Contract) OriginalId() string {
	return c.originalId
}
