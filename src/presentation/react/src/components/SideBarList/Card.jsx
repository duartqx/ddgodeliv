import React from "react";
import RoundImage from "../RoundImage";
import "./Card.css";

/** @typedef {{ value: any, label: any, border: boolean }} CardPartProps */

/** @param {CardPartProps} props */
function CardPart({ value, label, border }) {
  return (
    <div className={border ? "border-bottom" : ""}>
      <span className="text-muted fw-medium">{label}</span>{" "}
      <span className="fw-light text-muted">{value}</span>
    </div>
  );
}

/** @param {{
 *  title: string
 *  parts: CardPartProps[]
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
        <div className="d-flex flex-column flex-grow-1">
          <div className="fw-semibold" style={{ width: "100%" }}>
            {title}
          </div>
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
    </div>
  );
}
