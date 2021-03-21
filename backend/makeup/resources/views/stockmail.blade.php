<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>
        <title>今日股票資訊</title>
    </head>
    <body>
        <table style="border:2px #cccccc solid;" cellpadding="5" border="1">
            <thead>
                <tr>
                    <td style="width:">名稱(代號)</td>
                    <td>股數</td>
                    <td>成本</td>
                    <td>加總損益</td>
                    <td>今日損益</td>
                </tr>
            </thead>
            <tbody>
                @foreach ($result['data'] as $stock)
                    <tr>
                        <td>{{ $stock['stock_name'].'('.$stock['stock_number'].')' }}</td>
                        <td>{{ number_format($stock['stock_shares'], 2) }}</td>
                        <td>{{ number_format($stock['stock_cost'], 2) }}</td>
                        <td>{{ number_format($stock['stock_income'], 2) }}</td>
                        <td>{{ number_format($stock['today_income'], 2) }}</td>
                    </tr>
                @endforeach
            </tbody>
        </table>
        <p>累計加總成本: {{ number_format($result['total_stock_cost'], 2) }}
        <p>累計加總損益: {{ number_format($result['total_stock_income'], 2) }}</p>
        <p>累計加總價值: {{ number_format($result['total_stock_value'], 2) }}</p>
        <p>累計今日損益: {{ number_format($result['total_today_income'], 2) }}</p>
    </body>
</html>
