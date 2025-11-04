import { TextField, Button } from "@mui/material";
import "./App.css";
import { useState } from "react";
import "./internal/ApiClient.ts";
import ApiGatewayClient from "./internal/ApiClient.ts";
import type CreateUserResponse from "./types/responses/createUserResponse.ts";

function App() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  function createAccountCallback() {
    return (response: void | CreateUserResponse) => {
      if (response == null) {
        return;
      }

      console.log(response);
    };
  }

  return (
    <>
      <div className={"flexContainerColumn"}>
        <TextField
          placeholder={"Username"}
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        ></TextField>
        <TextField
          placeholder={"Password"}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        ></TextField>
        <Button>
          Log In
        </Button>
        <Button onClick={ () => {
          ApiGatewayClient.createAccount(email, password).then(
            createAccountCallback()
          )
        }
        }>Sign Up</Button>
      </div>
    </>
  );
}

export default App;
