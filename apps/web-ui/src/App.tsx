import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import HostParty from "./pages/HostParty";
import JoinParty from "./pages/JoinParty";
import History from "./pages/History";
import Profile from "./pages/Profile";
import AdminDashboard from "./pages/AdminDashboard";
import "./App.css";
import Welcome from "./pages/Welcome";


const App: React.FC = () => {
  return (
    <Router>
      <Routes>
      <Route path="/" element={<Welcome />} />
      <Route path="/home" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/host" element={<HostParty />} />
        <Route path="/join" element={<JoinParty />} />
        <Route path="/history" element={<History />} />
        
        <Route path="/profile" element={<Profile />} />
        <Route path="/admin" element={<AdminDashboard />} />
      </Routes>
    </Router>
  );
};


export default App
