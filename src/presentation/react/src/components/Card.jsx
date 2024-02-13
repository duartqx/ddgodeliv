import React from "react";

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
          backgroundColor: selected && "#f8f0fa",
        }}
      >
        <img
          src="https://images.assetsdelivery.com/compings_v2/tanyadanuta/tanyadanuta1910/tanyadanuta191000003.jpg"
          className="rounded-circle img-thumbnail mx-2"
          style={{
            objectFit: "cover",
            width: "2.2rem",
            height: "2.2rem",
            minWidth: "2.2rem",
          }}
        />
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
