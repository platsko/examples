<?php

namespace App\Http\Controllers;

use App\Models\Goods;
use App\Models\Wallet;
use App\Services\VmService;
use Artisan;
use Illuminate\Http\Request;

class VmController extends Controller
{
    protected $goods;

    protected $wallets;

    public function __construct(Goods $goods, Wallet $wallet)
    {
        $this->goods = $goods->get()->keyBy('id');
        $this->wallets = $wallet->with('coins', 'coins.coin')->get()->keyBy('type');
    }

    /**
     * Display a index page
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $depositAmount = $this->wallets->get('deposit')->coins->sum(function ($item) {
            return $item->quantity * $item->coin->cost;
        });

        return view('index', [
            'goods'   => $this->goods,
            'wallets' => $this->wallets,
            'deposit' => $depositAmount,
        ]);
    }

    /**
     * Response data as JSON
     *
     * @return \Illuminate\Http\Response
     */
    public function refresh()
    {
        $depositAmount = $this->wallets->get('deposit')->coins->sum(function ($item) {
            return $item->quantity * $item->coin->cost;
        });

        return response()->json([
                                    'goods'   => $this->goods,
                                    'wallets' => $this->wallets,
                                    'deposit' => $depositAmount,
                                ]);
    }

    /**
     * Put coins to Deposit.
     *
     * @param  \Illuminate\Http\Request $request
     * @param \App\Services\VmService   $vmService
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function putCoin(Request $request, VmService $vmService)
    {
        $userWallet = $this->wallets->get('user');
        $depositWallet = $this->wallets->get('deposit');

        $coin = $userWallet->coins->keyBy('coin_id')->get(
            $request->get('id')
        );

        return $vmService->putCoin($coin, $userWallet, $depositWallet);
    }

    /**
     * Make purchase a selected goods.
     *
     * @param \Illuminate\Http\Request $request
     * @param \App\Services\VmService  $vmService
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function purchase(Request $request, VmService $vmService)
    {
        $goods = $this->goods->get(
            $request->get('id')
        );

        $vmWallet = $this->wallets->get('vm');
        $depositWallet = $this->wallets->get('deposit');

        return $vmService->purchase($goods, $depositWallet, $vmWallet);
    }

    /**
     * Refund deposit.
     *
     * @param \Illuminate\Http\Request $request
     * @param \App\Services\VmService  $vmService
     *
     * @return \Illuminate\Http\RedirectResponse
     * @throws \Exception
     * @throws \Throwable
     */
    public function refund(Request $request, VmService $vmService)
    {
        $vmWallet = $this->wallets->get('vm');
        $userWallet = $this->wallets->get('user');
        $depositWallet = $this->wallets->get('deposit');

        return $vmService->refund($depositWallet, $userWallet, $vmWallet);
    }

    /**
     * Reset the VM to initial condition.
     *
     * @return \Illuminate\Http\RedirectResponse
     */
    public function reset()
    {
        Artisan::call('migrate:refresh', [
            '--seed' => true,
        ]);

        return redirect()->route('index');
    }
}
