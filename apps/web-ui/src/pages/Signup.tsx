import React from "react";
import { Link } from "react-router-dom";

const Signup: React.FC = () => {
  return (
    <div className="container">
      <h2>Signup</h2>
      <form className="form">
        <input type="text" placeholder="Full Name" required />
        <input type="email" placeholder="Email" required />
        <input type="password" placeholder="Password" required />
        <button type="submit" className="btn">Signup</button>
      </form>

      <div className="social-login">
        <p>or</p>
        <button className="btn">Sign up with Google</button>
      </div>

      <p>
        Already have an account? <Link to="/login">Login here</Link>
      </p>
    </div>
  );
};

export default Signup;
