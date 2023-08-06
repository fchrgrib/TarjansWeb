import React, {ChangeEvent, useState, useRef, useEffect} from 'react';
import './App.css';
import axios from "axios";
import "."
import Graph from "react-graph-vis"

function App() {
  const [fileContent, setFileContent] = useState<[string,string][]>([]);
  const [resultGraph, setResultGraph] = useState<[string,string][]>([]);
  const [type, setType] = useState("scc")
  const [isButtonClicked, setIsButtonClicked] = useState(false)

  const handleFileChange = async (event: ChangeEvent<HTMLInputElement>) => {
    const file = event?.target?.files?.[0];
    if (file) {
      const text = await file.text();
      const lines = text.split('\n');
      const pairs = lines.map(line => line.replace('\r','').split(' ') as [string, string]);
      setFileContent(pairs);
    }
  };

  const postButton = async () =>{
    if (!isButtonClicked && fileContent){
      setIsButtonClicked(true)
      try {
        const response = await axios.post("http://localhost:8080/post-graph",{
          input:fileContent,
          type: type
        })
        if (response.status<400) {
          console.log(response.data.graph)
          setResultGraph(response.data.graph)
        }
        setIsButtonClicked(false)
      }catch (e) {
        console.log(e)
        setIsButtonClicked(false)
      }
    }
  }

  return (
      <div className="w-full h-screen flex justify-center mt-[5rem]">
        <div className="w-[30rem] h-auto flex-col">
          <div className="w-[30rem] h-auto flex-col justify-between">
            <input type="file" className="w-1/2 select-none" onChange={handleFileChange} />
            <button className="w-1/2 px-6 mt-4 mb-4-2 py-2 bg-fuchsia-400 rounded-lg text-white font-semibold select-none" onClick={postButton}>
                POST
            </button>
          </div>
          <select value={type} onChange={e=>{setType(e.target.value)}} className="border border-gray-500 p-1 w-[10rem] rounded-lg select-none">
            <option key="scc" value="scc">scc</option>
            <option key="bridge" value="bridge">bridge</option>
          </select>
          <button className="w-[10rem] p-2 bg-pink-400 rounded-lg text-white font-semibold select-none mt-4 ml-2" onClick={()=>{window.location.reload()}}>
            CLEAR GRAPH
          </button>
          <div>{(type === "scc")? "SCC:":"Bridges:"}</div>
          {resultGraph&&resultGraph.map((item,index)=>{
            const edges = []
            if (item.length>1){
              for(let i=0; i<item.length-1;i++){
                edges.push({from: item[i], to:item[i+1]})
              }
            }
            let nodes = []
            if (type==="scc"){
              for(let i=0;i<item.length-1;i++){
                nodes.push({ id: item[i], label: item[i]})
              }
            }else{
              nodes = item.map((node,idx) => ({ id: node, label: node}))
            }
            const graph = {
              nodes: nodes,
              edges: edges,
            }

            return <Graph key={index} graph={graph}/>;
          })}
        </div>
      </div>
  );
}

export default App;
