import React, { useState } from "react";
import CardForm from "./CardForm";
import { createDriver } from "../services/driver/driver";

/** @param {{ appendDriver: Function, closeForm: Function }} props */
export default function DriverCardForm({ appendDriver, closeForm }) {
  const [driverName, setDriverName] = useState("");
  const [driverEmail, setDriverEmail] = useState("");
  const [driverLicense, setDriverLicense] = useState("");
  const [error, setError] = useState("")

  const handleSubmit = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    const driver = await createDriver({
      name: driverName,
      email: driverEmail,
      license: driverLicense,
    });

    if (driver?.id) {
      appendDriver(driver)
      closeForm()
    } else {
      setError(`Could not create driver ${driverEmail}`)
      setTimeout(() => setError(""), 5000)
    }
  };

  return (
    <CardForm
      error={error}
      handleSubmit={handleSubmit}
      title="Create New Driver"
      inputs={[
        {
          label: "Name",
          type: "text",
          placeholder: "José Silva",
          value: driverName,
          onChangeHandler: (e) => setDriverName(e.target.value),
        },
        {
          label: "Email",
          type: "email",
          placeholder: "example@email.com",
          value: driverEmail,
          onChangeHandler: (e) => setDriverEmail(e.target.value),
        },
        {
          label: "License",
          type: "text",
          placeholder: "123-19219-12xb",
          value: driverLicense,
          onChangeHandler: (e) => setDriverLicense(e.target.value),
        },
      ]}
    />
  );
}
