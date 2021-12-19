import React from 'react';
import './App.css';
import { Login } from './Login/Main'
import { Form } from './Form/Main'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

function App() {
  return (
    <div>
      <Router>
          <Routes>
            <Route path="/" element={<Login/>}/>
            <Route path="/form" element={<Form/>}/>
          </Routes>
      </Router>
    </div>

  );
}

export default App;
