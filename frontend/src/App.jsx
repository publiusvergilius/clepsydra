import { useEffect, useState } from "react";
import { Greet, GetAllQuarta, CreateQuartum } from "../wailsjs/go/main/App";
import classes from "./App.module.css";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [resultQuarta, setResultQuarta] = useState([]);
  const [name, setName] = useState("");
  const updateName = (e) => setName(e.target.value);
  const updateResultText = (result) => setResultText(result);

  const updateResultQuarta = (result) => setResultQuarta(result);

  function getAllQuarta() {
    // CreateQuartum()
    console.log("executando...");
    const res = GetAllQuarta();
    Promise.resolve(res).then(JSON.parse).then(updateResultQuarta);
  }

  useEffect(() => {
    getAllQuarta();
  }, []);

  console.log(resultQuarta);
  /*return (
        <div id="App">
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={()=> {
			getAllQuarta();
		}}>Greet</button>
            </div>
	    <div>
	    {resultQuarta.map(q => (
		    <>
		    <p className={classes.primary_text} key={`id_${q?.id}`}>{q?.titulum}</p>
		    <p key={`pars_${q?.id}`}>{
			    q?.pars === 1 ? "Manhã"
			    : "Noite"
		    }</p>
		    <p key={q?.id}>{q?.hora}</p>
		    </>
	    ))}
	    </div>
        </div>
    )*/

  return (
    <div className={classes.layout}>
      <div className={classes.box} style={{ backgroundColor: "yellow" }}>
        {resultQuarta
          .filter((q) => q.pars === 1)
          .map((q) => (
            <>
              <p
                className={classes.primary_text}
                key={`
				    id_${q?.id}`}
              >
                {q?.titulum}
              </p>
              <p key={q?.id}>{q?.hora}</p>
            </>
          ))}
      </div>
      <div className={classes.box} style={{ backgroundColor: "green" }}></div>
      <div className={classes.box} style={{ backgroundColor: "red" }}></div>
      <div className={classes.box} style={{ backgroundColor: "blue" }}></div>
    </div>
  );
}

export default App;
