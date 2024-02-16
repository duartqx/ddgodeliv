import React from "react";

export default function CardFormCreateButton({
  label,
  height,
  onClickHandler,
}) {
  return (
      <div
        className="p-3 d-flex justify-content-center"
        style={{
          position: "absolute",
          bottom: 0,
          background: "linear-gradient(transparent 0%, white 50%)",
          width: "22rem",
          height: `${height}`,
        }}
      >
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
      </div>
  );
}