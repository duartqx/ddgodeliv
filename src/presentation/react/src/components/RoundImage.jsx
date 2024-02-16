import React from "react";

/** @param {{ size: string, src: string }} props */
export default function RoundImage({ size, src }) {
    return (
        <img
            src={
                src ||
                "https://images.assetsdelivery.com/compings_v2/tanyadanuta/tanyadanuta1910/tanyadanuta191000003.jpg"
            }
            className="rounded-circle img-thumbnail mx-2 align-self-center flex-shrink-1"
            style={{
                objectFit: "cover",
                width: size,
                height: size,
            }}
        />
    );
}
