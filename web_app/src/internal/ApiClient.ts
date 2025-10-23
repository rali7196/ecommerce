export default class ApiGatewayClient {
  static ApiUrl: string = "http://localhost:3000";

  // static createAccount(username: string, password: string, callback: function): void {
  //   const requestBody: string = {
  //     "username":username,
  //     "password":password
  //   }
  //
  //   const requestInit: RequestInit = {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json"
  //     },
  //     body: JSON.stringify(requestBody)
  //   }
  //
  //   const route = ApiGatewayClient.ApiUrl + "/users";
  //
  //   const response = await fetch(route, requestInit).t
  //   return
  // }

  static async helloWorld(): Promise<any>{
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
