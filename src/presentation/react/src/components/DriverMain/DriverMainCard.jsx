import React from "react";
import RoundImage from "../RoundImage";
import { Link } from "react-router-dom";
import DottedLink from "../DottedLink";

function DriverMainInnerCard({ label, value }) {
  return (
    <>
      <div className="col-6 d-flex flex-column justify-content-end">
        <div className="fw-light">{label}</div>
        <div className="fw-bold">{value}</div>
      </div>
    </>
  );
}

/** @param {{ group: { label: string, value: any }[]}} props */
function DriverMainInnerCardRow({ group }) {
  return (
    <div className="row p-2">
      {group.map((ic, i) => (
        <DriverMainInnerCard
          label={ic.label}
          value={ic.value}
          key={`drivermaininnercardrow__${i}__${
            ic.label?.replace(" ", "") || `obj__${i}`
          }`}
        />
      ))}
    </div>
  );
}

/** @param {{ driver: import("../../services/driver/driver").Driver }} props */
export default function DriverMainCard({ driver }) {
  const cardGroups = [
    [
      { label: "Vehicle:", value: "Ford E250" },
      { label: "Since:", value: "2006-02-01" },
    ],
    [
      { label: "Deliveries:", value: "33" },
      { label: "Info:", value: "Something" },
    ],
    [
      { label: "Vehicle License:", value: "badsf-1212" },
      { label: "", value: <DottedLink to={"#"} label="Documents" /> },
    ],
  ];

  return (
    <div className="card bg-body-tertiary">
      <div className="card-body row">
        <div className="col-lg-3 text-center">
          <RoundImage size="13rem" />
        </div>
        <div className="col-lg-9 p-2">
          <div className="row">
            <div className="col-12 rounded-2">
              {cardGroups.map((g) => (
                <DriverMainInnerCardRow group={g} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
