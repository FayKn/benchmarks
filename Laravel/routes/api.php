<?php

use App\Http\Controllers\SteamController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/steamUsers', [SteamController::class, 'index']);
