import _axios, { AxiosInstance } from "axios"
import { getToken, removeToken }  from "../utils/Token"
import { useNavigate } from 'react-router-dom';
import histroy from '../History'

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
        // If 401 is mean, token invalid
        if(error.response.status === 401) {
            // remove token
            removeToken()
            histroy.replace("/login");
        }
        return Promise.reject(error)
    }
)
export default axios;
