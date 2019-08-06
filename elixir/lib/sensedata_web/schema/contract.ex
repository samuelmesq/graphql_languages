defmodule SensedataWeb.Schema.Contract do
  use Absinthe.Schema.Notation

  @contracts [
    %{original_id: "1", amount: 11.11, status: :active, owner_id: "2" },
    %{original_id: "2", amount: 22.22, status: :active, owner_id: "2" },
    %{original_id: "3", amount: 33.33, status: :active, owner_id: "1" },
  ]

  enum :contract_status, values: [:active, :inactive]

  object :contract do
    field :original_id, :string
    field :amount, :float
    field :start, :string
    field :end, :string
    field :status, :contract_status
  end

  enum :contract_outer, values: [:profile, :all]

  input_object :contract_filter do
    field :outer, non_null(:contract_outer)
  end

  object :contract_queries do
    field :contracts, list_of(:contract) do
      arg :filter, :contract_filter
      resolve &list_contracts/2
    end

  end


  def list_contracts(%{ filter: %{ outer: :profile } }, _resolution) do
    contracts = Enum.filter(@contracts, fn c -> c.owner_id == "2" end)
    {:ok, contracts}
  end

  def list_contracts(_args, _resolution) do
    {:ok, @contracts}
  end
end
