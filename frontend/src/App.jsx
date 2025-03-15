import {useState} from 'react';
import './App.css';
import {Greet, GetAllQuarta, CreateQuartum} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [resultQuarta, setResultQuarta] = useState([]);
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    const updateResultQuarta = (result) => setResultQuarta(result);

    function getAllQuarta (){
	    	// CreateQuartum()	
	    	console.log("executando...")
		const res = GetAllQuarta()
	    	Promise.resolve(res).then(JSON.parse).then(updateResultQuarta);
    }

console.log(resultQuarta)
    return (
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
		    <p key={q?.id}>{q?.titulum}</p>
		    <p key={q?.id}>{q?.pars}</p>
		    <p key={q?.id}>{q?.hora}</p>
		    </>
	    ))}
	    </div>
        </div>
    )
}

export default App
