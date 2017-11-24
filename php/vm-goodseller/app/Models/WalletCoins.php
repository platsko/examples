<?php

namespace App\Models;

class WalletCoins extends AppModel
{
    public function coin()
    {
        return $this->hasOne('App\Models\Coin', 'id', 'coin_id');
    }
}
