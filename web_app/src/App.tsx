import { TextField, Button } from "@mui/material";
import "./App.css";
import { useState } from "react";

function App() {
  const [userName, setUserName] = useState<String>("");
  const [password, setPassword] = useState<String>("");
  //
  return (
    <>
      <div>
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
        <Button>Log In</Button>
      </div>
    </>
  );
}

export default App;
