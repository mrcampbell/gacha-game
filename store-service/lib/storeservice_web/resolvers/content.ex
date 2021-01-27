defmodule StoreserviceWeb.Resolvers.Content do
  alias Storeservice.{Repo, Store}
  import Ecto.Query

  def list_stores(_parent, _args, _resolution) do
    {:ok, Repo.all(Store)}
  end

  def find_store(_parent, %{id: id}, _resolution) do
    case Repo.one(from s in Store, where: s.id == ^id) do
      nil ->
        {:error, "Store with ID [#{id}] not found"}
      s ->
        {:ok, s}
    end
  end
end
