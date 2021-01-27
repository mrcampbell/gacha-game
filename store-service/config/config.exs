# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :storeservice,
  ecto_repos: [Storeservice.Repo]

# Configures the endpoint
config :storeservice, StoreserviceWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "ZMkO5Cwucda/TyuuKbfo4WnrEM1tyZOVwopFM77zOxPh2tsgYHDjeTg1fJoXlHUE",
  render_errors: [view: StoreserviceWeb.ErrorView, accepts: ~w(json), layout: false],
  pubsub_server: Storeservice.PubSub,
  live_view: [signing_salt: "11XMG/3D"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
