import Head from "next/head";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import MainNavbar from "../components/MainNavbar";
import React, { useState } from "react";
import { postLogin } from "../services/login_service";
import { redirectToRequestForm } from "../services/redirect_service";

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
      <Head>
        <title>UNR - Log In</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainNavbar loggedIn={false} />
      <div
        style={{
          boxShadow: "2px 2px 5px 4px rgba(0, 0, 0, 0.07)",
          marginLeft: "30vw",
          marginRight: "30vw",
          paddingTop: "2vh",
          display: "flex",
          marginTop: "15vh",
          flexDirection: "column",
          alignItems: "center",
          backgroundColor: "#FFFFFF",
        }}
      >
        <h3>Login</h3>
        <Form
          onSubmit={handleSubmit}
          style={{
            display: "flex",
            justifyContent: "flex-star",
            marginTop: "2vh",
            flexDirection: "column",
            alignItems: "center",
            marginBottom: "2vh",
          }}
        >
          <Form.Group>
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
          <Form.Group>
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
          <span>
            <Button type="submit">Login</Button>{" "}
            <Button variant="success" onClick={() => redirectToRequestForm()}>
              Request
            </Button>{" "}
          </span>
        </Form>
      </div>
    </>
  );
};

export default Login;
