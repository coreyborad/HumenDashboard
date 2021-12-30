import {axios} from "./Axios";
import { User } from "../interfaces/Interface"
import { setToken }  from "../utils/Token"

export const login = async (data: User) =>{
    const result = await axios.post("/makeup/api/v1/login", {
        grant_type: "password",
        email: data.email,
        password: data.password
    })
    setToken(result.data.access_token)
    return result
};
