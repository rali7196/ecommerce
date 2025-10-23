import { TextField, Button } from "@mui/material";
import "./App.css";
import { useState } from "react";
import "./internal/ApiClient.ts";
import ApiGatewayClient from "./internal/ApiClient.ts";

function App() {
  const [userName, setUserName] = useState<String>("");
  const [password, setPassword] = useState<String>("");

  return (
    <>
      <div className={"flexContainerColumn"}>
        <TextField
          placeholder={"Username"}
          value={userName}
          onChange={(e) => setUserName(e.target.value)}
        ></TextField>
        <TextField
          placeholder={"Password"}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        ></TextField>
        <Button
          onClick={() => {
            ApiGatewayClient.helloWorld()
              .then((response) => response.json())
              .then((json) => console.log(json));
          }}
        >
          Log In
        </Button>
        <Button>Sign Up</Button>
      </div>
    </>
  );
}

export default App;
