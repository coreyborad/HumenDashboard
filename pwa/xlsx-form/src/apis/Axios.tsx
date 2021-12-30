import _axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios"
import { CircularProgress, Snackbar, Alert} from '@mui/material';
import { getToken, removeToken } from "../utils/Token"
import histroy from '../History'
import { useMemo, useState, useEffect } from 'react';

const ax: AxiosInstance = _axios.create({
    baseURL: 'https://stock.cabbageattic.com/service',
    timeout: 5000,
});

export const axios = ax;

const useAxiosLoader = () => {
    const [counter, setCounter] = useState(0);
    const [showError, setShowError] = useState(false);
    const [errorMsg, setErrorMsg] = useState("");

    const closeError = (_event: any, reason: string) => {
        if (reason === 'clickaway') {
          return;
        }
        setShowError(false);
    };

    const interceptors = useMemo(() => {
        const inc = () => setCounter((counter: number) => counter + 1);
        const dec = () => setCounter((counter: number) => counter - 1);

        return ({
            request: (config: AxiosRequestConfig) => {
                // If has token, bring it on the headers
                inc()
                const token = getToken()
                if (token && config.headers) {
                    config.headers['Authorization'] = 'Bearer ' + token
                }
                return config
            },
            requestErr: (error: any) => {
                dec()
                console.log('request', error)
                setErrorMsg(error)
                setShowError(true)
                return Promise.reject(error)
            },
            response: (response: AxiosResponse) => {
                dec()
                return response
            },
            responseErr: (error:any) => {
                dec()
                setErrorMsg(error.response.data.error || error.response.data.error_description)
                setShowError(true)
                // If 401 is mean, token invalid
                if (error.response.status === 401) {
                    // remove token
                    removeToken()
                    histroy.replace("/login");
                }
                return Promise.reject(error)
            }
        });
    }, []);

    useEffect(() => {
        // add request interceptors
        const reqInterceptor = ax.interceptors.request.use(interceptors.request, interceptors.requestErr);
        // add response interceptors
        const resInterceptor = ax.interceptors.response.use(interceptors.response, interceptors.responseErr);
        return () => {
            // remove all intercepts when done
            ax.interceptors.request.eject(reqInterceptor);
            ax.interceptors.response.eject(resInterceptor);
        };
    }, [interceptors]);

    return [
        counter > 0,
        (
        <Snackbar
            anchorOrigin={{
                vertical: "top",
                horizontal: "center",
            }}
            open={showError}
            autoHideDuration={3000}
            onClose={closeError}
        >
            <Alert severity="error">{errorMsg}</Alert>
        </Snackbar>
        )
    ];
};

export const GlobalLoader = () => {
    const [loading, showError] = useAxiosLoader();

    return (
        <div>
            {/* Global Loading */}
            <div
                style={{
                    position: "fixed",
                    display: loading ? "block":"none",
                    width: "100%",
                    height: "100%",
                    top: 0,
                    left: 0,
                    opacity: 0.7,
                    backgroundColor: "#000",
                    zIndex: 99
                }}
            >
                {loading }
                <CircularProgress style={{
                    position: "relative",
                    top: "calc(50% - 40px)",
                    left: "calc(50% - 40px)",

                }}/>
            </div>
            {showError}
        </div>

    );
}
