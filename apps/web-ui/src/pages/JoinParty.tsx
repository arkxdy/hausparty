import React, { useState } from "react";

const JoinParty: React.FC = () => {
  const [privateCode, setPrivateCode] = useState<string>("");

  return (
    <div className="container">
      <h2>Join a Party</h2>
      <div className="form">
        <h3>Join Public Party</h3>
        <input type="text" placeholder="Search party by name or location" />
        <button className="btn">Browse Public Parties</button>

        <h3>Join Private Party</h3>
        <input
          type="text"
          maxLength={6}
          value={privateCode}
          onChange={(e) => setPrivateCode(e.target.value)}
          placeholder="Enter 6-digit code"
        />
        <button className="btn">Join Private Party</button>
      </div>
    </div>
  );
};

export default JoinParty;
