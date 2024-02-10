import React, { useState } from "react";
import { useLocation, useNavigate, Link } from "react-router-dom";
import { login } from "../services/auth";
import useAuth from "../hooks/useAuth";

/**
 * @typedef {{
 *  token: ?string
 *  expiresAt: ?string
 *  status: ?string
 * }} AuthData
 */

export default function Login() {
  const navigate = useNavigate();
  const location = useLocation();
  const { from } = location.state || { from: { pathname: "/dashboard" } };

  const { setAuth } = useAuth();
  const [error, setError] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    try {
      /** @type {AuthData} */
      const authData = await login({ email, password });
      setAuth(authData);
      if (authData?.status) {
        navigate(from);
      } else {
        setError("Login Error");
      }
    } catch (e) {
      console.log(e);
    }
  };

  return (
    <>
      <div className="content">
        <div className="p-5">
          <div
            className="col-md-6"
            style={{ maxWidth: "24rem", paddingTop: "14rem" }}
          >
            <form
              onSubmit={handleLogin}
              method="post"
              className="flex flex-column align-items-center"
              style={{ height: "100%" }}
            >
              <div className="form-group">
                <label>Email</label>
                <input
                  type="email"
                  className="form-control"
                  placeholder="example@email.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div className="form-group">
                <label>Password</label>
                <input
                  type="password"
                  className="form-control"
                  placeholder="**********"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <div className="mt-3 d-flex col-md-12">
                <button type="submit" className="col-md-5 btn btn-primary">
                  Login
                </button>
                <Link
                  to="/register"
                  className="col-md-5 ms-auto btn btn-outline-primary"
                >
                  Create Account
                </Link>
              </div>
            </form>
            {error ? (
              <div className="my-3 py-2 alert alert-danger">{error}</div>
            ) : (
              ""
            )}
          </div>
        </div>
      </div>
    </>
  );
}
