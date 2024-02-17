import React, { useEffect, useState } from "react";
import DriverMainHeader from "./Header/DriverMainHeader";
import DriverMainCard from "./DriverMainCard";
import DriverMainTabbed from "./DriverMainTabbed";
import useWidthHeight from "../../middlewares/useWidthHeight";

/**
 * @param {{
 * driver: import("../../services/driver/driver").Driver
 * deleteHandler: (id: number) => void
 * }}
 * props */
export default function DriverMain({ driver, deleteHandler }) {
  const { windowWidth } = useWidthHeight();
  const [activeTab, setActiveTab] = useState(
    /** @type {{ tab: "inf"|"del"|"sta", isActive: (tab: string) => boolean }} */ ({
      tab: "inf",
      isActive: function (tab) {
        return this.tab === tab;
      },
    })
  );

  if (driver && windowWidth <= 1151) {
    return (
      <>
        <div
          className="d-flex flex-column mx-4 flex-grow-1"
        >
          <DriverMainHeader driver={driver} deleteHandler={deleteHandler} />
          <DriverMainTabbed
            driver={driver}
            activeTab={activeTab}
            setActiveTab={setActiveTab}
          />
        </div>
      </>
    );
  }

  return (
    driver && (
      <>
        <div
          className="d-flex flex-column mx-4 flex-grow-1"
        >
          <div className="px-5">
            <DriverMainHeader driver={driver} deleteHandler={deleteHandler} />
            <DriverMainCard driver={driver} />
          </div>
        </div>
      </>
    )
  );
}
