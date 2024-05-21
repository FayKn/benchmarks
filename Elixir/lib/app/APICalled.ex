defmodule App.APICalled do
  use Ecto.Schema
  import Ecto.Changeset

  schema "APIsCalled" do
    field :api_url, :string
    field :count, :integer, default: 0
  end

  def changeset(api_called, attrs) do
    api_called
    |> cast(attrs, [:api_url, :count])
    |> validate_required([:api_url, :count])
  end
end