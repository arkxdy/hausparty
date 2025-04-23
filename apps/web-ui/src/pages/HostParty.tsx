import React, { useState } from "react";

const HostParty: React.FC = () => {
  const [isPrivate, setIsPrivate] = useState<boolean>(false);

  return (
    <div className="container">
      <h2>Host a Party</h2>
      <form className="form">
        <input type="text" placeholder="Party Name" required />
        <textarea placeholder="Description" required />
        <input type="datetime-local" required />
        <input type="text" placeholder="Location" required />
        <input type="number" placeholder="Max Guests" required />
        <label>
          <input
            type="checkbox"
            onChange={() => setIsPrivate(!isPrivate)}
          />
          Private Party
        </label>
        {isPrivate && (
          <input type="text" placeholder="6-digit Code" maxLength={6} />
        )}
        <button type="submit" className="btn">Create Party</button>
      </form>
    </div>
  );
};

export default HostParty;
