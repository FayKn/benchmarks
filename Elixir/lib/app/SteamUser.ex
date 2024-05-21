defmodule App.SteamUser do
  use Ecto.Schema
  @derive {Jason.Encoder, only: [:VanityName, :SteamID]}

  @primary_key false
  schema "SteamUsers" do
    field :VanityName, :string, primary_key: true
    field :SteamID, :string
  end
end