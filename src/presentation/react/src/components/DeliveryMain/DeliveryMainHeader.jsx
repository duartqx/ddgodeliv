import React, { useEffect, useState } from "react";
import DriverMainHeaderButton from "../MainHeaderButton";
import BlackButton from "../CardForm/BlackButton";
import * as driverService from "../../services/driver/driver";
import CardForm from "../CardForm/CardForm";

/**
 * @param {{
 *  delivery: import("../../services/deliveries/deliveries").Delivery
 * }} props
 */
function AssignToDriverForm({ delivery }) {
  const [driverId, setDriverId] = useState(null);
  const [drivers, setDrivers] = useState(
    /** @type {driverService.Driver[]} */ ([])
  );

  useEffect(() => {
    driverService.companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const handleSubmit = (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();

    alert(`assign ${delivery.id} ${driverId}`);
  };

  return (
    <>
      <CardForm
        handleSubmit={handleSubmit}
        title="Assign to Driver"
        inputs={[
          {
            label: "Driver",
            type: "select",
            options: drivers.map((d) => ({
              value: d.id,
              label: `${d.user.email} ${d.user.name}`,
            })),
            onChangeHandler: (e) => setDriverId(e.target.value),
          },
        ]}
      />
    </>
  );
}

function AssignToDriverFormWithBackdrop({ delivery, handleBackdropClick }) {
  return (
    <div
      style={{
        position: "absolute",
        top: "55vh",
        right: "50vw",
      }}
    >
      <AssignToDriverForm delivery={delivery} />
      <div
        onClick={handleBackdropClick}
        style={{
          top: 0,
          left: 0,
          position: "fixed",
          background: `linear-gradient(
                      rgba(255, 255, 255, 0.75) 0%,
                      rgba(255, 255, 255, 0.98) 25%,
                      rgba(255, 255, 255, 0.98) 75%,
                      rgba(255, 255, 255, 0.75) 100%
                    )`,
          height: "100vh",
          width: "100vw",
        }}
      ></div>
    </div>
  );
}

/**
 * @param {{
 *  delivery: import("../../services/deliveries/deliveries").Delivery
 * }} props
 * */
export default function DriverMainHeader({ delivery }) {
  const [showForm, setShowForm] = useState(false);

  const handleSetShowForm = () => setShowForm(!showForm);

  return (
    <>
      <div className="my-4 d-flex" style={{ height: "8vh" }}>
        <div className="d-flex flex-column justify-content-center">
          <div className="fw-bold">{delivery.sender.name}</div>
          <div className="fw-light">{delivery.sender.email}</div>
        </div>
        <div className="ms-auto d-flex">
          <div
            className="align-self-center mx-3"
            style={{ position: "relative" }}
          >
            <BlackButton
              onClickHandler={handleSetShowForm}
              label="Assign Driver"
            />
            {showForm && (
              <AssignToDriverFormWithBackdrop
                delivery={delivery}
                handleBackdropClick={handleSetShowForm}
              />
            )}
          </div>
          <DriverMainHeaderButton
            icon="bi-chat-left-text-fill"
            onClickHandler={() => alert("button chat")}
          />

          <DriverMainHeaderButton
            icon="bi-telephone-fill"
            onClickHandler={() => alert("button phone")}
          />
        </div>
      </div>
    </>
  );
}
