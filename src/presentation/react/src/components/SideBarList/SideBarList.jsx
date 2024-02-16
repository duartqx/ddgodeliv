import React from "react";
import "./SideBarList.css";

export default function SideBarList({
    listing,
    createButton,
    filterValue,
    filterOnChangeHandler,
}) {
    return (
        <>
            <div
                className="d-flex flex-column flex-shrink-0 border-right"
                style={{
                    maxHeight: "100vh",
                    height: "100vh",
                    position: "relative",
                    borderRight: "solid 1px #f0f2f7"
                }}
            >
                <div className="sidebarlisting__scrollbar">
                    {filterOnChangeHandler && (
                        <div
                            className="p-3"
                            style={{
                                position: "absolute",
                                height: "7vh",
                                width: "19rem",
                                background:
                                    "linear-gradient(white 40%, transparent 100%)",
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
                {createButton}
            </div>
        </>
    );
}
