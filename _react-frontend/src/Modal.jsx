import classes from "./App.module.css";

export default function Modal({ children, isOpen, onClose }) {
  return (
    <div
      style={{ display: isOpen ? "flex" : "none" }}
      className={classes.modal}
    >
      <div className="modal-content">
        <span className={classes.close} onClick={onClose}>
          &times;
        </span>
        {children}
      </div>
    </div>
  );
}
