import React from "react";

/**
 * @typedef {{
 *  label: string
 *  onChangeHandler: Function
 *  options: { value: string | number, label: string}
 *  required: boolean
 * }} CardFormSelectObject
 */

/** @param {CardFormSelectObject} props */
export default function CardFormSelect({
  label,
  onChangeHandler,
  options,
  required = true,
}) {
  return (
    <div className="form-group mb-3">
      <label className="text-body-tertiary fw-light">{label}</label>
      <select
        onChange={onChangeHandler}
        required={required}
        className="form-select"
      >
        <option value="">Select One</option>
        {options.map((o) => (
          <option
            value={o.value}
            key={`select__option__${o.value}__${o.label.replaceAll(" ", "")}`}
          >
            {o.label}
          </option>
        ))}
      </select>
    </div>
  );
}
