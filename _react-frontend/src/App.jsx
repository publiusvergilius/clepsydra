import { useEffect, useReducer, useState } from "react";
import { Greet, GetAllQuarta, CreateQuartum } from "../wailsjs/go/main/App";
import classes from "./App.module.css";
import Modal from "./Modal";

/**
 * @typedef {Object} Quartum
 * @property {number} id
 * @property {string} titulum
 * @property {number} pars
 * @property {string} hora
 * @property {number} dies_id
 */

/**
 *
 * @param {*} state
 * @param {*} action
 * @returns
 */
function reducer(state, action) {
  switch (action.type) {
    case "titulum":
      return { ...state, titulum: action.payload };
    case "pars":
      return { ...state, pars: action.payload };
    case "hora":
      return { ...state, hora: action.payload };
    case "min":
      return { ...state, min: action.payload };
    default:
      return state;
  }
}

const initialState = {
  titulum: "",
  pars: 1,
  hora: "01",
  min: "00",
};

function App() {
  const [resultQuarta, setResultQuarta] = useState([]);
  const [isOpen, setIsOpen] = useState(false);
  const [state, dispatch] = useReducer(reducer, initialState);

  const updateResultQuarta = (result) => setResultQuarta(result);

  function getAllQuarta() {
    const res = GetAllQuarta();
    Promise.resolve(res).then(JSON.parse).then(updateResultQuarta);
  }

  function CreateQuartum() {
    const timestamp = `${state.hora}:${state.min}:00`;

    const reqBody = {
      titulum: state.titulum,
      pars: state.pars,
      hora: timestamp,
      dies_id: 1,
    };

    console.log(reqBody);
    const res = CreateQuartum(JSON.stringify(reqBody));
    console.log(res);
    // Promise.resolve(res).then(JSON.parse).then(updateResultQuarta);
  }

  function openModal() {
    setIsOpen(!isOpen);
  }

  /**
   * @param {React.FormEvent} e
   */
  function handleSubmit(e) {
    e.preventDefault();
    CreateQuartum();
    setIsOpen(false);
  }

  useEffect(() => {
    getAllQuarta();
  }, []);

  return (
    <>
      <div className={classes.add_button} onClick={openModal}>
        <span>+</span>
      </div>
      <Modal onClose={openModal} isOpen={isOpen}>
        <form action="" onSubmit={handleSubmit} className={classes.form}>
          <input
            type="text"
            placeholder="Titulum"
            value={state.titulum}
            required
            onChange={(e) =>
              dispatch({ type: "titulum", payload: e.target.value })
            }
          />
          <select
            name="pars"
            id="pars"
            style={{ width: "100%" }}
            defaultValue={1}
            onChange={(e) =>
              dispatch({ type: "pars", payload: e.target.value })
            }
          >
            <option value="1">Mane</option>
            <option value="2">Tarde</option>
            <option value="3">Nox</option>
            <option value="4">Meridianum</option>
          </select>
          <div
            style={{
              width: "100%",
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <div style={{ flex: 1 }}>
              <label htmlFor="hours" style={{ color: "black", marginRight: 5 }}>
                Horas
              </label>
              <select
                name="hours"
                id=""
                defaultValue="01"
                onChange={(e) =>
                  dispatch({ type: "hora", payload: e.target.value })
                }
              >
                <option value="00">00</option>
                <option value="01" defaultValue>
                  01
                </option>
                <option value="02">02</option>
                <option value="03">03</option>
                <option value="04">04</option>
                <option value="05">05</option>
                <option value="06">06</option>
              </select>
            </div>
            <div style={{ flex: 1 }}>
              <label
                htmlFor="minutes"
                style={{ color: "black", marginRight: 5 }}
              >
                Min
              </label>
              <select
                name="minutes"
                id=""
                onChange={(e) =>
                  dispatch({ type: "min", payload: e.target.value })
                }
              >
                <option value="00" defaultValue={"00"}>
                  00
                </option>
                <option value="15">15</option>
                <option value="30">30</option>
                <option value="45">45</option>
              </select>
            </div>
          </div>
          <button type="submit">Adicionar</button>
        </form>
      </Modal>
      <div className={classes.layout}>
        <Box num={1} color="yellow" data={resultQuarta} />
        <Box num={2} color="green" data={resultQuarta} />
        <Box num={3} color="red" data={resultQuarta} />
        <Box num={4} color="blue" data={resultQuarta} />
      </div>
    </>
  );
}

function Box({ num, color, data }) {
  return (
    <div className={classes.box} style={{ backgroundColor: color }}>
      {data
        .filter((q) => q.pars === num)
        .map((q) => (
          <div>
            <p
              className={classes.primary_text}
              key={`
				    id_${q?.id}`}
            >
              {q?.titulum}
            </p>
            <p key={q?.id}>{q?.hora}</p>
          </div>
        ))}
    </div>
  );
}

export default App;
