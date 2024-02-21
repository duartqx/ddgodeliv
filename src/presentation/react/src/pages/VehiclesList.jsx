import React, { useEffect, useState } from "react";
import VehicleCard from "../components/VehicleCard";
import SideBarList from "../components/SideBarList/SideBarList";
import * as vehicleService from "../services/vehicles/vehicles";
import VehicleCardForm from "../components/CardForm/Vehicle/VehicleCardForm";

export default function VehiclesList() {
  const [vehicles, setVehicles] = useState(
    /** @type {import("../services/vehicles/vehicles").Vehicle[]} */ ([])
  );
  const [filterVehicle, setFilterVehicle] = useState("");
  const [selectedVehicle, setSelectedVehicle] = useState(0);

  useEffect(() => {
    vehicleService.companyVehicles().then((vehicles) => setVehicles(vehicles));
  }, []);

  const deleteClickHandler = async (/** @type {number} */ id) => {
    if (await vehicleService.deleteVehicle({ id })) {
      setVehicles(vehicles.filter((v) => v.id !== id));
    }
  };

  const filteredVehicles = filterVehicle
    ? vehicles.filter((v) =>
        v.model.name.toLowerCase().includes(filterVehicle.toLowerCase())
      )
    : vehicles;

  const vehicleCards = filteredVehicles.map((v, idx) => (
    <VehicleCard
      vehicle={v}
      selected={selectedVehicle === idx}
      key={`vehicle__${v.id}__${v.model.name}`}
      onClickHandler={() => setSelectedVehicle(idx)}
      deleteHandler={() => deleteClickHandler(v.id)}
    />
  ));

  return (
    <>
      <VehicleCardForm
        appendVehicle={(vehicle) => setVehicles(vehicles.concat(vehicle))}
      />
      <SideBarList
        listing={vehicleCards}
        filterValue={filterVehicle}
        filterOnChangeHandler={(e) => setFilterVehicle(e.target.value)}
      />
    </>
  );
}
