import Cookies from "js-cookie"

const key:string = "xlsx-form-pwa-token"

export const getToken = () => {
    return Cookies.get(key)
};

export const removeToken = () => {
    return Cookies.remove(key)
};

export const setToken = (token:string) => {
    return Cookies.set(key, token)
};
