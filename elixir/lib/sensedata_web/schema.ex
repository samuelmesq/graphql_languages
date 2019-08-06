defmodule SensedataWeb.Schema do
  use Absinthe.Schema

  object :image do
    field :href, :string
    field :title, :string
  end

  import_types SensedataWeb.Schema.Viewer
  import_types SensedataWeb.Schema.Contract

  query do
    import_fields :viewer_queries
    import_fields :contract_queries
  end
end
