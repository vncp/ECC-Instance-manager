import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
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
    console.log(e);
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
        style={{
          marginLeft: "30vw",
          marginRight: "30vw",
          paddingBottom: "2vh",
          paddingTop: "1vh",
          display: "flex",
          justifyContent: "center",
          marginTop: "15vh",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <h1>Login</h1>
        <Form onSubmit={handleSubmit}>
          <Form.Group controlId="netid">
            <Form.Label>NetID</Form.Label>
            <Form.Control
              required
              type="netid"
              id="netid"
              name="netid"
              placeholder="NetID"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Form.Group controlId="password">
            <Form.Label>Password</Form.Label>
            <Form.Control
              required
              type="password"
              id="password"
              name="password"
              placeholder="Password"
              onChange={handleInputChange}
            />
          </Form.Group>
          <Button type="submit">Login</Button>
        </Form>
      </div>
    </>
  );
};

export default Login;
