import React, { useState } from "react";
import DriversList from "../components/DriversList";
import VehiclesList from "../components/VehiclesList";
import NavSideBar from "../components/NavSideBar";

export default function Dashboard() {
  const [sidebar, setSidebar] = useState(false);

  return (
    <main className="d-flex">
      <NavSideBar
        onClickDrivers={() => setSidebar(false)}
        onClickVehicles={() => setSidebar(true)}
      />
      {!sidebar && <DriversList />}
      {sidebar && <VehiclesList />}
    </main>
  );
}
