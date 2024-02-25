import React, { useEffect, useState } from "react";
import DriverMainHeaderButton from "../MainHeaderButton";
import BlackButton from "../CardForm/BlackButton";
import AssignToDriverFormWithBackdrop from "./AssignToDriverFormWithBackdrop";

/**
 * @param {{
 *  delivery: import("../../services/deliveries/deliveries").Delivery
 *  handleFilterOutDelivery: (delivery: import("../../services/deliveries/deliveries").Delivery) => void
 * }} props
 * */
export default function DeliveryMainHeader({
  delivery,
  handleFilterOutDelivery,
}) {
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
                handleFilterOutDelivery={handleFilterOutDelivery}
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
