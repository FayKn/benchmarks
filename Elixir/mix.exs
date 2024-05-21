defmodule App.MixProject do
  use Mix.Project

  def project do
    [
      app: :app,
      version: "1.0.0",
      elixir: "~> 1.10",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {App.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:plug_cowboy, "~> 2.0"},
      {:jason, "~> 1.2"},
      {:ecto_sql, "~> 3.5"},
      {:myxql, ">= 0.0.0"},
      {:httpoison, "~> 2.2"},
      {:poison, "~> 5.0"},
      {:dotenv, "~> 3.1"},
      {:cachex, "~> 3.1"}
    ]
  end
end
