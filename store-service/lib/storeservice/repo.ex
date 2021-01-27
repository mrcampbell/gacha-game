defmodule Storeservice.Repo do
  use Ecto.Repo,
    otp_app: :storeservice,
    adapter: Ecto.Adapters.Postgres
end
