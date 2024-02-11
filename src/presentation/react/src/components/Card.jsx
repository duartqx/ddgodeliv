import React from "react";

/** @param {{ label: string, value: any }} props */
function CardPart({ label, value }) {
  return (
    <div
      className="border-bottom"
      style={{ paddingBottom: "0.3rem", marginTop: "0.3rem" }}
    >
      <span className="text-body-secondary fw-light">{label}</span> {value}
    </div>
  );
}

/** @param {{ title: string, parts: {label: string, value: any }[]}} props */
export default function Card({ title, parts }) {
  return (
    <div className="card m-3">
      <div className="card-header px-2 d-flex align-items-center text-center">
        <div
          className="rounded-circle img-thumbnail"
          style={{
            backgroundColor: "#000",
            width: "2.2rem",
            height: "2.2rem",
            minWidth: "2.2rem",
          }}
        ></div>
        <strong className="text-center mx-auto">{title}</strong>
      </div>
      <div className="card-body px-3">
        {parts.map(({ label, value }) => (
          <CardPart
            label={label}
            value={value}
            key={`cardPart__${label}__${value}`}
          />
        ))}
      </div>
    </div>
  );
}
