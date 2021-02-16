import React, { useEffect, useState } from "react";
import MainNavbar from "../components/MainNavbar";
import { AuthProps, privateRoute } from "../components/privateRoute";
import getInstances from "../services/instances_service";
import Head from "next/head";
import styles from "../styles/Dashboard.module.css";
import Container from "react-bootstrap/Container";
import RequestTable from "../components/Table";
import Spinner from "react-bootstrap/Spinner";

type Props = AuthProps;

const ACTION_URL = "http://localhost:3001/api/action";

const Dashboard = ({ auth }: Props) => {
  const [error, setError] = useState(null);
  const [requests, setRequests] = useState([]);
  const [requestState, setRequestState] = useState("Idle");
  const [isLoaded, setIsLoaded] = useState(false);

  const actionHandler = async (e: any) => {
    const data = JSON.stringify({
      netid: e.target.dataset.name,
      task: e.target.dataset.task,
    });
    const params = {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: data,
    };

    setRequestState("Waiting on response from server..");
    const response = await fetch(ACTION_URL, params);
    if (response.ok) {
      return response;
    } else {
      setError(true);
    }
  };

  useEffect(() => {
    getInstances(auth.authorizationString)
      .then((instances) => {
        setRequests(instances);
        setIsLoaded(true);
      })
      .catch((error) => {
        setError(error);
        setIsLoaded(true);
      });
  }, []);

  if (error) {
    return (
      <>
        <MainNavbar />
        <div className={styles.container}>
          <b>Error: </b>Failed to retrieve data
        </div>
      </>
    );
  } else if (!isLoaded) {
    return (
      <>
        <MainNavbar />
        <div className={styles.container}>
          <Head>
            <title>Loading...</title>
          </Head>
          <Spinner animation="boder" />
          <b>Loading...</b>
        </div>
      </>
    );
  } else {
    return (
      <div className={styles.container}>
        <Head>
          <title>Dashboard</title>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <MainNavbar />
        <div style={{ height: "100px" }}></div>
        <div className={styles.status}>{requestState}</div>
        <Container fluid>
          <RequestTable
            requests={requests}
            clickHandler={actionHandler}
          ></RequestTable>
        </Container>
      </div>
    );
  }
};

export default privateRoute(Dashboard);
