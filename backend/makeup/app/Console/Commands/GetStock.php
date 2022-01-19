<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Http;
use App\Repositories\Mongo\StockHistoryRepository as MongoStockHistoryRepository;
use App\Repositories\StockInfoRepository;
use Carbon\Carbon;
use Exception;

class GetStock extends Command
{
    protected $stockInfoRepository;
    protected $mongoStockHistoryRepository;

    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'stock:get';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Get stock';

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

        $response = Http::get('https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL?response=json');
        $response = $response->json();
        foreach ($response['data'] as $stock) {
            try {
                $stockInfo = $this->stockInfoRepository->findByField('stock_number', $stock[0])->first();
                if($stockInfo === null){
                    $stockInfo = $this->stockInfoRepository->create([
                        'stock_number' => $stock[0],
                        'stock_name' => $stock[1]
                    ]);
                }
                $deal_date = Carbon::parse($response['date'])->setTimezone('Asia/Taipei');
                $this->mongoStockHistoryRepository->create([
                    'stock_number' => $stock[0],
                    'deal_date' => $deal_date,
                    'deal_count' => intval(str_replace(',', '', $stock[2])),
                    'price_on_open' => floatval(str_replace(',', '', $stock[4])),
                    'price_on_highest' => floatval(str_replace(',', '', $stock[5])),
                    'price_on_lowest' => floatval(str_replace(',', '', $stock[6])),
                    'price_on_close' => floatval(str_replace(',', '', $stock[7])),
                ]);
                
            } catch (Exception $e) {
                // var_dump($e->getMessage());
            }
        }
    }
}
