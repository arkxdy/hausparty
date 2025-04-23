import React from "react";
import { Link } from "react-router-dom";

const Welcome: React.FC = () => {
  return (
    <div className="container">
      <h2>🎉 Welcome to HausParty!</h2>
      <p>
        HausParty is your one-stop platform to <strong>host</strong>, <strong>join</strong>, and <strong>rate</strong> house parties in your area!
      </p>
      <p>
        🎈 Create exclusive or public parties <br />
        🧩 Join private events with a unique 6-digit code <br />
        🌟 Rate and review parties you attend <br />
        🧑‍💼 Admins help keep the vibes safe and fun!
      </p>

      <div style={{ display: "flex", flexDirection: "column", gap: "1rem", marginTop: "2rem" }}>
        <Link to="/login" className="btn">Login</Link>
        <Link to="/signup" className="btn">Signup</Link>
      </div>
    </div>
  );
};

export default Welcome;
