import Config

# Development configurations goes here
config :app, port: 8000

# Example of database configuration
config :app, App.Repo,
       database: "SteamUsers",
       username: "root",
       password: "",
       hostname: "localhost",
       port: 3306,
       pool_size: 200

config :plug_cowboy,
  log_exceptions_with_status_code: [400..599]
