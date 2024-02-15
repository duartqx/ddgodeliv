import React from "react";
import RoundImage from "../RoundImage";

/** @param {{ label: string, value: any, border: boolean }} props */
function CardPart({ label, value, border }) {
  return (
    <div className={border ? "border-bottom" : ""}>
      <span className="text-body-secondary fw-light">{label}</span> {value}
    </div>
  );
}

/** @param {{
 * title: string
 * parts: {
 *   label: string
 *   value: any
 *   border: boolean
 *  }[]
 *  onClickHandler: () => void
 * }} props */
export default function Card({ title, parts, onClickHandler, selected }) {
  return (
    <div className="card m-3">
      <div
        className="card-header px-2 d-flex align-items-center text-center"
        onClick={onClickHandler}
        style={{
          cursor: "pointer",
          backgroundColor: selected && "#deeafc"
        }}
      >
        <RoundImage size="2.2rem" />
        <div className="pl-4 fw-semibold" style={{ width: "100%" }}>
          {title}
        </div>
      </div>
      <div className="card-body px-3">
        {parts.map(({ label, value, border }) => (
          <CardPart
            label={label}
            value={value}
            border={border}
            key={`cardPart__${label}__${value}`}
          />
        ))}
      </div>
    </div>
  );
}
