from starlette.applications import Starlette
from starlette.graphql import GraphQLApp
from schema import schema
import uvicorn

app = Starlette()
app.add_route('/', GraphQLApp(schema=schema))


if __name__ == '__main__':
    uvicorn.run(app, host='0.0.0.0', port=4004)
