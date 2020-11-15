<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Http;
use App\Repositories\Mongo\StockHistoryRepository as MongoStockHistoryRepository;
use App\Repositories\StockInfoRepository;
use Carbon\Carbon;
use Exception;

class GetStockHistory extends Command
{
    protected $stockInfoRepository;
    protected $mongoStockHistoryRepository;

    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'stock:gethistory';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Get stock history';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct(
        StockInfoRepository $stockInfoRepository,
        MongoStockHistoryRepository $mongoStockHistoryRepository
    )
    {
        parent::__construct();
        $this->stockInfoRepository = $stockInfoRepository;
        $this->mongoStockHistoryRepository = $mongoStockHistoryRepository;
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $stock_info = $this->stockInfoRepository->get();
        foreach ($stock_info as $stock) {
            $stock_number = $stock->stock_number;
            var_dump($stock_number);
            try {
                $response = Http::get('https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20201101&stockNo=' . $stock_number);
                $response = $response->json();
                foreach ($response['data'] as $record) {
                    $record[0] = str_replace("/", "-", $record[0]);
                    $insert = [
                        'stock_number' => $stock_number,
                        'deal_date' => Carbon::parse($record[0])->addyear(1911),
                        'deal_count' => intval(str_replace(',', '', $record[1])),
                        'price_on_open' => floatval(str_replace(',', '', $record[3])),
                        'price_on_highest' => floatval(str_replace(',', '', $record[4])),
                        'price_on_lowest' => floatval(str_replace(',', '', $record[5])),
                        'price_on_close' => floatval(str_replace(',', '', $record[6])),
                    ];
                    $this->mongoStockHistoryRepository->create($insert);
                }
                sleep(5);
            } catch (\Throwable $th) {
            }
        }
    }
}
