import React from "react";

/** @typedef {{
 *  label: string
 *  type: string
 *  placeholder: string
 *  value: string
 *  onChangeHandler: Function
 *  required?: boolean
 * }} CardFormInputObject
 */

/** @param {CardFormInputObject} props */
export default function CardFormInput({
  label,
  type,
  placeholder,
  value,
  onChangeHandler,
  required = true,
}) {
  return (
    <div className="form-group mb-3">
      <label className="text-body-tertiary fw-light">{label}</label>
      <input
        type={type}
        className="form-control"
        placeholder={placeholder}
        value={value}
        onChange={onChangeHandler}
        required={required}
      />
    </div>
  );
}
