defmodule Storeservice.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    children = [
      # Start the Ecto repository
      Storeservice.Repo,
      # Start the Telemetry supervisor
      StoreserviceWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: Storeservice.PubSub},
      # Start the Endpoint (http/https)
      StoreserviceWeb.Endpoint
      # Start a worker by calling: Storeservice.Worker.start_link(arg)
      # {Storeservice.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Storeservice.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  def config_change(changed, _new, removed) do
    StoreserviceWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
