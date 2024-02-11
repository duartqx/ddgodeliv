import React, { useState } from "react";
import Logout from "../components/Logout";
import { Link } from "react-router-dom";
import { companyDrivers } from "../services/driver/driver";
import { useEffect } from "react";

export default function Dashboard() {
  const [sidebar, setSidebar] = useState(false);
  const [drivers, setDrivers] = useState(
    /** @type {import("../services/driver/driver").Driver[]} */ ([]),
  );

  useEffect(() => {
    companyDrivers().then((drivers) => setDrivers(drivers));
  }, []);

  return (
    <main className="d-flex">
      <nav
        className="d-flex flex-column flex-shrink-0"
        style={{ width: "4.5rem", height: "100vh", backgroundColor: "#000" }}
        data-bs-theme="dark"
      >
        <Link
          to="/dashboard"
          className="d-block p-3 link-body-emphasis text-center"
        >
          <i className="bi bi-house"></i>
        </Link>
        <ul className="nav nav-pills nav-flush flex-column mb-auto text-center">
          <li className="nav-link">
            <Link to="/dashboard">
              <i className="bi bi-bell"></i>
            </Link>
          </li>

          <li className="nav-link">
            <Link to="#" onClick={() => setSidebar(true)}>
              <i className="bi bi-truck"></i>
            </Link>
          </li>

          <li className="nav-link">
            <Link to="#" onClick={() => setSidebar(false)}>
              <i className="bi bi-person"></i>
            </Link>
          </li>
        </ul>
        <div className="d-flex align-items-center justify-content-center p-3">
          <Logout />
        </div>
      </nav>

      {!sidebar && (
        <div
          className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary shadow"
          style={{ width: "19rem", maxHeight: "100vh", position: "relative" }}
        >
          <div style={{ overflowY: "auto", paddingBottom: "6rem" }}>
            {drivers.map((d) => {
              return (
                <div
                  className="card m-3"
                  key={`driver__${d.id}__${d.user.email}`}
                >
                  <div className="card-header px-2 d-flex align-items-center text-center">
                    <div
                      className="rounded-circle img-thumbnail"
                      style={{
                        backgroundColor: "#000",
                        width: "2.2rem",
                        height: "2.2rem",
                        minWidth: "2.2rem",
                      }}
                    ></div>
                    <strong className="text-center mx-auto">
                      {d.user.name}
                    </strong>
                  </div>
                  <div className="card-body px-3">
                    <div>
                      <div
                        className="border-bottom"
                        style={{ paddingBottom: "0.3rem" }}
                      >
                        {d.user.email}
                      </div>
                      <div className="" style={{ marginTop: "0.3rem" }}>
                        {d.license_id}
                      </div>
                    </div>
                  </div>
                </div>
              );
            })}
            <div
              className="p-3 d-flex align-items-center"
              style={{
                background: "linear-gradient(transparent 0%, #f8f0fa 60%)",
                position: "absolute",
                bottom: 0,
                width: "100%",
                height: "7rem",
              }}
            >
              <button
                className="btn mx-auto align-self-end text-white p-3 drop-shadow"
                style={{ backgroundColor: "#000", width: "90%" }}
              >
                Create New Driver
              </button>
            </div>
          </div>
        </div>
      )}
      {sidebar && (
        <div
          className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary shadow"
          style={{ width: "19rem" }}
        >
          Vehicles
        </div>
      )}
    </main>
  );
}
