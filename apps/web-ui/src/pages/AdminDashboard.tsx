import React from "react";

const AdminDashboard: React.FC = () => {
  // Mock data
  const pendingParties = [
    { id: 1, name: "Secret Rave", host: "Jane" },
    { id: 2, name: "Silent Disco", host: "Bob" },
  ];

  return (
    <div className="container">
      <h2>Admin Dashboard</h2>

      <section>
        <h3>Pending Party Approvals</h3>
        <ul>
          {pendingParties.map((party) => (
            <li key={party.id}>
              {party.name} by {party.host}
              <div>
                <button className="btn small">Approve</button>
                <button className="btn small">Reject</button>
              </div>
            </li>
          ))}
        </ul>
      </section>
    </div>
  );
};

export default AdminDashboard;
