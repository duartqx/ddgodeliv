import React from "react";
import "./SideBarList.css";

export default function SideBarList({
  listing,
  filterValue,
  filterOnChangeHandler,
}) {
  return (
    <>
      <div
        className="d-flex flex-column flex-shrink-0 border-right"
        style={{
          maxHeight: "calc(100vh - 4rem)",
          height: "calc(100vh - 4rem)",
          width: "22rem",
          position: "relative",
          borderRight: "solid 1px #f0f2f7",
        }}
      >
        <div className="sidebarlisting__scrollbar">
          {filterOnChangeHandler && (
            <div
              className="p-3"
              style={{
                position: "absolute",
                height: "7vh",
                width: "22rem",
                zIndex: 2,
                background: "linear-gradient(white 40%, transparent 100%)",
              }}
            >
              <input
                className="form-control"
                type="text"
                placeholder="Search"
                value={filterValue}
                onChange={filterOnChangeHandler}
              />
            </div>
          )}
          <div style={{ paddingTop: "4rem" }}>{listing}</div>
        </div>
      </div>
    </>
  );
}
