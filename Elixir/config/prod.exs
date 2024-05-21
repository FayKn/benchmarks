import Config

# Production configuration goes here.
config :app, port: 8000

# Example of database configuration
config :app, App.Repo,
       database: "SteamUsers",
       username: "root",
       password: "",
       hostname: "db",
       port: 3306,
       pool_size: 2000