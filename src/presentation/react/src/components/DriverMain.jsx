import React from "react";
import { DeliveryStatus } from "../domains/deliveries/status";

/** @param {{ driver: import("../services/driver/driver").Driver }} props */
export default function DriverMain({ driver }) {
  return (
    driver && (
      <>
        <div className="d-flex flex-column flex-grow-1">
          <div className="p-3">
            <div className="my-4">
              <div className="d-flex">
                <div className="d-flex flex-column">
                  <div className="fw-bold">{driver.user.name}</div>
                  <div className="fw-light">ID: {driver.license_id}</div>
                </div>
                <div className="ms-auto d-flex">
                  <button className="btn" onClick={() => alert("button chat")}>
                    <i className="bi bi-chat-left-text-fill"></i>
                  </button>
                  <button className="btn" onClick={() => alert("button phone")}>
                    <i className="bi bi-telephone-fill"></i>
                  </button>
                  <button className="btn" onClick={() => alert("button options")}>
                    <i className="bi bi-three-dots-vertical"></i>
                  </button>
                </div>
              </div>
            </div>
            <div className="card bg-body-tertiary">
              <div className="card-body">
                <div className="d-flex">
                  <img
                    src="https://images.assetsdelivery.com/compings_v2/tanyadanuta/tanyadanuta1910/tanyadanuta191000003.jpg"
                    className="rounded-circle img-thumbnail mx-2 align-self-center flex-shrink-1"
                    style={{
                      objectFit: "cover",
                      width: "8rem",
                      height: "8rem",
                    }}
                  />
                  <div className="p-5">
                    <div>{driver.user.email}</div>
                    <div>{DeliveryStatus[driver.status]}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </>
    )
  );
}
