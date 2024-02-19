import React from "react";
import { Link } from "react-router-dom";

export default function DottedLink({ to, label, color="#9e9fa1" }) {
  return (
    <Link
      to={to}
      style={{
        color: color,
        textDecoration: `underline dotted ${color}`,
      }}
    >
      {label}
    </Link>
  );
}
