import React, { useState } from "react";
import CardForm from "./CardForm";

export default function DriverCardForm() {
  const [driverName, setDriverName] = useState("");
  const [driverEmail, setDriverEmail] = useState("");
  const [driverLicense, setDriverLicense] = useState("");

  const handleSubmit = (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    alert(`Create Driver Form Submit: ${driverName} / ${driverEmail} / ${driverLicense}`);
  };

  return (
    <CardForm
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
