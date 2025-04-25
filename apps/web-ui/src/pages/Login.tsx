import React from "react";
import { Link } from "react-router-dom";

const Login: React.FC = () => {
  return (
    <div className="container">
      <h2>Login</h2>
      <form className="form">
        <input type="email" placeholder="Email" required />
        <input type="password" placeholder="Password" required />
        <button type="submit" className="btn">Login</button>
      </form>

      <div className="social-login">
        <p>or</p>
        <button className="btn">Sign in with Google</button>
      </div>

      <p>
        Don't have an account? <Link to="/signup">Signup here</Link>
      </p>
    </div>
  );
};

export default Login;
