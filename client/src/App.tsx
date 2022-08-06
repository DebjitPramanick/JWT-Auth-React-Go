import React from "react";
import { Link, Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { Login } from "./views/Login";
import "./styles.css"
import Register from "./views/Register";
import Home from "./views/Home";

export const App = () => {
  return (
    <div className="App">
      <Router>
        <div className="nav-bar">
          <Link to="/">Home</Link>
          <Link to="/login">Login</Link>
          <Link to="/register">Register</Link>
        </div>
        <Routes>
          <Route path="/" element={<Home />}></Route>
          <Route path="/register" element={<Register />}></Route>
          <Route path="/login" element={<Login />}></Route>
        </Routes>
      </Router>
    </div>
  );
}

export default App;
