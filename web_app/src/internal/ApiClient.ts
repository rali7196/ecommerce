import {type CreateUserRequest, type CreateUserResponse} from "../types/createUser.ts";

export default class ApiGatewayClient {
  static ApiUrl: string = "http://localhost:3000";

  static async createAccount(email: string, password: string): Promise<CreateUserResponse | void> {
    const requestBody: CreateUserRequest = {
      email: email,
      password: password
    }

    const requestInit: RequestInit = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(requestBody)
    }

    const route: string = ApiGatewayClient.ApiUrl + "/users";

    try {
      const response: Response = await fetch(route, requestInit)
      return await response.json() as CreateUserResponse;
    }
    catch (e) {
      console.log(e)
    }
  }

  static async helloWorld(): Promise<Response | void> {
    const requestInit: RequestInit = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };

    const route: string = ApiGatewayClient.ApiUrl + "/"

    try {
      return await fetch(route, requestInit)
    }
    catch (e) {
      console.log(e)
    }
  }
}
