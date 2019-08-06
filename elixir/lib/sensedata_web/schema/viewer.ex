defmodule SensedataWeb.Schema.Viewer do
  use Absinthe.Schema.Notation
  alias SensedataWeb.Schema.Notification
  import_types SensedataWeb.Schema.Notification

  object :viewer do
    field :id, :id
    field :name, :string
    field :avatar, :image
    field :notifications, non_null(list_of(non_null(:notification))) do
      resolve &Notification.list_user_notifications/3
    end
  end

  @viewer %{
    id: "2",
    name: "Samuel",
    avatar: %{
      href: "https://avatar.com/1.jpg",
      title: "Samuel Avatar",
    }
  }

  object :viewer_queries do
    field :viewer, :viewer do
      resolve fn _, _ ->
        {:ok, @viewer}
      end
    end
  end
end
