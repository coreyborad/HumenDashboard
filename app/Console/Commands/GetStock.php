<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Ixudra\Curl\Facades\Curl;
use App\Repositories\StockHistoryRepository;

class GetStock extends Command
{
    protected $stockHistoryRepository;

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
    public function __construct(StockHistoryRepository $stockHistoryRepository)
    {
        parent::__construct();
        $this->stockHistoryRepository = $stockHistoryRepository;
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $response = Curl::to('https://www.twse.com.tw/exchangeReport/STOCK_DAY_ALL?response=json')
            ->asJson()
            ->get();
        foreach ($response->data as $stock) {
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
