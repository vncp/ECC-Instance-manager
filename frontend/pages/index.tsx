import { useState, useEffect } from "react";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import RequestTable from "../components/Table";
import Spinner from "react-bootstrap/Spinner";

const BACKEND_URL = "http://localhost:3001/api/";
const TEST_URL = BACKEND_URL + "test";
const ACTION_URL = BACKEND_URL + "action/";

export default function Home() {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [requests, setRequests] = useState([]);
  const [requestState, setRequestState] = useState("Idle");

  const actionHandler = async (e: any) => {
    let data = JSON.stringify({
      netid: e.target.dataset.name,
      task: e.target.dataset.task,
    });
    setRequestState("Waiting on response from server..");
    const query_params = "?netid=" + e.target.dataset.name;
    const response = await fetch(
      ACTION_URL + e.target.dataset.task + query_params
    );
    console.log(response);
  };

  useEffect(() => {
    fetch(TEST_URL)
      .then((res: Response) => res.json())
      .then((result) => {
        setIsLoaded(true);
        setRequests(result);
      })
      .catch((error) => {
        setIsLoaded(true);
        setError(error);
        console.error(error);
      });
  }, []);

  if (error) {
    return (
      <div className={styles.container}>
        <b>Error:</b>Failed to retrieve data
      </div>
    );
  } else if (!isLoaded) {
    return (
      <div className={styles.container}>
        <Head>
          <title>Loading...</title>
        </Head>
        <Spinner animation="border"></Spinner>
        <b>Loading...</b>
      </div>
    );
  } else {
    return (
      <div className={styles.container}>
        <Head>
          <title>Remote Instance Management</title>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <Navbar
          collapseOnSelect
          expand="xl"
          bg="dark"
          variant="dark"
          fixed="top"
        >
          <Navbar.Brand href="#home">
            Remote Linux Instance Management
          </Navbar.Brand>
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />
          <Navbar.Collapse id="Responsive-navbar-nav">
            <Nav className="ml-auto" fill>
              <Nav.Item>
                <Nav.Link href="https:/remote.engr.unr.edu/">
                  Remote Home
                </Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="#Help">Documentation</Nav.Link>
              </Nav.Item>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
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
}
