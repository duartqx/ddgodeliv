import React from "react";
import { Link, useLocation } from "react-router-dom";
import Logout from "./Logout";
import { Paths as P } from "../../path";

function NavBarLi({ path, atPath, icon }) {
  return (
    <Link to={path || "#"}>
      <li
        className="nav-link py-3"
        style={{ backgroundColor: atPath && "#2388fd" }}
      >
        <i className={`bi bi-${icon} text-white`}></i>
      </li>
    </Link>
  );
}

function NavBarLiGroup({ group }) {
  return (
    <>
      {group.map((p) => (
        <div
          key={`${p.path?.replaceAll("/", "")}__${P.icon?.replaceAll(" ", "")}}`}
        >
          <NavBarLi path={p.path} atPath={p.atPath} icon={p.icon} />
        </div>
      ))}
    </>
  );
}

export default function NavBar() {
  const location = useLocation();

  const at = (/** @type {string} */ path) => location.pathname === path;

  return (
    <>
      <nav
        className="d-flex flex-column flex-shrink-0"
        style={{
          width: "4.5rem",
          height: "100vh",
          backgroundColor: "#000",
        }}
        data-bs-theme="dark"
      >
        <ul className="nav nav-pills nav-flush flex-column mb-auto text-center pt-2">
          <NavBarLiGroup
            group={[
              {
                path: P.root,
                atPath: at(P.root),
                icon: "house",
              },
              { path: "#", atPath: false, icon: "bell" },
              {
                path: P.drivers,
                atPath: at(P.drivers),
                icon: "person",
              },
              {
                path: P.vehicles,
                atPath: at(P.vehicles),
                icon: "truck",
              },
              {
                path: P.delivery.available,
                atPath: at(P.delivery.available),
                icon: "hourglass",
              },
              {
                path: P.delivery.company,
                atPath: at(P.delivery.company),
                icon: "tags-fill",
              },
            ]}
          />
        </ul>

        <hr className="mx-2 text-white" />

        <div className="d-flex align-items-center justify-content-center p-3">
          <Logout />
        </div>
      </nav>
    </>
  );
}
