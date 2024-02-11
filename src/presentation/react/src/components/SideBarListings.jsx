import React from "react";

export default function SideBarListings({ listing, createButton }) {
  return (
    <div
      className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary shadow"
      style={{ maxHeight: "100vh", height: "100vh" }}
    >
      <div style={{ overflowY: "auto", paddingBottom: "6rem" }}>{listing}</div>
      {createButton}
    </div>
  );
}
