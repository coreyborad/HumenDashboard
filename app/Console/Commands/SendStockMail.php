<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use App\Repositories\UserRepository;
use App\Repositories\StockHistoryRepository;

class SendStockMail extends Command
{

    protected $userRepository;
    protected $stockHistoryRepository;

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
        StockHistoryRepository $stockHistoryRepository
    )
    {
        parent::__construct();
        $this->userRepository = $userRepository;
        $this->stockHistoryRepository = $stockHistoryRepository;
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
            
            $historys = $this->stockHistoryRepository
                ->scopeQuery(function($query) use($stock_list){
                    $d = mktime(0, 0, 0, 7, 31, 2020);
                    return $query->whereIn('stock_number', $stock_list)
                                ->where('deal_date', date("Y-m-d", $d));
                })
                ->get();
            foreach ($historys as $key => $value) {
                var_dump($value->deal_count);
            }
        });
        return 0;
    }
}
