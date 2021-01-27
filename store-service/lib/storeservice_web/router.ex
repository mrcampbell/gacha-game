defmodule StoreserviceWeb.Router do
  use StoreserviceWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api" do
    pipe_through :api

    forward "/graphiql", Absinthe.Plug.GraphiQL,
      schema: StoreserviceWeb.Schema

    forward "/", Absinthe.Plug,
      schema: StoreserviceWeb.Schema

  end

end
