import React from "react";

export default function CreateNewButton({ label, onClickHandler }) {
  return (
    <div
      className="p-3 d-flex align-items-center"
      style={{
        background: "linear-gradient(transparent 0%, #f8f0fa 60%)",
        position: "absolute",
        bottom: 0,
        width: "100%",
        height: "7rem",
      }}
    >
      <button
        className="btn mx-auto align-self-end text-white p-3 drop-shadow"
        style={{ backgroundColor: "#000", width: "90%" }}
        onClick={onClickHandler}
      >
        {label}
      </button>
    </div>
  );
}
