import React from "react";

/**
 * @param {{
 *  name: string
 *  setName: React.Dispatch<React.SetStateAction<string>>
 * }} props
 **/
export default function NameInput({ name, setName }) {
  return (
    <div className="form-group">
      <div className="input-group mb-3">
        <label className="input-group-text" style={{ minWidth: "6rem" }}>
          Name
        </label>
        <input
          type="text"
          className="form-control"
          placeholder="Maria Silva"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </div>
    </div>
  );
}
