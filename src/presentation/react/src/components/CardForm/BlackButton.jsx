import React from "react";

export default function BlackButton({ onClickHandler, label }) {
  return (
    <button
      className="btn align-self-end text-white p-3 shadow"
      style={{
        backgroundColor: "#000",
        width: "100%",
      }}
      onClick={onClickHandler}
    >
      {label}
    </button>
  );
}
