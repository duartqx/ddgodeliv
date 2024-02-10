import React from "react";

/** @param {{ err: string }} props */
export default function Error({ err }) {
  return <div className="my-3 py-2 alert alert-danger">{err}</div>;
}
