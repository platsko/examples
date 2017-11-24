<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

abstract class AppModel extends Model
{
    /**
     * Indicates if the model should be timestamped.
     *
     * @var bool
     */
    public $timestamps = false;
}
