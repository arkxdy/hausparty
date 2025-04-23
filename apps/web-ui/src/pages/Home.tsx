import React from "react";
import { Link } from "react-router-dom";

const Home: React.FC = () => {
  return (
    <div className="container">
      <h2>Hey party animal 🕺💃</h2>
      <p>What would you like to do today?</p>

      <div style={{ display: "flex", flexDirection: "column", gap: "1rem", marginTop: "1.5rem" }}>
        <Link to="/host" className="btn">Host a Party</Link>
        <Link to="/join" className="btn">Join a Party</Link>
        <Link to="/history" className="btn">My Party History</Link>
        <Link to="/profile" className="btn">My Profile</Link>
      </div>
    </div>
  );
};

export default Home;
