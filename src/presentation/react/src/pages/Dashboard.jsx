import React, { useState } from "react";
import Logout from "../components/Logout";
import { Link } from "react-router-dom";

export default function Dashboard() {
  const [sidebar, setSidebar] = useState(false);

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

      {sidebar && (
        <div
          className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary"
          style={{ width: "18rem" }}
        >
          Drivers
        </div>
      )}
      {!sidebar && (
        <div
          className="d-flex flex-column flex-shrink-0 border-right bg-body-tertiary"
          style={{ width: "18rem" }}
        >
          Vehicles
        </div>
      )}
    </main>
  );
}
