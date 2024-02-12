import React from "react";

/** @param {{ label: string, value: any, border: boolean }} props */
function CardPart({ label, value, border }) {
  return (
    <div
      className={border ? "border-bottom" : ""}
      style={{ paddingBottom: "0.3rem", marginTop: "0.3rem" }}
    >
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
 * deleteHandler: Function
 * }} props */
export default function Card({ title, parts, deleteHandler }) {
  return (
    <div className="card m-3">
      <div className="card-header px-2 d-flex align-items-center text-center">
        <div
          className="rounded-circle img-thumbnail mx-2"
          style={{
            backgroundColor: "#000",
            width: "2.2rem",
            height: "2.2rem",
            minWidth: "2.2rem",
          }}
        ></div>
        <div className="pl-4 fw-semibold" style={{ width: "100%" }}>
          {title}
        </div>
        {deleteHandler && (
          <i
            className="btn bi bi-trash-fill align-self-middle ms-auto"
            style={{ color: "#6ea8fe" }}
            onClick={deleteHandler}
          ></i>
        )}
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
