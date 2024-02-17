import React from "react";
import RoundImage from "../RoundImage";
import DottedLink from "../DottedLink";

function DriverMainInnerCard({ label, value }) {
  return (
    <>
      <div className="col-6 d-flex flex-column justify-content-end">
        <div className="fw-light px-4">{label}</div>
        <div className="flex-shrink-1">
          <div className="fw-bold px-4">{value}</div>
        </div>
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
          key={`drivermaininnercard__${i}__${
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
    <div className="card bg-body-tertiary flex-grow-1 p-2 my-4">
      <div className="row">
        <div className="col-xxl-3 col-xl-4 d-flex align-items-center justify-content-center p-4">
          <RoundImage src="" size="13rem" />
        </div>
        <div
          className="col-xxl-9 col-xl-8 p-2 d-flex flex-column justify-content-center"
          style={{ height: "80%" }}
        >
          <div className="col-12">
            {cardGroups.map((g, i) => (
              <DriverMainInnerCardRow
                group={g}
                key={`drivermaininnercardrow__${i}__${
                  g.label?.replace(" ", "") || `obj__${i}`
                }`}
              />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
