import React, { useState } from "react";

const Profile: React.FC = () => {
  const [name, setName] = useState<string>("John Doe");
  const [email, setEmail] = useState<string>("john@example.com");
  const [editMode, setEditMode] = useState<boolean>(false);

  const handleUpdate = (e: React.FormEvent) => {
    e.preventDefault();
    setEditMode(false);
    // API call can go here
    alert("Profile updated!");
  };

  return (
    <div className="container">
      <h2>Your Profile</h2>
      {editMode ? (
        <form className="form" onSubmit={handleUpdate}>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <input type="email" value={email} disabled />
          <button type="submit" className="btn">Save</button>
        </form>
      ) : (
        <div className="form">
          <p><strong>Name:</strong> {name}</p>
          <p><strong>Email:</strong> {email}</p>
          <button className="btn" onClick={() => setEditMode(true)}>
            Edit Profile
          </button>
        </div>
      )}
    </div>
  );
};

export default Profile;
