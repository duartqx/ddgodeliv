import React, { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import EmailInput from "../components/EmailInput";
import PasswordInput from "../components/PasswordInput";
import SubmitAndLinkButton from "../components/SubmitAndLinkButton";
import Error from "../components/Error";
import { AuthContext } from "../middlewares/AuthContext";
import * as Types from "../middlewares/authTypes";

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

  const { login } = useContext(AuthContext);
  const [error, setError] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    try {
      if (await login({ email, password })) {
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
            <h2 className="mb-3">Sign In</h2>
            <form
              onSubmit={handleLogin}
              method="post"
              className="flex flex-column align-items-center"
              style={{ height: "100%" }}
            >
              <EmailInput email={email} setEmail={setEmail} />
              <PasswordInput password={password} setPassword={setPassword} />
              <SubmitAndLinkButton
                submitLabel="Submit"
                linkTo="/register"
                linkLabel="Sign up Instead"
              />
            </form>
            {error && <Error err={error} />}
          </div>
        </div>
      </div>
    </>
  );
}
