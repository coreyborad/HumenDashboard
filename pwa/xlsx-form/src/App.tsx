import React from 'react';
import './App.css';
import { Login } from './Login/Main'
import { Form } from './views/Form/Main'
import Auth from './utils/Auth';
import { BrowserRouter as Router} from 'react-router-dom'
import { useRoutes } from "react-router-dom";

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

// function App() {
//   return (
//     <div>
//       <Router>
//           <Routes>
//             <Route path="/login" element={<Login/>}></Route>
//             <Route path="/" element={<Login/>}>
//               <Route element={<Auth/>}>
//                 <Route path="form" element={<Form/>}/>
//               </Route>
//             </Route>
//           </Routes>
//       </Router>
//     </div>

//   );
// }

const AppWrapper = () => {
  return (
    <Router>
      <App />
    </Router>
  );
};

export default AppWrapper;
