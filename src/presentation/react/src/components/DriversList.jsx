import React, { useEffect, useState } from "react";
import DriverCard from "./DriverCard";
import SideBarListings from "./SideBarListings";
import CreateNewButton from "./CreateNewButton";
import { companyDrivers } from "../services/driver/driver";
import DriverCardForm from "./DriverCardForm";

export default function DriversList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([])
  );
  const [showForm, setShowForm] = useState(false);

  useEffect(() => {
    companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const driversCards = drivers.map((d) => (
    <DriverCard driver={d} key={`driver__${d.id}__${d.user.email}`} />
  ));

  return (
    <>
      <div style={{ width: "19rem", height: "100vh" }}>
        {showForm && <DriverCardForm />}

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
