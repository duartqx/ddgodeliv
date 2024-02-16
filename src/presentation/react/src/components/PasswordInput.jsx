import React from "react";

/**
 * @param {{
 *  password: string
 *  setPassword: React.Dispatch<React.SetStateAction<string>>
 * }} props
 **/
export default function PasswordInput({ password, setPassword }) {
    return (
        <div className="form-group">
            <div className="input-group mb-3">
                <label
                    className="input-group-text"
                    style={{ minWidth: "8rem" }}
                >
                    Password
                </label>
                <input
                    type="password"
                    className="form-control"
                    placeholder="**********"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
            </div>
        </div>
    );
}
