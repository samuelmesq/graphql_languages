const fs = require('fs');
const Koa = require('koa');
const { ApolloServer, gql } = require('apollo-server-koa');

const resolvers = require('./resolvers');


const typeDefs = gql(fs.readFileSync(__dirname.concat('/schema.graphql'), 'utf8'))

const server = new ApolloServer({ typeDefs, resolvers });

const app = new Koa();
server.applyMiddleware({ app });

app.listen({ port: 4001 }, () => {
  console.log(`🚀  Server ready at http://localhost:4001${server.graphqlPath}`);
});
