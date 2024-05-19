<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class SteamUsers extends Model
{
    protected $table = 'SteamUsers';
    public $timestamps = false;

    protected $fillable = [
        'VanityName',
        'SteamID'
    ];
    protected $primaryKey = 'VanityName';
    public $incrementing = false;
}
