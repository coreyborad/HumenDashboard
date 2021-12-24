import _axios, { AxiosInstance } from "axios"
import { getToken }  from "../utils/Token"

const axios:AxiosInstance = _axios.create({
    baseURL: 'https://stock.cabbageattic.com/service',
    timeout: 5000,
});

axios.interceptors.request.use(
    config => {
        // If has token, bring it on the headers
        const token = getToken()
        if(token && config.headers){
            config.headers['Authorization'] = 'Bearer ' + token
        }
        return config
    },
    error => {
        console.log('request', error)
        return Promise.reject(error)
    }
)

axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        console.log('response', error)
        return Promise.reject(error)
    }
)
export default axios;
