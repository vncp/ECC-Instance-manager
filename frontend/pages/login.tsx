import MainNavbar from "../components/MainNavbar";
import React, { useState } from "react";
import { postLogin } from "../services/login_service";

export type LoginInput = {
  netid: string;
  password: string;
};

const Login = () => {
  const defaultValues: LoginInput = {
    netid: "",
    password: "",
  };

  const [inputs, setInputs] = useState(defaultValues);

  const handleSubmit = async (e: React.ChangeEvent<any>) => {
    e.preventDefault();
    postLogin(inputs);
  };

  const handleInputChange = (e: React.ChangeEvent<any>) => {
    e.persist();
    setInputs({
      ...inputs,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <>
      <MainNavbar />
      <div
        style={{ display: "flex", justifyContent: "center", marginTop: "10vh" }}
      >
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="netid">Netid</label>
            <input
              type="netid"
              id="netid"
              name="netid"
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              onChange={handleInputChange}
            />
          </div>
          <button type="submit">Login</button>
        </form>
      </div>
    </>
  );
};

export default Login;
