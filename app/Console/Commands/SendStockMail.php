<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Repositories\UserRepository;
use App\Repositories\StockHistoryRepository;
use App\Repositories\StockInfoRepository;
use Carbon\Carbon;
use App\Mail\StockDaily as DailyMail;
use Illuminate\Support\Facades\Mail;

class SendStockMail extends Command
{

    protected $userRepository;
    protected $stockHistoryRepository;
    protected $stockInfoRepository;

    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'send:stockmail';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Command description';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct(
        UserRepository $userRepository,
        StockHistoryRepository $stockHistoryRepository,
        StockInfoRepository $stockInfoRepository
    )
    {
        parent::__construct();
        $this->userRepository = $userRepository;
        $this->stockHistoryRepository = $stockHistoryRepository;
        $this->stockInfoRepository = $stockInfoRepository;
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        $this->userRepository->with(['stocks'])
        ->all()
        ->each(function($user){
            $stock_list = [];
            foreach ($user->stocks as $stock) {
                array_push($stock_list, $stock->stock_number);
            }
            
            $lasttwodays = $this->stockHistoryRepository
                ->select('deal_date')
                ->groupBy('deal_date')
                ->orderBy('deal_date', 'desc')
                ->limit(2)
                ->get();

            $historys = $this->stockHistoryRepository
                ->scopeQuery(function($query) use($stock_list, $lasttwodays){
                    return $query->whereIn('stock_number', $stock_list)
                                ->whereIn(
                                    'deal_date', 
                                    [
                                        $lasttwodays[0]->deal_date,
                                        $lasttwodays[1]->deal_date
                                    ]
                                );
                })
                ->get()
                ->groupBy('stock_number');
            
            // stock_number
            // stock_name
            // stock_shares
            // stock_cost
            // stock_income
            // today_income
            $result = [
                'data' => [],
                'total_stock_income' => 0,
                'total_today_income' => 0,
                'total_stock_cost' => 0,
                'total_stock_value' => 0
            ];
            foreach ($historys as $stock) {
                $this_stock_number = $stock[0]->stock_number;
                $user_stock = $user->stocks->filter(
                    function ($item) use($this_stock_number) {
                        return $this_stock_number === $item->stock_number;
                    }
                )->first();
                $stock_income = bcsub(bcmul($stock[1]->price_on_close, $user_stock->shares, 2), $user_stock->cost, 2);
                $today_income = bcmul(bcsub($stock[1]->price_on_close, $stock[0]->price_on_close, 2), $user_stock->shares, 2);
                array_push($result['data'], [
                    'stock_number' => $this_stock_number,
                    'stock_name' => $this->stockInfoRepository
                                    ->findByField('stock_number', $this_stock_number)
                                    ->first()
                                    ->stock_name,
                    'stock_shares' => $user_stock->shares,
                    'stock_cost' => $user_stock->cost,
                    'stock_income' => $stock_income,
                    'today_income' => $today_income,
                ]);
                $result['total_stock_income'] = bcadd($result['total_stock_income'], $stock_income, 2);
                $result['total_today_income'] = bcadd($result['total_today_income'], $today_income, 2);
                $result['total_stock_cost'] = bcadd($result['total_stock_cost'], $user_stock->cost, 2);
            }
            $result['total_stock_value'] = bcadd($result['total_stock_cost'], $result['total_stock_income'], 2);
            Mail::to($user->email)->send(new DailyMail($result));
        });
        return 0;
    }
}
