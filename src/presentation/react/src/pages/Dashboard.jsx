import React, { useState } from "react";
import DriversList from "../components/DriversList";
import VehiclesList from "../components/VehiclesList";
import NavSideBar from "../components/NavSideBar";

/**
 * @typedef {{
 *  selected: "drivers"|"vehicles"
 *  isSelected: (field: string) => boolean
 * }} SideBar
 */

export default function Dashboard() {
  const [sidebar, setSidebar] = useState(
    /** @type {SideBar} */ ({
      selected: "drivers",
      isSelected: function (/** @type {string} */ field) {
        return this.selected === field;
      },
    })
  );

  return (
    <main className="d-flex">
      <NavSideBar
        sidebar={sidebar}
        onClickDrivers={() => setSidebar({ ...sidebar, selected: "drivers" })}
        onClickVehicles={() => setSidebar({ ...sidebar, selected: "vehicles" })}
      />
      {sidebar.isSelected("drivers") && <DriversList />}
      {sidebar.isSelected("vehicles") && <VehiclesList />}
    </main>
  );
}
