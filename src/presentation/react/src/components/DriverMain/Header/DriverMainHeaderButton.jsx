import React from "react";

/** @param {{ icon: string, onClickHandler: () => void }} */
export default function DriverMainHeaderButton({ icon, onClickHandler }) {
  return (
    <button className="btn" onClick={onClickHandler}>
      <i className={`bi ${icon}`}></i>
    </button>
  );
}
