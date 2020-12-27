<?php

namespace App\Services;

use App\Exceptions\ErrorException;
use App\Repositories\MakeupInfoRepository;
use App\Repositories\MakeupPriceRepository;


class MakeupService
{
    protected $makeupPriceRepository;
    protected $makeupInfoRepository;

    public function __construct(
        MakeupPriceRepository $makeupPriceRepository,
        MakeupInfoRepository $makeupInfoRepository
    )
    {
        $this->makeupPriceRepository = $makeupPriceRepository;
        $this->makeupInfoRepository = $makeupInfoRepository;
    }

    public function getUserMakeupList(int $user_id)
    {
        // $user_Makeup_list = $this->userHasMakeupRepository->with('Makeup_info')->findByField('user_id', $user_id);
        // $user_Makeup_list = $user_Makeup_list->map(function($Makeup){
        //     $last_Makeup = $this->MakeupHistoryRepository
        //         ->where('Makeup_number', $Makeup->Makeup_number)
        //         ->orderBy('deal_date', 'desc')
        //         ->first();
        //     $Makeup->last_Makeup = $last_Makeup;
        //     return $Makeup;
        // });
        // return $user_Makeup_list;
    }

    public function createUserMakeup(int $user_id, array $data)
    {
        // try {
        //     $data = $this->userHasMakeupRepository->create([
        //         'user_id' => $user_id,
        //         'Makeup_number' => $data['Makeup_number'],
        //         'shares' => $data['shares'],
        //         'cost' => $data['cost'],
        //     ]);
        // } catch (\Throwable $th) {
        //     throw new ErrorException(500, 'error');
        // }
        // return $data;
    }

    public function deleteUserMakeup(int $id)
    {
        // try {
        //     $data = $this->userHasMakeupRepository->delete($id);
        // } catch (\Throwable $th) {
        //     throw new ErrorException(500, 'error');
        // }
        // return $data;
    }

}
