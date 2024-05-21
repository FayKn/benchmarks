defmodule Routes.ApiRouter do
  use Routes.Base
  use HTTPoison.Base

  get "/steamUsers" do
    import Ecto.Query
    # query only the VanityName
    query = from u in App.SteamUser, select: [:VanityName]
    result = App.Repo.all(query)
    vanityNamesArr = Enum.map(result, fn user ->
      Map.get(user, :VanityName)
      |> String.split("/")
      |> Enum.reject(&(&1 == ""))
      |> List.last()
    end)

    vanityNamesArr
    |> Enum.map(fn vanity_name ->
      Task.async(fn ->
        case fetch_steam(vanity_name) do
          {:ok, response} ->
            steamid = response["response"]["steamid"] # Extract the steamid from the response
            case App.Repo.query("UPDATE SteamUsers SET SteamID = ? WHERE VanityName LIKE ?", [steamid, "%#{vanity_name}%"]) do
              {:ok, _result} -> :ok
              {:error, _reason} -> {:error, :internal_server_error}
            end
          {:error, reason} -> {:error, reason}
        end
      end)
    end)
    |> Enum.map(&Task.await/1)

    send(conn, 200, vanityNamesArr)
  end

  # 404
  match _ do
    send(conn, 404, "Not found")
  end


  def fetch_steam(vanity_name) do
    steam_api_key = System.get_env("STEAM_API_KEY") |> String.trim("\r")
    url = "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=#{steam_api_key}&vanityurl=#{vanity_name}"

    case Cachex.get(:steam_cache, url) do
      {:ok, nil} ->
        case get(url) do
          {:ok, %HTTPoison.Response{status_code: 200, body: body}} ->
            result = {:ok, Poison.decode!(body)}
            Cachex.put(:steam_cache, url, result)
            updateAPIsCalled(url)
            result
          {:ok, %HTTPoison.Response{status_code: status_code}} ->
            IO.inspect(status_code)
            {:error, "Received status code #{status_code}"}
          {:error, %HTTPoison.Error{reason: reason}} ->
            IO.inspect(reason)
            {:error, reason}
        end
      {:ok, result} -> result
      {:error, _} -> {:error, "Failed to fetch from cache"}
    end
  end

  def updateAPIsCalled(url) do
    import Ecto.Query

    query = from a in App.APICalled, where: a.api_url == ^url, select: a.count

    result = App.Repo.one(query)

    case result do
      nil ->
        # URL does not exist, create a new record
        %App.APICalled{api_url: url, count: 1}
        |> App.Repo.insert()
      _ ->
        # URL exists, increment the count
        query = from(a in App.APICalled, where: a.api_url == ^url, update: [inc: [count: 1]])
        App.Repo.update_all(query, [])
    end
  end


end