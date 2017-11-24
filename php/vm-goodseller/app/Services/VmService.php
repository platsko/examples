<?php

namespace App\Services;

use App\Models\Goods;
use App\Models\Wallet;
use App\Models\WalletCoins;
use Closure;
use DB;
use Exception;
use Illuminate\Support\Collection;

class VmService
{
    /**
     * @param \App\Models\WalletCoins $coin
     * @param \App\Models\Wallet      $userWallet
     * @param \App\Models\Wallet      $depositWallet
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function putCoin(WalletCoins $coin, Wallet $userWallet, Wallet $depositWallet)
    {
        if (! $coin->quantity > 0) {
            return response()->json([
                                        'status'  => 'error',
                                        'message' => 'Недостаточно средств',
                                    ]);
        }

        $coins = collect()->push([
                                     'id'       => $coin->coin->id,
                                     'cost'     => $coin->coin->cost,
                                     'quantity' => 1,
                                 ]);

        $this->makeTransaction($coins, $userWallet, $depositWallet);

        return redirect()->route('refresh');
    }

    /**
     * @param \App\Models\Goods  $goods
     * @param \App\Models\Wallet $vmWallet
     * @param \App\Models\Wallet $depositWallet
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function purchase(Goods $goods, Wallet $depositWallet, Wallet $vmWallet)
    {
        if (! $goods->quantity > 0) {
            return response()->json([
                                        'status'  => 'error',
                                        'message' => 'Товар закончился',
                                    ]);
        }

        $amount = $depositWallet->coins->sum(function ($item) {
            return $item->quantity * $item->coin->cost;
        });

        if ($goods->cost > $amount) {
            return response()->json([
                                        'status'  => 'error',
                                        'message' => 'Недостаточно средств',
                                    ]);
        }

        $coins = $this->prepareCoinsByAmount($depositWallet->coins, $amount);
        $this->makeTransaction($coins, $depositWallet, $vmWallet, function () use ($goods) {
            $goods->decrement('quantity');
        });

        $odd = $amount - $goods->cost;
        if ($odd > 0) {
            $coins = $this->prepareCoinsByAmount($vmWallet->coins, $odd);
            $this->makeTransaction($coins, $vmWallet, $depositWallet);
        }

        return redirect()->route('refresh');
    }

    /**
     * @param \App\Models\Wallet $payer
     * @param \App\Models\Wallet $payee
     * @param \App\Models\Wallet $through
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function refund(Wallet $payer, Wallet $payee, Wallet $through)
    {
        $amount = $payer->coins->sum(function ($item) {
            return $item->quantity * $item->coin->cost;
        });

        if (! $amount > 0) {
            return response()->json([
                                        'status'  => 'error',
                                        'message' => 'Средства не вносились',
                                    ]);
        }

        $coins = $this->prepareCoinsByAmount($payer->coins, $amount);
        $this->makeTransaction($coins, $payer, $through);

        $coins = $this->prepareCoinsByAmount($through->coins, $amount);
        $this->makeTransaction($coins, $through, $payee);

        return redirect()->route('refresh');
    }

    /**
     * @param \Illuminate\Support\Collection $walletCoins
     * @param int                            $amount
     *
     * @return \Illuminate\Support\Collection
     */
    protected function prepareCoinsByAmount(Collection $walletCoins, int $amount)
    {
        $coins = new Collection;

        $walletCoins->sortByDesc('coin.cost')->each(function ($item) use ($coins, &$amount) {
            while ($item->quantity > 0 && $amount > 0 && $item->coin->cost <= $amount) {
                $item->quantity--;
                $amount -= $item->coin->cost;
                $coins->put($item->coin->id, [
                    'id'       => $item->coin->id,
                    'cost'     => $item->coin->cost,
                    'quantity' => $coins->has($item->coin->id) ? $coins->get($item->coin->id)['quantity'] + 1 : 1,
                ]);
            }
        });

        return $coins;
    }

    protected function makeTransfer(Wallet $payer, Wallet $payee, Wallet $through = null)
    {

    }

    /**
     * @param \Illuminate\Support\Collection $coins
     * @param \App\Models\Wallet             $payer
     * @param \App\Models\Wallet             $payee
     * @param \Closure|null                  $extra
     *
     * @throws \Exception
     * @throws \Throwable
     */
    protected function makeTransaction(Collection $coins, Wallet $payer, Wallet $payee, Closure $extra = null)
    {
        DB::transaction(function () use ($coins, $payer, $payee, $extra) {

            if ($coins->isEmpty()) {
                throw new Exception('Unable to perform operation');
            }

            $coins->each(function ($coin) use ($payer, $payee) {

                DB::table('wallet_coins')
                  ->where([
                              'coin_id'   => $coin['id'],
                              'wallet_id' => $payee->id,
                          ])
                  ->increment('quantity', $coin['quantity']);

                DB::table('wallet_coins')
                  ->where([
                              'coin_id'   => $coin['id'],
                              'wallet_id' => $payer->id,
                          ])
                  ->decrement('quantity', $coin['quantity']);

                DB::table('transactions')
                  ->insert([
                               'amount'          => $coin['cost'],
                               'payer_wallet_id' => $payer->id,
                               'payee_wallet_id' => $payee->id,
                           ]);
            });

            if ($extra instanceof Closure) {
                $extra->call($this);
            }
        });
    }
}