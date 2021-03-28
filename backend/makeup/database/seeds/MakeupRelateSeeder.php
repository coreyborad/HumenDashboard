<?php

use Illuminate\Database\Seeder;
use App\Models\MakeupSale;
use App\Models\MakeupCost;
use App\Models\MakeupSaleCostRelate;

class MakeupRelateSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $sales = MakeupSale::all();
        $sales->each(function ($sale, $key) {
            $remain = $sale->count;
            $cost_list = MakeupCost::with(['hadSold'])
                ->orderBy('order_date', 'asc')
                ->where('makeup_id', '=', $sale->makeup_id)
                ->get();
            $cost_list = $cost_list->filter(function($cost){
                $rem = ($cost->count - $cost->hadSold->sum('relate_count'));
                if($rem <= 0){
                    return false;
                }else{
                    return true;
                }
            });
            // 新增至銷售成本關係表
            foreach ($cost_list as $cost) {
                if($remain > $cost->count){
                    $remain = $remain - $cost->count;
                    MakeupSaleCostRelate::create([
                        'cost_id' => $cost->id,
                        'sale_id' => $sale->id,
                        'relate_count' => $cost->count
                    ]);
                // 小於或等於時代表銷售數量已攤平完畢
                } else {
                    MakeupSaleCostRelate::create([
                        'cost_id' => $cost->id,
                        'sale_id' => $sale->id,
                        'relate_count' => $remain
                    ]);
                    break;
                }
            }
        });
    }
}