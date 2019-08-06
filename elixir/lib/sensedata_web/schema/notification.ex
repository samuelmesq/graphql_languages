defmodule SensedataWeb.Schema.Notification do
  use Absinthe.Schema.Notation

  @notifications [
    %{owner_id: "2", kind: :error, title: "i am a error"},
    %{owner_id: "2", kind: :info, title: "i am a info"},
    %{owner_id: "1", kind: :warning, title: "i am a warning"},
  ]

  enum :notification_kind, values: [:info, :error, :warning]

  object :notification do
    field :title, non_null(:string)
    field :kind, non_null(:notification_kind)
  end

  def list_user_notifications(parent, _args, _resolution) do
    notifications = Enum.filter(@notifications, fn n -> n.owner_id == parent.id end)
    {:ok, notifications}
  end
end
