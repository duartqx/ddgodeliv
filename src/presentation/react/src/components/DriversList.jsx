import React, { useEffect, useState } from "react";
import DriverCard from "./DriverCard";
import SideBarListings from "./SideBarListings";
import CreateNewButton from "./CreateNewButton";
import { companyDrivers } from "../services/driver/driver";

export default function DriversList() {
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([])
  );
  const [showForm, setShowForm] = useState(false);

  // Create New Driver State
  const [driverName, setDriverName] = useState("");
  const [driverEmail, setDriverEmail] = useState("");
  const [driverLicense, setDriverLicense] = useState("");

  useEffect(() => {
    companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  const driversCards = drivers.map((d) => (
    <DriverCard driver={d} key={`driver__${d.id}__${d.user.email}`} />
  ));

  const handleSubmit = (/** @type {React.FormEvent} */ e) => {
    e.preventDefault();
    alert("Create Driver Submit");
  };

  return (
    <>
      <div style={{ width: "19rem", height: "100vh" }}>
        {showForm && (
          <div
            className="card mx-auto"
            style={{
              position: "absolute",
              zIndex: 2,
              bottom: "5.6rem",
              left: "5rem",
              width: "17rem",
            }}
          >
            <div
              className="card-header text-white text-center"
              style={{ backgroundColor: "#000" }}
            >
              Create New Driver
            </div>
            <div className="card-body shadow">
              <div className="p-4"></div>
              <form action="post" onSubmit={handleSubmit}>
                <div className="form-group mb-3">
                  <label className="text-body-tertiary fw-light">Name</label>
                  <input
                    type="text"
                    className="form-control"
                    placeholder="José Silva"
                    value={driverName}
                    onChange={(e) => setDriverName(e.target.value)}
                  />
                </div>

                <div className="form-group mb-3">
                  <label className="text-body-tertiary fw-light">Email</label>
                  <input
                    type="email"
                    className="form-control"
                    placeholder="example@email.com"
                    value={driverEmail}
                    onChange={(e) => setDriverEmail(e.target.value)}
                  />
                </div>

                <div className="form-group mb-3">
                  <label className="text-body-tertiary fw-light">License</label>
                  <input
                    type="text"
                    className="form-control"
                    placeholder="123-19219-12xb"
                    value={driverLicense}
                    onChange={(e) => setDriverLicense(e.target.value)}
                  />
                </div>

                <button
                  className="btn text-center col-12 p-3 mt-5"
                  style={{ backgroundColor: "#f8f0fa" }}
                >
                  Submit
                </button>
              </form>
            </div>
          </div>
        )}

        <SideBarListings listing={driversCards} />

        <CreateNewButton
          label={!showForm ? "Create New Driver" : "Cancel"}
          height={!showForm ? "6rem" : "90vh"}
          onClickHandler={() => setShowForm(!showForm)}
        />
      </div>
    </>
  );
}
