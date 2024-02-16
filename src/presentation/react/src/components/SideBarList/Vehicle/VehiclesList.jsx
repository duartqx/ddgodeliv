import React, { useEffect, useState } from "react";
import VehicleCard from "./VehicleCard";
import SideBarList from "../SideBarList";
import * as vehicleService from "../../../services/vehicles/vehicles";
import VehicleCardForm from "../../CardForm/Vehicle/VehicleCardForm";

export default function VehiclesList() {
  const [vehicles, setVehicles] = useState(
    /** @type {import("../../../services/vehicles/vehicles").Vehicle[]} */ ([]),
  );

  useEffect(() => {
    vehicleService.companyVehicles().then((vehicles) => setVehicles(vehicles));
  }, []);

  const deleteClickHandler = async (/** @type {number} */ id) => {
    if (await vehicleService.deleteVehicle({ id })) {
      setVehicles(vehicles.filter((v) => v.id !== id));
    }
  };

  const vehicleCards = vehicles.map((v) => (
    <VehicleCard
      vehicle={v}
      key={`vehicle__${v.id}__${v.model.name}`}
      deleteHandler={() => deleteClickHandler(v.id)}
    />
  ));

  return (
    <>
      <VehicleCardForm
        appendVehicle={(vehicle) => setVehicles(vehicles.concat(vehicle))}
      />
      <SideBarList listing={vehicleCards} />
    </>
  );
}
