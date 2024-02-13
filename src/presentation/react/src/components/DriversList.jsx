import React, { useEffect, useState } from "react";
import DriverCard from "./DriverCard";
import SideBarListings from "./SideBarListings";
import * as driverService from "../services/driver/driver";
import DriverCardForm from "./DriverCardForm";
import DriverMain from "./DriverMain";

export default function DriversList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([])
  );
  const [filterDriver, setFilterDriver] = useState("");
  const [selectedDriver, setSelectedDriver] = useState(0);

  useEffect(() => {
    driverService.companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const deleteClickHandler = async (/** @type {number} */ id) => {
    if (await driverService.deleteDriver({ id })) {
      setDrivers(drivers.filter((d) => d.id !== id));
    }
  };

  const filteredDrivers = filterDriver
    ? drivers.filter((d) => d.user.name.includes(filterDriver))
    : drivers;

  const driversCards = filteredDrivers.map((d, idx) => (
    <DriverCard
      onClickHandler={() => setSelectedDriver(idx)}
      selected={selectedDriver === idx}
      driver={d}
      key={`driver__${d.id}__${d.user.email}`}
    />
  ));

  return (
    <>
      <div className="d-flex" style={{ width: "76.5rem" }}>
        <div style={{ minWidth: "19rem", height: "100vh" }}>
          <DriverCardForm
            appendDriver={(driver) => setDrivers(drivers.concat(driver))}
          />
          <SideBarListings
            listing={driversCards}
            filterValue={filterDriver}
            filterOnChangeHandler={(e) => setFilterDriver(e.target.value)}
          />
        </div>
        {filteredDrivers[selectedDriver] && (
          <DriverMain driver={filteredDrivers[selectedDriver]} />
        )}
      </div>
    </>
  );
}
