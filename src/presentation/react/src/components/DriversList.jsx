import React, { useEffect, useState } from "react";
import DriverCard from "./DriverCard";
import SideBarListings from "./SideBarListings";
import CreateNewButton from "./CreateNewButton";
import { companyDrivers } from "../services/driver/driver";

export default function DriversList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([]),
  );

  useEffect(() => {
    companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const driversCards = drivers.map((d) => (
    <DriverCard driver={d} key={`driver__${d.id}__${d.user.email}`} />
  ));

  return (
    <SideBarListings
      listing={driversCards}
      createButton={
        <CreateNewButton label="Create New Driver" onClickHandler={() => {}} />
      }
    />
  );
}
