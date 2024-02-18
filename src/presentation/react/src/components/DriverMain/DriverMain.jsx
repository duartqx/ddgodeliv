import React, { useState } from "react";
import DriverMainHeader from "./Header/DriverMainHeader";
import DriverMainTabbed from "./DriverMainTabbed";
import useWidthHeight from "../../middlewares/useWidthHeight";
import DriverMainCarded from "./DriverMainCarded";

/**
 * @param {{
 * driver: import("../../services/driver/driver").Driver
 * deleteHandler: (id: number) => void
 * }}
 * props */
export default function DriverMain({ driver, deleteHandler }) {
  const { windowWidth } = useWidthHeight();

  return (
    driver && (
      <>
        <div
          className="d-flex flex-column mx-4 flex-grow-1"
          style={{ overflow: "hidden", maxHeight: "100vh" }}
        >
          <DriverMainHeader driver={driver} deleteHandler={deleteHandler} />
          {windowWidth <= 1200 ? (
            <DriverMainTabbed driver={driver} />
          ) : (
            <DriverMainCarded driver={driver} />
          )}
        </div>
      </>
    )
  );
}
