import React from "react";
import DriverMainCard from "./DriverMainCard";

export default function DriverMainTabbed({ activeTab, setActiveTab, driver }) {
  return (
    <div className="col-12 d-flex flex-column flex-grow-1">
      <ul className="nav nav-tabs">
        <li
          className="nav-item"
          onClick={() => setActiveTab({ ...activeTab, tab: "inf" })}
        >
          <a
            href="#"
            className={`nav-link ${activeTab.isActive("inf") && "active"}`}
          >
            Driver
          </a>
        </li>
        <li
          className="nav-item"
          onClick={() => setActiveTab({ ...activeTab, tab: "del" })}
        >
          <a
            href="#"
            className={`nav-link ${activeTab.isActive("del") && "active"}`}
          >
            Deliveries
          </a>
        </li>
        <li
          className="nav-item"
          onClick={() => setActiveTab({ ...activeTab, tab: "sta" })}
        >
          <a
            href="#"
            className={`nav-link ${activeTab.isActive("sta") && "active"}`}
          >
            Statistics
          </a>
        </li>
      </ul>

      {activeTab.isActive("inf") && <DriverMainCard driver={driver} />}
      {activeTab.isActive("del") && (
        <div
          className="mt-4 p-4"
          style={{ backgroundColor: "black", height: "100%" }}
        ></div>
      )}
      {activeTab.isActive("sta") && (
        <div
          className="mt-4 p-4"
          style={{ backgroundColor: "pink", height: "100%" }}
        ></div>
      )}
    </div>
  );
}
