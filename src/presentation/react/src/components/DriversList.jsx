import React, { useEffect, useState } from "react";
import DriverCard from "./DriverCard";
import SideBarListings from "./SideBarListings";
import CreateNewButton from "./CreateNewButton";
import * as driverService from "../services/driver/driver";
import DriverCardForm from "./DriverCardForm";

export default function DriversList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([])
  );
  const [showForm, setShowForm] = useState(false);

  useEffect(() => {
    driverService.companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const deleteClickHandler = async (/** @type {number} */ id) => {
    if (await driverService.deleteDriver({ id })) {
      setDrivers(drivers.filter((d) => d.id !== id));
    }
  };

  const driversCards = drivers.map((d) => (
    <DriverCard
      deleteHandler={() => deleteClickHandler(d.id)}
      driver={d}
      key={`driver__${d.id}__${d.user.email}`}
    />
  ));

  return (
    <>
      <div style={{ width: "19rem", height: "100vh" }}>
        {showForm && (
          <DriverCardForm
            closeForm={() => setShowForm(false)}
            appendDriver={(driver) => setDrivers(drivers.concat(driver))}
          />
        )}

        <SideBarListings listing={driversCards} />

        <CreateNewButton
          label={!showForm ? "Create New Driver" : "Cancel"}
          height={!showForm ? "6rem" : "90vh"}
          onClickHandler={() => setShowForm(!showForm)}
        />
      </div>
    </>
  );
}
