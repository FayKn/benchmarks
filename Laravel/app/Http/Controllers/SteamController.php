<?php

namespace App\Http\Controllers;

use App\Models\APIsCalled;
use App\Models\SteamUsers;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class SteamController extends Controller
{
    public function index() {
        $returnArray = $this->getVanityNamesArr();

        foreach ($returnArray as $vanityName) {
            $response = $this->fetchSteam($vanityName);

            $user = SteamUsers::where('VanityName', 'like', '%'. $vanityName .'%')->first();

            $user->SteamID = $response ?? 'Not Found';
            $user->save();
        }

        return response()->json($returnArray);
    }

    private function getVanityNamesArr()
    {
        $users = SteamUsers::get();

        $usernamesArray = [];

        foreach ($users as $user) {
            $username = $this->getUsernameFromSteamUrl($user->VanityName);
            array_push($usernamesArray, $username);
        }

        return $usernamesArray;
    }

    private function fetchSteam($vanityName) {
        $url = 'http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=' . config('app.steamApiKey'). '&vanityurl='. $vanityName;
        // check if the response is cached
        if (cache()->has($url)) {
            return cache($url);
        }

        $response = Http::get($url);

        if(APIsCalled::where('Api_url', $url)->exists()){
            APIsCalled::where('Api_url', $url)->increment('count');
        }
        else {
            $apiCalled = new APIsCalled();
            $apiCalled->Api_url = $url;
            $apiCalled->count = 1;
            $apiCalled->save();
        }

        // cache the response in redis
        $response = $response->json();
        $response = $response['response']['steamid'] ?? null;

        // cache the response in redis
        cache([$url => $response], now()->addMinutes(10));

        return $response;
    }

    private function getUsernameFromSteamUrl($url) {
        $path = parse_url($url, PHP_URL_PATH);
        $parts = explode('/', $path);

        if(end($parts) == '') {
            array_pop($parts);
        }

        return end($parts);
    }
}
