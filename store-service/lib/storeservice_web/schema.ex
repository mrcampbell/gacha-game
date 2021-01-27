defmodule StoreserviceWeb.Schema do
  use Absinthe.Schema
  import_types StoreserviceWeb.Schema.ContentTypes

  alias StoreserviceWeb.Resolvers

  query do
    @desc "Get all stores"
    field :stores, list_of(:store) do
      resolve &Resolvers.Content.list_stores/3
    end

    @desc "Get single store"
    field :store, :store do
      arg :id, non_null(:id)
      resolve &Resolvers.Content.find_store/3
    end
  end

end
