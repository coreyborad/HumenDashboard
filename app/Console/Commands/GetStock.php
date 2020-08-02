<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Http;
use App\Repositories\StockHistoryRepository;
use App\Repositories\StockInfoRepository;

class GetStock extends Command
{
    protected $stockHistoryRepository;
    protected $stockInfoRepository;

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
        StockHistoryRepository $stockHistoryRepository,
        StockInfoRepository $stockInfoRepository
    )
    {
        parent::__construct();
        $this->stockHistoryRepository = $stockHistoryRepository;
        $this->stockInfoRepository = $stockInfoRepository;
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
        foreach ($response->data as $stock) {
            $stockInfo = $this->stockHistoryRepository->findByField('stock_name', $stock[0]);
            if(!$stockInfo){
                $stockInfo = $this->stockHistoryRepository->create([
                    'stock_number' => $stock[0],
                    'stock_name' => $stock[1]
                ]);
            }
            $this->stockHistoryRepository->create([
                'stock_number' => $stock[0],
                'deal_date' => $response->date,
                'deal_count' => intval(str_replace(',', '', $stock[2])),
                'price_on_open' => floatval(str_replace(',', '', $stock[4])),
                'price_on_highest' => floatval(str_replace(',', '', $stock[5])),
                'price_on_lowest' => floatval(str_replace(',', '', $stock[6])),
                'price_on_close' => floatval(str_replace(',', '', $stock[7])),
            ]);
        }
    }
}
