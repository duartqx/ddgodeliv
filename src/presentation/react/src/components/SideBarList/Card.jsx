import React from "react";
import RoundImage from "../RoundImage";
import "./Card.css";

/** @param {{ value: any, border: boolean }} props */
function CardPart({ value, border }) {
  return (
    <div className={border ? "border-bottom" : ""}>
      <span className="text-muted fw-medium">{value}</span>
    </div>
  );
}

/** @param {{
 * title: string
 * parts: {
 *   value: any
 *   border: boolean
 *  }[]
 *  onClickHandler: () => void
 * }} props */
export default function Card({ title, parts, onClickHandler, selected }) {
  return (
    <div className="card m-2">
      <div
        className="p-2 d-flex align-items-center hover-card"
        onClick={onClickHandler}
        style={{
          backgroundColor: selected && "#f0f2f7",
        }}
      >
        <RoundImage size="3.0rem" />
        <div className="d-flex flex-column">
          <div className="fw-semibold text-center" style={{ width: "100%" }}>
            {title}
          </div>
          {parts.map(({ label, value, border }) => (
            <CardPart
              value={value}
              border={border}
              key={`cardPart__${label}__${value}`}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
