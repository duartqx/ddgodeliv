import React, { useState } from "react";
import DriverMainDriverInfo from "./DriverMainDriverInfo";
import DriverMainDeliveries from "./DriverMainDeliveries";
import DriverMainStatistics from "./DriverMainStatistics";

function DriverMainTab({ setActive, isActive, label }) {
  return (
    <li className="nav-item" onClick={setActive}>
      <a href="#" className={`nav-link ${isActive && "active"}`}>
        {label}
      </a>
    </li>
  );
}

export default function DriverMainTabbed({ driver }) {
  const [activeTab, setActiveTab] = useState(
    /** @type {{ tab: "inf"|"del"|"sta", isActive: (tab: string) => boolean }} */ ({
      tab: "inf",
      isActive: function (tab) {
        return this.tab === tab;
      },
    })
  );

  return (
    <div
      className="d-flex flex-column flex-grow-1"
      style={{ maxHeight: "87vh" }}
    >
      <ul className="nav nav-tabs flex-nowrap">
        <DriverMainTab
          setActive={() => setActiveTab({ ...activeTab, tab: "inf" })}
          isActive={activeTab.isActive("inf")}
          label="Driver"
        />
        <DriverMainTab
          setActive={() => setActiveTab({ ...activeTab, tab: "del" })}
          isActive={activeTab.isActive("del")}
          label="Deliveries"
        />
        <DriverMainTab
          setActive={() => setActiveTab({ ...activeTab, tab: "sta" })}
          isActive={activeTab.isActive("sta")}
          label="Statistics"
        />
      </ul>

      {activeTab.isActive("inf") && <DriverMainDriverInfo driver={driver} />}
      {activeTab.isActive("del") && <DriverMainDeliveries driver={driver} />}
      {activeTab.isActive("sta") && <DriverMainStatistics driver={driver} />}
    </div>
  );
}
