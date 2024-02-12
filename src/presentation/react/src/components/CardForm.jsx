import React from "react";
import CardFormInput from "./CardFormInput";
import Error from "./Error"

/** @param {{
 *  title: string
 *  error: string
 *  handleSubmit: Function
 *  inputs: import("./CardFormInput").CardFormInputObject[]
 * }} props
 */
export default function CardForm({ title, error, handleSubmit, inputs }) {
  return (
    <div
      className="card mx-auto"
      style={{
        position: "absolute",
        zIndex: 2,
        bottom: "5.6rem",
        left: "5rem",
        width: "17rem",
      }}
    >
      <div
        className="card-header text-white text-center p-3 fw-semibold"
        style={{ backgroundColor: "#000" }}
      >
        {title}
      </div>
      <div className="card-body shadow">
        <div className="py-4">{error && <Error err={error} noPadding={true} />}</div>
        <form action="post" onSubmit={handleSubmit}>

          {inputs.map((props) => <CardFormInput {...props} />)}

          <button
            className="btn text-center col-12 p-3 mt-5"
            style={{ backgroundColor: "#f8f0fa" }}
          >
            Submit
          </button>
        </form>
      </div>
    </div>
  );
}
