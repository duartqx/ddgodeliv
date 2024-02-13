import React from "react";

export default function SideBarListings({
  listing,
  createButton,
  filterValue,
  filterOnChangeHandler,
}) {
  return (
    <div
      className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary shadow"
      style={{ maxHeight: "100vh", height: "100vh" }}
    >
      <div style={{ overflowY: "auto", paddingBottom: "6rem" }}>
        {filterOnChangeHandler && (
          <div className="p-3">
            <label className="fw-light">Filter</label>
            <input
              className="form-control"
              type="text"
              value={filterValue}
              onChange={filterOnChangeHandler}
            />
          </div>
        )}
        {listing}
      </div>
      {createButton}
    </div>
  );
}
