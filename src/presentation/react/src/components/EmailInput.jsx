import React from "react";

/**
 * @param {{
 *  email: string
 *  setEmail: React.Dispatch<React.SetStateAction<string>>
 * }} props
 **/
export default function EmailInput({ email, setEmail }) {
  return (
    <div className="form-group">
      <div className="input-group mb-3">
        <label className="input-group-text" style={{ minWidth: "6rem" }}>
          Email
        </label>
        <input
          type="email"
          className="form-control"
          placeholder="example@email.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>
    </div>
  );
}
