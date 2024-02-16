import React from "react";

/** @param {{ err: string }} props */
export default function Error({ err, noPadding }) {
    return (
        <div className={`${noPadding ? "" : "my-3 py-2"} alert alert-danger`}>
            {err}
        </div>
    );
}
