import React, { useEffect, useState } from "react";
import CardForm from "../CardForm";
import CardFormCreateButton from "../CardFormCreateButton";
import { getVehicleModels } from "../../../services/vehicles/vehicleModels";
import { createVehicle } from "../../../services/vehicles/vehicles";

/** @param {{ appendVehicle: Function }} props */
export default function VehicleCardForm({ appendVehicle }) {
  const [showForm, setShowForm] = useState(false);
  const [vehicleModelId, setVehicleModelId] = useState(0);
  const [vehicleLicenseId, setVehicleLicenseId] = useState("");
  const [error, setError] = useState("");
  const [createdAlert, setCreatedAlert] = useState("");
  const [vehicleModels, setVehicleModels] = useState(
    /** @type {import("../../../services/vehicles/vehicleModels").VehicleModel[]} */ ([]),
  );

  useEffect(() => {
    getVehicleModels().then((models) => setVehicleModels(models));
  }, []);

  const setCreatedAlertTimeout = (createdAlert) => {
    setCreatedAlert(createdAlert);
    setTimeout(() => setCreatedAlert(""), 5000);
  };

  const handleSubmit = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    const vehicle = await createVehicle({
      license: vehicleLicenseId,
      model: Number(vehicleModelId),
    });

    if (vehicle?.id) {
      appendVehicle(vehicle);
      setShowForm(false);
      setCreatedAlertTimeout(`${vehicle.model.name} created`);
    } else {
      setError(`Could not create vehicle`);
      setTimeout(() => setError(""), 5000);
    }
  };

  return (
    <>
      <div
        style={{
          zIndex: 20,
          position: "relative",
          bottom: 0,
        }}
      >
        {createdAlert && (
          <div
            className="alert alert-success text-center"
            style={{
              width: "16rem",
              left: "1rem",
              top: "2.2rem",
              padding: "0.8rem 0.8rem",
            }}
          >
            {createdAlert}
          </div>
        )}
        <CardFormCreateButton
          label={!showForm ? "Create New Vehicle" : "Cancel"}
          height={!showForm ? "6rem" : "100vh"}
          onClickHandler={() => setShowForm(!showForm)}
        />
        {showForm && (
          <CardForm
            error={error}
            handleSubmit={handleSubmit}
            title="Create New Vehicle"
            inputs={[
              {
                label: "License",
                type: "text",
                placeholder: "123-19219-12xb",
                value: vehicleLicenseId,
                onChangeHandler: (e) => setVehicleLicenseId(e.target.value),
              },
              {
                label: "Model",
                type: "select",
                options: vehicleModels.map((m) => ({
                  value: m.id,
                  label: `${m.manufacturer} ${m.name} ${m.year}`,
                })),
                onChangeHandler: (e) => setVehicleModelId(e.target.value),
              },
            ]}
          />
        )}
      </div>
    </>
  );
}
