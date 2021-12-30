import React, { useLayoutEffect, useState } from 'react';
import './App.css';
import { Login } from './Login/Main'
import { Form } from './views/Form/Main'
import Auth from './utils/Auth';
import { useRoutes, Router } from "react-router-dom";
import customHistory from "./History"


const App = () => {
  let element = useRoutes([
    {
      path: "/",
      element: <Auth/>,
      children: [
        {
          path: "form",
          element: <Form />
        },
      ]
    },
    {
      path: "/login",
      element: <Login />,
    }
  ]);
  return element;
}

const CustomRouter = ({ history, ...props }:{ children: any; history: any; }) => {
  const [state, setState] = useState({
    action: history.action,
    location: history.location
  });

  useLayoutEffect(() => history.listen(setState), [history]);

  return (
    <Router
      {...props}
      location={state.location}
      navigationType={state.action}
      navigator={history}
    />
  );
};

const AppWrapper = () => {
  return (
    <CustomRouter history={customHistory}>
      <App />
    </CustomRouter>
  );
};

export default AppWrapper;
