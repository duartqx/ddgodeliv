import React, { useEffect, useState } from "react";
import VehicleCard from "./VehicleCard";
import SideBarListings from "./SideBarListings";
import CreateNewButton from "./CreateNewButton";
import { companyVehicles } from "../services/vehicles/vehicles";

export default function VehiclesList() {
  const [vehicles, setVehicles] = useState(
    /** @type {import("../services/vehicles/vehicles").Vehicle[]} */ ([]),
  );

  useEffect(() => {
    companyVehicles().then((vehicles) => setVehicles(vehicles));
  }, []);

  const vehicleCards = vehicles.map((d) => (
    <VehicleCard vehicle={d} key={`vehicle__${d.id}__${d.model.name}`} />
  ));

  return (
    <SideBarListings
      listing={vehicleCards}
      createButton={
        <CreateNewButton label="Create New Vehicle" onClickHandler={() => {}} />
      }
    />
  );
}
