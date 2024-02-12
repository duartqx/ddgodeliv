import React, { useState } from "react";
import CardForm from "./CardForm";
import { createDriver } from "../services/driver/driver";
import CreateNewButton from "./CreateNewButton";

/** @param {{ appendDriver: Function }} props */
export default function DriverCardForm({ appendDriver }) {
  const [showForm, setShowForm] = useState(false);
  const [driverName, setDriverName] = useState("");
  const [driverEmail, setDriverEmail] = useState("");
  const [driverLicense, setDriverLicense] = useState("");
  const [error, setError] = useState("");
  const [createdAlert, setCreatedAlert] = useState("");

  const setCreatedAlertTimeout = (createdAlert) => {
    setCreatedAlert(createdAlert);
    setTimeout(() => setCreatedAlert(""), 5000);
  };

  const handleSubmit = async (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    const driver = await createDriver({
      name: driverName,
      email: driverEmail,
      license: driverLicense,
    });

    if (driver?.id) {
      appendDriver(driver);
      setShowForm(false);
      setCreatedAlertTimeout(`${driver.user.email} created`);
    } else {
      setError(`Could not create driver ${driverEmail}`);
      setTimeout(() => setError(""), 5000);
    }
  };

  return (
    <>
      <div
        style={{
          zIndex: 20,
          position: "absolute",
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
        <CreateNewButton
          label={!showForm ? "Create New Driver" : "Cancel"}
          height={!showForm ? "6rem" : "100vh"}
          onClickHandler={() => setShowForm(!showForm)}
        />
        {showForm && (
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
        )}
      </div>
    </>
  );
}
