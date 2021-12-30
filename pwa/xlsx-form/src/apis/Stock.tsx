import {axios} from "./Axios";

export const hello = async () => {
    return axios.get("/stock/api/v1")
};
