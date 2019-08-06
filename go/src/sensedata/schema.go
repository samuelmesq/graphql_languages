package sensedata

const Schema = `
schema {
  query: Query
}

scalar Date

type Query {
  viewer: Viewer
  contracts(filter: ContractFilter): PaginatedContracts
}

type Image {
  href: String!
  title: String!
}

type Viewer {
  name: String!
  avatar: Image!
  notifications: [Notification!]!
}

enum NotificationKind { WARNING, ERROR, INFO }

type Notification {
  title: String!
  kind: NotificationKind!
}

enum ListFilterOuter { ALL, PROFILE }

input ContractFilter {
  outer: ListFilterOuter
}

type PaginatedContracts {
  totalCount: Int!
  nodes: [Contract!]!
}

type Contract {
  amount: Float!
  originalId: String!
}
`
