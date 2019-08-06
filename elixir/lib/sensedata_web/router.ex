defmodule SensedataWeb.Router do
  #   use SensedataWeb, :router
  use Phoenix.Router

  scope "/graphql" do
    forward "/", Absinthe.Plug,
      schema: SensedataWeb.Schema
  end
end
