import React from "react";
import { Link } from "react-router-dom";

export default function SubmitAndLinkButton({
  submitLabel,
  linkTo,
  linkLabel,
}) {
  return (
    <div className="mt-3 d-flex col-md-12">
      <button type="submit" className="col-md-5 btn btn-primary">
        {submitLabel}
      </button>
      <Link to={linkTo} className="col-md-5 ms-auto btn btn-outline-primary">
        {linkLabel}
      </Link>
    </div>
  );
}
