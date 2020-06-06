import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

function App() {
  const [inventoryState,setInventoryState] = useState([])


const handleRoot = () => {
  axios.get("http://localhost:8081/")
  .then((res) => console.log(res))
}
const handleAllItems = () => {
  axios.get("http://localhost:8081/all/")
  .then((res) => {
    setInventoryState(res.data);
    console.log(inventoryState);
  })
};

const handleAddItems = (name, itemtype) => {
  axios.post(`http://localhost:8081/add/${name}/${itemtype}`)
  .then((res) => console.log(res))
};

const handleRemove = (name) => {
  axios.delete(`http://localhost:8081/remove/${name}`)
  .then((res) => console.log(res))
};

const handleUpdate = (itemtype, name) => {
  axios.get(`http://localhost:8081/update/${itemtype}/${name}`)
  .then((res) => console.log(res))
};




  return (
    <div className="App">
      <button onClick={handleRoot}>Click me for rooot</button>
     <button onClick={handleAllItems}>Click me for all</button>
     <button onClick={handleAddItems("axe", "weapon")}>Click me to add</button>
     <button>Click me for remove</button>
     <button>Click me for update</button>

      <div className="map-division">
        {inventoryState.map((x) => <div style={{
        border:"2px solid black",
        height:"100px",
        width:"100px",
        }}><p>{x.Name}</p><p>{x.ItemType}</p></div>)}
      </div>

    </div>
  );
}

export default App;
