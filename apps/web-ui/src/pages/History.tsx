import React, { useState } from "react";

const History: React.FC = () => {
  const [activeTab, setActiveTab] = useState<"hosted" | "joined">("hosted");

  const hostedParties = [
    { name: "Rooftop Bash", date: "2025-03-21" },
    { name: "BBQ & Beats", date: "2025-02-14" },
  ];

  const joinedParties = [
    { name: "Karaoke Night", date: "2025-03-02" },
    { name: "Neon Party", date: "2025-01-30" },
  ];

  return (
    <div className="container">
      <h2>📜 My Party History</h2>

      <div className="tabs">
        <button
          className={activeTab === "hosted" ? "tab active" : "tab"}
          onClick={() => setActiveTab("hosted")}
        >
          🏠 Hosted
        </button>
        <button
          className={activeTab === "joined" ? "tab active" : "tab"}
          onClick={() => setActiveTab("joined")}
        >
          🎉 Joined
        </button>
      </div>

      <ul className="party-list">
        {(activeTab === "hosted" ? hostedParties : joinedParties).map((party, index) => (
          <li key={index}>
            <strong>{party.name}</strong> <br />
            <small>{party.date}</small>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default History;
