import React, { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

import EmailInput from "../components/EmailInput";
import NameInput from "../components/NameInput";
import PasswordInput from "../components/PasswordInput";
import SubmitAndLinkButton from "../components/SubmitAndLinkButton";
import Error from "../components/Error";
import { AuthContext } from "../middlewares/AuthContext";

export default function Register() {
  const navigate = useNavigate();
  const location = useLocation();
  const { from } = location.state || { from: { pathname: "/dashboard" } };
  const { register } = useContext(AuthContext);

  const [error, setError] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleRegister = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();

    try {
      /** @type {boolean} */
      const registered = await register({ name, email, password });

      if (registered) {
        navigate(from);
      } else {
        setError("Error, could not create user");
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
            <h2 className="mb-3">Sign Up</h2>
            <form
              onSubmit={handleRegister}
              method="post"
              className="flex flex-column align-items-center"
              style={{ height: "100%" }}
            >
              <NameInput name={name} setName={setName} />
              <EmailInput email={email} setEmail={setEmail} />
              <PasswordInput password={password} setPassword={setPassword} />
              <SubmitAndLinkButton
                submitLabel="Create Account"
                linkTo="/login"
                linkLabel="Sign In Instead"
              />
            </form>
            {error ? <Error err={error} /> : ""}
          </div>
        </div>
      </div>
    </>
  );
}
