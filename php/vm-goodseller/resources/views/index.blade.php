<!doctype html>
<html lang="{{ app()->getLocale() }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>VM Goods Seller - Laravel WEB Application</title>

    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <style type="text/css">
        button {
            cursor: pointer;
        }
        button:disabled {
            cursor: default;
        }
    </style>
</head>
<body>
<div class="container-fluid">

    <div class="row">
        <div class="col-sm-12">
            <div class="alert text-center" role="alert">
                <strong id="response_message">&nbsp;</strong>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="col-sm-12">
            <div class="card text-center">
                <div class="card-body">
                    <h5 class="card-title">Депозит</h5>
                    <h6 class="card-subtitle mb-2 text-muted">
                        <strong id="deposit_amount">{{ $deposit }}</strong> руб
                    </h6>
                    <button type="button" data-id="0" data-route="{!! route('refund') !!}">
                        Сдача / Возврат
                    </button>
                </div>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="col-sm-6">
            <div class="card">
                <div class="card-body">
                    <h5 class="card-title">Кошелек пользователя</h5>
                </div>
                <ul class="list-group list-group-flush">
                    @foreach($wallets->get('user')->coins as $item)
                    <li class="list-group-item">
                        Номинал:
                        <button type="button"
                                data-id="{{ $item->coin->id }}"
                                data-route="{!! route('putcoin') !!}"
                                style="width: 5rem;"
                        >
                            {{ $item->coin->cost }} руб
                        </button>
                        Доступно: <strong id="wallets_user_coins_coin_id_{{ $item->coin->id }}">
                            {{ $item->quantity }}
                        </strong> шт
                    </li>
                    @endforeach
                </ul>
            </div>
        </div>
        <div class="col-sm-6">
            <div class="card">
                <div class="card-body">
                    <h5 class="card-title">VM</h5>
                </div>
                <ul class="list-group list-group-flush">
                    @foreach($wallets->get('vm')->coins as $item)
                    <li class="list-group-item">
                        <button type="button" style="width: 5rem;" disabled>
                            {{ $item->coin->cost }} руб
                        </button>
                        Остаток: <strong id="wallets_vm_coins_coin_id_{{ $item->coin->id }}">
                            {{ $item->quantity }}
                        </strong> шт
                    @endforeach
                </ul>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="col-sm-12">
            <div class="card text-center">
                <div class="card-body">
                    <h5 class="card-title">Товары</h5>
                </div>
                <ul class="list-group list-group-flush">
                    @foreach($goods as $item)
                        <li class="list-group-item">
                            <button type="button"
                                    data-id="{{ $item->id }}"
                                    data-route="{!! route('purchase') !!}"
                                    style="width: 15rem;"
                            >
                                {{ $item->name }}: {{ $item->cost }} руб
                            </button>
                            Доступно: <strong id="goods_id_{{ $item->id }}">
                                {{ $item->quantity }}
                            </strong> шт
                        </li>
                    @endforeach
                </ul>
            </div>
        </div>
    </div>

    <div class="row">
        <div class="col-sm-12">
            <div class="card text-center">
                <div class="card-body">
                    <form method="post" action="{!! route('reset') !!}">
                        @csrf
                        <button type="submit">Сброс в начальное состояние</button>
                    </form>
                </div>
            </div>
        </div>
    </div>

</div>

<!-- Optional JavaScript -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js" crossorigin="anonymous"></script>

<script type="application/javascript">
$(function() {
    var buttons = $('button[data-id][data-route]'),
        message = $('#response_message'),
        msgClear = function () {
            message.html('&nbsp;').parent('div').removeClass('alert-success alert-danger');
        };
    buttons.each(function () {
        var item = $(this);
        item.on('click', function (e) {
            msgClear();
            buttons.attr('disabled', true);
            $.post(item.data('route'),{ id: item.data('id') }, function (resp) {
                if (resp.status === 'error') {
                    message.text(resp.message).parent('div').addClass('alert-danger');
                } else {
                    $('#deposit_amount').text(resp.deposit);
                    $.each(resp.goods, function (idx, goods) {
                        $('#goods_id_'+goods.id).text(goods.quantity);
                    });
                    $.each(resp.wallets, function (key, item) {
                        $.each(item.coins, function (idx, coin) {
                            $('#wallets_'+key+'_coins_coin_id_'+coin.coin_id).text(coin.quantity);
                        });
                    });
                }
            })
            .fail(function(resp) {
                message.text(resp.status + ': ' + resp.statusText).parent('div').addClass('alert-danger');
            })
            .always(function() {
                buttons.attr('disabled', false);
            });
        });
    });
    $('button[type="submit"]').on('click', function (e) {
        $('button').attr('disabled', true);
        $(this).closest('form').submit();
    });
});
</script>

</body>
</html>
