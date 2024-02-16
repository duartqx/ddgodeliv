import React from "react";
import { Link } from "react-router-dom";

export default function DottedLink({ to, label }) {
  return (
    <Link
      to={to}
      style={{
        color: "#9e9fa1",
        textDecoration: "underline dotted #9e9fa1",
      }}
    >
      {label}
    </Link>
  );
}
