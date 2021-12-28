import { Accounting } from "../interfaces/Interface";
import axios from "./Axios";

export const appendRecord = async (info:Accounting) => {
    return axios.post("/stock/api/v1/xlsx", info)
};
