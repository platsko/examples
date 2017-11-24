<?php

namespace App\Models;

class Wallet extends AppModel
{
    public function coins()
    {
        return $this->hasMany('App\Models\WalletCoins');
    }
}
