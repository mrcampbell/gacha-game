defmodule StoreserviceWeb.Schema.ContentTypes do
  use Absinthe.Schema.Notation

  object :store do
    field :id, :id
    field :name, :string
  end
end
