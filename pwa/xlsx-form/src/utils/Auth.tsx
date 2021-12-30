import React from "react";
import { getToken }  from "../utils/Token"
import { useNavigate, Outlet, useLocation } from 'react-router-dom';

function Auth (){
    const navigate = useNavigate();
    const { pathname } = useLocation();
    React.useEffect(() => {
        const token = getToken()
        if (!token){
            navigate("/login");
        }
        if (pathname === "/"){
            navigate("/form");
        }
    }, [navigate, pathname]);
    return <Outlet />
}
export default Auth;