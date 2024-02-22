import React from "react";
import CardFormInput from "./CardFormInput";
import CardFormSelect from "./CardFormSelect";
import Error from "../Error";

/**
 * @param {{
 *  title: string
 *  error: string
 *  handleSubmit: () => void
 *  inputs: import("./CardFormInput").CardFormInputObject[]
 *  handleDismiss: () => void
 * }} props
 */
export default function CardForm({
  title,
  error,
  handleSubmit,
  inputs,
  handleDismiss = () => {},
}) {
  return (
    <div
      className="card mx-auto"
      style={{
        position: "absolute",
        zIndex: 2,
        bottom: "5.6rem",
        left: "0.5rem",
        width: "21rem",
      }}
    >
      <div
        className="card-header text-white text-center p-3 fw-semibold"
        style={{ backgroundColor: "#000" }}
      >
        {title}
      </div>
      <div className="card-body shadow">
        <div className="py-4">
          {error && <Error err={error} noPadding={true} />}
        </div>
        <form action="post" onSubmit={handleSubmit}>
          {inputs.map((props) => {
            if (props.type === "select") {
              return (
                <CardFormSelect
                  {...props}
                  key={`${props.label.replaceAll(" ", "")}__${props.type}`}
                />
              );
            }
            return (
              <CardFormInput
                {...props}
                key={`${props.label.replaceAll(" ", "")}__${props.type}`}
              />
            );
          })}

          <button
            className="btn text-center col-12 p-3 mt-5"
            style={{ backgroundColor: "#f0f2f7" }}
          >
            Submit
          </button>
        </form>
      </div>
    </div>
  );
}
